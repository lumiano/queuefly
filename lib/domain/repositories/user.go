package repositories

import (
	"queuefly/lib/data"
	"queuefly/lib/infra"
)

type UserRepository struct {
	data.Database
	*infra.EchoHandler
}

func NewUserRepository(database data.Database, logger *infra.EchoHandler) UserRepository {
	return UserRepository{
		database,
		logger,
	}
}
