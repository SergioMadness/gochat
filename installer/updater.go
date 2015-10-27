package installer

import (
	"chat/config"
	"fmt"

	"github.com/mattes/migrate/migrate"
)

func Update() bool {
	result := false

	allErrors, ok := migrate.UpSync("mysql://"+config.GetDSN(), "./migrations")
	if !ok {
		fmt.Println("Update failed")
		for _, errorO := range allErrors {
			fmt.Println(errorO.Error())
		}
		result = false
	} else {
		fmt.Println("Update successful")
	}

	return result
}
