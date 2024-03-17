class GitTrace
  def initialize
    @path = "/Users/weijie/code"
    @git_folders = []
    @splited_paths = []
    @git_paths = []
  end

  def list
    recursive_read
    truncate_git_folders
    join_string_paths

    @git_paths.each { |git_path| p git_path }
  end

  def recursive_read
    Dir.glob("**/*/", File::FNM_DOTMATCH, base: @path) do |entry_name|
      @git_folders << entry_name if entry_name.include?(".git/")
    end
  end

  def truncate_git_folders
    @git_folders.each do |entry_name|
      dot_path = entry_name.split("/")
      @splited_paths << dot_path if dot_path.last == ".git"
    end
  end

  def join_string_paths
    @splited_paths.each do |git_path|
      git_path.pop
      @git_paths << git_path.join("/")
    end
  end

end

GitTrace.new.list
