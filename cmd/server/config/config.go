package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
)

// Config конфигурация сервера
type Config struct {
	Dsn            string `mapstructure:"dsn"`         // Строка подключения к БД
	ServerHostPort string `mapstructure:"server-addr"` // Хост и порт для запуска grpc Сервера
	Authkey        string `mapstructure:"authkey"`     // Ключ для создания и валидации jwt токена
}

func ReadConfig() (*Config, error) {
	var configPath string

	flag.StringVar(&configPath, "config", "./cmd/server/conf.yaml", "Путь до файла конфигурации")

	flag.Parse()

	v := viper.New()

	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config

	err := v.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return &config, nil
}
