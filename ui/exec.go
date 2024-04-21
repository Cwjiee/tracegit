package ui

import (
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type editorFinishedMsg struct{ err error }

var choices = []string{"nvim", "vim", "vscode", "nano", "zed"}

func openEditor(choice string) tea.Cmd {
	// editor := os.Getenv("EDITOR")
	// err := os.Chdir(redirectRepo)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	editor := choice
	if choice == "" {
		editor = "vim"
	}

	if choice == "vscode" {
		editor = "code"
	}

	c := exec.Command(editor, redirectRepo)
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return editorFinishedMsg{err}
	})
}

type execModel struct {
	cursor int
	choice string
}

func (m *execModel) Init() tea.Cmd {
	return nil
}

func (m *execModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			m.choice = choices[m.cursor]
			return m, openEditor(m.choice)

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		case tea.KeyLeft.String():
			listScreen := listScreen()
			return RootScreen().SwitchScreen(&listScreen)
		}
	}

	return m, nil
}

func (m *execModel) View() string {
	s := strings.Builder{}
	i := inputTitle.Copy().Render("Which editor do you want to use?")

	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return i + "\n\n" + s.String()
}

func execScreen() execModel {
	return execModel{}
}
