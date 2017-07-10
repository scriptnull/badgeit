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
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
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

func executeTask(message []byte) {
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

	dir, err := ioutil.TempDir("", "repo")
	if err != nil {
		log.Fatalln("Error creating temporary folder: ", err)
	}
	defer os.RemoveAll(dir)

	d := downloader.NewDownloader(downloader.DownloaderOptions{
		Type:   payload.Download,
		Remote: payload.Remote,
		Path:   dir,
	})
	log.Println("Downloading the repository: ", payload.Remote)
	err = d.Download()
	if err != nil {
		log.Println("Error Downloading repository: ", err)
		// report error
	}
	log.Println("Downloading complete @ ", dir)

	wd, err := os.Getwd()
	if err != nil {
		log.Println("Error Getting Working Directory: ", err)
		// report error
	}

	result, err := exec.Command(filepath.Join(wd, "badgeit"), "-f", "all-json", dir).Output()
	if err != nil {
		log.Println("Error Executing badgeit: ", err)
		// report error
	}

	err = callback(payload.Callback, result)
	if err != nil {
		log.Println("Error While Posting callback: ", err)
	}
}

func callback(responseURL string, buf []byte) error {
	jsonPayload, err := json.Marshal(map[string]string{
		"badges": string(buf),
		"error":  "hello",
	})
	if err != nil {
		return err
	}
	_, err = http.Post(responseURL, "application/json", strings.NewReader(string(jsonPayload)))
	if err != nil {
		return err
	}
	return nil
}
