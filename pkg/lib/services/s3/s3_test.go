package s3

import (
	"github.com/integr8ly/erd-operator/pkg/lib/services/s3/backend"
	"github.com/integr8ly/erd-operator/pkg/lib/services/s3/backend/fake"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	cases := []struct {
		Name        string
		Data        func() (region string, bucket string, apiKey string, apiToken string)
		ExpectError bool
		Validate    func(t *testing.T, service *S3Service)
	}{
		{
			Name: "Should instantiate service",
			Data: func() (region string, bucket string, apiKey string, apiToken string) {
				region = "eu-west-1"
				bucket = "bkp"
				apiKey = "12345"
				apiToken = "abcde"

				return
			},
			ExpectError: false,
			Validate: func(t *testing.T, service *S3Service) {
				assert.NotNil(t, service)
				assert.Equal(t, service.bucket, "bkp")
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			region, bucket, apiKey, apiToken := tc.Data()

			svc, err := New(bucket, apiKey, apiToken, region)

			if tc.ExpectError && err == nil {
				assert.NotNil(t, err)
			}

			if !tc.ExpectError && err != nil {
				assert.Nil(t, err)
			}

			assert.NotNil(t, svc)

			tc.Validate(t, svc)
		})
	}
}

func TestS3Service_Validate(t *testing.T) {
	cases := []struct {
		Name        string
		Bucket      string
		Svc         func() *backend.Backend
		ExpectError bool
	}{
		{
			Name: "Should pass validation",
			Svc: func() *backend.Backend {
				b := &backend.Backend{}
				_ = b.New("", "", "")

				return b
			},
			Bucket:      "bkp",
			ExpectError: false,
		},
		{
			Name: "Should not pass validation due to empty bucket",
			Svc: func() *backend.Backend {
				b := &backend.Backend{}
				_ = b.New("", "", "")

				return b
			},
			Bucket:      "",
			ExpectError: true,
		},
		{
			Name: "Should not pass validation due to nil service",
			Svc: func() *backend.Backend {
				return nil
			},
			Bucket:      "bkp",
			ExpectError: true,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			svc := S3Service{
				bucket:  tc.Bucket,
				service: tc.Svc(),
			}

			err := svc.Validate()

			if tc.ExpectError && err == nil {
				assert.NotNil(t, err)
			}
			if !tc.ExpectError && err != nil {
				assert.Nil(t, err)
			}
		})
	}
}

func TestS3Service_Assert(t *testing.T) {
	cases := []struct {
		Name        string
		Bucket      string
		Key         string
		Svc         func() *S3Service
		ExpectError bool
	}{
		{
			Name:   "Should fail to assert s3 service (403)",
			Bucket: "bkp",
			Key:    " ",
			Svc: func() *S3Service {
				fakeBackend := fake.FakeBackend{}
				_ = fakeBackend.New("", "", "")

				fakeBackend.AddResponses(fake.Response{
					Bucket: "bkp",
					Key:    " ",
					Err:    fake.New("forbidden", 403, ""),
				})

				return &S3Service{
					bucket:  "bkp",
					service: &fakeBackend,
				}
			},
			ExpectError: true,
		},
		{
			Name:   "Should fail to assert s3 service (401)",
			Bucket: "bkp",
			Key:    " ",
			Svc: func() *S3Service {
				fakeBackend := fake.FakeBackend{}
				_ = fakeBackend.New("", "", "")

				fakeBackend.AddResponses(fake.Response{
					Bucket: "bkp",
					Key:    " ",
					Err:    fake.New("forbidden", 401, ""),
				})

				return &S3Service{
					bucket:  "bkp",
					service: &fakeBackend,
				}
			},
			ExpectError: true,
		},
		{
			Name:   "Should assert s3 service",
			Bucket: "bkp",
			Key:    " ",
			Svc: func() *S3Service {
				fakeBackend := fake.FakeBackend{}
				_ = fakeBackend.New("", "", "")

				return &S3Service{
					bucket:  "bkp",
					service: &fakeBackend,
				}
			},
			ExpectError: false,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			svc := tc.Svc()

			err := svc.Assert()

			if tc.ExpectError && err == nil {
				assert.NotNil(t, err)
			}
			if !tc.ExpectError && err != nil {
				assert.Nil(t, err)
			}
		})
	}
}

func TestS3Service_isValidCode(t *testing.T) {
	cases := []struct{
		Name string
		Code int
		ExpectResult bool
	}{
		{
			Name:         "Should validate code",
			Code:         200,
			ExpectResult: true,
		},
		{
			Name:         "Should validate code (404)",
			Code:         404,
			ExpectResult: true,
		},
		{
			Name:         "Should not validate code (403)",
			Code:         403,
			ExpectResult: false,
		},
		{
			Name:         "Should not validate code (401)",
			Code:         401,
			ExpectResult: false,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			s := &S3Service{}

			assert.Equal(t, tc.ExpectResult, s.isValidCode(tc.Code))
		})
	}
}
