package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

func ExtractList() []string {
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println("error finding current direcotory: ", err)
		return nil
	}

	rubyPath := path.Join(currentPath, "readgit.rb")
	cmd := exec.Command("ruby", rubyPath)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe:", err)
		return nil
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("error starting command:", err)
		return nil
	}

	scanner := bufio.NewScanner(stdout)

	var outputLines []string

	for scanner.Scan() {
		outputLines = append(outputLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from scanner:", err)
		return nil
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for command:", err)
		return nil
	}

	var repos []string

	for _, lines := range outputLines {
		lines = strings.Trim(lines, "\"")
		repos = append(repos, lines)
	}

	return repos
}
