package tfoperator

import (
	configv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/config/v1alpha1"
	operatorv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/operator/v1alpha1"
	betav1 "k8s.io/api/apps/v1beta1"
	corev1 "k8s.io/api/core/v1"
	extbetav1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// newCRDForConfig returns CustomResourceDefinition object for TF-Config operator
func newCRDForConfig(cr *operatorv1alpha1.TFOperator) *extbetav1.CustomResourceDefinition {
	return &extbetav1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "tfconfigs.config.tf.mirantis.com",
		},
		Spec: extbetav1.CustomResourceDefinitionSpec{
			Group: "config.tf.mirantis.com",
			Names: extbetav1.CustomResourceDefinitionNames{
				Kind:     "TFConfig",
				ListKind: "TFConfigList",
				Plural:   "tfconfigs",
				Singular: "tfconfig",
			},
			Scope:   "Namespaced",
			Version: "v1alpha1",
			Subresources: &extbetav1.CustomResourceSubresources{
				Status: &extbetav1.CustomResourceSubresourceStatus{},
			},
		},
	}
}

// newOperatorForConfig returns Deployment of TF-Config operator
func newOperatorForConfig(cr *operatorv1alpha1.TFOperator) *betav1.Deployment {
	labels := map[string]string{
		"name": "tungstenfabric-config-operator",
	}
	var replicas int32
	replicas = 1
	return &betav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "tungstenfabric-config-operator",
			Namespace: cr.Namespace,
		},
		Spec: betav1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: "tungstenfabric-operator",
					Containers: []corev1.Container{
						{
							Name:  "tungstenfabric-config-operator",
							Image: "willco/tf-config-operator",
							Ports: []corev1.ContainerPort{
								{
									Name:          "metrics",
									ContainerPort: 60000,
								},
							},
							Command: []string{"tungstenfabric-config-operator"},
							Env: []corev1.EnvVar{
								{
									Name: "WATCH_NAMESPACE",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.namespace",
										},
									},
								},
								{
									Name: "POD_NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.name",
										},
									},
								},
								{
									Name:  "OPERATOR_NAME",
									Value: "tungstenfabric-config-operator",
								},
							},
						},
					},
				},
			},
		},
	}
}

// newCRForConfig returns CR object for Config
func newCRForConfig(cr *operatorv1alpha1.TFOperator) *configv1alpha1.TFConfig {
	var replicas int32
	replicas = 3
	return &configv1alpha1.TFConfig{
		TypeMeta: metav1.TypeMeta{
			Kind:       "TFConfig",
			APIVersion: "config.tf.mirantis.com/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "tf-config",
			Namespace: cr.Namespace,
		},
		Spec: configv1alpha1.TFConfigSpec{
			APISpec: configv1alpha1.TFConfigAPISpec{
				Enabled:  true,
				Replicas: &replicas,
				Image:    "willco/opencontrail-config-api:r5.1",
				Ports: []configv1alpha1.Port{
					{
						Name: "api",
						Port: 9100,
					},
					{
						Name: "introspect",
						Port: 8084,
					},
				},
			},
			SVCMonitorSpec: configv1alpha1.TFConfigSVCMonitorSpec{
				Enabled:  true,
				Replicas: &replicas,
				Image:    "willco/opencontrail-config-svc-monitor:r5.1",
				Ports: []configv1alpha1.Port{
					{
						Name: "introspect",
						Port: 8088,
					},
				},
			},
			SchemaSpec: configv1alpha1.TFConfigSchemaSpec{
				Enabled:  true,
				Replicas: &replicas,
				Image:    "willco/opencontrail-config-schema:r5.1",
			},
			DeviceMgrSpec: configv1alpha1.TFConfigDeviceMgrSpec{
				Enabled:  true,
				Replicas: &replicas,
				Image:    "willco/opencontrail-config-devicemgr:r5.1",
				Ports: []configv1alpha1.Port{
					{
						Name: "introspect",
						Port: 8087,
					},
				},
			},
		},
	}
}
