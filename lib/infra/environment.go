package infra

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENV"`
	DBUrl       string `mapstructure:"DB_URL"`
}

const (
	EnvFile     = "../.env"
	Production  = "production"
	Development = "development"
)

func LoadConfig() (config Config, err error) {

	viper.SetConfigFile(EnvFile)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}

func NewConfig() Config {
	config, err := LoadConfig()

	if err != nil {
		log.Fatal("Environment can't be loaded ☠ ", err)

	}

	return config
}
