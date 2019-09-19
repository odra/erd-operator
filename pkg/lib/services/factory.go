package services

import (
	"errors"
	"github.com/integr8ly/erd-operator/pkg/lib/services/mapbox"
	"github.com/integr8ly/erd-operator/pkg/lib/services/s3"
)

const (
	MapBox Type = 0
	S3     Type = 1
)

func Build(st Type, data ...interface{}) (Service, error) {
	switch st {
	case MapBox:
		mpToken := data[0].(string)
		return mapbox.New(mpToken), nil
	case S3:
		bucket := data[0].(string)
		apiKey := data[1].(string)
		apiToken := data[2].(string)
		region := data[3].(string)
		return s3.New(bucket, apiKey, apiToken, region), nil
	default:
		return nil, errors.New("invalid service type")
	}
}
