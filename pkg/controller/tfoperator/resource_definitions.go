package tfoperator

import (
	"fmt"

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

func convertPortsToConfigPorts(svc []Service) []configv1alpha1.Port {
	var out []configv1alpha1.Port
	for _, s := range svc {
		for _, p := range s.Ports {
			out = append(out, configv1alpha1.Port{Name: p.Name, Port: p.Port})
		}
	}
	return out
}

// newCRForConfig returns CR object for Config
func newCRForConfig(cr *operatorv1alpha1.TFOperator, defaults *Entities) *configv1alpha1.TFConfig {
	replicas := make(map[string]*int32)
	image := make(map[string]string)
	ports := make(map[string][]configv1alpha1.Port)

	// API
	if &cr.Spec.TFConfig.APISpec != nil {
		if replicas["api"] = &defaults.Get("tf-config-api").Size; cr.Spec.TFConfig.APISpec.Replicas != nil {
			replicas["api"] = cr.Spec.TFConfig.APISpec.Replicas
		}
		if image["api"] = defaults.Get("tf-config-api").Image; len(cr.Spec.TFConfig.APISpec.Image) > 0 {
			image["api"] = cr.Spec.TFConfig.APISpec.Image
		}
		if ports["api"] = convertPortsToConfigPorts(defaults.Get("tf-config-api").Services); len(cr.Spec.TFConfig.APISpec.Ports) > 0 {
			ports["api"] = cr.Spec.TFConfig.APISpec.Ports
		}
	}
	// SVC monitor
	if &cr.Spec.TFConfig.SVCMonitorSpec != nil {
		if replicas["svc-monitor"] = &defaults.Get("tf-config-svc-monitor").Size; cr.Spec.TFConfig.SVCMonitorSpec.Replicas != nil {
			replicas["svc-monitor"] = cr.Spec.TFConfig.SVCMonitorSpec.Replicas
		}
		if image["svc-monitor"] = defaults.Get("tf-config-svc-monitor").Image; len(cr.Spec.TFConfig.SVCMonitorSpec.Image) > 0 {
			image["svc-monitor"] = cr.Spec.TFConfig.SVCMonitorSpec.Image
		}
		if ports["svc-monitor"] = convertPortsToConfigPorts(defaults.Get("tf-config-svc-monitor").Services); len(cr.Spec.TFConfig.SVCMonitorSpec.Ports) > 0 {
			ports["svc-monitor"] = cr.Spec.TFConfig.SVCMonitorSpec.Ports
		}
	}
	// Schema
	if &cr.Spec.TFConfig.SchemaSpec != nil {
		if replicas["schema"] = &defaults.Get("tf-config-schema").Size; cr.Spec.TFConfig.SchemaSpec.Replicas != nil {
			replicas["schema"] = cr.Spec.TFConfig.SchemaSpec.Replicas
		}
		if image["schema"] = defaults.Get("tf-config-schema").Image; len(cr.Spec.TFConfig.SchemaSpec.Image) > 0 {
			image["schema"] = cr.Spec.TFConfig.SchemaSpec.Image
		}
	}
	// Device mgr
	if &cr.Spec.TFConfig.DeviceMgrSpec != nil {
		if replicas["devicemgr"] = &defaults.Get("tf-config-devicemgr").Size; cr.Spec.TFConfig.DeviceMgrSpec.Replicas != nil {
			replicas["devicemgr"] = cr.Spec.TFConfig.DeviceMgrSpec.Replicas
		}
		if image["devicemgr"] = defaults.Get("tf-config-devicemgr").Image; len(cr.Spec.TFConfig.DeviceMgrSpec.Image) > 0 {
			image["devicemgr"] = cr.Spec.TFConfig.DeviceMgrSpec.Image
		}
		if ports["devicemgr"] = convertPortsToConfigPorts(defaults.Get("tf-config-devicemgr").Services); len(cr.Spec.TFConfig.DeviceMgrSpec.Ports) > 0 {
			ports["devicemgr"] = cr.Spec.TFConfig.DeviceMgrSpec.Ports
		}
	}

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
			ConfigMapList: []string{"tf-rabbitmq-cfgmap", "tf-zookeeper-cfgmap"},
			APISpec: configv1alpha1.TFConfigAPISpec{
				Enabled:  true,
				Replicas: replicas["api"],
				Image:    image["api"],
				Ports:    ports["api"],
			},
			SVCMonitorSpec: configv1alpha1.TFConfigSVCMonitorSpec{
				Enabled:  true,
				Replicas: replicas["svc-monitor"],
				Image:    image["svc-monitor"],
				Ports:    ports["svc-monitor"],
			},
			SchemaSpec: configv1alpha1.TFConfigSchemaSpec{
				Enabled:  true,
				Replicas: replicas["schema"],
				Image:    image["schema"],
			},
			DeviceMgrSpec: configv1alpha1.TFConfigDeviceMgrSpec{
				Enabled:  true,
				Replicas: replicas["devicemgr"],
				Image:    image["devicemgr"],
				Ports:    ports["devicemgr"],
			},
		},
	}
}

func getConfigMapForRabbitMQ(cr *operatorv1alpha1.TFOperator, defaults *Entities) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "tf-rabbitmq-cfgmap",
			Namespace: cr.Namespace,
		},
		Data: map[string]string{
			"RABBITMQ_NODES":     defaults.Get("rabbitmq").Services[0].Name,
			"RABBITMQ_NODE_PORT": fmt.Sprintf("%d", defaults.Get("rabbitmq").Services[0].Ports[0].Port),
			"RABBITMQ_SERVERS":   fmt.Sprintf("%s:%d", defaults.Get("rabbitmq").Services[0].Name, defaults.Get("rabbitmq").Services[0].Ports[0].Port),
		},
	}
}

func getConfigMapForZookeeper(cr *operatorv1alpha1.TFOperator, defaults *Entities) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "tf-zookeeper-cfgmap",
			Namespace: cr.Namespace,
		},
		Data: map[string]string{
			"ZOOKEEPER_NODES":               defaults.Get("zookeeper").Services[0].Name,
			"ZOOKEEPER_PORT":                fmt.Sprintf("%d", defaults.Get("zookeeper").Services[0].Ports[0].Port),
			"ZOOKEEPER_PORTS":               "2888:3888",
			"ZOOKEEPER_SERVERS":             fmt.Sprintf("%s:%d", defaults.Get("zookeeper").Services[0].Name, defaults.Get("zookeeper").Services[0].Ports[0].Port),
			"ZOOKEEPER_SERVERS_SPACE_DELIM": fmt.Sprintf("%s:%d", defaults.Get("zookeeper").Services[0].Name, defaults.Get("zookeeper").Services[0].Ports[0].Port),
		},
	}
}

func getConfigMapForCassandraConfig(cr *operatorv1alpha1.TFOperator, defaults *Entities) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "tf-rabbitmq-cfgmap",
			Namespace: cr.Namespace,
		},
		Data: map[string]string{
			"RABBITMQ_NODES":     defaults.Get("rabbitmq").Services[0].Name,
			"RABBITMQ_NODE_PORT": fmt.Sprintf("%d", defaults.Get("rabbitmq").Services[0].Ports[0].Port),
			"RABBITMQ_SERVERS":   fmt.Sprintf("%s:%d", defaults.Get("rabbitmq").Services[0].Name, defaults.Get("rabbitmq").Services[0].Ports[0].Port),
		},
	}
}
