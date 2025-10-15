package main

import (
	"gateway/internal/database"
	"gateway/internal/service/config"
	"log/slog"
	"os"

	"github.com/pkg/errors"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	configService := config.NewConfigService()

	if err := configService.LoadConfig(); err != nil {
		panic(errors.Wrap(err, "could not load config"))
	}

	db := database.NewDatabase(&configService.Config.Database, log)
	
	closeDb, err := db.Connect()
	if err != nil {
		panic(err)
	}

	defer closeDb()

	// Migrate the schema
	db.DB.AutoMigrate(&database.WebhookEvent{})
}
