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

	category1 := dataGenerator.CreateCategory()
	pet1.CategoryID = &category1.ID

	tag1 := dataGenerator.CreateTag()
	tag2 := dataGenerator.CreateTag()

	petTag1 := dataGenerator.CreatePetTag()
	petTag1.PetID = pet1.ID
	petTag1.TagID = tag1.ID

	petTag2 := dataGenerator.CreatePetTag()
	petTag2.PetID = pet1.ID
	petTag2.TagID = tag2.ID

	err = dataBuilder.
		Reset().
		AddPets(pet1, pet2).
		AddCategories(category1).
		AddTags(tag1, tag2).
		AddPetTags(petTag1, petTag2).
		Submit()
	if err != nil {
		panic(err)
	}
}
