package xfile

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// TODO 头部固定

// ReadXLSX ...
func ReadXLSX(ctx context.Context, fpath string, sheets ...string) ([][]interface{}, error) {
	f, err := excelize.OpenFile(fpath)
	if err != nil {
		return nil, err
	}
	if len(sheets) == 0 {
		sheetMap := f.GetSheetMap()
		for _, sheet := range sheetMap {
			sheets = append(sheets, sheet)
		}
	}
	list := make([][]interface{}, 0)
	for _, sheet := range sheets {
		rows := f.GetRows(sheet)
		if len(rows) == 0 {
			continue
		}
		for _, row := range rows {
			line := make([]interface{}, 0)
			for _, v := range row {
				line = append(line, v)
			}
			list = append(list, line)
		}
	}
	return list, nil
}

// WriteXLSX ...
func WriteXLSX(
	ctx context.Context,
	fpath string,
	list [][]interface{},
	rowColors ...map[int]string,
) error {
	if !strings.HasSuffix(fpath, ".xlsx") {
		fpath += ".xlsx"
	}
	f := excelize.NewFile()
	sheet1 := "Sheet1"
	f.SetActiveSheet(f.NewSheet(sheet1))
	for k, v := range list {
		// 通过这种方式区分 float64 类型字段设置值为 0 和默认值为 0
		for kk, vv := range v {
			if vvv, ok := vv.(float64); ok {
				if vvv == math.MaxFloat64 {
					v[kk] = ""
				}
			}
		}
		f.SetSheetRow(sheet1, fmt.Sprintf("A%d", k+1), &v)
		// 设置行表格颜色
		if len(rowColors) > 0 && k > 0 && k-1 < len(rowColors) {
			rowColor := rowColors[k-1]
			for colIdx, color := range rowColor {
				cell := excelize.ToAlphaString(colIdx) + fmt.Sprintf("%d", k+1)
				style, _ := f.NewStyle(fmt.Sprintf(`{"fill":{"type":"pattern","color":["%s"],"pattern":1}}`, color))
				f.SetCellStyle(sheet1, cell, cell, style)
			}
		}
	}
	return f.SaveAs(fpath)
}
