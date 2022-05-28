package database

import "database/sql"

var connector DBConnection //Acts as tunnel to GetConnection method if implemented


type DBConnection interface{
	GetConnection() (*sql.DB, error)
}
