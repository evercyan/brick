package xtea

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// SelectOption 选项
type SelectOption struct {
	Label string      `json:"title"`
	Value interface{} `json:"value"`
	Desc  string      `json:"desc"`
}

func (t SelectOption) Title() string       { return t.Label }
func (t SelectOption) Description() string { return t.Desc }
func (t SelectOption) FilterValue() string { return t.Label }

// ----------------------------------------------------------------

// selectModel ...
type selectModel struct {
	list  list.Model
	value *SelectOption
}

// Init ...
func (t selectModel) Init() tea.Cmd {
	return nil
}

// Update ...
func (t selectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch mt := msg.(type) {
	case tea.KeyMsg:
		if IsQuit(mt) {
			t.value = nil
			return t, tea.Quit
		}
		if IsEnter(mt) {
			v, ok := t.list.SelectedItem().(*SelectOption)
			if ok {
				t.value = v
			}
			return t, tea.Quit
		}
	case tea.WindowSizeMsg:
		t.list.SetSize(mt.Width, mt.Height)
	}
	var cmd tea.Cmd
	t.list, cmd = t.list.Update(msg)
	return t, cmd
}

// View ...
func (t selectModel) View() string {
	return t.list.View()
}

// ----------------------------------------------------------------

// Select ...
func Select(title string, options []*SelectOption) (*SelectOption, error) {
	if len(options) == 0 {
		return nil, fmt.Errorf("未提供选项")
	}
	items := make([]list.Item, 0)
	hasDesc := false
	for _, v := range options {
		items = append(items, list.Item(v))
		if v.Desc != "" {
			hasDesc = true
		}
	}
	// 是否显示描述
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = hasDesc
	l := list.New(items, delegate, 0, 0)
	l.SetShowFilter(false)       // 是否显示过滤项
	l.SetFilteringEnabled(false) // 是否显示过滤项
	l.SetShowHelp(false)         // 是否显示帮助信息
	l.SetShowStatusBar(false)    // 是否显示状态栏
	l.SetShowTitle(title != "")  // 是否显示标题
	l.SetShowPagination(false)   // 是否显示分页
	sm := selectModel{list: l}
	sm.list.Title = title
	tp, err := tea.NewProgram(&sm).Run()
	if err != nil {
		return nil, err
	}
	m := tp.(selectModel)
	if m.value == nil {
		return nil, fmt.Errorf("未选中选项")
	}
	return m.value, nil
}
