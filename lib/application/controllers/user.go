package controllers

import (
	"github.com/gin-gonic/gin"
	"queuefly/lib/domain/services"
	"queuefly/lib/infra"

	"net/http"
	"queuefly/lib/domain/models"
)

type UserController struct {
	services.UserService
	logger *infra.EchoHandler
}

func NewUserController(service services.UserService, logger *infra.EchoHandler) UserController {
	return UserController{service, logger}
}

func (u UserController) CreateUser(c *gin.Context) {
	user := models.User{}

	u.logger.Info("test")

	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	u.logger.Info("teste")

	c.JSON(200, gin.H{"data": "user created"})

}
