package config

type RootConfig struct {
	Hosts HTTP `yaml:"http"`
}

type HTTP struct {
	Servers []Server `yaml:"server"`
}
