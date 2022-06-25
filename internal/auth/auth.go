package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func (v *Version) FullyQualified() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *Version) MajorMinor() string {
	return fmt.Sprintf("%d.%d", v.Major, v.Minor)
}

type API struct {
	Version  Version
	BasePath string
	Log      logrus.FieldLogger
}

func (API *API) RegisterHandlers(router gin.IRouter) {
	API.RegisterHandlersWithMiddleware(router)
}

func (API *API) RegisterHandlersWithMiddleware(router gin.IRouter, middleware ...gin.HandlerFunc) {
	basePath := fmt.Sprintf("/v%s/%s", API.Version.MajorMinor(), API.BasePath)
	baseRouter := router.Group(basePath, middleware...)
	API.registerAllHandlers(baseRouter)
}

func noOpHandler(context *gin.Context) {}

func (API *API) registerAllHandlers(router gin.IRouter) {
	// TODO add handlers
	router.POST("/authorize", noOpHandler)
	router.POST("/", noOpHandler)
	router.GET("/token", noOpHandler)
	router.POST("/token/revoke", noOpHandler)
}
