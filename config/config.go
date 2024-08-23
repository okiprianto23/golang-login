package config

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Database struct {
	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	DBPort     int    `envconfig:"DB_PORT" default:"5432"`
	DBUser     string `envconfig:"DB_USER"`
	DBPassword string `envconfig:"DB_PASSWORD"`
	DBName     string `envconfig:"DB_NAME"`
}

type Server struct {
	Application string `envconfig:"application" default:"auth"`
	Version     string `envconfig:"version" default:"1.0.0"`
	Host        string `envconfig:"host" default:"localhost"`
	Port        int    `envconfig:"port" default:"8080"`
}

type Config struct {
	Server   Server
	Database Database
}

func InitiateConfig() (config Config, err error) {
	// Mengisi config dari variabel lingkungan
	err = envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
		return
	}

	PrintConfig(config)
	return
}

func PrintConfig(c Config) {
	data, _ := json.MarshalIndent(c, "", "\t")
	fmt.Println(string(data))
}
