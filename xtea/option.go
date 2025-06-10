package xtea

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evercyan/brick/xlodash"
)

// optionValue ...
type optionValue string

func (t optionValue) FilterValue() string { return "" }

// ----------------------------------------------------------------

// optionDelegate ...
type optionDelegate struct{}

func (t optionDelegate) Height() int                             { return 1 }
func (t optionDelegate) Spacing() int                            { return 0 }
func (t optionDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (t optionDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	value, ok := listItem.(optionValue)
	if !ok {
		return
	}
	str := fmt.Sprintf("> %s", value)
	fn := defaultStyle.Render
	if index == m.Index() {
		fn = activeStyle.Render
	}
	fmt.Fprint(w, fn(str))
	return
}

// ----------------------------------------------------------------

// optionModel ...
type optionModel struct {
	list  list.Model
	title string
	value string
}

// Init ...
func (t optionModel) Init() tea.Cmd {
	return nil
}

// Update ...
func (t optionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch mt := msg.(type) {
	case tea.KeyMsg:
		if IsQuit(mt) {
			t.value = ""
			return t, tea.Quit
		}
		if IsEnter(mt) {
			if v, ok := t.list.SelectedItem().(optionValue); ok {
				t.value = string(v)
			}
			return t, tea.Quit
		}
	case tea.WindowSizeMsg:
		t.list.SetWidth(mt.Width)
		return t, nil
	}
	var cmd tea.Cmd
	t.list, cmd = t.list.Update(msg)
	return t, cmd
}

// View ...
func (t optionModel) View() string {
	title := ""
	if t.title != "" {
		title = fmt.Sprintf("%s\n", titleStyle.Render(t.title))
	}
	return fmt.Sprintf("%s%s\n", title, t.list.View())
}

// ----------------------------------------------------------------

// Option ...
func Option(title string, options []string) (string, error) {
	if len(options) == 0 {
		return "", fmt.Errorf("未提供选项")
	}
	items := make([]list.Item, 0)
	for _, v := range options {
		items = append(items, optionValue(v))
	}
	l := list.New(items, optionDelegate{}, 0, xlodash.Min(len(options), 6))
	l.SetShowFilter(false)       // 是否显示过滤项
	l.SetFilteringEnabled(false) // 是否显示过滤项
	l.SetShowHelp(false)         // 是否显示帮助信息
	l.SetShowStatusBar(false)    // 是否显示状态栏
	l.SetShowTitle(false)        // 是否显示标题
	l.SetShowPagination(false)   // 是否显示分页
	sm := optionModel{list: l, title: title}
	sm.list.Title = title
	tp, err := tea.NewProgram(&sm).Run()
	if err != nil {
		return "", err
	}
	m := tp.(optionModel)
	if m.value == "" {
		return "", fmt.Errorf("未选中选项")
	}
	return m.value, nil
}
