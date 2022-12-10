package data

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"queuefly/lib/infra"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(config infra.Config, logger *infra.EchoHandler) Database {

	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})

	if err != nil {
		logger.Info("Url: ")
		logger.Panic(err.Error())
	}

	logger.Info("Database connection established")

	return Database{
		DB: db,
	}
}
