package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ExtractList() ([]string, []string) {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	binPath := filepath.Dir(ex)

	pathExist := pathExist(binPath)

	workingDir := getPath(pathExist, binPath)

	currentpath := binPath + "/../trace.rb"
	cmd := exec.Command("ruby", currentpath, workingDir)

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

func pathExist(binPath string) bool {

	f, err := os.Stat(binPath + "/../data/path")
	if err != nil {
		log.Fatal(err)
		return false
	}

	if f.Size() > 0 {
		return true
	}

	return false
}

func getPath(pathExist bool, binPath string) string {

	var path string

	if !pathExist {
		fmt.Println("Enter a path:")
		_, err := fmt.Scan(&path)
		if err != nil {
			log.Fatal(err)
			return ""
		}

		data := []byte(path)
		err2 := os.WriteFile(binPath+"/../data/path", data, 0644)
		if err2 != nil {
			log.Fatal(err)
		}

	} else {
		data, err := os.ReadFile(binPath + "/../data/path")
		if err != nil {
			log.Fatal(err)
		}

		path = string(data)
		path = strings.TrimSpace(path)
	}

	return path
}
