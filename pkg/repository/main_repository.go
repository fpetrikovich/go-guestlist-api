package repository

import (
	"database/sql"
	"fmt"
	"log"
)

/*
Implementation of a MySQL repository.

It creates a new instance of the `MySQLRepository` and sets up a connection to a MySQL database.

The connection is specified by the values of the environment variables.

The code logs a message indicating if the MySQL connection was successful or not.
*/
type MySQLRepository struct {
	Connection *sql.DB
}

func NewMySQLRepository() *MySQLRepository {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", "user", "password", "guestlist-mysql", "3306", "database")
	connection, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("[INFO] MySQL connection successful.")

	return &MySQLRepository{
		Connection: connection,
	}
}
