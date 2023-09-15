package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg time.Duration

type model struct {
	spinner   spinner.Model
	timeLeft  time.Duration
	totalTime time.Duration
	mode      string
	gomos     int
	quitting  bool
}

var quitKeys = key.NewBinding(
	key.WithKeys("q", "esc", "ctrl+c"),
	key.WithHelp("", "press q to quit"),
)

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return model{
		spinner:   s,
		timeLeft:  25 * time.Minute,
		totalTime: 25 * time.Minute,
		mode:      "Work",
		gomos:     0,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(1 * time.Second)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, quitKeys) {
			m.quitting = true
			return m, tea.Quit
		}
		return m, nil

	case tickMsg:
		m.timeLeft -= time.Duration(msg)
		// newPercent := float64(100 - int((m.timeLeft*100)/m.totalTime))
		if m.timeLeft <= 0 {
			if m.mode == "Work" {
				m.gomos++
				if m.gomos%4 == 0 {
					m.mode = "Long Rest"
					m.timeLeft = 15 * time.Minute
					m.totalTime = 15 * time.Minute
				} else {
					m.mode = "Rest"
					m.timeLeft = 5 * time.Minute
					m.totalTime = 5 * time.Minute
				}
			} else {
				m.mode = "Work"
				m.timeLeft = 25 * time.Minute
				m.totalTime = 25 * time.Minute
			}
		}
		return m, tea.Tick(time.Second, func(t time.Time) tea.Msg {
			return tickMsg(1 * time.Second)
		})

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	str := fmt.Sprintf("\n\n   %s %s Time Left: %s %s\n\n", m.spinner.View(), m.mode, m.timeLeft, quitKeys.Help().Desc)
	if m.quitting {
		return str + "\n"
	}
	return str
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
