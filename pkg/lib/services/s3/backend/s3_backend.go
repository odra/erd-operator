package backend

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Backend struct {
	session *session.Session
}

func (b *Backend) New(region string, apiKey string, apiToken string) error {
	s, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(apiKey, apiToken, ""),
	})
	if err != nil {
		return err
	}

	b.session = s

	return nil
}

func (b *Backend) CheckObject(bucket string, key string) error {
	_, err := s3.New(b.session).GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err == nil {
		return nil
	}

	return err.(awserr.RequestFailure)
}

func (b *Backend) IsNil() bool {
	return b == nil
}
