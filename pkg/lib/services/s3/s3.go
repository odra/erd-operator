package s3

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/integr8ly/erd-operator/pkg/lib/services/s3/backend"
)

type S3Service struct {
	bucket  string
	service client
}

func New(bucket string, apiKey string, apiToken string, region string) (*S3Service, error) {
	svc := &backend.Backend{}

	err := svc.New(region, apiKey, apiToken)
	if err != nil {
		return nil, err
	}

	return &S3Service{
		bucket:  bucket,
		service: svc,
	}, nil
}

func (s *S3Service) Assert() error {
	err := s.service.CheckObject(s.bucket, " ")

	if err == nil {
		return nil
	}

	parsedErr := err.(awserr.RequestFailure)
	if !s.isValidCode(parsedErr.StatusCode()) {
		return errors.New(parsedErr.Message())
	}

	return nil
}

func (s *S3Service) Validate() error {
	if s.bucket == "" {
		return errors.New("bucket is empty")
	}

	if s.service.IsNil() {
		return errors.New("service backend is nil")
	}

	return nil
}

func (s *S3Service) isValidCode(code int) bool {
	errorCodes := []int{401, 403}

	for _, errorCode := range errorCodes {
		if code == errorCode {
			return false
		}
	}

	return true
}
