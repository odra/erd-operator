package services

import (
	"errors"
	"fmt"
	"github.com/Emergency-Response-Demo/erd-operator/pkg/lib/services/mapbox"
	"github.com/Emergency-Response-Demo/erd-operator/pkg/lib/services/s3"
)

const (
	MapBox Type = 0
	S3     Type = 1
)

func Build(st Type, data ...interface{}) (Service, error) {
	argSize := len(data)
	switch st {
	case MapBox:
		if argSize < 1 {
			return nil, errors.New("expected 1 argument but got 0")
		}
		mpToken := data[0].(string)
		return mapbox.New(mpToken), nil
	case S3:
		if argSize < 4 {
			return nil, fmt.Errorf("expected 4 arguments but got %d", argSize)
		}
		bucket := data[0].(string)
		apiKey := data[1].(string)
		apiToken := data[2].(string)
		region := data[3].(string)
		return s3.New(bucket, apiKey, apiToken, region)
	default:
		return nil, errors.New("invalid service type")
	}
}
