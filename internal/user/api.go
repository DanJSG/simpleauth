package user

import (
	"fmt"
	"github.com/danjsg/simpleauth/internal/web"
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
	Service  Service
}

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (API *API) RegisterHandlers(router gin.IRouter, middleware ...gin.HandlerFunc) {
	basePath := fmt.Sprintf("/v%s/%s", API.Version.MajorMinor(), API.BasePath)
	baseRouter := router.Group(basePath, middleware...)
	API.registerAllHandlers(baseRouter)
}

func (API *API) registerAllHandlers(router gin.IRouter) {
	// TODO add handlers
	router.POST("/authorize", func(context *gin.Context) {
		var creds credentials
		err := context.ShouldBindJSON(&creds)
		if err != nil {
			web.HandleJsonPostError(context, err)
			return
		}
		// TODO finish implementation
		//refreshToken := API.Service.authorizeUser(context, &creds)
	})
	router.POST("/", noOpHandler)
	router.GET("/token", noOpHandler)
	router.POST("/token/revoke", noOpHandler)
}

func noOpHandler(context *gin.Context) {}
