package main

import (
	"os"

	"backend/internal/util/db"
	"backend/test"
)

func main() {
	session, err := db.Connect(os.Getenv("POSTGRES_DB_DSN"))
	if err != nil {
		panic("could not connect to the database")
	}

	dataGenerator := test.NewGenerator()
	dataBuilder := test.NewBuilder(session)

	user1 := dataGenerator.CreateUser()
	user1.Sup = "auth0:19482"

	image1 := dataGenerator.CreateImage()
	image1.CreatedBy = user1.ID

	image2 := dataGenerator.CreateImage()
	image2.CreatedBy = user1.ID

	err = dataBuilder.Reset().AddImages(image1, image2).AddUsers(user1).Submit()
	if err != nil {
		panic(err.Error())
	}
}
