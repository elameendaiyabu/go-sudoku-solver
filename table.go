package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

var stringBoard [9][9]string

func Table(board *Board) {
	for i, col := range board {
		for j, v := range col {
			stringBoard[i][j] = strconv.Itoa(v)
		}
	}

	columns := []table.Column{
		{Title: " ", Width: 2},
		{Title: " ", Width: 2},
		{Title: " ", Width: 2},
		{Title: " ", Width: 2},
		{Title: " ", Width: 2},
		{Title: " ", Width: 2},
		{Title: " ", Width: 2},
		{Title: " ", Width: 2},
		{Title: " ", Width: 2},
	}

	rows := []table.Row{
		stringBoard[0][:],
		stringBoard[1][:],
		stringBoard[2][:],
		stringBoard[3][:],
		stringBoard[4][:],
		stringBoard[5][:],
		stringBoard[6][:],
		stringBoard[7][:],
		stringBoard[8][:],
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	s := table.DefaultStyles()

	s.Cell = s.Cell.Foreground(lipgloss.
		Color("229")).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderTop(false)
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(false).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		// Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
