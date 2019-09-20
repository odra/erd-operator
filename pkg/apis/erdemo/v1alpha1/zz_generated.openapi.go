// +build !

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/Emergency-Response-Demo/erd-operator/pkg/apis/erdemo/v1alpha1.EmergencyResponseDemo":       schema_pkg_apis_erdemo_v1alpha1_EmergencyResponseDemo(ref),
		"github.com/Emergency-Response-Demo/erd-operator/pkg/apis/erdemo/v1alpha1.EmergencyResponseDemoSpec":   schema_pkg_apis_erdemo_v1alpha1_EmergencyResponseDemoSpec(ref),
		"github.com/Emergency-Response-Demo/erd-operator/pkg/apis/erdemo/v1alpha1.EmergencyResponseDemoStatus": schema_pkg_apis_erdemo_v1alpha1_EmergencyResponseDemoStatus(ref),
	}
}

func schema_pkg_apis_erdemo_v1alpha1_EmergencyResponseDemo(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EmergencyResponseDemo is the Schema for the emergencyresponsedemos API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Emergency-Response-Demo/erd-operator/pkg/apis/erdemo/v1alpha1.EmergencyResponseDemoSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Emergency-Response-Demo/erd-operator/pkg/apis/erdemo/v1alpha1.EmergencyResponseDemoStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/Emergency-Response-Demo/erd-operator/pkg/apis/erdemo/v1alpha1.EmergencyResponseDemoSpec", "github.com/Emergency-Response-Demo/erd-operator/pkg/apis/erdemo/v1alpha1.EmergencyResponseDemoStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_erdemo_v1alpha1_EmergencyResponseDemoSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EmergencyResponseDemoSpec defines the desired state of EmergencyResponseDemo",
				Properties: map[string]spec.Schema{
					"secretName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"selfSignedCerts": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"subDomain": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"masterUrl": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"selfSignedCerts", "subDomain", "masterUrl"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_erdemo_v1alpha1_EmergencyResponseDemoStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EmergencyResponseDemoStatus defines the observed state of EmergencyResponseDemo",
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"lastHeartbeatTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"lastTransitionTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
				},
				Required: []string{"type", "status"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}
