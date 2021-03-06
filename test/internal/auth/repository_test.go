package auth

import (
	"fmt"
	"github.com/danjsg/simpleauth/internal/auth"
	"sync"
	"testing"
)

func TestMemoryRepository_CreateUser(t *testing.T) {
	repo := auth.NewUserRepository()
	e1 := []string{
		"a@test.com",
		"b@test.com",
		"c@test.com",
		"d@test.com",
	}
	e2 := []string{
		"e@test.com",
		"f@test.com",
		"g@test.com",
		"h@test.com",
	}
	e3 := []string{
		"i@test.com",
		"j@test.com",
		"k@test.com",
		"l@test.com",
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		createUsers(repo, &e1)
		wg.Done()
	}()
	go func() {
		createUsers(repo, &e2)
		wg.Done()
	}()
	createUsers(repo, &e3)
	wg.Wait()
	fmt.Printf("Length: %d\n", len(repo.GetUsers()))
}

func createUsers(repo auth.UserRepository, emails *[]string) {
	for _, email := range *emails {
		credentials := auth.Credentials{
			Email:    email,
			Password: "password",
		}
		repo.CreateUser(&credentials)
	}
}

func TestMemoryRepository_UpdateUser(t *testing.T) {
	repo := auth.NewUserRepository()
	email := "test@test.com"
	credentials := auth.Credentials{Email: email, Password: "password"}
	createdUser := repo.CreateUser(&credentials)
	updatedUser := auth.User{
		Id:       createdUser.Id,
		Email:    createdUser.Email,
		Password: "newPassword",
		Tokens:   createdUser.Tokens,
	}
	repo.UpdateUser(&updatedUser)
	userRetrievedByEmail := repo.GetUserByEmail(&email)
	fmt.Printf("Email: %s; Password: %s\n", userRetrievedByEmail.Email, userRetrievedByEmail.Password)
	userRetrievedById := repo.GetUserById(&createdUser.Id)
	fmt.Printf("Email: %s; Password: %s\n", userRetrievedById.Email, userRetrievedById.Password)

}
