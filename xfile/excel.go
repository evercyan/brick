package xfile

import (
	"context"
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/evercyan/brick/xlodash"
	"github.com/evercyan/brick/xtype"
)

// ----------------------------------------------------------------

// interface2string ...
func interface2string(list [][]interface{}) [][]string {
	lines := make([][]string, 0)
	for _, v := range list {
		line := make([]string, 0)
		for _, vv := range v {
			line = append(line, fmt.Sprint(vv))
		}
		lines = append(lines, line)
	}
	return lines
}

// readExcel ...
func readExcel(ctx context.Context, fpath string) ([][]string, error) {
	if strings.HasSuffix(fpath, ".csv") {
		return ReadCSV(ctx, fpath)
	} else if strings.HasSuffix(fpath, ".xlsx") {
		list, err := ReadXLSX(ctx, fpath)
		if err != nil {
			return nil, err
		}
		return interface2string(list), nil
	}
	return nil, fmt.Errorf("invalid file ext")
}

// writeExcel ...
func writeExcel(ctx context.Context, fpath string, list [][]interface{}, forces ...bool) error {
	fdir := filepath.Dir(fpath)
	if !IsDir(fdir) {
		if err := os.MkdirAll(fdir, os.ModePerm); err != nil {
			return err
		}
	}
	if strings.HasSuffix(fpath, ".csv") {
		return WriteCSV(ctx, fpath, list, forces...)
	} else if strings.HasSuffix(fpath, ".xlsx") {
		return WriteXLSX(ctx, fpath, list, forces...)
	}
	return fmt.Errorf("invalid file ext")
}

// ----------------------------------------------------------------

// ReadCSV ...
func ReadCSV(ctx context.Context, fpath string) ([][]string, error) {
	file, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return csv.NewReader(file).ReadAll()
}

// WriteCSV ...
func WriteCSV(ctx context.Context, fpath string, list [][]interface{}, forces ...bool) error {
	if len(list) == 0 {
		return fmt.Errorf("emtpy record")
	}
	if !strings.HasSuffix(fpath, ".csv") {
		fpath += ".csv"
	}
	if !xlodash.First(forces) && IsExist(fpath) {
		return fmt.Errorf("file exist")
	}
	file, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, v := range list {
		row := make([]string, 0)
		for _, vv := range v {
			row = append(row, xtype.ToString(vv))
		}
		writer.Write(row)
	}
	return nil
}

// ----------------------------------------------------------------

// ReadXLSX ...
func ReadXLSX(ctx context.Context, fpath string, sheets ...string) ([][]interface{}, error) {
	if !IsExist(fpath) {
		return nil, fmt.Errorf("file not exist")
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
func WriteXLSX(ctx context.Context, fpath string, list [][]interface{}, forces ...bool) error {
	if len(list) == 0 {
		return fmt.Errorf("emtpy record")
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
		// 通过这种方式区分 float64 类型字段设置值为 0 和默认值为 0
		for kk, vv := range v {
			if vvv, ok := vv.(float64); ok {
				if vvv == math.MaxFloat64 {
					v[kk] = ""
				}
			}
		}
		f.SetSheetRow(sheet1, fmt.Sprintf("A%d", k+1), &v)
	}
	return f.SaveAs(fpath)
}

// ----------------------------------------------------------------

// ReadExcel ...
func ReadExcel(ctx context.Context, fpath string, list interface{}) error {
	if list == nil {
		return fmt.Errorf("list cannot be nil")
	}
	// 查询原始数据
	records, err := readExcel(ctx, fpath)
	if err != nil {
		return err
	}
	if len(records) == 0 {
		return nil
	}
	// 类型分支处理, 支持 [][]string 和 []*T
	switch t := list.(type) {
	case *[][]string:
		*t = records
		return nil
	default:
		listValue := reflect.ValueOf(list)
		if listValue.Kind() != reflect.Ptr ||
			listValue.Elem().Kind() != reflect.Slice ||
			listValue.Elem().Type().Elem().Elem().Kind() != reflect.Struct {
			return fmt.Errorf("only support []*struct")
		}
		var (
			elemType  = listValue.Elem().Type().Elem().Elem()
			elemValue = listValue.Elem()
		)
		// 解析结构体标签映射
		headerMap, fieldMap := make(map[string]string), make(map[string]reflect.Kind)
		for i := 0; i < elemType.NumField(); i++ {
			field := elemType.Field(i)
			tag := field.Tag.Get("excel")
			if tag == "" {
				continue
			}
			headerMap[tag] = field.Name
			fieldMap[field.Name] = field.Type.Kind()
		}
		if len(headerMap) == 0 {
			return fmt.Errorf("tag excel not found")
		}
		// 匹配表头与标签以及转换数据行
		header := records[0]
		columnMap := make(map[int]string)
		for k, v := range header {
			if vv, ok := headerMap[strings.TrimSpace(v)]; ok {
				columnMap[k] = vv
			}
		}
		for _, row := range records[1:] {
			newItem := reflect.New(elemType).Interface()
			itemValue := reflect.ValueOf(newItem).Elem()
			for k, v := range row {
				fieldName, ok := columnMap[k]
				if !ok {
					continue
				}
				field := itemValue.FieldByName(fieldName)
				if !field.IsValid() || !field.CanSet() {
					continue
				}
				if err := fillExcelFieldValue(field, v, fieldMap[fieldName]); err != nil {
					return fmt.Errorf("column %s: %w", header[k], err)
				}
			}
			elemValue.Set(reflect.Append(elemValue, reflect.ValueOf(newItem)))
		}
		return nil
	}
}

// fillExcelFieldValue 填充字段值并做类型转换
func fillExcelFieldValue(field reflect.Value, value string, kind reflect.Kind) error {
	if value == "" {
		return nil
	}
	switch kind {
	case reflect.String:
		field.SetString(value)
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		field.SetInt(xtype.ToInt64(value))
	case reflect.Float32, reflect.Float64:
		field.SetFloat(xtype.ToFloat64(value))
	case reflect.Bool:
		field.SetBool(xtype.ToBool(value))
	default:
		// 兼容 time.Time
		if field.Type() == reflect.TypeOf(time.Time{}) {
			t := xtype.ToTime(value)
			if t.IsZero() {
				return fmt.Errorf("无效的时间格式: %s", value)
			}
			field.Set(reflect.ValueOf(t))
			return nil
		}
		return fmt.Errorf("unsupported type: %s", kind)
	}
	return nil
}

// WriteExcel ...
func WriteExcel(ctx context.Context, fpath string, list interface{}, forces ...bool) error {
	if list == nil {
		return fmt.Errorf("list cannot be nil")
	}
	switch list.(type) {
	case [][]interface{}:
		return writeExcel(ctx, fpath, list.([][]interface{}), forces...)
	default:
		listValue := reflect.ValueOf(list)
		if listValue.Kind() != reflect.Slice {
			return fmt.Errorf("only support []interface{}")
		}
		//if listValue.Len() == 0 {
		//	return fmt.Errorf("emtpy record")
		//}
		lines := make([][]interface{}, 0)
		header := make([]interface{}, 0)
		// 用于存储每列的颜色信息，key 为列索引，value 为颜色值
		colColors := make(map[int]string)
		// 用于存储每行每列的颜色信息
		rowColors := make([]map[int]string, 0)
		// 字段名到列索引的映射
		fieldColIdx := make(map[string]int)
		for i := 0; i < listValue.Len(); i++ {
			item := listValue.Index(i)
			if item.Kind() == reflect.Ptr {
				item = item.Elem()
			}
			line := make([]interface{}, 0)
			// 当前行的颜色
			rowColor := make(map[int]string)
			for j := 0; j < item.NumField(); j++ {
				fieldType := item.Type().Field(j)
				// 非导出数据
				if !fieldType.IsExported() {
					continue
				}
				// 检测 Colors 字段
				if fieldType.Name == "Colors" && fieldType.Type.Kind() == reflect.Map {
					colorMap := item.Field(j).Interface().(map[string]string)
					if len(colorMap) > 0 {
						// 遍历颜色 map，通过 fieldColIdx 直接找到列索引
						for fieldName, color := range colorMap {
							if colIdx, ok := fieldColIdx[fieldName]; ok {
								rowColor[colIdx] = color
							}
						}
					}
					continue
				}
				// 第一行时处理标题写入
				if i == 0 {
					var colIdx int
					if tag := fieldType.Tag.Get("excel"); tag != "" {
						colIdx = len(header)
						header = append(header, tag)
						fieldColIdx[fieldType.Name] = colIdx
					} else {
						colIdx = len(header)
						header = append(header, fieldType.Name)
						fieldColIdx[fieldType.Name] = colIdx
					}
				}
				line = append(line, item.Field(j).Interface())
			}
			lines = append(lines, line)
			rowColors = append(rowColors, rowColor)
		}
		lines = append([][]interface{}{header}, lines...)
		// 处理颜色
		if strings.HasSuffix(fpath, ".csv") {
			return writeExcel(ctx, fpath, lines, forces...)
		} else if strings.HasSuffix(fpath, ".xlsx") {
			return writeExcelWithColors(ctx, fpath, lines, colColors, rowColors, forces...)
		}
		return fmt.Errorf("invalid file ext")
	}
}

// writeExcelWithColors ...
func writeExcelWithColors(ctx context.Context, fpath string, list [][]interface{}, colColors map[int]string, rowColors []map[int]string, forces ...bool) error {
	if len(list) == 0 {
		return fmt.Errorf("emtpy record")
	}
	if !strings.HasSuffix(fpath, ".xlsx") {
		fpath += ".xlsx"
	}
	if !xlodash.First(forces) && IsExist(fpath) {
		return fmt.Errorf("file exist")
	}
	fdir := filepath.Dir(fpath)
	if !IsDir(fdir) {
		if err := os.MkdirAll(fdir, os.ModePerm); err != nil {
			return err
		}
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

		// 应用行颜色
		if k > 0 && k-1 < len(rowColors) {
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
