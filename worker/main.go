package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/garyburd/redigo/redis"
	"github.com/scriptnull/badgeit/common"
	"github.com/scriptnull/badgeit/contracts"
	"github.com/scriptnull/badgeit/worker/downloader"
)

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

	listenRedisTaskQueue(redisConn)
}

func listenRedisTaskQueue(redisConn redis.Conn) {
	log.Println("==== ==== ==== ====")
	log.Println("Waiting for message to arrive on badge:worker redis queue")
	log.Println("==== ==== ==== ====")
	payload, err := redis.Strings(redisConn.Do("BRPOP", "badge:worker", 0))
	if err != nil {
		log.Fatalln("Failed to do blocking RPOP on badge:worker", err)
	}
	log.Printf("Recieved Payload %+v", payload)
	taskPayload := payload[1]

	log.Printf("Starting Task for message: %s \n", taskPayload)
	executeTask([]byte(taskPayload), redisConn)
	log.Printf("Finished Task for message: %s \n", taskPayload)

	listenRedisTaskQueue(redisConn)
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

func executeTask(message []byte, redisConn redis.Conn) {

	// Parse input message
	payload := struct {
		Remote   string
		Download string
		Callback string
	}{}
	err := json.Unmarshal(message, &payload)
	if err != nil {
		log.Printf("Error Parsing the payload %d", err)
		return
	}

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
