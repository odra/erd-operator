package service

import (
	"fmt"
	"github.com/Emergency-Response-Demo/erd-operator/pkg/lib/helpers"
	"github.com/Emergency-Response-Demo/erd-operator/pkg/lib/services"
	corev1 "k8s.io/api/core/v1"
)

func (h *Helper) CheckServices(secret *corev1.Secret) error {
	var svc services.Service
	var err error

	//mapbox service
	svc, err = h.getMapBoxService(secret)
	if err != nil {
		return err
	}
	err = h.validateService(svc)
	if err != nil {
		return fmt.Errorf("[MAPBOX] %s", err.Error())
	}

	//aws_s3 service
	svc, err = h.getS3Service(secret)
	if err != nil {
		return err
	}
	err = h.validateService(svc)
	if err != nil {
		return fmt.Errorf("[AWS_S3] %s", err.Error())
	}

	return nil
}

func (h *Helper) getMapBoxService(secret *corev1.Secret) (services.Service, error) {
	value, err := helpers.GetKey(secret, "mapbox_api_key")
	if err != nil {
		return nil, fmt.Errorf("[MAPBOX] %s", err.Error())
	}

	svc, err := services.Build(services.MapBox, value)
	if err != nil {
		return nil, fmt.Errorf("[MAPBOX] %s", err.Error())
	}

	return svc, nil
}

func (h *Helper) getS3Service(secret *corev1.Secret) (services.Service, error) {
	bucket, err := helpers.GetKey(secret, "s3_bucket_name")
	if err != nil {
		return nil, fmt.Errorf("[AWS_S3] %s", err.Error())
	}

	apiKey, err := helpers.GetKey(secret, "s3_api_key")
	if err != nil {
		return nil, fmt.Errorf("[AWS_S3] %s", err.Error())
	}

	apiToken, err := helpers.GetKey(secret, "s3_api_token")
	if err != nil {
		return nil, fmt.Errorf("[AWS_S3] %s", err.Error())
	}

	region, err := helpers.GetKey(secret, "s3_region")
	if err != nil {
		return nil, fmt.Errorf("[AWS_S3] %s", err.Error())
	}

	svc, err := services.Build(services.S3, bucket, apiKey, apiToken, region)
	if err != nil {
		return nil, fmt.Errorf("[AWS_S3] %s", err.Error())
	}

	return svc, nil
}

func (h *Helper) validateService(svc services.Service) error {
	var err error

	err = svc.Validate()
	if err != nil {
		return err
	}

	err = svc.Assert()
	if err != nil {
		return err
	}

	return nil
}
