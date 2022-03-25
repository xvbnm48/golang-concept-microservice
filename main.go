package main

import (
	"github.com/xvbnm48/golang-concept-microservice/app"
	"github.com/xvbnm48/golang-concept-microservice/logger"
)

func main() {
	// log.Println("Starting server on port 8080")
	logger.Log.Info("Starting server on port 8080")
	app.Start()
}
