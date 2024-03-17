package utils

import (
	"bufio"
	"fmt"
	"os/exec"
)

func ExtractList() ([]string, []string) {
	cmd := exec.Command("ruby", "/Users/weijie/code/GitTrace/trace.rb")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe:", err)
		return nil, nil
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("error starting command:", err)
		return nil, nil
	}

	scanner := bufio.NewScanner(stdout)

	var outputLines []string

	for scanner.Scan() {
		outputLines = append(outputLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from scanner:", err)
		return nil, nil
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for command:", err)
		return nil, nil
	}

	var data []string
	var repos []string
	var desc []string

	for _, lines := range outputLines {
		data = append(data, lines)
	}

	for _, lines := range data {
		if lines == "DescSec" {
			data = data[1:]
			break
		}

		repos = append(repos, lines)
		data = data[1:]
	}

	for _, lines := range data {
		desc = append(desc, lines)
	}

	return repos, desc
}
