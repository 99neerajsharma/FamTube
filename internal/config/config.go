package config

import (
	"go.uber.org/config"
	"log"
	"path/filepath"
	"runtime"
)

func ConfigInitializer() *config.YAML {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	configFilePath := filepath.Join(basePath, "..", "..", "config.yml")
	option := config.File(configFilePath)
	yaml, err := config.NewYAML(option)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return yaml
}
