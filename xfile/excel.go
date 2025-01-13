package xfile

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/evercyan/brick/xlodash"
)

// ToCSV ...
func ToCSV(ctx context.Context, fpath string, list [][]string) error {
	if len(list) == 0 {
		return nil
	}
	if !strings.HasSuffix(fpath, ".csv") {
		fpath += ".csv"
	}
	if IsExist(fpath) {
		return fmt.Errorf("file exist")
	}
	lines := xlodash.Map(list, func(k int, v []string) string {
		for kk, vv := range v {
			v[kk] = strings.ReplaceAll(vv, ",", " ")
		}
		return strings.Join(v, ", ")
	})
	return os.WriteFile(fpath, []byte(strings.Join(lines, "\n")), 0755)
}

// ToXLSX ...
func ToXLSX(ctx context.Context, fpath string, list [][]string, forces ...bool) error {
	if len(list) == 0 {
		return nil
	}
	if !strings.HasSuffix(fpath, ".xlsx") {
		fpath += ".xlsx"
	}
	if !xlodash.First(forces) && IsExist(fpath) {
		return fmt.Errorf("file exist")
	}
	f := excelize.NewFile()
	sheet1 := "Sheet1"
	f.SetActiveSheet(f.NewSheet(sheet1))
	for k, v := range list {
		f.SetSheetRow(sheet1, fmt.Sprintf("A%d", k+1), &v)
	}
	return f.SaveAs(fpath)
}

// ReadXLSX 读取 xlsx 纪录
func ReadXLSX(ctx context.Context, fpath string, sheets ...string) ([][]string, error) {
	if !IsExist(fpath) {
		return nil, fmt.Errorf("文件不存在: %s", fpath)
	}
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
	list := make([][]string, 0)
	for _, sheet := range sheets {
		rows := f.GetRows(sheet)
		if len(rows) == 0 {
			continue
		}
		list = append(list, rows...)
	}
	return list, nil
}

// ReadCSV 读取 csv 纪录
func ReadCSV(ctx context.Context, fpath string) ([][]string, error) {
	file, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return csv.NewReader(file).ReadAll()
}
