package main

import (
	"github.com/danjsg/simpleauth/internal/auth"
	"github.com/danjsg/simpleauth/internal/logging"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	log := logging.DefaultLogger()
	ginLogger := logging.GinLogger(log)
	router.Use(ginLogger, gin.Recovery())
	baseRouter := router.Group("/api")

	API := auth.UserAPI{
		Version: auth.Version{
			Major: 1,
			Minor: 0,
			Patch: 0,
		},
		BasePath: "/auth",
		Log:      log,
	}
	API.RegisterHandlers(baseRouter)

	err := router.Run()
	if err != nil {
		log.Fatalf("Fatal error encountered whilst running server. Error: %s", err)
	}
}
