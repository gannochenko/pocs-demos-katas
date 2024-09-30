package main

import (
	"api/internal/dto"
	"api/internal/service/config"
	"api/internal/util/db"
)

func main() {
	configService := config.NewConfigService()
	config, err := configService.GetConfig()
	if err != nil {
		panic(err)
	}

	connection, err := db.Connect(config.Postgres.DatabaseDSN)
	if err != nil {
		panic(err)
	}

	connection.Create(&dto.Pet{})
}
