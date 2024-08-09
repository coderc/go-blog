package config

import (
	"context"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/sethvargo/go-envconfig"
	"github.com/spf13/viper"
)

func Init() {
	InitEnvParam()
	InitConfig()
}

func InitEnvParam() {
	if err := envconfig.Process(context.TODO(), envParam); err != nil {
		log.Fatalf("failed to process env: %s", err)
	}
}

func InitConfig() {
	viper.AddConfigPath(GetEnv().GetConfig())
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config file: %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("failed to unmarshal config: %s", err)
	}

	v := validator.New()
	if err := v.Struct(config); err != nil {
		log.Fatalf("failed to validate config: %s", err)
	}

	GetConfig().Print()
}
