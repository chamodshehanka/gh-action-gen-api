package config

import (
	"github.com/spf13/viper"
	"log"
	"path/filepath"
)

var config *viper.Viper

func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("default")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing default configuration file")
	}

	envConfig := viper.New()
	envConfig.SetConfigType("yaml")
	envConfig.AddConfigPath("config/")
	envConfig.SetConfigName(env)
	err = envConfig.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing env configuration file")
	}

	config.MergeConfigMap(envConfig.AllSettings())
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *viper.Viper {
	return config
}
