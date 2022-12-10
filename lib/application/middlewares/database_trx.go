package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"queuefly/lib/application/constants"
	"queuefly/lib/data"
	"queuefly/lib/infra"
)

type DatabaseTrx struct {
	handler infra.RequestHandler
	logger  *infra.EchoHandler
	db      data.Database
}

func statusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

func NewDatabaseTrx(
	handler infra.RequestHandler,
	logger *infra.EchoHandler,
	db data.Database,
) DatabaseTrx {
	return DatabaseTrx{
		handler: handler,
		logger:  logger,
		db:      db,
	}
}

// Setup sets up database transaction middleware
func (m DatabaseTrx) Setup() {
	m.logger.Info("setting up database transaction middleware")

	m.handler.Gin.Use(func(c *gin.Context) {
		txHandle := m.db.DB.Begin()
		m.logger.Info("beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		c.Set(constants.DBTransaction, txHandle)
		c.Next()

		// commit transaction on success status
		if statusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated, http.StatusNoContent}) {
			m.logger.Info("committing transactions")
			if err := txHandle.Commit().Error; err != nil {
				m.logger.Error("trx commit error: ")
			}
		} else {
			m.logger.Info("rolling back transaction due to status code: 500")
			txHandle.Rollback()
		}
	})
}
