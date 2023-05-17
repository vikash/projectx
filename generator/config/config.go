package config

type Config struct {
	Domains []Domain `yaml:"domains"`
	Global  Global   `yaml:"global"`
}

type Global struct {
	PackagePrefix string `yaml:"packagePrefix"`
	GenFolder     string `yaml:"genFolder"`
}

type Domain struct {
	Name     string                 `yaml:"name"`
	Entities []Entity               `yaml:"entities"`
	Database map[string]interface{} `yaml:"database"`
}

type Entity struct {
	Name   string
	Fields []Field
}

type Field map[string]string
