package config

import (
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server *Server
		Db     *Db
		Rabbit *Rabbit
	}

	Server struct {
		Port int
	}

	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		Charset  string
	}

	Rabbit struct {
		Host     string
		Port     int
		User     string
		Password string
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
