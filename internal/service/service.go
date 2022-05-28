package service

import (
	"go-testing-mocks/internal/users"

	"go-testing-mocks/internal/users/repository"
)


//A simple function for changing a state to on if a user is not in DB yet
func makeUserAvailable (user users.User) string{
	state := "off"

	found := repository.UserExist(user.Name) //Real call to DB, this is where we need to mock, because we don't want to depend on external service for our testing

	if found{
		return "User already exists"
	}

	state = "on"

	return state
}

//Simply checks if an user exists, and if not adds the user to DB
func RegisterUser(user users.User) string{
	found := repository.UserExist(user.Name) //We need to mock this check, because we don't want to depend on external services for our testing

	if found{
		return "User already exists"
	}

	repository.AccessRepo.CreateUser(user) //We need to mock this too, for the reasons mentioned before

	return user.Id
}
