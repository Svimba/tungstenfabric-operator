package tfconfig

import (
	configv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/config/v1alpha1"
	betav1 "k8s.io/api/apps/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getConfigMapsObject(cr *configv1alpha1.TFConfig) []corev1.EnvFromSource {
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

func getContainerPortObject(portList []configv1alpha1.Port) []corev1.ContainerPort {
	var list []corev1.ContainerPort
	for _, p := range portList {
		pobj := &corev1.ContainerPort{
			Name:          p.Name,
			ContainerPort: p.Port,
		}
		list = append(list, *pobj)
	}
	return list
}

func getEnvVariablesObject(envs []configv1alpha1.EnvVar) []corev1.EnvVar {
	var list []corev1.EnvVar
	for _, e := range envs {
		eobj := &corev1.EnvVar{
			Name:  e.Name,
			Value: e.Value,
		}
		list = append(list, *eobj)
	}
	return list
}

func getServicePortObject(portList []configv1alpha1.Port) []corev1.ServicePort {
	var list []corev1.ServicePort
	for _, p := range portList {
		pobj := &corev1.ServicePort{
			Name: p.Name,
			Port: p.Port,
		}
		list = append(list, *pobj)
	}
	return list
}

// newDeploymentForAPI retuns deployment definition
func newDeploymentForAPI(cr *configv1alpha1.TFConfig) *betav1.Deployment {
	labels := map[string]string{
		"app": cr.Name + "-api",
	}
	// var cmd []string
	// cmd = append(cmd, "env")

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
	ports := getContainerPortObject(cr.Spec.APISpec.Ports)
	if len(ports) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Ports = ports
	} else {
		// If ports are not defined
		deploy.Spec.Template.Spec.Containers[0].Ports = []corev1.ContainerPort{
			{Name: "api", ContainerPort: 9100},
			{Name: "introspect", ContainerPort: 8084},
		}
	}
	// set environment variable(s) if defined in spec
	envs := getEnvVariablesObject(cr.Spec.APISpec.EnvList)
	if len(envs) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Env = envs
	}
	return deploy
}

func newServicesForAPI(cr *configv1alpha1.TFConfig) *corev1.Service {
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
	ports := getServicePortObject(cr.Spec.APISpec.Ports)
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

// newDeploymentForSVCMonitor retuns deployment definition
func newDeploymentForSVCMonitor(cr *configv1alpha1.TFConfig) *betav1.Deployment {
	labels := map[string]string{
		"app": cr.Name + "-svc-monitor",
	}

	deploy := &betav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-svc-monitor",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: betav1.DeploymentSpec{
			Replicas: cr.Spec.SVCMonitorSpec.Replicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  cr.Name + "-svc-monitor",
							Image: cr.Spec.SVCMonitorSpec.Image,
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
	ports := getContainerPortObject(cr.Spec.SVCMonitorSpec.Ports)
	if len(ports) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Ports = ports
	} else {
		// If ports are not defined
		deploy.Spec.Template.Spec.Containers[0].Ports = []corev1.ContainerPort{
			{Name: "introspect", ContainerPort: 8088},
		}
	}

	return deploy
}

func newServicesForSVCMonitor(cr *configv1alpha1.TFConfig) *corev1.Service {
	labels := map[string]string{
		"app": cr.Name + "-svc-monitor",
	}
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-svc-monitor",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
		},
	}
	ports := getServicePortObject(cr.Spec.SVCMonitorSpec.Ports)
	if len(ports) > 0 {
		service.Spec.Ports = ports
	} else {
		// If ports are not defined
		service.Spec.Ports = []corev1.ServicePort{
			{Name: "introspect", Port: 8088},
		}
	}
	return service
}

// newDeploymentForSVCMonitor retuns deployment definition
func newDeploymentForSchema(cr *configv1alpha1.TFConfig) *betav1.Deployment {
	labels := map[string]string{
		"app": cr.Name + "-schema",
	}

	deploy := &betav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-schema",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: betav1.DeploymentSpec{
			Replicas: cr.Spec.SchemaSpec.Replicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  cr.Name + "-schema",
							Image: cr.Spec.SchemaSpec.Image,
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

	return deploy
}

// newDeploymentForDeviceMgr retuns deployment definition
func newDeploymentForDeviceMgr(cr *configv1alpha1.TFConfig) *betav1.Deployment {
	labels := map[string]string{
		"app": cr.Name + "-devicemgr",
	}

	deploy := &betav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-devicemgr",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: betav1.DeploymentSpec{
			Replicas: cr.Spec.DeviceMgrSpec.Replicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  cr.Name + "-devicemgr",
							Image: cr.Spec.DeviceMgrSpec.Image,
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
	ports := getContainerPortObject(cr.Spec.DeviceMgrSpec.Ports)
	if len(ports) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Ports = ports
	} else {
		// If ports are not defined
		deploy.Spec.Template.Spec.Containers[0].Ports = []corev1.ContainerPort{
			{Name: "introspect", ContainerPort: 8087},
		}
	}
	return deploy
}

func newServicesForDeviceMgr(cr *configv1alpha1.TFConfig) *corev1.Service {
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
	ports := getServicePortObject(cr.Spec.DeviceMgrSpec.Ports)
	if len(ports) > 0 {
		service.Spec.Ports = ports
	} else {
		// If ports are not defined
		service.Spec.Ports = []corev1.ServicePort{
			{Name: "introspect", Port: 8087},
		}
	}
	return service
}
