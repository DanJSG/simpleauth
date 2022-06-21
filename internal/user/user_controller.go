package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandlers(engine gin.IRouter, log *logrus.Logger) {
	v1 := engine.Group("/v1/user")
	registerHandlers(v1, log)
}

func registerHandlers(router gin.IRouter, log *logrus.Logger) {
	router.POST("/authorize", func(context *gin.Context) {
		var authPair credentials
		err := context.ShouldBindJSON(&authPair)
		if err != nil {
			if errors.Is(err, io.EOF) {
				context.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
					"message": "Request body cannot be empty.",
				})
			} else {
				context.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
					"message": "An unexpected error occurred.",
				})
			}
			return
		}
		log.Infof("authPair: %+v", authPair)
	})
}
