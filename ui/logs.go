package ui

import (
	"github.com/Cwjiee/tracegit/utils"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240")).
	Margin(0, 2)

type logModel struct {
	log, contribution table.Model
}

func (m *logModel) Init() tea.Cmd { return nil }

func (m *logModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.log.Focused() {
				m.log.Blur()
			} else {
				m.log.Focus()
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
	m.log, cmd = m.log.Update(msg)
	return m, cmd
}

func (m *logModel) View() string {
	return baseStyle.Render(m.log.View()) + "\n\n\n\tContribution\n" + baseStyle.Render(m.contribution.View())
}

func logScreen() logModel {
	// commit log
	logColumns := []table.Column{
		{Title: "Hash", Width: 60},
		{Title: "Message", Width: 50},
		{Title: "Author", Width: 30},
	}

	ContrColomns := []table.Column{
		{Title: "Contributors", Width: 40},
		{Title: "Contribution count", Width: 40},
		{Title: "Percentage", Width: 40},
	}

	rows, commiters := utils.GetLogs(redirectRepo)

	t1 := table.New(
		table.WithColumns(logColumns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(25),
	)

	t2 := table.New(
		table.WithColumns(ContrColomns),
		table.WithRows(commiters),
		table.WithFocused(false),
		table.WithHeight(10),
	)

	s1 := table.DefaultStyles()
	s1.Header = s1.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s1.Selected = s1.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	s2 := table.DefaultStyles()
	s2.Header = s2.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)

	t1.SetStyles(s1)
	t2.SetStyles(s2)

	return logModel{
		log:          t1,
		contribution: t2,
	}
}
