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

func handleJsonPostError(context *gin.Context, err error) {
	if errors.Is(err, io.EOF) {
		context.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"message": "Request body cannot be empty.",
		})
	} else {
		context.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": "An unexpected error occurred.",
		})
	}
}

func authorizeUser(context *gin.Context, log *logrus.Logger) {
	var authPair credentials
	err := context.ShouldBindJSON(&authPair)
	if err != nil {
		handleJsonPostError(context, err)
		return
	}
	emailMissing := authPair.Email == ""
	passwordMissing := authPair.Password == ""
	if emailMissing || passwordMissing {
		var message string
		if emailMissing && passwordMissing {
			message = "Authorization failed: No email or password provided."
		} else if emailMissing {
			message = "Authorization failed: No email provided."
		} else if passwordMissing {
			message = "Authorization failed: No password provided."
		}
		context.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
			"message": message,
		})
		return
	}
	log.Infof("authPair: %+v", authPair)
}
