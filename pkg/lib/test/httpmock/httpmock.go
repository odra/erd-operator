package httpmock

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func New(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func SimpleMock(statusCode int, body string) *http.Client {
	client := New(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: statusCode,
			Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
			Header:     make(http.Header),
		}
	})

	return client
}
