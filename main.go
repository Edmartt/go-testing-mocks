package main

import (
	"fmt"

	conn "go-testing-mocks/database"
	"go-testing-mocks/internal/users"
	"go-testing-mocks/internal/users/repository"
)

func main() {

	conn.CreateTable()

	user := users.User{
		Name: "edmartt",
	}

	id, err := repository.CreateUser(user)

	if err != nil{
		fmt.Println("user exists")
	}

	fmt.Println(*id) 
}
