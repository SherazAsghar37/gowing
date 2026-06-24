package config

import "time"

type Server struct {
	Listen       string        `yaml:"listen"`
	Name         string        `yaml:"name"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	Routes       []Route       `yaml:"routes"`
}
