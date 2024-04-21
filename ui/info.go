package ui

import (
	"github.com/Cwjiee/tracegit/utils"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

type infoModel struct {
	viewport viewport.Model
}

func infoScreen() infoModel {
	var width = 150

	content := utils.GetContent(redirectRepo)

	view := viewport.New(width, 40)
	view.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	renderer, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)

	str, _ := renderer.Render(content)
	view.SetContent(str)

	return infoModel{
		viewport: view,
	}
}

func (m *infoModel) Init() tea.Cmd {
	return nil
}

func (m *infoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	// case tea.WindowSizeMsg:
	// 	h, v := docStyle.GetFrameSize()
	// 	m.list.SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEsc:
			listScreen := listScreen()
			return RootScreen().SwitchScreen(&listScreen)
		case tea.KeyLeft:
			logScreen := logScreen()
			return RootScreen().SwitchScreen(&logScreen)
		}
	default:
		return m, nil
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m *infoModel) View() string {
	return docStyle.Render(m.viewport.View() + m.infoView())
}

func (m infoModel) infoView() string {
	return helpStyle("\n  ↑/↓: Navigate • esc: Quit\n")
}
