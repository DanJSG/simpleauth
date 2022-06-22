package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RegisterHandlers(engine gin.IRouter, log *logrus.Logger) {
	v1 := engine.Group("/v1/user")
	registerHandlers(v1, log)
}

func registerHandlers(router gin.IRouter, log *logrus.Logger) {
	router.POST("/authorize", func(context *gin.Context) { authorizeUser(context, log) })
}
