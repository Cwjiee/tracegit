package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/Cwjiee/tracegit/utils"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)
var items []list.Item
var redirectRepo string
var currentPath string
var editMode bool

const (
	purple   = lipgloss.Color("#8F00FF")
	darkGray = lipgloss.Color("#767676")
)

var (
	inputTitle = lipgloss.NewStyle().Foreground(purple)
	inputHelp  = lipgloss.NewStyle().Foreground(darkGray)
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type rootScreenModel struct {
	model tea.Model
}

func RootScreen() rootScreenModel {
	var rootModel tea.Model

	if !editMode {
		screen_one := listScreen()
		rootModel = &screen_one
	} else {
		screen_two := editScreen()
		rootModel = &screen_two
	}

	return rootScreenModel{
		model: rootModel,
	}
}

func (m rootScreenModel) Init() tea.Cmd {
	return m.model.Init()
}

func (m rootScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.model.Update(msg)
}

func (m rootScreenModel) View() string {
	return m.model.View()
}

func (m rootScreenModel) SwitchScreen(model tea.Model) (tea.Model, tea.Cmd) {
	m.model = model
	return m.model, m.model.Init()
}

func InitializeScreen() {
	getFormatedData()
	currentPath = utils.GetPath(true)

	p := tea.NewProgram(RootScreen(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting program:", err)
		os.Exit(1)
	}
}

func getFormatedData() {
	items = items[:0]
	pathExist := utils.DotpathExist()
	prefix := utils.GetPath(pathExist)

	repos := utils.ExtractList(prefix)
	descs := utils.ExtractDesc(repos)

	for i, repo := range repos {
		repo, _ = strings.CutPrefix(repo, prefix+"/")
		items = append(items, item{title: repo, desc: descs[i]})
	}
}
