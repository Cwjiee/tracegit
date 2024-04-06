package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func GetPath(pathExist bool) string {

	var path string

	if !pathExist {
		path = PathPrompt()
		WritePath(path)
	} else {

		homeDir := getHomeDir()

		data, err := os.ReadFile(homeDir)
		if err != nil {
			log.Fatal(err)
		}

		path = string(data)
		path = strings.TrimSpace(path)
	}

	return path
}

func WritePath(path string) {
	homeDir := getHomeDir()
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
}
