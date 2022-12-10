package services

import (
	"queuefly/lib/domain/models"
	"queuefly/lib/domain/repositories"
	"queuefly/lib/domain/usecases"

	"queuefly/lib/infra"
)

type UserService struct {
	*infra.EchoHandler
	repositories.UserRepository
}

func NewUserService(logger *infra.EchoHandler, repository repositories.UserRepository) UserService {
	return UserService{logger, repository}
}

func (s UserService) Create(user models.User) error {

	var UserUseCase usecases.UserUseCase

	newUser, _ := UserUseCase.Create(user)

	return s.Database.Create(&newUser).Error

}

func (s UserService) Update(user models.User) error {

	var UserUseCase usecases.UserUseCase

	updateUser, _ := UserUseCase.Update(user)

	return s.Database.Save(&updateUser).Error

}
