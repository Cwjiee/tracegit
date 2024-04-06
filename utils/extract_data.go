package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func ExtractList(workingDir string) []string {

	var gitRepos []string
	err := filepath.Walk(workingDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && strings.HasSuffix(path, ".git") {
			gitDir := filepath.Dir(path)
			gitRepos = append(gitRepos, gitDir)
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		fmt.Println("err walking path ", workingDir)
	}

	return gitRepos
}

func ExtractDesc(repos []string) []string {
	var descs []string

	for _, repo := range repos {
		path := repo + "/.git/description"
		desc, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("error reading description file", err)
		}
		descs = append(descs, string(desc))
	}

	return descs
}
