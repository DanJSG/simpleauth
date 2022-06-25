package auth

import (
	"github.com/danjsg/simpleauth/internal/auth"
	"github.com/danjsg/simpleauth/internal/collections"
	"github.com/danjsg/simpleauth/internal/logging"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestTemplating(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	baseRouter := router.Group("/api")

	log := logging.DefaultLogger()
	API := &auth.API{
		Log:      log,
		BasePath: "user",
		Version: auth.Version{
			Major: 1,
			Minor: 0,
			Patch: 0,
		},
		RelativePaths: collections.HashSet[string](),
	}
	API.RegisterHandlers(baseRouter)

	routes := router.Routes()
	for _, route := range routes {
		log.Infof("Registered: %s %s", route.Method, route.Path)
	}
}
