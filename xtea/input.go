package xtea

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// inputModel ...
type inputModel struct {
	textInput textinput.Model
	label     string
}

// Init ...
func (t inputModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update ...
func (t inputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch mt := msg.(type) {
	case tea.KeyMsg:
		if IsQuit(mt, true) {
			t.textInput.SetValue("")
			return t, tea.Quit
		}
		if IsEnter(mt) {
			return t, tea.Quit
		}
	}
	t.textInput, cmd = t.textInput.Update(msg)
	return t, cmd
}

// View ...
func (t inputModel) View() string {
	if t.label == "" {
		return t.textInput.View() + "\n"
	}
	return fmt.Sprintf("%s\n%s\n", t.label, t.textInput.View())
}

// ----------------------------------------------------------------

// Input ...
func Input(label, placeholder string) (string, error) {
	ti := textinput.New()
	ti.Focus()
	ti.Width = 1024
	ti.ShowSuggestions = true
	if placeholder != "" {
		ti.Placeholder = placeholder
		ti.SetSuggestions([]string{placeholder})
	}
	im := inputModel{
		textInput: ti,
		label:     label,
	}
	tp, err := tea.NewProgram(im).Run()
	if err != nil {
		return "", err
	}
	m := tp.(inputModel)
	value := m.textInput.Value()
	if value == "" {
		return "", fmt.Errorf("未输入文本")
	}
	return value, nil
}
