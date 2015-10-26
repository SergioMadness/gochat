package installer

import (
	"chat/config"
	"fmt"

	"github.com/SergioMadness/migrate/migrate"
)

func Install() {
	Config()
}

func installDb() bool {
	result := true

	fmt.Println("mysql://" + config.GetDSN())
	allErrors, ok := migrate.UpSync("mysql://"+config.GetDSN(), "./migrations")
	if !ok {
		fmt.Println("No connection to database")
		for _, errorO := range allErrors {
			fmt.Println(errorO.Error())
		}
		result = false
	}

	return result
}
