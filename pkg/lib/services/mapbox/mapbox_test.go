package mapbox

import (
	"github.com/integr8ly/erd-operator/pkg/lib/test/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	cases := []struct{
		Name string
		Token string
		Validate func(t *testing.T, mb *mapBox)
	}{
		{
			Name:      "Should match token",
			Token:     "12345",
			Validate: func(t *testing.T, mb *mapBox) {
				assert.IsType(t, &mapBox{}, mb)
				assert.Equal(t, mb.token, "12345")
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			mb := New(tc.Token)
			tc.Validate(t, mb)
		})
	}
}

func TestMapBox_Validate(t *testing.T) {
	cases := []struct{
		Name string
		Token string
		Validate func(t *testing.T, mb *mapBox)
	}{
		{
			Name:      "Should validate mapbox",
			Token:     "12345",
			Validate: func(t *testing.T, mb *mapBox) {
				assert.Nil(t, mb.Validate())
			},
		},
		{
			Name:      "Should invalidate mapbox",
			Token:     "",
			Validate: func(t *testing.T, mb *mapBox) {
				err := mb.Validate()
				assert.NotNil(t, err)
				assert.Error(t, err, "MapBox token is empty")
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			mb := &mapBox{token:tc.Token}
			tc.Validate(t, mb)
		})
	}
}

func TestMapBox_Assert(t *testing.T) {
	cases := []struct{
		Name string
		Token string
		Client func() *http.Client
		Validate func(t *testing.T, mb *mapBox)
	}{
		{
			Name:     "Should assert mapbox service",
			Token:    "12345",
			Client: func() *http.Client {
				return httpmock.SimpleMock(404, `NOT FOUND`)
			},
			Validate: func(t *testing.T, mb *mapBox) {
				err := mb.Assert()
				assert.Nil(t, err)
			},
		},
		{
			Name:     "Should not assert mapbox service",
			Token:    "12345",
			Client: func() *http.Client {
				return httpmock.SimpleMock(401, `UNAUTHORIZED`)
			},
			Validate: func(t *testing.T, mb *mapBox) {
				err := mb.Assert()
				assert.Error(t, err, "token not authorized to perform request")
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			client := tc.Client()
			mb := &mapBox{token:tc.Token, httpClient: client}
			tc.Validate(t, mb)
		})
	}
}