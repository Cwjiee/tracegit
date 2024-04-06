package utils

import (
	"github.com/charmbracelet/huh"
)

func PathPrompt() string {
	var path string

	huh.NewInput().
		Title("Enter your code directory path").
		Placeholder("/Users/weijie/code").
		Prompt("path: ").
		Value(&path).
		Run()

	return path
}
