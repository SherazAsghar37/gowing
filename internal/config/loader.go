package config

import (
	"os"

	"github.com/SherazAsghar37/gowing/internal/utils"
	"gopkg.in/yaml.v3"
)

func LoadConfigFile(path string) *RootConfig {
	bytes, err := os.ReadFile(path)
	if err != nil {
		utils.HandleError("Unable to read config file", err)
	}

	var rootConfig RootConfig
	err = yaml.Unmarshal(bytes, &rootConfig)
	if err != nil {
		utils.HandleError("Unable to decode config file", err)
	}
	return &rootConfig
}
