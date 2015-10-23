package config

/**
* Different common configuration functions
 */

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

import _ "github.com/go-sql-driver/mysql"

type DBConfiguration struct {
	Host, Login, Password, DBname string
}

var connection *sql.DB
var DBConfig DBConfiguration
var isConfigLoaded bool

/**
* Returnes main DB connection
 */
func GetConnection() *sql.DB {
	var err error

	if connection == nil && ((!isConfigLoaded && LoadConfig()) || isConfigLoaded) {
		dsn := GetDSN()
		fmt.Println("DSN:" + dsn)
		connection, err = sql.Open("mysql", dsn)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return connection
}

/**
* Prepare DSN
 */
func GetDSN() string {
	var dsn string
	
	if isConfigLoaded || (!isConfigLoaded && LoadConfig()) {
		dsn = DBConfig.Login
		if DBConfig.Password != "" {
			dsn += ":" + DBConfig.Password
		}
		dsn += "@" + DBConfig.Host + "/" + DBConfig.DBname + "?charset=utf8"
	}

	return dsn
}

/**
* Load configuration from yml file
 */
func LoadConfig() bool {
	result := false

	conf, err := ioutil.ReadFile("./gochat.yml")
	if err == nil && yaml.Unmarshal(conf, &DBConfig) == nil {
		result = true
		isConfigLoaded = true
	}

	return result
}

/**
* Write DB configuration to yml file
 */
func SaveConfig() bool {
	result := false

	d, err := yaml.Marshal(&DBConfig)
	if err == nil && ioutil.WriteFile("./gochat.yml", d, 0644) == nil {
		result = true
	}

	return result
}
