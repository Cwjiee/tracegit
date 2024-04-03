package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strings"
)

const rubyScript = `
    class Trace
      def initialize()
        @path = $work_dir
        @git_folders = []
        @git_paths = []
      end

      def list
        recursive_read
        truncate_git_folders

        @git_paths
          .map! { |git_path| git_path.join("/") }
          .each { |git_path| puts git_path }

        puts "DescSec"

        @git_paths.each do |git_path|
          pathname = "#{@path}/#{git_path}/.git/description"
          puts File.read(pathname)
        end
      end

      def recursive_read
        Dir.glob("**/*/", File::FNM_DOTMATCH, base: @path) do |entry_name|
          @git_folders << entry_name if entry_name.include?(".git/")
        end
      end

      def truncate_git_folders
        @git_folders.each do |entry_name|
          dot_path = entry_name.split("/")

          if dot_path.last == ".git"
            dot_path.pop # get rid of ".git"
            @git_paths << dot_path
          end
        end
      end

    end

    $work_dir = ARGV[0]
    Trace.new.list
`

func ExtractList() ([]string, []string) {

	pathExist := pathExist()

	workingDir := getPath(pathExist)

	cmd := exec.Command("ruby", "-e", rubyScript, workingDir)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing Ruby script:", err)
	}

	outputLines := strings.Split(string(output), "\n")

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

func pathExist() bool {

	homeDir := getHomeDir()
	f, err := os.Stat(homeDir)
	if errors.Is(err, fs.ErrNotExist) {
		return false
	}

	if f.Size() > 0 {
		return true
	}

	return false
}
