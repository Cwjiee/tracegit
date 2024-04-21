// package ui

// import (
// 	"github.com/charmbracelet/bubbles/key"
// 	"github.com/charmbracelet/bubbles/list"
// 	tea "github.com/charmbracelet/bubbletea"
// )

// type listKeyMap struct {
// 	editPath key.Binding
// }

// func newListKeyMap() *listKeyMap {
// 	return &listKeyMap{
// 		editPath: key.NewBinding(
// 			key.WithKeys("ctrl + e"),
// 			key.WithHelp("ctrl + e", "edit code path"),
// 		),
// 	}
// }

// type listModel struct {
// 	list list.Model
// 	keys *listKeyMap
// }

// func listScreen() listModel {
// 	var listKeys = newListKeyMap()

// 	repoList := list.New(items, list.NewDefaultDelegate(), 0, 0)
// 	repoList.Title = "Git Directories"
// 	repoList.AdditionalShortHelpKeys = func() []key.Binding {
// 		return []key.Binding{
// 			listKeys.editPath,
// 		}
// 	}

// 	return listModel{
// 		list: repoList,
// 		keys: listKeys,
// 	}
// }

// func (m listModel) Init() tea.Cmd {
// 	return nil
// }

// func (m *listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	var cmd tea.Cmd

// 	switch msg := msg.(type) {
// 	case tea.WindowSizeMsg:
// 		h, v := docStyle.GetFrameSize()
// 		m.list.SetSize(msg.Width-h, msg.Height-v)
// 	case tea.KeyMsg:
// 		if m.list.FilterState() == list.Filtering {
// 			break
// 		}
// 		switch msg.Type {
// 		case tea.KeyCtrlC, tea.KeyEsc:
// 			return m, tea.Quit
// 		case tea.KeyCtrlE:
// 			editScreen := editScreen()
// 			return RootScreen().SwitchScreen(&editScreen)
// 		case tea.KeyEnter:
// 			itemIndex := m.list.Index()
// 			repo := items[itemIndex].FilterValue()
// 			redirectRepo = currentPath + "/" + repo
// 			infoScreen := infoScreen()
// 			return RootScreen().SwitchScreen(&infoScreen)
// 		}
// 	}

// 	m.list, cmd = m.list.Update(msg)
// 	return m, cmd
// }

// func (m listModel) View() string {
// 	return docStyle.Render(m.list.View())
// }

package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type listModel struct {
	list     list.Model
	quitting bool
}

func (m listModel) Init() tea.Cmd {
	return nil
}

func (m listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "e":
			editScreen := editScreen()
			return RootScreen().SwitchScreen(&editScreen)
		case "enter":
			itemIndex := m.list.Index()
			repo := items[itemIndex].FilterValue()
			redirectRepo = currentPath + "/" + repo
			logScreen := logScreen()
			return RootScreen().SwitchScreen(&logScreen)
		}
	}
	// 	var cmd tea.Cmd

	// 	switch msg := msg.(type) {
	// 	case tea.WindowSizeMsg:
	// 		h, v := docStyle.GetFrameSize()
	// 		m.list.SetSize(msg.Width-h, msg.Height-v)
	// 	case tea.KeyMsg:
	// 		if m.list.FilterState() == list.Filtering {
	// 			break
	// 		}
	// 		switch msg.Type {
	// 		case tea.KeyCtrlC, tea.KeyEsc:
	// 			return m, tea.Quit
	// 		case tea.KeyCtrlE:
	// 			editScreen := editScreen()
	// 			return RootScreen().SwitchScreen(&editScreen)
	// 		case tea.KeyEnter:
	// 			itemIndex := m.list.Index()
	// 			repo := items[itemIndex].FilterValue()
	// 			redirectRepo = currentPath + "/" + repo
	// 			infoScreen := infoScreen()
	// 			return RootScreen().SwitchScreen(&infoScreen)
	// 		}
	// 	}

	// 	m.list, cmd = m.list.Update(msg)
	// 	return m, cmd

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m listModel) View() string {
	return "\n" + m.list.View()
}

func listScreen() listModel {
	// items := []list.Item{
	// 	item("Ramen"),
	// 	item("Tomato Soup"),
	// 	item("Hamburgers"),
	// 	item("Cheeseburgers"),
	// 	item("Currywurst"),
	// 	item("Okonomiyaki"),
	// 	item("Pasta"),
	// 	item("Fillet Mignon"),
	// 	item("Caviar"),
	// 	item("Just Wine"),
	// }

	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "What do you want for dinner?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle

	return listModel{list: l}
}
