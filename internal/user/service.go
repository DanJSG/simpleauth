package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Service interface {
	createUser(*gin.Context, *credentials) string
	authorizeUser(*gin.Context, *credentials) string
	getAccessToken(*gin.Context, string) string
	logoutUser(*gin.Context, string, string) bool
}

type repositoryBackedService struct {
	log        logrus.FieldLogger
	repository *repository
}

func (s *repositoryBackedService) createUser(context *gin.Context, creds *credentials) string {
	emailMissing := creds.Email == ""
	passwordMissing := creds.Password == ""
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
		return ""
	}
	s.log.Infof("authPair: %#v", creds)
	//TODO finish implementing
	panic("implement me")
}

func (s *repositoryBackedService) authorizeUser(context *gin.Context, creds *credentials) string {
	//TODO implement me
	panic("implement me")
}

func (s *repositoryBackedService) getAccessToken(context *gin.Context, refreshToken string) string {
	//TODO implement me
	panic("implement me")
}

func (s *repositoryBackedService) logoutUser(context *gin.Context, accessToken string, refreshToken string) bool {
	//TODO implement me
	panic("implement me")
}
