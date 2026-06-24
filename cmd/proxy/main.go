package main

import (
	"github.com/SherazAsghar37/gowing/internal/config"
)

func main() {
	config.LoadConfigFile(config.ConfigurationFilePath)
}
