package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Get() *Config {

	if err := godotenv.Load(); err != nil {
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
