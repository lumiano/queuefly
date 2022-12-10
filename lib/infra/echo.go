package infra

import "go.uber.org/zap"

type EchoHandler struct {
	*zap.Logger
}

func getEchoEnvironment() *zap.Logger {

	config := NewConfig()

	if config.Environment == Development {
		log, _ := zap.NewDevelopment()

		return log

	}

	log, _ := zap.NewProduction()

	return log

}

func NewEchoHandler() *EchoHandler {

	logger := getEchoEnvironment()

	return &EchoHandler{
		logger,
	}
}
