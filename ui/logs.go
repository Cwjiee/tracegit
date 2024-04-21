package ui

import (
	"github.com/Cwjiee/tracegit/utils"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type logModel struct {
	table table.Model
	view  viewport.Model
}

func (m *logModel) Init() tea.Cmd { return nil }

func (m *logModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter", tea.KeyRight.String():
			infoScreen := infoScreen()
			return RootScreen().SwitchScreen(&infoScreen)
			// return m, tea.Batch(
			// 	tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			// )
		case tea.KeyLeft.String():
			listScreen := listScreen()
			return RootScreen().SwitchScreen(&listScreen)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m *logModel) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func logScreen() logModel {
	// commit log
	columns := []table.Column{
		{Title: "Hash", Width: 20},
		{Title: "Message", Width: 60},
		{Title: "Author", Width: 60},
	}

	rows := utils.GetLogs(redirectRepo)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(30),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	// contribution
	v := viewport.New(140, 10)
	v.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)
	v.SetContent("s string")

	return logModel{
		table: t,
		view:  v,
	}
}
