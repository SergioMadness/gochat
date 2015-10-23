package installer

import (
	"chat/config"
	"fmt"
	"os"

	"github.com/mattes/migrate/migrate"
)

func Uninstall() {
	if unstallDb() {
		fmt.Println("Database is cleared")

		err := os.Remove("./gochat.yml")
		if err != nil {
			fmt.Println("Can't remove configuration file")
		}
	}
}

func unstallDb() bool {
	result := true

	fmt.Println("mysql://" + config.GetDSN())
	allErrors, ok := migrate.DownSync("mysql://"+config.GetDSN(), "./migrations")
	if !ok {
		fmt.Println("No connection to database")
		for _, errorO := range allErrors {
			fmt.Println(errorO.Error())
		}
		result = false
	}

	return result
}
