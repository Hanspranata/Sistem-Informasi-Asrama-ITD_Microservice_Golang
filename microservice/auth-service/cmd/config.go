package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
	// tambahkan konfigurasi lain yang diperlukan
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %s", err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %s", err)
	}

	return &config, nil
}

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	// Gunakan config.DatabaseURL, config.ServerPort, dan konfigurasi lain sesuai kebutuhan

	// Contoh penggunaan:
	fmt.Printf("Database URL: %s\n", config.DatabaseURL)
	fmt.Printf("Server Port: %s\n", config.ServerPort)
}
