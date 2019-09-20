package fake

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type FakeBackend struct {
	session   *session.Session
	responses map[string]*fakeError
}

type Response struct {
	Bucket string
	Key    string
	Err    *fakeError
}

func (b *FakeBackend) New(region string, apiKey string, apiToken string) error {
	s, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(apiKey, apiToken, ""),
	})
	if err != nil {
		return err
	}

	b.session = s
	b.responses = make(map[string]*fakeError, 0)

	return nil
}

func (b *FakeBackend) AddResponses(responses ...Response) {
	for _, res := range responses {
		key := fmt.Sprintf("%s/%s", res.Bucket, res.Key)
		b.responses[key] = res.Err
	}
}

func (b *FakeBackend) CheckObject(bucket string, key string) error {
	responseKey := fmt.Sprintf("%s/%s", bucket, key)

	err, found := b.responses[responseKey]
	if !found {
		return nil
	}

	fmt.Printf("FOUND: ---%v---", err)


	return err
}

func (b *FakeBackend) IsNil() bool {
	return b == nil
}
