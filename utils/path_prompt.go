package utils

import (
	"github.com/charmbracelet/huh"
)

func pathPrompt() string {
	var path string

	huh.NewInput().
		Title("Enter your code directory path").
		Prompt("path: ").
		Value(&path).
		Run()

	return path
}
