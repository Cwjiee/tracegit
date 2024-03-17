class GitTrace
  def initialize
    @path = "/Users/weijie/code"
    @git_folders = []
    @git_paths = []
  end

  def list
    recursive_read
    truncate_git_folders
    join_string_paths
  end

  def recursive_read
    Dir.glob("**/*/", File::FNM_DOTMATCH, base: @path) do |entry_name|
      @git_folders << entry_name if entry_name.include?(".git/")
    end
  end

  def truncate_git_folders
    @git_folders.each do |entry_name|
      dot_path = entry_name.split("/")
      @git_paths << dot_path if dot_path.last == ".git"
    end
  end

  def join_string_paths
    @git_paths.map do |git_path|
      git_path.pop
      p git_path.join("/")
    end
  end

end

GitTrace.new.list
