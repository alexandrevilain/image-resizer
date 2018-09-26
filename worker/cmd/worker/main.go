package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/alexandrevilain/image-resizer/worker/pkg/config"
	"github.com/alexandrevilain/image-resizer/worker/pkg/database"
	"github.com/alexandrevilain/image-resizer/worker/pkg/image"
	"github.com/alexandrevilain/image-resizer/worker/pkg/queue"
	"github.com/alexandrevilain/image-resizer/worker/pkg/storage"
)

func main() {
	config, err := config.GetFromEnv()
	if err != nil {
		// If we can't parse configuration, stop app:
		log.Fatalf(err.Error())
	}

	jobs := make(chan []byte)
	storage, err := storage.NewClient(config.StorageServer, config.StorageBucket, config.StorageAccessKeyID, config.StorageSecretAccessKey, config.StorageSSL)
	if err != nil {
		log.Fatalf(err.Error())
	}

	db, err := database.Connect(config.PostgresConnectionString)

	queueConnection, err := queue.Connect(config.NatsServers, config.NatsQueue)
	if err != nil {
		log.Fatalf(err.Error())
	}
	queueConnection.GetIncomingMessages(jobs)

	go func() {
		for {
			select {
			case job := <-jobs:
				imageCropper := image.NewResizerJob(job)
				path, err := imageCropper.Run()
				if err != nil {
					log.Printf("Error while cropping image: %s", err.Error())
				}
				url, err := storage.UploadFile(path)
				if err != nil {
					log.Printf("Error while saving image: %s", err.Error())
				}
				imageCropper.ResultURL = url
				err = db.Create(*imageCropper)
				if err != nil {
					log.Printf("Error while saving image infos in database: %s", err.Error())
				}
			}
		}
	}()

	// Check for a closing message:
	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, os.Interrupt, os.Kill)
	sig := <-sigquit
	log.Printf("Gracefully shutting down worker, caught: %+v", sig)
	close(jobs)
	if err := queueConnection.Close(); err != nil {
		log.Printf("Unable to close connection to worker: %v", err)
	} else {
		log.Println("Worker stopped")
	}
}
