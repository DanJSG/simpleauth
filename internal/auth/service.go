package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goland-playground/internal/web"
	"net/http"
)

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func authorizeUser(context *gin.Context, log *logrus.Logger) {
	var authPair credentials
	err := context.ShouldBindJSON(&authPair)
	if err != nil {
		web.HandleJsonPostError(context, err)
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
