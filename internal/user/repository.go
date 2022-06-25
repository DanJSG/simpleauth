package user

type mongoRepository struct {
}

type repository interface {
	getUserByEmail() string
	getUserByID() string
	createUser()
	authorizeUser() string
}
