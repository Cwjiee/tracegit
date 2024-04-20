package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func GetContent(repo string) string {
	var content string

	err := filepath.WalkDir(repo, func(path string, d fs.DirEntry, err error) error {
		if d.Name() == "README.md" {
			fmt.Println("found it")
			data, err := os.ReadFile(path)
			if err != nil {
				fmt.Println("error reading file", err)
				return err
			}

			content = string(data)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking repo dir", err)
	}
	return content
}
