package storages

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var s3sess *s3.S3
var sesss *session.Session
var REGION, BUCKET string

func init() {

	rand.Seed(time.Now().UnixNano())

	REGION = os.Getenv("REGION")
	BUCKET := os.Getenv("BUCKET")
	fmt.Println(REGION, BUCKET)

	// start aws sessions
	sesss := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	}))

	// Setup the S3 Upload Manager. Also see the SDK doc for the Upload Manager
	// for more information on configuring part size, and concurrency.
	//
	// http://docs.aws.amazon.com/sdk-for-go/api/service/s3/s3manager/#NewUploader

	// initiate s3 with created session object
	s3sess = s3.New(sesss)

}

func create_file() (file os.File, name string, err error) {
	f, err := os.OpenFile("out.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return *f, "", err
	}
	for _, v := range "AP" {
		fmt.Fprintf(f, "%b\n", v)
	}
	defer f.Close()
	return *f, f.Name(), nil
}

// function to upload file to s3
func AwsStorageCheck(st *Storage) map[string]string {

	file, filename, err := create_file()

	if err != nil {
		return map[string]string{"status": "Fail", "service": "aws", "error": err.Error()}
	}

	uploader := s3manager.NewUploader(sesss)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(st.BUCKET),
		Key:    aws.String(filename),
		Body:   &file,
	})
	if err != nil {
		// Print the error and exit.
		return map[string]string{"status": "Fail", "service": "aws", "error": err.Error()}
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, BUCKET)
	return map[string]string{"status": "ok", "service": "aws"}
	/*
		_, err := s3sess.PutObject(&s3.PutObjectInput{
			Bucket:               aws.String(BUCKET),
			Key:                  aws.String(imageName),
			ACL:                  aws.String("private"),
			Body:                 bytes.NewReader(body),
			ContentLength:        aws.Int64(contentLength),
			ContentType:          aws.String(http.DetectContentType(body)),
			ContentDisposition:   aws.String("attachment"),
			ServerSideEncryption: aws.String("AES256"),
		})
	*/
}
