package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort       string `mapstructure:"SERVER_PORT"`
	EnvType          string `mapstructure:"ENV_TYPE"`
	DBSource         string `mapstructure:"DB_SOURCE"`
	WebappBaseUrl    string `mapstructure:"WEBAPP_BASE_URL"`
	SurveyAPIHost    string `mapstructure:"SURVEY_API_HOST"`
	SurveyBaseURL    string `mapstructure:"SURVEY_BASE_URL"`
	SecretSessionKey string `mapstructure:"SECRET_SESSION_KEY"`
}

var prodRequiredVariables = [7]string{
	"SERVER_PORT",
	"ENV_TYPE",
	"DB_SOURCE",
	"WEBAPP_BASE_URL",
	"SURVEY_API_HOST",
	"SURVEY_BASE_URL",
	"SECRET_SESSION_KEY",
}

var devRequiredVariables = [7]string{
	"SERVER_PORT",
	"ENV_TYPE",
	"DB_SOURCE",
	"WEBAPP_BASE_URL",
	"SURVEY_API_HOST",
	"SURVEY_BASE_URL",
	"SECRET_SESSION_KEY",
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(path + ".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("could not loadconfig: %v", err)
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("could not loadconfig: %v", err)
	}

	const devEnvType = "dev"
	const prodEnvType = "prod"

	if config.EnvType != devEnvType && config.EnvType != prodEnvType {
		fmt.Println("!!!!!!!!!!!!!!!!!!!")
		log.Fatalf(`ENV_TYPE variable should be "dev" or "prod"`)
	}

	if config.EnvType == devEnvType {
		for _, key := range devRequiredVariables {
			if !viper.IsSet(key) {
				fmt.Println("!!!!!!!!!!!!!!!!!!!")
				log.Fatalf("required environment variable %s is not set", key)
			}
		}
	}

	if config.EnvType == prodEnvType {
		for _, key := range prodRequiredVariables {
			if !viper.IsSet(key) {
				fmt.Println("!!!!!!!!!!!!!!!!!!!")
				log.Fatalf("required environment variable %s is not set", key)
			}
		}
	}

	return
}
