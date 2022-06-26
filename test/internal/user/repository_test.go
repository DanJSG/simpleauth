package user

import (
	"fmt"
	"github.com/danjsg/simpleauth/internal/user"
	"sync"
	"testing"
)

func TestMemoryRepository_CreateUser(t *testing.T) {
	repo := user.NewMemoryRepository()
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
	fmt.Printf("Length: %d\n", len(repo.getUsers()))

	//fmt.Printf("Users: %v\n", users(repo.getUsers()))
}

func createUsers(repo user.Repository, emails *[]string) {
	for _, email := range *emails {
		credentials := user.credentials{
			Email:    email,
			Password: "password",
		}
		repo.CreateUser(&credentials)
	}
}
