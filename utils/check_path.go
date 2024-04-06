package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/user"
)

func DotpathExist() bool {

	homeDir := getHomeDir()
	f, err := os.Stat(homeDir)
	if errors.Is(err, fs.ErrNotExist) {
		return false
	}

	if f.Size() > 0 {
		return true
	}

	return false
}

func getHomeDir() string {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error finding user", err)
	}

	return currentUser.HomeDir + "/.path"
}
