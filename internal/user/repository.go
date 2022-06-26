package user

import (
	"fmt"
	"github.com/danjsg/simpleauth/internal/collections"
	"math"
	"sync"
)

type User struct {
	Id    [12]byte
	Email string `json:"email"`
	// TODO replace this with a hash of the password
	Password string                  `json:"password"`
	Tokens   collections.Set[string] `json:"tokens"`
}

type Updater interface {
	CreateUser(*Credentials) *User
	UpdateUser(*User)
}

type Retriever interface {
	GetUserByEmail(*string) *User
	GetUserById(*[12]byte) *User
	GetUsers() []*User
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

func (m *memoryRepository) CreateUser(credentials *Credentials) *User {
	user := User{
		Id:       createIDFromCount(m.count),
		Email:    credentials.Email,
		Password: credentials.Password,
		Tokens:   collections.HashSet[string](),
	}
	m.writeMutex.Lock()
	fmt.Printf("Adding user with email %s\n", user.Email)
	m.usersById[user.Id] = &user
	m.usersByEmail[user.Email] = &user
	m.users = append(m.users, &user)
	m.writeMutex.Unlock()
	return &user
}

func (m *memoryRepository) UpdateUser(user *User) {
	*m.usersByEmail[user.Email] = *user
}

func (m *memoryRepository) GetUserByEmail(email *string) *User {
	return m.usersByEmail[*email]
}

func (m *memoryRepository) GetUserById(id *[12]byte) *User {
	return m.usersById[*id]
}

func (m *memoryRepository) GetUsers() []*User {
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
