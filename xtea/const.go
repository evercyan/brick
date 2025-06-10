package xtea

import (
	"github.com/charmbracelet/lipgloss"
)

// ...
var (
	defaultStyle = lipgloss.NewStyle()
	titleStyle   = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#8E8E8E",
		Dark:  "#747373",
	})
	activeStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#5DF586"))
)
