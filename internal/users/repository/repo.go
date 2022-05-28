package repository

import (
	"fmt"

	"log"

	"github.com/google/uuid"

	"go-testing-mocks/internal/users"

	conn "go-testing-mocks/database"

)


type Repo struct{}

func init(){
	AccessRepo = Repo{}
}

func(r Repo) CreateUser(user users.User)(*string, error){
	return AccessRepo.CreateUser(user)
}


func CreateUser (user users.User) (*string, error){
	log.Println("======Creating User======")
	
	connection := conn.GetNewCon()

	query := "INSERT INTO users (id, name) VALUES(?, ?)"

	prStatement, pError := connection.Prepare(query)

	if pError != nil{
		return nil, fmt.Errorf("Error ocurred %s", pError.Error())
	}

	found := UserExist(user.Name)

	if found{
		return nil, fmt.Errorf("name %s is already registered", user.Name)
	}
	user.Id = uuid.NewString()

	prStatement.Exec(user.Id, user.Name)

	log.Println("======User created======")


	return &user.Id, nil
}


var UserExist = func(username string) bool{
	con := conn.GetNewCon()

	query := "SELECT id FROM users WHERE name = ?"

	ps, psErr := con.Query(query, username)

	if psErr != nil{
		log.Println(psErr.Error())
	}

	if ps.Next(){
		return true
	}

	return false
}
