package user

import (
	"fmt"
	"github.com/danjsg/simpleauth/internal/collections"
	"math"
	"sync"
)

type User struct {
	Id           [12]byte
	Email        string                  `json:"email"`
	PasswordHash string                  `json:"passwordHash"`
	Tokens       collections.Set[string] `json:"tokens"`
}

type Updater interface {
	CreateUser(*credentials) *User
	AuthorizeUser(*credentials) *User
}

type Retriever interface {
	GetUserByEmail(*string) *User
	GetUserById(*[12]byte) *User
	getUsers() []*User
}

type Repository interface {
	Updater
	Retriever
}

type memoryRepository struct {
	writeMutex   sync.Mutex
	usersByEmail map[string]*User
	usersById    map[[12]byte]*User
	users        []*User
	count        int
}

func NewMemoryRepository() Repository {
	return &memoryRepository{
		usersByEmail: make(map[string]*User),
		usersById:    make(map[[12]byte]*User),
		users:        make([]*User, 0, 32),
		count:        0,
	}
}

func (m *memoryRepository) CreateUser(credentials *credentials) *User {
	user := User{
		Id:    createIDFromCount(m.count),
		Email: credentials.Email,
		// TODO add hash function
		PasswordHash: credentials.Password,
		Tokens:       collections.HashSet[string](),
	}
	m.writeMutex.Lock()
	fmt.Printf("Adding user with email %s\n", user.Email)
	m.usersById[user.Id] = &user
	m.usersByEmail[user.Email] = &user
	m.users = append(m.users, &user)
	m.writeMutex.Unlock()
	return &user
}

func (m *memoryRepository) AuthorizeUser(credentials *credentials) *User {
	//TODO implement me
	panic("implement me")
}

func (m *memoryRepository) GetUserByEmail(email *string) *User {
	//TODO implement me
	panic("implement me")
}

func (m *memoryRepository) GetUserById(id *[12]byte) *User {
	//TODO implement me
	panic("implement me")
}

func (m *memoryRepository) getUsers() []*User {
	return m.users
}

func createIDFromCount(count int) [12]byte {
	intPow := func(x, y int) int { return int(math.Pow(float64(x), float64(y))) }
	return [12]byte{
		byte((count / intPow(math.MaxUint8, 0)) % 255),
		byte((count / intPow(math.MaxUint8, 1)) % 255),
		byte((count / intPow(math.MaxUint8, 2)) % 255),
		byte((count / intPow(math.MaxUint8, 3)) % 255),
		byte((count / intPow(math.MaxUint8, 4)) % 255),
		byte((count / intPow(math.MaxUint8, 5)) % 255),
		byte((count / intPow(math.MaxUint8, 6)) % 255),
		byte((count / intPow(math.MaxUint8, 7)) % 255),
		byte((count / intPow(math.MaxUint8, 8)) % 255),
		byte((count / intPow(math.MaxUint8, 9)) % 255),
		byte((count / intPow(math.MaxUint8, 10)) % 255),
		byte((count / intPow(math.MaxUint8, 11)) % 255),
	}
}
