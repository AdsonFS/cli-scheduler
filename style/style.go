package style

import "github.com/charmbracelet/lipgloss"

var Base = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

var Help = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

