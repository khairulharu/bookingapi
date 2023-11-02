package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Get() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error when load env: %v", err.Error())
	}
	return &Config{
		Database{
			Host: "ep-cold-pine-25979554.ap-southeast-1.aws.neon.fl0.io",
			Port: "5432",
			User: "fl0user",
			Pass: "4WDZnMg0TEsq",
			Name: "booking",
		},
		Server{
			Host: "",
			Port: "8080",
		},
	}
}
