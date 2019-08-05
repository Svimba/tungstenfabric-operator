package main

import (
	extbetav1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// newCRDForConfig returns CustomResourceDefinition object for TF-Config operator
func newCRDForAnalytics() *extbetav1.CustomResourceDefinition {
	return &extbetav1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "tfanalytics.analytics.tf.mirantis.com",
		},
		Spec: extbetav1.CustomResourceDefinitionSpec{
			Group: "analytics.tf.mirantis.com",
			Names: extbetav1.CustomResourceDefinitionNames{
				Kind:     "TFAnalytics",
				ListKind: "TFAnalyticsList",
				Plural:   "tfanalytics",
				Singular: "tfanalytics",
			},
			Scope:   "Namespaced",
			Version: "v1alpha1",
			Subresources: &extbetav1.CustomResourceSubresources{
				Status: &extbetav1.CustomResourceSubresourceStatus{},
			},
		},
	}
}
