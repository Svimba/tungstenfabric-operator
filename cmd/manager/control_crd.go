package main

import (
	extbetav1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// newCRDForControl returns CustomResourceDefinition object for TF-Control operator
func newCRDForControl() *extbetav1.CustomResourceDefinition {
	return &extbetav1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "tfcontrols.control.tf.mirantis.com",
		},
		Spec: extbetav1.CustomResourceDefinitionSpec{
			Group: "control.tf.mirantis.com",
			Names: extbetav1.CustomResourceDefinitionNames{
				Kind:     "TFControl",
				ListKind: "TFControlList",
				Plural:   "tfcontrols",
				Singular: "tfcontrol",
			},
			Scope:   "Namespaced",
			Version: "v1alpha1",
			Subresources: &extbetav1.CustomResourceSubresources{
				Status: &extbetav1.CustomResourceSubresourceStatus{},
			},
		},
	}
}
