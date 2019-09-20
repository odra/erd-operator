package helpers

import (
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"testing"
)

func TestGetKey(t *testing.T) {
	cases := []struct {
		Name        string
		Key         string
		Secret      *corev1.Secret
		ExpectValue string
		ExpectError bool
	}{
		{
			Name: "Should retrieve key",
			Key:  "name",
			Secret: &corev1.Secret{
				Data: map[string][]byte{
					"name": []byte("myname"),
				},
				Type: corev1.SecretTypeOpaque,
			},
			ExpectValue: "myname",
			ExpectError: false,
		},
		{
			Name: "Should not retrieve key",
			Key:  "name2",
			Secret: &corev1.Secret{
				Data: map[string][]byte{
					"name": []byte("myname"),
				},
				Type: corev1.SecretTypeOpaque,
			},
			ExpectValue: "",
			ExpectError: true,
		},
		{
			Name:        "Should not retrieve key from empty secret",
			Key:         "name",
			Secret:      &corev1.Secret{},
			ExpectValue: "",
			ExpectError: true,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			value, err := GetKey(tc.Secret, tc.Key)

			if tc.ExpectError && err == nil {
				assert.NotNil(t, err)
			}

			if !tc.ExpectError && err != nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, tc.ExpectValue, value)
		})
	}
}
