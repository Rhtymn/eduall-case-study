package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var errMissingKey = fmt.Errorf("error missing key")

type Config struct {
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

	postgreUrl := os.Getenv("POSTGRE_URL")
	if postgreUrl == "" {
		return c, errMissingKey
	}
	c.PostgreURL = postgreUrl

	return c, nil
}
