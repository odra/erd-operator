package services

import (
	"github.com/integr8ly/erd-operator/pkg/lib/services/mapbox"
	"github.com/integr8ly/erd-operator/pkg/lib/services/s3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuild(t *testing.T) {
	cases := []struct{
		Name string
		ExpectError bool
		Build func() (Service, error)
		Validate func(t *testing.T, service Service)
	}{
		{
			Name:        "Should retrieve mapbox service",
			ExpectError: false,
			Build: func() (service Service, e error) {
				return Build(0, "token")
			},
			Validate: func(t *testing.T, service Service) {
				assert.IsType(t, &mapbox.MapBox{}, service)
			},
		},
		{
			Name:        "Should fail to retrieve mapbox service",
			ExpectError: true,
			Build: func() (service Service, e error) {
				return Build(0)
			},
			Validate: func(t *testing.T, service Service) {
				assert.Nil(t, service)
			},
		},
		{
			Name:        "Should retrieve s3 service",
			ExpectError: false,
			Build: func() (service Service, e error) {
				return Build(1, "bucket", "apikey", "apitoken", "region")
			},
			Validate: func(t *testing.T, service Service) {
				assert.IsType(t, &s3.S3Service{}, service)
			},
		},
		{
			Name:        "Should fail retrieve s3 service",
			ExpectError: true,
			Build: func() (service Service, e error) {
				return Build(1, "bucket", "apikey", "apitoken")
			},
			Validate: func(t *testing.T, service Service) {
				assert.Nil(t, service)
			},
		},
		{
			Name:        "Should fail retrieve any service at all",
			ExpectError: true,
			Build: func() (service Service, e error) {
				return Build(50)
			},
			Validate: func(t *testing.T, service Service) {
				assert.Nil(t, service)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			svc, err := tc.Build()

			if tc.ExpectError && err == nil {
				assert.NotNil(t, err)
			}

			if !tc.ExpectError && err != nil {
				assert.Nil(t, err)
			}

			tc.Validate(t, svc)
		})
	}
}
