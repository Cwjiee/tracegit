package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

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

type listModel struct {
	list list.Model
	keys *listKeyMap
}

func listScreen() listModel {
	var listKeys = newListKeyMap()

	repoList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	repoList.Title = "Git Directories"
	repoList.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.editPath,
		}
	}

	return listModel{
		list: repoList,
		keys: listKeys,
	}
}

func (m *listModel) Init() tea.Cmd {
	return nil
}

func (m *listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

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
			editScreen := editScreen()
			return RootScreen().SwitchScreen(&editScreen)
		case tea.KeyEnter, tea.KeyRight:
			itemIndex := m.list.Index()
			repo := items[itemIndex].FilterValue()
			redirectRepo = currentPath + "/" + repo
			logScreen := logScreen()
			return RootScreen().SwitchScreen(&logScreen)
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *listModel) View() string {
	return docStyle.Render(m.list.View())
}
