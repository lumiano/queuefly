package usecases

import "queuefly/lib/domain/models"

type UserUseCase interface {
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
}
