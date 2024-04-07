package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Cwjiee/tracegit/utils"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)
var items []list.Item
var redirectRepo string

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type listKeyMap struct {
	editPath key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		editPath: key.NewBinding(
			key.WithKeys("ctrl + e"),
			key.WithHelp("ctrl + e", "edit code path"),
		),
	}
}

type model struct {
	list      list.Model
	textInput textinput.Model
	keys      *listKeyMap
	EditMode  bool
}

func newModel(items []list.Item, currentPath string) model {
	var listKeys = newListKeyMap()

	repoList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	repoList.Title = "Git Directories"
	repoList.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.editPath,
		}
	}

	textInput := textinput.New()
	textInput.Placeholder = "/Users/weijie/code"
	textInput.SetValue(currentPath)
	textInput.Focus()
	textInput.CharLimit = 35
	textInput.Width = 25

	return model{
		list:      repoList,
		keys:      listKeys,
		textInput: textInput,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			return m, tea.Quit
		}
	}

	if !m.EditMode {
		return updateMain(msg, m)
	}
	return updateEdit(msg, m)
}

func updateMain(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyMsg:
		if m.list.FilterState() == list.Filtering {
			break
		}
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyCtrlE:
			m.EditMode = true
			return m, nil
		case tea.KeyEnter:
			itemIndex := m.list.Index()
			redirectRepo = items[itemIndex].FilterValue()
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func updateEdit(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			utils.WritePath(m.textInput.Value())
			getFormatedData()
			cmd = m.list.SetItems(items)
			m.EditMode = false
			return m, cmd
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if !m.EditMode {
		return docStyle.Render(m.list.View())
	} else {
		return docStyle.Render("\n" + editView(m) + "\n")
	}
}

func editView(m model) string {
	return fmt.Sprintf(
		"Edit your path\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func getFormatedData() []list.Item {
	items = nil
	pathExist := utils.DotpathExist()
	prefix := utils.GetPath(pathExist)

	repos := utils.ExtractList(prefix)
	descs := utils.ExtractDesc(repos)

	for i, repo := range repos {
		repo, _ = strings.CutPrefix(repo, prefix+"/")
		items = append(items, item{title: repo, desc: descs[i]})
	}
	return items
}

func main() {
	getFormatedData()
	prefix := utils.GetPath(true)

	p := tea.NewProgram(newModel(items, prefix), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	if redirectRepo != "" {
		redirectRepoPath := prefix + "/" + redirectRepo
		cmd := exec.Command("zed", redirectRepoPath)
		if err := cmd.Run(); err != nil {
			cmd = exec.Command("vim", redirectRepoPath)
			cmd.Run()
		}
	}
}
