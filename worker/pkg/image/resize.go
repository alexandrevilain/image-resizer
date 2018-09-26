package image

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/labstack/gommon/log"
)

type ResizeImageJob struct {
	ID          string `json:"id"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	OriginalURL string `json:"originalUrl"`
	ResultURL   string `json:"resultUrl"`
}

func NewResizerJob(data []byte) *ResizeImageJob {
	var job ResizeImageJob
	dec := json.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&job); err != nil {
		log.Errorf("Unable to decode message!")
	}
	return &job
}

func (job *ResizeImageJob) Run() (string, error) {
	// Create temp directory and files
	dir, err := ioutil.TempDir("", "resizer")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(dir)
	originalFile, err := ioutil.TempFile(dir, "original")
	if err != nil {
		return "", err
	}
	if err := downloadImage(job.OriginalURL, originalFile); err != nil {
		return "", err
	}

	originalFilePath := originalFile.Name()
	extension := filepath.Ext(job.OriginalURL)
	name := fmt.Sprintf("%s_%s", strings.Replace(filepath.Base(job.OriginalURL), extension, "", -1), "resized")
	destinationFilePath := dir + name + extension
	err = job.resizeImage(originalFilePath, destinationFilePath)
	if err != nil {
		return "", err
	}
	return destinationFilePath, nil
}

func (job *ResizeImageJob) resizeImage(originalFilePath, destinationFilePath string) error {
	log.Infof("Cropping image: %s", originalFilePath)
	src, err := imaging.Open(originalFilePath)
	if err != nil {
		return fmt.Errorf("failed to open orginal image: %v", err)
	}
	resizedImage := imaging.Resize(src, job.Width, job.Height, imaging.Lanczos)
	err = imaging.Save(resizedImage, destinationFilePath)
	if err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}
	return nil
}

func downloadImage(url string, destinationFile *os.File) error {
	log.Infof("Downloading: %s", url)
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Write the body to file
	_, err = io.Copy(destinationFile, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
