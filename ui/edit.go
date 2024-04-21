package ui

import (
	"fmt"

	"github.com/Cwjiee/tracegit/utils"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type editModel struct {
	textInput textinput.Model
}

func editScreen() editModel {
	textInput := textinput.New()
	textInput.Placeholder = "/Users/weijie/code"
	textInput.SetValue(currentPath)
	textInput.Focus()
	textInput.CharLimit = 35
	textInput.Width = 25

	return editModel{
		textInput: textInput,
	}

}

func (m *editModel) Init() tea.Cmd {
	return nil
}

func (m *editModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			utils.WritePath(m.textInput.Value())
			getFormatedData()
			currentPath = utils.GetPath(true)
			// listScreen := listScreen()
			// return RootScreen().SwitchScreen(&listScreen)
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m *editModel) View() string {
	return fmt.Sprintf(
		"\n%s\n\n%s\n\n%s\n",
		inputTitle.Width(20).Render("Edit your path"),
		m.textInput.View(),
		inputHelp.Width(50).Render("(esc to quit)  (enter to confirm changes)"),
	) + "\n"
}
