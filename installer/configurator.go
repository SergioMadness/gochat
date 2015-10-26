package installer

import (
	"bufio"
	"chat/config"
	"fmt"
	"os"
	"strings"
)

func Config() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("This is GoChat installer. It needs some information. Please, answer next questions.")
	fmt.Println("Database configuration:")

	fmt.Print("Host: ")
	host, _ := reader.ReadString('\n')

	fmt.Print("Login: ")
	login, _ := reader.ReadString('\n')

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')

	fmt.Print("Database name: ")
	dbname, _ := reader.ReadString('\n')

	config.DBConfig.Host = strings.Trim(host, "\n\r")
	config.DBConfig.Login = strings.Trim(login, "\n\r")
	config.DBConfig.Password = strings.Trim(password, "\n\r")
	config.DBConfig.DBname = strings.Trim(dbname, "\n\r")

	configIsSaved := false
	if config.SaveConfig() {
		fmt.Println("Database configuration is saved")
		configIsSaved = true
	} else {
		fmt.Println("Database configuration could not be saved")
	}

	if configIsSaved && installDb() {
		fmt.Println("Database is installed")
	}
}
