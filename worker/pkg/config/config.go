package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Configuration is the struct used to store all configuration
type Configuration struct {
	NatsServers              string `default:"nats://localhost:4222" envconfig:"NATS_SERVERS"`
	NatsQueue                string `default:"jobs" envconfig:"NATS_QUEUE"`
	PostgresConnectionString string `default:"postgres://supinfo:supinfo@localhost:5432/images?sslmode=disable" envconfig:"POSTGRES_CONNECTION_STRING"`
	StorageServer            string `default:"localhost:9000" envconfig:"STORAGE_SERVER"`
	StorageBucket            string `default:"uploads" envconfig:"STORAGE_BUCKETNAME"`
	StorageAccessKeyID       string `default:"supinfo" envconfig:"STORAGE_ACCESSKEY"`
	StorageSecretAccessKey   string `default:"supinfo1234" envconfig:"STORAGE_SECRETKEY"`
	StorageSSL               bool   `default:"false" envconfig:"STORAGE_SSL"`
}

// GetFromEnv is returning the configuration of the worker
// populated with env variables
func GetFromEnv() (*Configuration, error) {
	var configuration Configuration
	err := envconfig.Process("", &configuration)
	if err != nil {
		return &configuration, err
	}
	return &configuration, nil
}
