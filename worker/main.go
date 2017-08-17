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
	log.Printf("Starting Task for message: %s", taskPayload)
	executeTask([]byte(taskPayload), redisConn)
	log.Printf("Finished Task for message: %s", taskPayload)
}

type taskResult struct {
	CallbackURL string
	RemoteURL   string
	// TODO: add callback headers
	Badges []common.Badge
	Error  string
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

	err = callback(callbackResponse)
	if err != nil {
		log.Println("Error While Posting callback: ", err)
	}

	listenRedisTaskQueue(redisConn)
}

func callback(result taskResult) error {
	if result.Error == "" {
		log.Println(result.Error)
	}
	jsonPayload, err := json.Marshal(map[string]interface{}{
		"badges": result.Badges,
		"error":  result.Error,
		"remote": result.RemoteURL,
	})
	if err != nil {
		return err
	}
	log.Println("Response Payload ", string(jsonPayload))
	_, err = http.Post(result.CallbackURL, "application/json", strings.NewReader(string(jsonPayload)))
	if err != nil {
		return err
	}
	return nil
}
