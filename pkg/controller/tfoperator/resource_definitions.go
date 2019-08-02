package tfoperator

import (
	"fmt"
	"reflect"

	configv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/config/v1alpha1"
	controlv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/control/v1alpha1"
	operatorv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/operator/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func convertPortsToConfigPorts(svc []Service) []configv1alpha1.Port {
	var out []configv1alpha1.Port
	for _, s := range svc {
		for _, p := range s.Ports {
			out = append(out, configv1alpha1.Port{Name: p.Name, Port: p.Port})
		}
	}
	return out
}

func convertEnvsToConfigEnvs(envs []Env) []configv1alpha1.EnvVar {
	var out []configv1alpha1.EnvVar
	for _, e := range envs {
		out = append(out, configv1alpha1.EnvVar{Name: e.Key, Value: e.Value})
	}
	return out
}

// newCRForConfig returns CR object for Config
func newCRForConfig(cr *operatorv1alpha1.TFOperator, defaults *Entities) *configv1alpha1.TFConfig {
	replicas := make(map[string]*int32)
	image := make(map[string]string)
	ports := make(map[string][]configv1alpha1.Port)
	envs := make(map[string][]configv1alpha1.EnvVar)

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
		if envs["api"] = convertEnvsToConfigEnvs(defaults.Get("tf-config-api").Envs); len(cr.Spec.TFConfig.APISpec.EnvList) > 0 {
			envs["api"] = cr.Spec.TFConfig.APISpec.EnvList
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
		if envs["svc-monitor"] = convertEnvsToConfigEnvs(defaults.Get("tf-config-svc-monitor").Envs); len(cr.Spec.TFConfig.SVCMonitorSpec.EnvList) > 0 {
			envs["svc-monitor"] = cr.Spec.TFConfig.SVCMonitorSpec.EnvList
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
		if envs["schema"] = convertEnvsToConfigEnvs(defaults.Get("tf-config-schema").Envs); len(cr.Spec.TFConfig.SchemaSpec.EnvList) > 0 {
			envs["schema"] = cr.Spec.TFConfig.SchemaSpec.EnvList
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
		if envs["devicemgr"] = convertEnvsToConfigEnvs(defaults.Get("tf-config-devicemgr").Envs); len(cr.Spec.TFConfig.DeviceMgrSpec.EnvList) > 0 {
			envs["devicemgr"] = cr.Spec.TFConfig.DeviceMgrSpec.EnvList
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
			ConfigMapList: []string{"tf-rabbitmq-cfgmap", "tf-zookeeper-cfgmap", "tf-cassandra-cfgmap"},
			APISpec: configv1alpha1.TFConfigAPISpec{
				Enabled:  true,
				Replicas: replicas["api"],
				Image:    image["api"],
				Ports:    ports["api"],
				EnvList:  envs["api"],
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

func convertPortsToControlPorts(svc []Service) []controlv1alpha1.Port {
	var out []controlv1alpha1.Port
	for _, s := range svc {
		for _, p := range s.Ports {
			out = append(out, controlv1alpha1.Port{Name: p.Name, Port: p.Port})
		}
	}
	return out
}

func convertEnvsToControlEnvs(envs []Env) []controlv1alpha1.EnvVar {
	var out []controlv1alpha1.EnvVar
	for _, e := range envs {
		out = append(out, controlv1alpha1.EnvVar{Name: e.Key, Value: e.Value})
	}
	return out
}

func convertSCToControlSC(sc SecCtx) controlv1alpha1.SecCtx {
	var out controlv1alpha1.SecCtx
	var capabilities []corev1.Capability
        for _, c := range sc.Capabilities {
		capabilities = append(capabilities, corev1.Capability(c))
	}
	out.Capabilities = capabilities
        return out 
}

// newCRForControl returns CR object for Control
func newCRForControl(cr *operatorv1alpha1.TFOperator, defaults *Entities) *controlv1alpha1.TFControl {
	replicas := make(map[string]*int32)
	image := make(map[string]string)
	secContexts := make(map[string]controlv1alpha1.SecCtx)

	// Control
	if &cr.Spec.TFControl.ControlSpec != nil {
		if replicas["control"] = &defaults.Get("tf-control-control").Size; cr.Spec.TFControl.ControlSpec.Replicas != nil {
			replicas["contrrol"] = cr.Spec.TFControl.ControlSpec.Replicas
		}
		if image["control"] = defaults.Get("tf-control-control").Image; len(cr.Spec.TFControl.ControlSpec.Image) > 0 {
			image["control"] = cr.Spec.TFControl.ControlSpec.Image
		}
	}
	// Named
	if &cr.Spec.TFControl.NamedSpec != nil {
		if replicas["named"] = &defaults.Get("tf-control-named").Size; cr.Spec.TFControl.NamedSpec.Replicas != nil {
			replicas["named"] = cr.Spec.TFControl.NamedSpec.Replicas
		}
		if image["named"] = defaults.Get("tf-control-named").Image; len(cr.Spec.TFControl.NamedSpec.Image) > 0 {
			image["named"] = cr.Spec.TFControl.NamedSpec.Image
		}
		if secContexts["named"] = convertSCToControlSC(defaults.Get("tf-control-named").SecContext); !reflect.DeepEqual(controlv1alpha1.SecCtx{}, cr.Spec.TFControl.NamedSpec.SecurityContext) {
			secContexts["named"] = cr.Spec.TFControl.NamedSpec.SecurityContext
		}
	}
	// Dns
	if &cr.Spec.TFControl.DnsSpec != nil {
		if replicas["dns"] = &defaults.Get("tf-control-dns").Size; cr.Spec.TFControl.DnsSpec.Replicas != nil {
			replicas["dns"] = cr.Spec.TFControl.DnsSpec.Replicas
		}
		if image["dns"] = defaults.Get("tf-control-dns").Image; len(cr.Spec.TFControl.DnsSpec.Image) > 0 {
			image["dns"] = cr.Spec.TFControl.DnsSpec.Image
		}
	}

	return &controlv1alpha1.TFControl{
		TypeMeta: metav1.TypeMeta{
			Kind:       "TFControl",
			APIVersion: "control.tf.mirantis.com/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "tf-control",
			Namespace: cr.Namespace,
		},
		Spec: controlv1alpha1.TFControlSpec{
			ConfigMapList: []string{"tf-rabbitmq-cfgmap", "tf-zookeeper-cfgmap", "tf-cassandra-config-cfgmap"},
			ControlSpec: controlv1alpha1.TFControlControlSpec{
				Enabled:  true,
				Replicas: replicas["control"],
				Image:    image["control"],
			},
			NamedSpec: controlv1alpha1.TFControlNamedSpec{
				Enabled:         true,
				Replicas:        replicas["named"],
				Image:           image["named"],
				SecurityContext: secContexts["named"],
			},
			DnsSpec: controlv1alpha1.TFControlDnsSpec{
				Enabled:  true,
				Replicas: replicas["dns"],
				Image:    image["dns"],
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

func getConfigMapForCassandra(cr *operatorv1alpha1.TFOperator, defaults *Entities) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "tf-cassandra-cfgmap",
			Namespace: cr.Namespace,
		},
		Data: map[string]string{
			"CONFIGDB_NODES":          defaults.Get("cassandra-config").Services[0].Name,
			"CASSANDRA_CQL_PORT":      fmt.Sprintf("%d", defaults.Get("cassandra-config").Services[0].Ports[0].Port),
			"CASSANDRA_PORT":          fmt.Sprintf("%d", defaults.Get("cassandra-config").Services[0].Ports[1].Port),
			"ANALYTICSDB_CQL_SERVERS": fmt.Sprintf("%s:%d", defaults.Get("cassandra-analytics").Services[0].Name, defaults.Get("cassandra-analytics").Services[0].Ports[0].Port),
			"ANALYTICSDB_CQL_PORT":    fmt.Sprintf("%d", defaults.Get("cassandra-analytics").Services[0].Ports[0].Port),
			"ANALYTICSDB_PORT":        fmt.Sprintf("%d", defaults.Get("cassandra-analytics").Services[0].Ports[1].Port),
		},
	}
}
