package v1alpha1

import (
	"gotest.tools/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"testing"
)

func TestEmergencyResponseDemo_SecretNamespacedName(t *testing.T) {
	cases := []struct {
		Name     string
		Instance EmergencyResponseDemo
		Expected types.NamespacedName
	}{
		{
			Name: "Should generate filled namespaced name",
			Instance: EmergencyResponseDemo{
				ObjectMeta: v1.ObjectMeta{
					Name:      "demo",
					Namespace: "erd",
				},
				Spec: EmergencyResponseDemoSpec{
					SecretName: "erd-credentials",
				},
			},
			Expected: types.NamespacedName{
				Namespace: "erd",
				Name:      "erd-credentials",
			},
		},
		{
			Name: "Should generate empty namespaced name",
			Instance: EmergencyResponseDemo{
				ObjectMeta: v1.ObjectMeta{
					Name:      "demo",
					Namespace: "erd",
				},
			},
			Expected: types.NamespacedName{
				Namespace: "erd",
				Name:      "",
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Expected, tc.Instance.SecretNamespacedName())
		})
	}
}
