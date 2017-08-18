package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nu7hatch/gouuid"

	"github.com/garyburd/redigo/redis"
	"github.com/scriptnull/badgeit/common"
	"github.com/scriptnull/badgeit/contracts"
	"github.com/scriptnull/badgeit/worker/downloader"
)

var workerID string

type workerStatus string

const (
	statusInitialized workerStatus = "INITIALIZED"
	statusWaiting     workerStatus = "WAITING"
	statusProcessing  workerStatus = "PROCESSING"
)

func setWorkerStatus(conn redis.Conn, status workerStatus) error {
	key := fmt.Sprintf("worker:%s:status", workerID)
	_, err := redis.String(conn.Do("SET", key, status))
	return err
}

func main() {
	log.Println("Booting Badgeit worker")

	log.Println("Checking status of clone dir")
	cloneDir := os.Getenv("CLONE_DIR")
	if _, err := os.Stat(cloneDir); os.IsNotExist(err) {
		log.Fatalln("CLONE_DIR is not present. Error: ", err)
	}
	log.Println("Cloning Dir found")

	redisConn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOSTNAME"), os.Getenv("REDIS_PORT")))
	if err != nil {
		log.Fatalln("Failed to initialize redis message queue", err)
	}
	defer redisConn.Close()
	log.Println("Initialized Redis Message Queue successfully")

	workerUUID, err := uuid.NewV4()
	if err != nil {
		log.Fatalln("Failed to generate workerId", err)
	}
	workerID = workerUUID.String()
	log.Println("worker ID is ", workerID)

	setWorkerStatus(redisConn, statusInitialized)

	listenRedisTaskQueue(redisConn)
}

func listenRedisTaskQueue(redisConn redis.Conn) {
	log.Println("==== ==== ==== ====")
	log.Println("Waiting for message to arrive on badge:worker redis queue")
	log.Println("==== ==== ==== ====")
	setWorkerStatus(redisConn, statusWaiting)

	payload, err := redis.Strings(redisConn.Do("BRPOP", "badge:worker", 0))
	if err != nil {
		log.Fatalln("Failed to do blocking RPOP on badge:worker", err)
	}
	log.Printf("Recieved Payload %+v", payload)
	taskPayload := payload[1]

	log.Printf("Starting Task for message: %s \n", taskPayload)
	workerTask := newTask([]byte(taskPayload), redisConn)
	if workerTask != nil {
		workerTask.removeFromQueuedSet()
		workerTask.addToProcessingSet()
		executeTask(workerTask, redisConn)
		workerTask.removeFromProcessingSet()
	}
	log.Printf("Finished Task for message: %s \n", taskPayload)

	listenRedisTaskQueue(redisConn)
}

type task struct {
	Remote    string
	Download  string
	Callback  string
	RedisConn redis.Conn
}

func newTask(message []byte, conn redis.Conn) *task {
	var t task
	err := json.Unmarshal(message, &t)
	if err != nil {
		log.Println("Error Parsing the payload", err)
		return nil
	}
	t.RedisConn = conn
	return &t
}

func (t *task) addToProcessingSet() error {
	_, err := t.RedisConn.Do("SADD", "badgeit:processingRemotes", t.Remote)
	if err != nil {
		log.Println("Error adding remote to processing set", err)
		return err
	}
	return nil
}

func (t *task) removeFromProcessingSet() error {
	_, err := t.RedisConn.Do("SREM", "badgeit:processingRemotes", t.Remote)
	if err != nil {
		log.Println("Error adding remote to processing set", err)
		return err
	}
	return nil
}

func (t *task) removeFromQueuedSet() error {
	_, err := t.RedisConn.Do("SREM", "badgeit:queuedRemotes", t.Remote)
	if err != nil {
		log.Println("Error adding remote to processing set", err)
		return err
	}
	return nil
}

type taskResult struct {
	CallbackURL string
	RemoteURL   string
	// TODO: add callback headers
	Badges []common.Badge
	Error  string
}

func (t taskResult) toJSON() (string, error) {
	jsonPayload, err := json.Marshal(map[string]interface{}{
		"badges": t.Badges,
		"error":  t.Error,
		"remote": t.RemoteURL,
	})

	return string(jsonPayload), err
}

func executeTask(payload *task, redisConn redis.Conn) {
	setWorkerStatus(redisConn, statusProcessing)

	// Create temporary directory for download operation
	cloneDir := os.Getenv("CLONE_DIR")
	dir, err := ioutil.TempDir(cloneDir, "repo")
	if err != nil {
		log.Fatalln("Error creating temporary folder: ", err)
	}
	defer os.RemoveAll(dir)

	// Initiate taskResult for reporting back to the callback server\
	callbackResponse := taskResult{
		CallbackURL: payload.Callback,
		RemoteURL:   payload.Remote,
	}

	// Download the repository
	d := downloader.NewDownloader(downloader.DownloaderOptions{
		Type:   payload.Download,
		Remote: payload.Remote,
		Path:   dir,
	})
	log.Println("Downloading the repository: ", payload.Remote, " @ ", dir)
	err = d.Download()
	if err != nil {
		errorStr := fmt.Sprintln("Error Downloading repository: ", err)
		callbackResponse.Error = errorStr
		callback(callbackResponse)
		return
	}
	log.Println("Downloading complete @ ", dir)

	callbackResponse.Badges = contracts.PossibleBadges(dir)
	log.Printf("Detected %d possible badges \n", len(callbackResponse.Badges))

	// send badges to callback URL
	log.Println("Sending badges to callback URL")
	err = callback(callbackResponse)
	if err != nil {
		log.Println("Error While Posting callback: ", err)
	}
	log.Println("Sending attempt to callback URL complete")

	// cache badges
	log.Println("Attempting to cache badges")
	jsonPayload, err := callbackResponse.toJSON()
	if err != nil {
		log.Println("Error while Marshalling JSON", err)
		return
	}
	_, err = redisConn.Do("SET", fmt.Sprintf("badge:%s", callbackResponse.RemoteURL), jsonPayload)
	if err != nil {
		log.Println("Cache Miss with error: ", err)
	}
	log.Println("Caching Attempt complete")
}

func callback(result taskResult) error {
	jsonPayload, err := result.toJSON()
	if err != nil {
		return err
	}

	log.Println("Response Payload ", jsonPayload)
	_, err = http.Post(result.CallbackURL, "application/json", strings.NewReader(jsonPayload))
	if err != nil {
		return err
	}
	return nil
}
