package tfoperator

import (
	"bytes"

	"gopkg.in/yaml.v2"
	corev1 "k8s.io/api/core/v1"
)

// Port - default port definition
type Port struct {
	Name string `yaml:"name"`
	Port int32  `yaml:"port"`
	// HostPort int32  `yaml:"host_port"`
}

// Service - default service definition
type Service struct {
	Name  string `yaml:"name"`
	Ports []Port `yaml:"ports"`
}

// SecCtx - security contect definition
type SecCtx struct {
	Capabilities []string `yaml:"capabilities"`
}

// Entity - generic entity for each record of default config
type Entity struct {
	Name       string          `yaml:"domain_name"`
	Size       int32           `yaml:"size"`
	Services   []Service       `yaml:"services"`
	Envs       []corev1.EnvVar `yaml:"envs"`
	Image      string          `yaml:"image"`
	SecContext SecCtx          `yaml:"securityContext"`
}

// Entities map of entities
type Entities map[string]*Entity

// Get returns entity by name or empty entity
func (s *Entities) Get(entity string) *Entity {
	if (*s)[entity] != nil {
		return (*s)[entity]
	}
	return &Entity{}
}

// Unmarshal - unmarshals entitym returns error or nil
func (s *Entities) Unmarshal(data []byte) error {
	err := yaml.NewDecoder(bytes.NewReader(data)).Decode(s)
	if err != nil {
		return err
	}
	for k, v := range *s {
		v.Name = k
	}
	return nil
}
