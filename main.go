package main

import (
	"adsons/cli-escalonador/model"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
  m := model.NewCliBubble()

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
