package config

import (
	"os"
	"strconv"
)

type config struct {
	Host       string
	Port       uint64
	User       string
	DBName     string
	DBPassword string
	SSLMode    bool
}

func GetConfig() (*config, error) {
	port, err := strconv.ParseUint(os.Getenv("APP_NAME_PORT"), 10, 64)

	if err != nil {
		return nil, err
	}

	return &config{
		Host:       os.Getenv("APP_NAME_HOST"), // TODO: Make this dynamic?
		Port:       port,
		User:       os.Getenv("APP_NAME_USER"),
		DBName:     os.Getenv("APP_NAME_DB_NAME"),
		DBPassword: os.Getenv("APP_NAME_DB_PASSWORD"),
		SSLMode:    false,
	}, nil
}
