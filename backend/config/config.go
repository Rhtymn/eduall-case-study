package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var errMissingKey = fmt.Errorf("missing key")

type Config struct {
	ServerAddr string
	PostgreURL string
}

func Init() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

func Load() (Config, error) {
	c := Config{}

	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		return c, errMissingKey
	}
	c.ServerAddr = serverAddr

	postgreUrl := os.Getenv("POSTGRE_URL")
	if postgreUrl == "" {
		return c, errMissingKey
	}
	c.PostgreURL = postgreUrl

	return c, nil
}
