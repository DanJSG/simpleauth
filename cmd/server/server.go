package main

import (
	"github.com/gin-gonic/gin"
	"goland-playground/internal/logging"
	"goland-playground/internal/user"
)

func main() {
	router := gin.New()

	log := logging.DefaultLogger()
	ginLogger := logging.GinLogger(log)
	router.Use(ginLogger, gin.Recovery())

	rootPath := router.Group("/api")
	user.RegisterHandlers(rootPath, log)

	err := router.Run()
	if err != nil {
		log.Fatalf("Fatal error encountered whilst running server. Error: %s", err)
	}
}
