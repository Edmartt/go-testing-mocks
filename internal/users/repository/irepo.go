package repository

import "go-testing-mocks/internal/users"


//For accessing CreateUser method according to the implementation desired
var AccessRepo UserRepository

type UserRepository interface{
	CreateUser(user users.User) (*string, error)
}
