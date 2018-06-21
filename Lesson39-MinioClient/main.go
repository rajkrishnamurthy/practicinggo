package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/minio/minio-go"
)

const (
	endpoint        = "localhost:9000"       // This should point to the real Minio end point
	accessKeyID     = "T1CSHX5X8E1CL2C6JFZL" // Dynamically generated Access Key and Secret Access Key
	secretAccessKey = "ejqzU8RZz7nXiCBPTmJvLK/R+w3acBygE6rRzJkr"
	useSSL          = false
	defBucketName   = "continube" // Name of the S3 like Bucket that you want to create in Minio
	downloadPath    = ""          // Default download path

)

func main() {

	// log.Println("server started")
	// http.HandleFunc("/store", handleStorageServices)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	bucketName := "demo"
	fileName := "./multipartclient.go"

	var err, errNew error
	minioClient, err := registerMinio(bucketName)
	if err != nil {
		errNew = errors.New("Cannot connect with the Minio Bucket")
		log.Printf("%v", errNew)
	}

	if err = uploadFile(minioClient, bucketName, fileName); err != nil {
		errNew = errors.New("Cannot Upload file into Minio")
		log.Printf("%v", errNew)
	}

}

func handleStorageServices(w http.ResponseWriter, r *http.Request) {
	var err, errNew error
	var jsonbody json.RawMessage
	switch r.Method {
	case "POST":
		if jsonbody, err = ioutil.ReadAll(r.Body); err != nil {
			errNew = errors.New("Cannot read the JSON Body on http/POST")
			log.Printf("%v", errNew)
			http.Error(w, errNew.Error(), http.StatusBadRequest)
		}
		switch {
		case len(jsonbody) < 1:
			errNew = errors.New("JSON Body not available on http/POST")
			log.Printf("%v", errNew)
			http.Error(w, errNew.Error(), http.StatusBadRequest)
		default:
			bucketName := r.URL.Query().Get("bucket")
			storeFile(w, jsonbody, bucketName)
		}

	}

}

func storeFile(w http.ResponseWriter, jsonbody json.RawMessage, bucketName string) {
	var err, errNew error
	minioClient, err := registerMinio(bucketName)
	if err != nil {
		errNew = errors.New("Cannot connect with the Minio Bucket")
		http.Error(w, errNew.Error(), http.StatusBadRequest)
		log.Printf("%v", errNew)
	}

	if err = uploadFile(minioClient, bucketName, ""); err != nil {
		errNew = errors.New("Cannot Upload file into Minio")
		http.Error(w, errNew.Error(), http.StatusBadRequest)
		log.Printf("%v", errNew)
	}

}

func registerMinio(bucketName string) (*minio.Client, error) {

	var err error
	if bucketName == "" {
		bucketName = defBucketName
	}
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	log.Printf("%#v\n", minioClient) // minioClient is now setup

	exists, err := minioClient.BucketExists(bucketName)
	if !exists && err == nil {
		// Bucket does not exist. Create one
		err = minioClient.MakeBucket(bucketName, "")
		if err != nil {
			log.Printf("%v", err)
			return nil, err
		}
	}
	if err == nil && exists {
		log.Printf("We already own %s\n", bucketName)
	}

	return minioClient, err
}

func uploadFile(minioClient *minio.Client, bucketName string, fileName string) (err error) {
	// objectName := "file"    // Name of the Object (File) that you want to store in the Minio Bucket
	// filePath := "./main.go" // Fully qualified File Path where the File is present, that you need to upload into Miino

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	objectName := fileInfo.Name()

	contentType, err := detectMimeType(fileName) // HTTP content-type of the file
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	// Upload the file with FPutObject
	n, err := minioClient.FPutObject(bucketName, objectName, fileName, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)

	return nil
}

func downloadFile(minioClient *minio.Client, bucketName string, objectName string) (err error) {

	fileName := strings.Join([]string{downloadPath, objectName}, string(os.PathSeparator))

	// Download the file with FGetObject
	err = minioClient.FGetObject(bucketName, objectName, fileName, minio.GetObjectOptions{})
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	log.Printf("Successfully downloaded %s\n", objectName)

	return nil
}

func detectMimeType(fileName string) (contentType string, err error) {

	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return "", err
	}
	// Reset the read pointer if necessary.
	file.Seek(0, 0)

	// Always returns a valid content-type and "application/octet-stream" if no others seemed to match.
	contentType = http.DetectContentType(buffer[:n])

	if contentType == "" {
		contentType = "application/octet-stream"
	}
	return contentType, nil
}
