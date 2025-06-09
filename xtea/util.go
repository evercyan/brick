package xtea

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evercyan/brick/xlodash"
)

// IsQuit ...
func IsQuit(mt tea.KeyMsg, ignores ...bool) bool {
	if mt.Type == tea.KeyCtrlC || mt.Type == tea.KeyEsc {
		return true
	}
	if !xlodash.First(ignores) && IsQ(mt) {
		return true
	}
	return false
}

// IsQ ...
func IsQ(mt tea.KeyMsg) bool {
	return mt.String() == "q" || mt.String() == "quit"
}

// IsEnter ...
func IsEnter(mt tea.KeyMsg) bool {
	return mt.Type == tea.KeyEnter
}
