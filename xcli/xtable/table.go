package xtable

// Package xtable 终端表格渲染
//  list := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
//  xtable.New(list).Render()

import (
	"fmt"
	"reflect"
)

type table struct {
	elem    interface{} // 原始数据
	style   int         // 边界样式枚举值
	border  bool        // 是否显示下边界
	headers []string    // 表格头部
	widths  []int       // 列宽
	rows    [][]string  // 显示数据
}

// New ...
func New(elem interface{}) Table {
	return &table{
		elem: elem,
	}
}

// ----------------------------------------------------------------

// Style 设置边界样式
func (t *table) Style(s int) Table {
	if _, ok := styles[s]; ok {
		t.style = s
	}
	return t
}

// Border 设置是否显示内容区下边界
func (t *table) Border(border bool) Table {
	t.border = border
	return t
}

// Header 设置表格头部数据
func (t *table) Header(v []string) Table {
	if len(v) > 0 {
		t.headers = v
	}
	return t
}

// Text 返回渲染后文本
func (t *table) Text() (content string) {
	err := t.parse()
	if err != nil {
		return err.Error()
	}
	if len(t.rows) == 0 {
		return content
	}
	b := styles[t.style]
	headerT, headerM, headerB, footer := []rune{b.DR}, []rune{b.V}, []rune{b.VR}, []rune{b.UR}
	for i, width := range t.widths {
		// 预留左右两空格长度
		headerT = append(headerT, []rune(repeat(b.H, width+2)+string(b.HD))...)
		headerB = append(headerB, []rune(repeat(b.H, width+2)+string(b.VH))...)
		footer = append(footer, []rune(repeat(b.H, width+2)+string(b.HU))...)
		if len(t.headers) > 0 {
			l := width - getLentgh([]rune(t.headers[i])) + 1
			headerM = append(headerM, []rune(" "+t.headers[i]+repeat(' ', l)+string(b.V))...)
		}
	}
	headerT[len(headerT)-1], headerB[len(headerB)-1], footer[len(footer)-1] = b.DL, b.VL, b.UL

	// 头部区域
	content += string(headerT) + "\n"
	if len(t.headers) > 0 {
		content += string(headerM) + "\n"
		if t.style != Markdown {
			content += string(headerB) + "\n"
		}
	}

	if t.style == Markdown {
		// markdown 表格表头下一行
		// | --- | --- | --- |
		content += string(b.V)
		for range t.widths {
			content += " --- " + string(b.V)
		}
		content += "\n"
	}

	// 内容区域
	for i, row := range t.rows {
		body := []rune{b.V}
		for i, width := range t.widths {
			l := width - getLentgh([]rune(row[i])) + 1
			body = append(body, []rune(" "+row[i]+repeat(' ', l)+string(b.V))...)
		}
		content += string(body) + "\n"
		if t.border && i != len(t.rows)-1 && t.style != Markdown {
			content += string(headerB) + "\n"
		}
	}

	// 底部区域
	content += string(footer)
	return content
}

// Render 输出渲染后文本
func (t *table) Render() {
	fmt.Println(t.Text() + "\n")
}

// ----------------------------------------------------------------

// parse 解析元数据
func (t *table) parse() (err error) {
	value := reflect.ValueOf(t.elem)
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		return errType
	}
	list := make([]interface{}, value.Len())
	for i := 0; i < value.Len(); i++ {
		list[i] = value.Index(i).Interface()
	}
	for index, item := range list {
		iv, it := reflect.ValueOf(item), reflect.TypeOf(item)
		if iv.Kind() == reflect.Ptr {
			iv = iv.Elem()
			it = it.Elem()
		}
		headerLen := len(t.headers)
		if iv.Kind() == reflect.Struct {
			if headerLen > 0 && headerLen != iv.NumField() {
				return errHeader
			}
			row := make([]string, 0)
			for n := 0; n < iv.NumField(); n++ {
				cn := it.Field(n).Name
				cv := fmt.Sprintf("%+v", iv.FieldByName(cn).Interface())
				row = append(row, cv)
				// 解析结构体中的 "table" 作为头部标题
				if index == 0 {
					ct := it.Field(n).Tag.Get("table")
					if ct == "" {
						ct = cn
					}
					if headerLen == 0 {
						t.headers = append(t.headers, ct)
					}
					t.widths = append(t.widths, len(ct))
				}
				if t.widths[n] < len(cv) {
					t.widths[n] = len(cv)
				}
				if headerLen > 0 && len(t.headers[n]) > t.widths[n] {
					t.widths[n] = len(t.headers[n])
				}
			}
			t.rows = append(t.rows, row)
		} else if iv.Kind() == reflect.Slice || iv.Kind() == reflect.Array {
			if headerLen > 0 && headerLen != iv.Len() {
				return errHeader
			}
			row := make([]string, 0)
			for n := 0; n < iv.Len(); n++ {
				cv := fmt.Sprintf("%+v", iv.Index(n).Interface())
				row = append(row, cv)
				if index == 0 {
					t.widths = append(t.widths, len(cv))
				}
				if len(cv) > t.widths[n] {
					t.widths[n] = len(cv)
				}
				if headerLen > 0 && len(t.headers[n]) > t.widths[n] {
					t.widths[n] = len(t.headers[n])
				}
			}
			t.rows = append(t.rows, row)
		} else {
			return errType
		}
	}
	return err
}
