package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Creator interface {
	CreateUser(*gin.Context, *Credentials)
}

type TokenIssuer interface {
	IssueRefreshToken(*gin.Context, *Credentials)
	IssueAccessToken(*gin.Context, *string)
}

type TokenRevoker interface {
	Revoke(*gin.Context, *string, *string)
}

type Service struct {
	repository *Repository
}

func (s *Service) CreateUser(context *gin.Context, credentials *Credentials) {
	if err := credentialsPresent(credentials); err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
			"message": err.Error(),
		})
		return
	}
	// TODO implement me
	panic("implement me")
}

func (s *Service) IssueRefreshToken(context *gin.Context, credentials *Credentials) {
	if err := credentialsPresent(credentials); err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
			"message": err.Error(),
		})
		return
	}
	//TODO implement me
	panic("implement me")
}

func (s *Service) IssueAccessToken(context *gin.Context, refreshToken *string) {
	if *refreshToken == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"message": "could not issue access token: refresh token is empty",
		})
		return
	}
	//TODO implement me
	panic("implement me")
}

func (s *Service) Revoke(context *gin.Context, accessToken *string, refreshToken *string) {
	accessTokenMissing := *accessToken == ""
	refreshTokenMissing := *refreshToken == ""
	if accessTokenMissing || refreshTokenMissing {
		var message string
		if accessTokenMissing && refreshTokenMissing {
			message = "could not revoke tokens: refresh and access tokens missing"
		} else if accessTokenMissing {
			message = "could not revoke tokens: access token missing"
		} else {
			message = "could not revoke tokens: refresh token missing"
		}
		context.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"message": message,
		})
	}
	//TODO implement me
	panic("implement me")
}

func credentialsPresent(credentials *Credentials) error {
	emailMissing := credentials.Email == ""
	passwordMissing := credentials.Password == ""
	credentialsMissing := emailMissing || passwordMissing
	if credentialsMissing {
		var message string
		if emailMissing && passwordMissing {
			message = "authorization failed: no email or password provided"
		} else if emailMissing {
			message = "authorization failed: no email provided"
		} else if passwordMissing {
			message = "authorization failed: no password provided"
		}
		return errors.New(message)
	}
	return nil
}
