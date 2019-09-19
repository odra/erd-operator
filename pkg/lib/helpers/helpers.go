package helpers

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
)

func GetKey(secret *corev1.Secret, key string) (string, error) {
	value, found := secret.Data[key]
	if !found {
		return "", fmt.Errorf("key %s not found in secret %s", key, secret.Name)
	}

	return string(value[:]), nil
}
