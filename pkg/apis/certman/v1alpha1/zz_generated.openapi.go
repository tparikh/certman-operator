// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/certman-operator/pkg/apis/certman/v1alpha1.Certificate":       schema_pkg_apis_certman_v1alpha1_Certificate(ref),
		"github.com/certman-operator/pkg/apis/certman/v1alpha1.CertificateSpec":   schema_pkg_apis_certman_v1alpha1_CertificateSpec(ref),
		"github.com/certman-operator/pkg/apis/certman/v1alpha1.CertificateStatus": schema_pkg_apis_certman_v1alpha1_CertificateStatus(ref),
	}
}

func schema_pkg_apis_certman_v1alpha1_Certificate(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Certificate is the Schema for the certificates API",
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
							Ref: ref("github.com/certman-operator/pkg/apis/certman/v1alpha1.CertificateSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/certman-operator/pkg/apis/certman/v1alpha1.CertificateStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/certman-operator/pkg/apis/certman/v1alpha1.CertificateSpec", "github.com/certman-operator/pkg/apis/certman/v1alpha1.CertificateStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_certman_v1alpha1_CertificateSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CertificateSpec defines the desired state of Certificate",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_certman_v1alpha1_CertificateStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CertificateStatus defines the observed state of Certificate",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
