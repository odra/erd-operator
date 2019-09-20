package s3

type client interface {
	New(region string, apiKey string, apiToken string) error
	CheckObject(bucket string, key string) error
	IsNil() bool
}
