package model

import (
	"adsons/cli-escalonador/style"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type CliBubble struct {
	scheduler1 *Scheduler
	scheduler2 *Scheduler
	scheduler3 *Scheduler
}
type tickMsg struct{}

func NewCliBubble() CliBubble {
	scheduler1 := NewScheduler(3)
	scheduler1.Start()

	scheduler2 := NewScheduler(2)
	scheduler2.Start()

	scheduler3 := NewScheduler(4)
	scheduler3.Start()

	return CliBubble{
		scheduler1: scheduler1,
		scheduler2: scheduler2,
		scheduler3: scheduler3,
	}
}

func newTickMsg() tea.Cmd {
	return tea.Tick(time.Millisecond * 100, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (m CliBubble) Init() tea.Cmd {
	return newTickMsg()
}

func (m CliBubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tickMsg:
		return m, newTickMsg()
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	table, cmd := m.scheduler1.Table.Update(msg)
	m.scheduler1.Table = &table
	return m, cmd
}

func (m CliBubble) View() string {
  s := fmt.Sprintf("Scheduler %d - Quantum: %ds\n%s\n\n", 1, m.scheduler1.quantum, style.Base.Render(m.scheduler1.Table.View()))
  s += fmt.Sprintf("Scheduler %d - Quantum: %ds\n%s\n\n", 2, m.scheduler2.quantum, style.Base.Render(m.scheduler2.Table.View()))
  s += fmt.Sprintf("Scheduler %d - Quantum: %ds\n%s\n\n", 3, m.scheduler3.quantum, style.Base.Render(m.scheduler3.Table.View()))

	s += style.Help("\nDeveloped by Adson Santos\nPress q to quit.")

	return s
}
