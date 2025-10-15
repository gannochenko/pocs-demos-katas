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
	err = db.DB.AutoMigrate(&database.WebhookEvent{})
	if err != nil {
		panic(errors.Wrap(err, "could not auto-migrate schema"))
	}

	// Create trigger function to clean up old webhook events
	createTriggerFunction := `
		CREATE OR REPLACE FUNCTION cleanup_old_webhook_events()
		RETURNS TRIGGER AS $$
		BEGIN
			DELETE FROM webhook_events
			WHERE created_at < NOW() - INTERVAL '3 days';
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;
	`

	err = db.DB.Exec(createTriggerFunction).Error
	if err != nil {
		panic(errors.Wrap(err, "could not create trigger function"))
	}

	// Drop trigger if exists and recreate it
	dropTrigger := `DROP TRIGGER IF EXISTS trigger_cleanup_old_webhook_events ON webhook_events;`
	err = db.DB.Exec(dropTrigger).Error
	if err != nil {
		panic(errors.Wrap(err, "could not drop existing trigger"))
	}

	createTrigger := `
		CREATE TRIGGER trigger_cleanup_old_webhook_events
		AFTER INSERT ON webhook_events
		FOR EACH STATEMENT
		EXECUTE FUNCTION cleanup_old_webhook_events();
	`

	err = db.DB.Exec(createTrigger).Error
	if err != nil {
		panic(errors.Wrap(err, "could not create trigger"))
	}

	log.Info("Migration completed successfully, including trigger setup")
}
