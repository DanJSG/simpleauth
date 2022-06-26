package auth

import (
	"fmt"
	"github.com/danjsg/simpleauth/pkg/collections"
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

type UserRepository interface {
	CreateUser(*Credentials) *User
	UpdateUser(*User)
	GetUserByEmail(*string) *User
	GetUserById(*[12]byte) *User
	GetUsers() []*User
}

type memoryBackedUserRepository struct {
	writeMutex   sync.Mutex
	usersByEmail map[string]*User
	usersById    map[[12]byte]*User
	users        []*User
	count        int
}

func NewUserRepository() UserRepository {
	return &memoryBackedUserRepository{
		usersByEmail: make(map[string]*User),
		usersById:    make(map[[12]byte]*User),
		users:        make([]*User, 0, 32),
		count:        0,
	}
}

func (m *memoryBackedUserRepository) CreateUser(credentials *Credentials) *User {
	user := User{
		Id:       createIDFromCount(m.count),
		Email:    credentials.Email,
		Password: credentials.Password,
		Tokens:   collections.HashSet[string](),
	}
	m.writeMutex.Lock()
	fmt.Printf("Adding auth with email %s\n", user.Email)
	m.usersById[user.Id] = &user
	m.usersByEmail[user.Email] = &user
	m.users = append(m.users, &user)
	m.writeMutex.Unlock()
	return &user
}

func (m *memoryBackedUserRepository) UpdateUser(user *User) {
	m.writeMutex.Lock()
	*m.usersByEmail[user.Email] = *user
	m.writeMutex.Unlock()
}

func (m *memoryBackedUserRepository) GetUserByEmail(email *string) *User {
	return m.usersByEmail[*email]
}

func (m *memoryBackedUserRepository) GetUserById(id *[12]byte) *User {
	return m.usersById[*id]
}

func (m *memoryBackedUserRepository) GetUsers() []*User {
	return m.users
}

func createIDFromCount(count int) [12]byte {
	intPow := func(x, y int) int { return int(math.Pow(float64(x), float64(y))) }
	return [12]byte{
		byte((count / intPow(math.MaxUint8, 0)) % math.MaxUint8),
		byte((count / intPow(math.MaxUint8, 1)) % math.MaxUint8),
		byte((count / intPow(math.MaxUint8, 2)) % math.MaxUint8),
		byte((count / intPow(math.MaxUint8, 3)) % math.MaxUint8),
		byte((count / intPow(math.MaxUint8, 4)) % math.MaxUint8),
		byte((count / intPow(math.MaxUint8, 5)) % math.MaxUint8),
		byte((count / intPow(math.MaxUint8, 6)) % math.MaxUint8),
		byte((count / intPow(math.MaxUint8, 7)) % math.MaxUint8),
		byte((count / intPow(math.MaxUint8, 8)) % math.MaxUint8),
		byte((count / intPow(math.MaxUint8, 9)) % math.MaxUint8),
		byte((count / intPow(math.MaxUint8, 10)) % math.MaxUint8),
		byte((count / intPow(math.MaxUint8, 11)) % math.MaxUint8),
	}
}
