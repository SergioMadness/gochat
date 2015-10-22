package config

/**
* Different common configuration functions
*/

import (
	"database/sql"
	"fmt"
)

import _ "github.com/go-sql-driver/mysql"

var connection *sql.DB

/**
* Returnes main DB connection
*/
func GetConnection() *sql.DB {
	var err error
	if connection == nil {
		connection, err = sql.Open("mysql", "root@/go_chat?charset=utf8")
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return connection
}
