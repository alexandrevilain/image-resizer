package storage

import (
	"fmt"
	"path/filepath"

	minio "github.com/minio/minio-go"
)

type StorageClient struct {
	client   *minio.Client
	endpoint string
	bucket   string
	ssl      bool
}

func NewClient(endpoint, bucket, accessKeyID, secretAccessKey string, useSSL bool) (*StorageClient, error) {
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		return nil, err
	}
	return &StorageClient{
		client:   minioClient,
		endpoint: endpoint,
		bucket:   bucket,
		ssl:      useSSL,
	}, nil
}

func (c *StorageClient) getPublicUrl(objectName string) string {
	protocol := "http"
	if c.ssl {
		protocol = "https"
	}
	return fmt.Sprintf("%s://%s/%s/%s", protocol, c.endpoint, c.bucket, objectName)
}

func (c *StorageClient) UploadFile(path string) (string, error) {
	objectName := filepath.Base(path)
	_, err := c.client.FPutObject(c.bucket, objectName, path, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return c.getPublicUrl(objectName), nil
}
