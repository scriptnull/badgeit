package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/garyburd/redigo/redis"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

const (
	downloadTypeGit  string = "git"
	downloadTypeCurl string = "curl"
)

func main() {

	log.Println("Booting Badgeit API server")

	log.Println("Initializing Message Queue")
	amqpConnection, workerQueue, workerChannel, err := initMessageQueue()
	if err != nil {
		log.Fatalln("Failed to initialize message queue", err)
	}
	defer amqpConnection.Close()
	defer workerChannel.Close()
	log.Println("Initialized Message Queue successfully")

	log.Println("Initializing Redis Message Queue")
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOSTNAME"), os.Getenv("REDIS_PORT")))
	if err != nil {
		log.Fatalln("Failed to initialize redis message queue", err)
	}
	defer conn.Close()
	log.Println("Initialized Redis Message Queue successfully")

	log.Println("Initializing API Server")
	initAPIServer(workerQueue, workerChannel, conn)
}

func initMessageQueue() (*amqp.Connection, *amqp.Queue, *amqp.Channel, error) {
	username := os.Getenv("RABBIT_USERNAME")
	password := os.Getenv("RABBIT_PASSWORD")
	hostname := os.Getenv("RABBIT_HOSTNAME")
	port := os.Getenv("RABBIT_PORT")
	conStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, hostname, port)
	conn, err := amqp.Dial(conStr)
	if err != nil {
		return nil, nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, nil, err
	}

	q, err := ch.QueueDeclare(
		"badgeit.worker", // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		return nil, nil, nil, err
	}

	return conn, &q, ch, nil
}

func initAPIServer(workerQueue *amqp.Queue, workerChannel *amqp.Channel, redisConn redis.Conn) {
	r := gin.Default()

	r.POST("/test/callback", func(c *gin.Context) {
		io.Copy(os.Stdout, c.Request.Body)
		defer c.Request.Body.Close()
		c.JSON(http.StatusOK, gin.H{
			"test": "ok",
		})
		return
	})

	r.GET("/badges", func(c *gin.Context) {
		downloadType, hasDownloadType := c.GetQuery("download")
		if !hasDownloadType || downloadType == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "download is a required",
			})
			return
		}
		if downloadType != downloadTypeGit && downloadType != downloadTypeCurl {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Allowed download types are %s, %s", downloadTypeGit, downloadTypeCurl),
			})
			return
		}

		remote, hasRemote := c.GetQuery("remote")
		if !hasRemote || remote == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "remote is a required",
			})
			return
		}

		callback, hasCallback := c.GetQuery("callback")
		if !hasCallback || callback == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "callback is a required",
			})
			return
		}

		payload := gin.H{
			"download": downloadType,
			"remote":   remote,
			"callback": callback,
		}

		// check for cached data

		// check if worker is already working on badge computation

		// queue a task for the worker
		jsonPayload, _ := json.Marshal(payload)
		err := workerChannel.Publish(
			"",               // exchange
			workerQueue.Name, // routing key
			false,            // mandatory
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(jsonPayload),
			})
		if err != nil {
			log.Println("Unable to queue request", err)
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": "Unable to queue request",
			})
			return
		}

		_, err = redisConn.Do("LPUSH", "badge:worker", []byte(jsonPayload))
		if err != nil {
			log.Println("Unable to queue request", err)
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": "Unable to queue request",
			})
			return
		}

		c.JSON(http.StatusAccepted, payload)
		return
	})
	r.Run(":8080")
}
