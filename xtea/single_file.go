package xtea

import (
	"fmt"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

// singleFileModel ...
type singleFileModel struct {
	filepicker filepicker.Model
	value      string
}

// Init ..
func (t singleFileModel) Init() tea.Cmd {
	return t.filepicker.Init()
}

// Update ...
func (t singleFileModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch mt := msg.(type) {
	case tea.KeyMsg:
		if IsQuit(mt) {
			t.value = ""
			return t, tea.Quit
		}
	}
	var cmd tea.Cmd
	t.filepicker, cmd = t.filepicker.Update(msg)
	if ok, fpath := t.filepicker.DidSelectFile(msg); ok {
		t.value = fpath
		return t, tea.Quit
	}
	return t, cmd
}

// View ...
func (t singleFileModel) View() string {
	return t.filepicker.View()
}

// ----------------------------------------------------------------

// SingleFile 选择单个文件
func SingleFile(targetDir string, allowTypes ...string) (string, error) {
	fp := filepicker.New()
	fp.AllowedTypes = allowTypes
	fp.CurrentDirectory = targetDir
	sfm := singleFileModel{filepicker: fp}
	tp, err := tea.NewProgram(&sfm).Run()
	if err != nil {
		return "", err
	}
	m := tp.(singleFileModel)
	if m.value == "" {
		return "", fmt.Errorf("未选择文件")
	}
	return m.value, nil
}
