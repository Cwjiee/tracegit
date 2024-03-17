class GitTrace
  def initialize
    @path = "/Users/weijie/code"
    @git_folders = []
    @git_paths = []
  end

  def list
    recursive_read
    truncate_git_folders
    @git_paths.each { |path| puts path.join("/") }
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

GitTrace.new.list
