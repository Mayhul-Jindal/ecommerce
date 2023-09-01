package bucket

import (
	"fmt"
	"io"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Bucket interface {
	UploadFile(fileReader io.Reader, fileHeader *multipart.FileHeader) (string, error)
}

type s3Bucket struct {
	AWSRegion     string
	AWSAccessKey  string
	AWSSecretKey  string
	AWSBucketName string
}

func NewS3Bucket(AWSRegion string, AWSAccessKey string, AWSSecretKey string, AWSBucketName string) Bucket {
	return &s3Bucket{
		AWSRegion:     AWSRegion,
		AWSAccessKey:  AWSAccessKey,
		AWSSecretKey:  AWSSecretKey,
		AWSBucketName: AWSBucketName,
	}
}

func (s *s3Bucket) UploadFile(fileReader io.Reader, fileHeader *multipart.FileHeader) (string, error) {
	awsSession, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String(s.AWSRegion),
			Credentials: credentials.NewStaticCredentials(s.AWSAccessKey, s.AWSSecretKey, ""),
		},
	})
	if err != nil {
		panic(err)
	}

	uploader := s3manager.NewUploader(awsSession)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.AWSBucketName),
		Key:    aws.String(fileHeader.Filename),
		Body:   fileReader,
	})
	if err != nil {
		return "", err
	}

	// Get the URL of the uploaded file
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s.AWSBucketName, fileHeader.Filename)

	return url, nil
}
