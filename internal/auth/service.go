package auth

import (
	"errors"
)

type UserService struct {
	Repository *UserRepository
}

func (s *UserService) CreateUser(credentials *Credentials) error {
	// validate email and password are both present
	// validate that user does not already exist with the provided email
	// create user in repository
	// TODO implement me
	panic("implement me")
}

func (s *UserService) IssueRefreshToken(credentials *Credentials) (string, error) {
	// validate email and password are both present
	// fetch user from repository based on provided email
	// validate that the user fetched from the repository has the same password as the provided credentials
	// generate a refresh token for the user
	// add the refresh token to the user in the repository
	// return the refresh token
	//TODO implement me
	panic("implement me")
}

func (s *UserService) IssueAccessToken(refreshToken *string) (string, error) {
	// validate that the refresh token is not empty or nil
	// check that the token is valid
	// get the user ID claim from the refresh token
	// fetch the user that matches the ID from the repository
	// validate that the token is contained in the user's list of tokens
	// generate an access token for the user
	// return the access token
	panic("implement me")
}

func (s *UserService) Revoke(accessToken *string, refreshToken *string) error {
	// validate that the access token is not empty or nil
	// validate that the refresh token is not empty or nil
	// check that the access token is valid
	// check that the refresh token is valid
	// get the user ID claim from the refresh token
	// fetch the user that matches the ID from the repository
	// validate that the token is contained in the user's list of tokens
	// delete the token from the user's list of tokens
	// update the repository with the edited user
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
