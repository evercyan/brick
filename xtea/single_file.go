package xtea

import (
	"fmt"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

// singleFileModel ...
type singleFileModel struct {
	filepicker   filepicker.Model
	selectedFile string
	quitting     bool
}

// Init ..
func (m singleFileModel) Init() tea.Cmd {
	return m.filepicker.Init()
}

// Update ...
func (m singleFileModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch mt := msg.(type) {
	case tea.KeyMsg:
		if IsQuit(mt) {
			m.quitting = true
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.filepicker, cmd = m.filepicker.Update(msg)
	if ok, fpath := m.filepicker.DidSelectFile(msg); ok {
		m.selectedFile = fpath
		return m, tea.Quit
	}
	return m, cmd
}

// View ...
func (m singleFileModel) View() string {
	if m.quitting {
		return ""
	}
	if m.selectedFile != "" {
		return "已选中: " + m.selectedFile
	}
	return m.filepicker.View()
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
	if m.selectedFile == "" {
		return "", fmt.Errorf("未选择文件")
	}
	return m.selectedFile, nil
}
