package main

import (
	"api/internal/service/config"
	"api/internal/util/db"
	"api/test"
)

func main() {
	configService := config.NewConfigService()
	conf, err := configService.GetConfig()
	if err != nil {
		panic(err)
	}

	session, err := db.Connect(conf.Postgres.DatabaseDSN)
	if err != nil {
		panic(err)
	}

	dataGenerator := test.NewGenerator()
	dataBuilder := test.NewBuilder(session)

	pet1 := dataGenerator.CreatePet()
	pet2 := dataGenerator.CreatePet()

	err = dataBuilder.
		Reset().
		AddPets(pet1, pet2).
		Submit()
	if err != nil {
		panic(err)
	}
}
