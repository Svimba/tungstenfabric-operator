package tfoperator

import (
	"bytes"

	"gopkg.in/yaml.v2"
)

// Port .
type Port struct {
	Name     string `yaml:"name"`
	Port     int32  `yaml:"port"`
	HostPort int32  `yaml:"host_port"`
}

// Service .
type Service struct {
	Name  string `yaml:"name"`
	Ports []Port `yaml:"ports"`
}

// Env .
type Env struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

// Entity .
type Entity struct {
	Name     string    `yaml:"domain_name"`
	Size     int32     `yaml:"size"`
	Services []Service `yaml:"services"`
	Envs     []Env     `yaml:"envs"`
	Image    string    `yaml:"image"`
}

// Entities .
type Entities map[string]*Entity

// Get .
func (s *Entities) Get(entity string) *Entity {
	return (*s)[entity]
}

// Unmarshal .
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
