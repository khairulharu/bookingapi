package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get() *Config {
	if os.Getenv("APP_ENV") == "prod" {
		return &Config{
			Database{
				Host: os.Getenv("PROD_DB_HOST"),
				Port: os.Getenv("PROD_DB_PORT"),
				User: os.Getenv("PROD_DB_USER"),
				Pass: os.Getenv("PROD_DB_PASS"),
				Name: os.Getenv("PROD_DB_NAME"),
				SSL:  os.Getenv("PROD_DB_SSL"),
			},
			Server{
				Host: os.Getenv("SRV_HOST"),
				Port: os.Getenv("PORT"),
			},
		}
	} else {

		if err := godotenv.Load(); err != nil {
			log.Printf("cannot load env : %s", err.Error())
		}
		return &Config{
			Database{
				Host: os.Getenv("DB_HOST"),
				Port: os.Getenv("DB_PORT"),
				User: os.Getenv("DB_USER"),
				Pass: os.Getenv("DB_PASS"),
				Name: os.Getenv("DB_NAME"),
				SSL:  os.Getenv("DB_SSL"),
			},
			Server{
				Host: os.Getenv("SRV_HOST"),
				Port: os.Getenv("PORT"),
			},
		}

	}

}
