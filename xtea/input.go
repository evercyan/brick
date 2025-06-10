package xtea

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// inputModel ...
type inputModel struct {
	textInput textinput.Model
	title     string
	value     string
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
			t.value = ""
			return t, tea.Quit
		}
		if IsEnter(mt) {
			t.value = t.textInput.Value()
			return t, tea.Quit
		}
	}
	t.textInput, cmd = t.textInput.Update(msg)
	return t, cmd
}

// View ...
func (t inputModel) View() string {
	label := ""
	if t.title != "" {
		label = fmt.Sprintf("%s\n", titleStyle.Render(t.title))
	}
	return fmt.Sprintf("%s%s\n", label, t.textInput.View())
}

// ----------------------------------------------------------------

// Input ...
func Input(title, placeholder string) (string, error) {
	ti := textinput.New()
	ti.Focus()
	ti.Width = 1024
	ti.ShowSuggestions = true
	ti.Cursor.Style = activeStyle
	ti.TextStyle = activeStyle
	if placeholder != "" {
		ti.Placeholder = placeholder
		ti.SetSuggestions([]string{placeholder})
	}
	im := inputModel{
		textInput: ti,
		title:     title,
	}
	tp, err := tea.NewProgram(im).Run()
	if err != nil {
		return "", err
	}
	m := tp.(inputModel)
	if m.value == "" {
		return "", fmt.Errorf("未输入文本")
	}
	return m.value, nil
}
