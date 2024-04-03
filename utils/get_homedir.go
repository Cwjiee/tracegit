package utils

import (
	"fmt"
	"os/user"
)

func getHomeDir() string {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error finding user", err)
	}

	return currentUser.HomeDir + "/.path"
}
