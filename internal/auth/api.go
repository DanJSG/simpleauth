package auth

import (
	"fmt"
	"github.com/danjsg/simpleauth/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserAPI struct {
	Version  web.Version
	BasePath string
	Log      logrus.FieldLogger
	Service  UserService
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (API *UserAPI) RegisterHandlers(router gin.IRouter, middleware ...gin.HandlerFunc) {
	basePath := fmt.Sprintf("/v%s/%s", API.Version.MajorMinor(), API.BasePath)
	baseRouter := router.Group(basePath, middleware...)
	API.registerAllHandlers(baseRouter)
}

func (API *UserAPI) registerAllHandlers(router gin.IRouter) {
	// TODO add handlers
	router.POST("/authorize", func(context *gin.Context) {
		var creds Credentials
		err := context.ShouldBindJSON(&creds)
		if err != nil {
			web.HandleJsonPostError(context, err)
			return
		}
		// TODO finish implementation
		//refreshToken := UserAPI.UserService.authorizeUser(context, &creds)
	})
	router.POST("/", noOpHandler)
	router.GET("/token", noOpHandler)
	router.POST("/token/revoke", noOpHandler)
}

func noOpHandler(context *gin.Context) {}
