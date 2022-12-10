package main

import (
	"github.com/joho/godotenv"
	"queuefly/lib/presentation"
)

func main() {

	_ = godotenv.Load()

	presentation.RootApp.Execute()

}
