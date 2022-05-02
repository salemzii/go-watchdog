package storages

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/salemzii/go-watchdog/service"
)

var s3sess *s3.S3
var sesss *session.Session
var REGION, BUCKET string

func create_file() (file os.File, name string, err error) {
	f, err := os.OpenFile("out.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return *f, "", err
	}
	for _, v := range "AP" {
		fmt.Fprintf(f, "%b\n", v)
	}

	return *f, f.Name(), nil
}

// function to upload file to s3
func AwsStorageCheck(st *Storage) service.ServiceCheck {

	rand.Seed(time.Now().UnixNano())

	REGION = st.Region
	BUCKET := st.BUCKET
	log.Println(REGION, BUCKET)

	// start aws sessions
	sesss := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	}))

	s3sess = s3.New(sesss)

	file, filename, err := create_file()

	if err != nil {
		return service.HandleError("aws", err)
	}

	defer file.Close()
	uploader := s3manager.NewUploader(sesss)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(st.BUCKET),
		Key:    aws.String(filename),
		Body:   &file,
	})
	if err != nil {
		// Print the error and exit.
		return service.HandleError("aws", err)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, BUCKET)
	return service.HandleSuccess("aws", err)
}
