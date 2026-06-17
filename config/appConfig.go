package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	ServerPort int    `mapstructure:"server_port"`
	DbHost     string `mapstructure:"host"`
	DbPort     int    `mapstructure:"port"`
	DbUser     string `mapstructure:"user"`
	DbPassword string `mapstructure:"password"`
	DbName     string `mapstructure:"db_name"`
}

func LoadConfig(path string) (*AppConfig, error) {
	viper.SetConfigName("config") // имя файла без расширения
	viper.SetConfigType("yaml")   // тип файла
	viper.AddConfigPath(path)     // путь к файлу
	viper.AddConfigPath(".")      // текущая директория
	viper.AutomaticEnv()          // автоматически читать переменные окружения

	// Приоритет: env vars > config file

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	viper.BindEnv("host", "DB_HOST")
	viper.BindEnv("port", "DB_PORT")
	viper.BindEnv("user", "DB_USER")
	viper.BindEnv("password", "DB_PASSWORD")
	viper.BindEnv("db_name", "DB_NAME")
	viper.BindEnv("server_port", "SERVER_PORT")

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
}
