package utils

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/table"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
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

func GetLogs(repo string) []table.Row {
	r, err := git.PlainOpen(repo)
	if err != nil {
		fmt.Println(err)
	}

	ref, err := r.Head()
	if err != nil {
		fmt.Println(err)
	}

	logs, err := r.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		log.Fatalf("Failed to retrieve commit history: %v", err)
	}

	// var commits []*object.Commit
	var commits []table.Row
	err = logs.ForEach(func(commit *object.Commit) error {
		commits = append(commits, table.Row{commit.Hash.String(), commit.Message, commit.Committer.Name})
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to iterate over commits: %v", err)
	}

	return commits
}
