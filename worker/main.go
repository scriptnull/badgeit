package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/scriptnull/badgeit/worker/downloader"
	"github.com/streadway/amqp"
)

func main() {
	log.Println("Booting Badgeit worker")

	log.Printf("Setting up connection to badgeit queue")
	username := os.Getenv("RABBIT_USERNAME")
	password := os.Getenv("RABBIT_PASSWORD")
	hostname := os.Getenv("RABBIT_HOSTNAME")
	port := os.Getenv("RABBIT_PORT")
	conStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, hostname, port)
	conn, err := amqp.Dial(conStr)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"badgeit.worker", // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Starting Task for message: %s", d.Body)
			executeTask(d.Body)
			log.Printf("Finished Task for message: %s", d.Body)
			d.Ack(false)
		}
	}()

	log.Printf("Booted Badgeit Worker. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

type taskResult struct {
	CallbackURL string
	// TODO: add callback headers
	Badges string
	Error  string
}

func executeTask(message []byte) {

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
	dir, err := ioutil.TempDir("", "repo")
	if err != nil {
		log.Fatalln("Error creating temporary folder: ", err)
	}
	defer os.RemoveAll(dir)

	// Initiate taskResult for reporting back to the callback server\
	callbackResponse := taskResult{
		CallbackURL: payload.Callback,
	}

	// Download the repository
	d := downloader.NewDownloader(downloader.DownloaderOptions{
		Type:   payload.Download,
		Remote: payload.Remote,
		Path:   dir,
	})
	log.Println("Downloading the repository: ", payload.Remote)
	err = d.Download()
	if err != nil {
		errorStr := fmt.Sprintln("Error Downloading repository: ", err)
		callbackResponse.Error = errorStr
		callback(callbackResponse)
		return
	}
	log.Println("Downloading complete @ ", dir)

	wd, err := os.Getwd()
	if err != nil {
		errorStr := fmt.Sprintln("Error Getting Working Directory: ", err)
		callbackResponse.Error = errorStr
		callback(callbackResponse)
		return
	}

	result, err := exec.Command(filepath.Join(wd, "badgeit"), "-f", "all-json", dir).Output()
	if err != nil {
		errorStr := fmt.Sprintln("Error Executing badgeit: ", err)
		callbackResponse.Error = errorStr
		callback(callbackResponse)
		return
	}

	callbackResponse.Badges = string(result)
	err = callback(callbackResponse)
	if err != nil {
		log.Println("Error While Posting callback: ", err)
	}
}

func callback(result taskResult) error {
	if result.Error == "" {
		log.Print(result.Error)
	}
	jsonPayload, err := json.Marshal(map[string]interface{}{
		"badges": result.Badges,
		"error":  result.Error,
	})
	if err != nil {
		return err
	}
	_, err = http.Post(result.CallbackURL, "application/json", strings.NewReader(string(jsonPayload)))
	if err != nil {
		return err
	}
	return nil
}
