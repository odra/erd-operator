package s3

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Service struct {
	bucket string
	apiKey string
	apiToken string
	region string
}

func New(bucket string, apiKey string, apiToken string, region string) *s3Service {
	return &s3Service{
		bucket:   bucket,
		apiKey:   apiKey,
		apiToken: apiToken,
		region: region,
	}
}

func (s *s3Service) bootstrap() (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Region: aws.String(s.region),
		Credentials: credentials.NewStaticCredentials(s.apiKey, s.apiToken, ""),
	})
}

func (s *s3Service) Assert() error {
	sess, err := s.bootstrap()
	if err != nil {
		return err
	}

	_, err = s3.New(sess).GetObject(&s3.GetObjectInput{
		Bucket:                     aws.String(s.bucket),
		Key:                        aws.String("  "),
	})

	if err == nil {
		return nil
	}

	parsedErr := err.(awserr.RequestFailure)
	if !s.isValidCode(parsedErr.StatusCode()) {
		return errors.New(parsedErr.Message())
	}

	return nil
}

func (s *s3Service) Validate() error {
	if s.bucket == "" {
		return errors.New("bucket is empty")
	}

	if s.apiKey == "" {
		return errors.New("apiKey is empty")
	}

	if s.apiToken == "" {
		return errors.New("apiToken is empty")
	}

	return nil
}

func (s *s3Service) isValidCode(code int) bool {
	errorCodes := []int{401, 403}

	for _, errorCode := range errorCodes {
		if code == errorCode {
			return false
		}
	}

	return true
}