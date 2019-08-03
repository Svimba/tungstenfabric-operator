package tfcontrol

import (
	controlv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/control/v1alpha1"
	betav1 "k8s.io/api/apps/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getConfigMapsObject(cr *controlv1alpha1.TFControl) []corev1.EnvFromSource {
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

func getContainerPortObject(portList []controlv1alpha1.Port) []corev1.ContainerPort {
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

func getEnvVariablesObject(envs []controlv1alpha1.EnvVar) []corev1.EnvVar {
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

func getServicePortObject(portList []controlv1alpha1.Port) []corev1.ServicePort {
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

// newDeploymentForControl retuns deployment definition
func newDeploymentForControl(cr *controlv1alpha1.TFControl) *betav1.Deployment {
	labels := map[string]string{
		"app": cr.Name + "-control",
	}
	// var cmd []string
	// cmd = append(cmd, "env")

	deploy := &betav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-control",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: betav1.DeploymentSpec{
			Replicas: cr.Spec.ControlSpec.Replicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  cr.Name + "-control",
							Image: cr.Spec.ControlSpec.Image,
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
	ports := getContainerPortObject(cr.Spec.ControlSpec.Ports)
	if len(ports) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Ports = ports
	} else {
		// If ports are not defined
		deploy.Spec.Template.Spec.Containers[0].Ports = []corev1.ContainerPort{
			{Name: "introspect", ContainerPort: 8083},
		}
	}
	// set environment variable(s) if defined in spec
	envs := getEnvVariablesObject(cr.Spec.ControlSpec.EnvList)
	if len(envs) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Env = envs
	}
	return deploy
}

// newDeploymentForNamed retuns deployment definition
func newDeploymentForNamed(cr *controlv1alpha1.TFControl) *betav1.Deployment {
	labels := map[string]string{
		"app": cr.Name + "-named",
	}

	deploy := &betav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-named",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: betav1.DeploymentSpec{
			Replicas: cr.Spec.NamedSpec.Replicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  cr.Name + "-named",
							Image: cr.Spec.NamedSpec.Image,
							SecurityContext: &corev1.SecurityContext{
								Capabilities: &corev1.Capabilities{
									Add: cr.Spec.NamedSpec.SecurityContext.Capabilities,
								},
							},
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
	ports := getContainerPortObject(cr.Spec.NamedSpec.Ports)
	if len(ports) > 0 {
		deploy.Spec.Template.Spec.Containers[0].Ports = ports
	} else {
		// If ports are not defined
		deploy.Spec.Template.Spec.Containers[0].Ports = []corev1.ContainerPort{
			{Name: "dns", ContainerPort: 53},
		}
	}

	return deploy
}

// newDeploymentForDns retuns deployment definition
func newDeploymentForDns(cr *controlv1alpha1.TFControl) *betav1.Deployment {
	labels := map[string]string{
		"app": cr.Name + "-dns",
	}

	deploy := &betav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-dns",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: betav1.DeploymentSpec{
			Replicas: cr.Spec.DnsSpec.Replicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  cr.Name + "-dns",
							Image: cr.Spec.DnsSpec.Image,
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
