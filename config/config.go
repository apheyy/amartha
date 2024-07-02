package config

import (
	"amarthaTest/model"
	"fmt"
	"os"
)

func InitConfig() model.AppConfig {
	dbCredential := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db := model.Database{
		Driver:      "postgres",
		Credential:  dbCredential,
		MaxOpenConn: 40,
		MaxIdleConn: 20,
		MaxIdleTime: 30000,
		MaxLifeTime: 1800000,
	}

	return model.AppConfig{
		Port:     8000,
		Database: db,
	}
}
