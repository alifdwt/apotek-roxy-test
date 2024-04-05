package main

import (
	"apotek-roxy/handler"
	"apotek-roxy/pkg/database/migration"
	"apotek-roxy/pkg/database/postgres"
	"apotek-roxy/pkg/dotenv"
	"apotek-roxy/pkg/logger"
	"apotek-roxy/repository"
	"apotek-roxy/service"

	"go.uber.org/zap"
)

func main() {
	log, err := logger.NewLogger()
	if err != nil {
		log.Error("Error while initializing logger", zap.Error(err))
	}

	config, err := dotenv.LoadConfig(".")
	if err != nil {
		log.Error("Error while loading config", zap.Error(err))
	}

	db, err := postgres.NewClient(
		config.DB_HOST,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_NAME,
		config.DB_PORT,
		config.APP_TIMEZONE,
	)
	if err != nil {
		log.Error("Error while connecting to database", zap.Error(err))
	}

	err = migration.RunMigration(db)
	if err != nil {
		log.Error("Error while running migration", zap.Error(err))
	}

	repository := repository.NewRepository(db)

	service := service.NewService(service.Deps{
		Repository: repository,
		Logger:     *log,
	})

	myHandler := handler.NewHandler(service)

	err = myHandler.Init().Run(config.SERVER_ADDRESS)
	if err != nil {
		log.Error("Error while running server", zap.Error(err))
	}
}
