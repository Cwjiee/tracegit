package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func getPath(pathExist bool) string {

	var path string

	homeDir := getHomeDir()

	if !pathExist {
		path = pathPrompt()

		f, err := os.Create(homeDir)
		if err != nil {
			fmt.Println("Error creating file", err)
		}

		defer f.Close()

		data := []byte(path)

		err = os.WriteFile(homeDir, data, 0644)
		if err != nil {
			fmt.Println("Error writing file", err)
		}

	} else {
		data, err := os.ReadFile(homeDir)
		if err != nil {
			log.Fatal(err)
		}

		path = string(data)
		path = strings.TrimSpace(path)
	}

	return path
}
