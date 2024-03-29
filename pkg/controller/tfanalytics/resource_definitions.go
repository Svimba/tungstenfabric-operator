package tfanalytics

import (
	analyticsv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/analytics/v1alpha1"
	betav1 "k8s.io/api/apps/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getConfigMapsObject(cr *analyticsv1alpha1.TFAnalytics) []corev1.EnvFromSource {
	var list []corev1.EnvFromSource

	for _, cfm := range cr.Status.ConfigMapList {
		cfobj := &corev1.EnvFromSource{
			ConfigMapRef: &corev1.ConfigMapEnvSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: cfm,
				},
			},
		}

		list = append(list, *cfobj)
	}
	return list
}

// newDeploymentForAPI retuns deployment definition
func newDeploymentForAPI(cr *analyticsv1alpha1.TFAnalytics) *betav1.Deployment {
	labels := map[string]string{
		"app": cr.Name + "-api",
	}

	deploy := &betav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-api",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: betav1.DeploymentSpec{
			Replicas: cr.Spec.APISpec.Replicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  cr.Name + "-api",
							Image: cr.Spec.APISpec.Image,
							// Command: cmd,
						},
					},
				},
			},
		},
	}
	// set co config maps
	configMaps := getConfigMapsObject(cr)
	if len(configMaps) > 0 {
		deploy.Spec.Template.Spec.Containers[0].EnvFrom = configMaps
	}
	// set ports if defined
	ports := cr.Spec.APISpec.Ports.GetContainerPortList()
	if len(ports) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Ports = ports
	} else {
		// If ports are not defined
		deploy.Spec.Template.Spec.Containers[0].Ports = []corev1.ContainerPort{
			{Name: "api", ContainerPort: 8081},
			{Name: "introspect", ContainerPort: 8090},
		}
	}
	// set environment variable(s) if defined in spec
	envs := cr.Spec.APISpec.EnvList
	if len(envs) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Env = envs
	}
	return deploy
}

func newServicesForAPI(cr *analyticsv1alpha1.TFAnalytics) *corev1.Service {
	labels := map[string]string{
		"app": cr.Name + "-api",
	}
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-api",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
		},
	}
	ports := cr.Spec.APISpec.Ports.GetServicePortList()
	if len(ports) > 0 {
		service.Spec.Ports = ports
	} else {
		// If ports are not defined
		service.Spec.Ports = []corev1.ServicePort{
			{Name: "api", Port: 9100},
			{Name: "introspect", Port: 8084},
		}
	}
	return service
}

// newDeploymentForAlarmGen retuns deployment definition
func newDeploymentForAlarmGen(cr *analyticsv1alpha1.TFAnalytics) *betav1.Deployment {
	labels := map[string]string{
		"app": cr.Name + "-alarm-gen",
	}

	deploy := &betav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-alarm-gen",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: betav1.DeploymentSpec{
			Replicas: cr.Spec.AlarmGenSpec.Replicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  cr.Name + "-alarm-gen",
							Image: cr.Spec.AlarmGenSpec.Image,
						},
					},
				},
			},
		},
	}
	// set co config maps
	configMaps := getConfigMapsObject(cr)
	if len(configMaps) > 0 {
		deploy.Spec.Template.Spec.Containers[0].EnvFrom = configMaps
	}
	// set ports if defined
	ports := cr.Spec.AlarmGenSpec.Ports.GetContainerPortList()
	if len(ports) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Ports = ports
	} else {
		// If ports are not defined
		deploy.Spec.Template.Spec.Containers[0].Ports = []corev1.ContainerPort{
			{Name: "introspect", ContainerPort: 5995},
		}
	}

	// set environment variable(s) if defined in spec
	envs := cr.Spec.AlarmGenSpec.EnvList
	if len(envs) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Env = envs
	}

	return deploy
}

func newServicesForAlarmGen(cr *analyticsv1alpha1.TFAnalytics) *corev1.Service {
	labels := map[string]string{
		"app": cr.Name + "-alarm-gen",
	}
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-alarm-gen",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
		},
	}
	ports := cr.Spec.AlarmGenSpec.Ports.GetServicePortList()
	if len(ports) > 0 {
		service.Spec.Ports = ports
	} else {
		// If ports are not defined
		service.Spec.Ports = []corev1.ServicePort{
			{Name: "introspect", Port: 5995},
		}
	}
	return service
}

// newDeploymentForCollector retuns deployment definition
func newDeploymentForCollector(cr *analyticsv1alpha1.TFAnalytics) *betav1.Deployment {
	labels := map[string]string{
		"app": cr.Name + "-collector",
	}

	deploy := &betav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-collector",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: betav1.DeploymentSpec{
			Replicas: cr.Spec.CollectorSpec.Replicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  cr.Name + "-collector",
							Image: cr.Spec.CollectorSpec.Image,
						},
					},
				},
			},
		},
	}
	// set co config maps
	configMaps := getConfigMapsObject(cr)
	if len(configMaps) > 0 {
		deploy.Spec.Template.Spec.Containers[0].EnvFrom = configMaps
	}
	// set ports if defined
	ports := cr.Spec.CollectorSpec.Ports.GetContainerPortList()
	if len(ports) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Ports = ports
	} else {
		// If ports are not defined
		deploy.Spec.Template.Spec.Containers[0].Ports = []corev1.ContainerPort{
			{Name: "collector", ContainerPort: 8086},
			{Name: "introspect", ContainerPort: 8089},
		}
	}
	// set environment variable(s) if defined in spec
	envs := cr.Spec.CollectorSpec.EnvList
	if len(envs) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Env = envs
	}

	return deploy
}

func newServicesForCollector(cr *analyticsv1alpha1.TFAnalytics) *corev1.Service {
	labels := map[string]string{
		"app": cr.Name + "-collector",
	}
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-collector",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
		},
	}
	ports := cr.Spec.CollectorSpec.Ports.GetServicePortList()
	if len(ports) > 0 {
		service.Spec.Ports = ports
	} else {
		// If ports are not defined
		service.Spec.Ports = []corev1.ServicePort{
			{Name: "collector", Port: 8086},
			{Name: "introspect", Port: 8089},
		}
	}
	return service
}

func newDeploymentForQueryEngine(cr *analyticsv1alpha1.TFAnalytics) *betav1.Deployment {
	labels := map[string]string{
		"app": cr.Name + "-snmp",
	}

	deploy := &betav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-snmp",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: betav1.DeploymentSpec{
			Replicas: cr.Spec.QueryEngine.Replicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  cr.Name + "-snmp",
							Image: cr.Spec.QueryEngine.Image,
						},
					},
				},
			},
		},
	}
	// set co config maps
	configMaps := getConfigMapsObject(cr)
	if len(configMaps) > 0 {
		deploy.Spec.Template.Spec.Containers[0].EnvFrom = configMaps
	}
	// set ports if defined
	ports := cr.Spec.QueryEngine.Ports.GetContainerPortList()
	if len(ports) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Ports = ports
	} else {
		// If ports are not defined
		deploy.Spec.Template.Spec.Containers[0].Ports = []corev1.ContainerPort{
			{Name: "introspect", ContainerPort: 8091},
		}
	}
	// set environment variable(s) if defined in spec
	envs := cr.Spec.QueryEngine.EnvList
	if len(envs) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Env = envs
	}

	return deploy
}

func newServicesForQueryEngine(cr *analyticsv1alpha1.TFAnalytics) *corev1.Service {
	labels := map[string]string{
		"app": cr.Name + "-devicemgr",
	}
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-devicemgr",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
		},
	}
	ports := cr.Spec.QueryEngine.Ports.GetServicePortList()
	if len(ports) > 0 {
		service.Spec.Ports = ports
	} else {
		// If ports are not defined
		service.Spec.Ports = []corev1.ServicePort{
			{Name: "introspect", Port: 8091},
		}
	}
	return service
}

func newDeploymentForSNMP(cr *analyticsv1alpha1.TFAnalytics) *betav1.Deployment {
	labels := map[string]string{
		"app": cr.Name + "-snmp",
	}

	deploy := &betav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-snmp",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: betav1.DeploymentSpec{
			Replicas: cr.Spec.SNMPSpec.Replicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  cr.Name + "-snmp",
							Image: cr.Spec.SNMPSpec.Image,
						},
					},
				},
			},
		},
	}
	// set co config maps
	configMaps := getConfigMapsObject(cr)
	if len(configMaps) > 0 {
		deploy.Spec.Template.Spec.Containers[0].EnvFrom = configMaps
	}
	// set ports if defined
	ports := cr.Spec.SNMPSpec.Ports.GetContainerPortList()
	if len(ports) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Ports = ports
	} else {
		// If ports are not defined
		deploy.Spec.Template.Spec.Containers[0].Ports = []corev1.ContainerPort{
			{Name: "introspect", ContainerPort: 5920},
		}
	}
	// set environment variable(s) if defined in spec
	envs := cr.Spec.SNMPSpec.EnvList
	if len(envs) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Env = envs
	}

	return deploy
}

func newServicesForSNMP(cr *analyticsv1alpha1.TFAnalytics) *corev1.Service {
	labels := map[string]string{
		"app": cr.Name + "-snmp",
	}
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-snmp",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
		},
	}
	ports := cr.Spec.SNMPSpec.Ports.GetServicePortList()
	if len(ports) > 0 {
		service.Spec.Ports = ports
	} else {
		// If ports are not defined
		service.Spec.Ports = []corev1.ServicePort{
			{Name: "introspect", Port: 5920},
		}
	}
	return service
}

func newDeploymentForTopology(cr *analyticsv1alpha1.TFAnalytics) *betav1.Deployment {
	labels := map[string]string{
		"app": cr.Name + "-topology",
	}

	deploy := &betav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-topology",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: betav1.DeploymentSpec{
			Replicas: cr.Spec.TopologySpec.Replicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  cr.Name + "-topology",
							Image: cr.Spec.TopologySpec.Image,
						},
					},
				},
			},
		},
	}
	// set co config maps
	configMaps := getConfigMapsObject(cr)
	if len(configMaps) > 0 {
		deploy.Spec.Template.Spec.Containers[0].EnvFrom = configMaps
	}
	// set ports if defined
	ports := cr.Spec.TopologySpec.Ports.GetContainerPortList()
	if len(ports) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Ports = ports
	} else {
		// If ports are not defined
		deploy.Spec.Template.Spec.Containers[0].Ports = []corev1.ContainerPort{
			{Name: "introspect", ContainerPort: 5921},
		}
	}
	// set environment variable(s) if defined in spec
	envs := cr.Spec.TopologySpec.EnvList
	if len(envs) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Env = envs
	}

	return deploy
}

func newServicesForTopology(cr *analyticsv1alpha1.TFAnalytics) *corev1.Service {
	labels := map[string]string{
		"app": cr.Name + "-topology",
	}
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-topology",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
		},
	}
	ports := cr.Spec.TopologySpec.Ports.GetServicePortList()
	if len(ports) > 0 {
		service.Spec.Ports = ports
	} else {
		// If ports are not defined
		service.Spec.Ports = []corev1.ServicePort{
			{Name: "introspect", Port: 5921},
		}
	}
	return service
}
