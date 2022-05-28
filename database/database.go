package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct{

}

func init(){
	connector = DB{}
}

func (db DB)GetConnection() (*sql.DB, error){
	connection, conError := sql.Open("sqlite3", "./testdb")
	
	if conError != nil{
		return nil, conError
	}

	return connection, nil
}


func CreateTable(){
	log.Println("======Creating Table=======")

	con := GetNewCon()

	query := "CREATE TABLE IF NOT EXISTS users (id text, name text)"

	con.Exec(query)
	log.Println("======Table created=======")
}

func GetNewCon() (*sql.DB){

	con, err := connector.GetConnection()

	if err != nil{
		log.Println(err.Error())
	}

	return con
}
