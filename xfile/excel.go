package xfile

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/evercyan/brick/xlodash"
	"github.com/evercyan/brick/xtype"
)

// ----------------------------------------------------------------

// readExcel ...
func readExcel(ctx context.Context, fpath string) ([][]string, error) {
	if strings.HasSuffix(fpath, ".csv") {
		return ReadCSV(ctx, fpath)
	} else if strings.HasSuffix(fpath, ".xlsx") {
		return ReadXLSX(ctx, fpath)
	}
	return nil, fmt.Errorf("invalid file ext")
}

// writeExcel ...
func writeExcel(ctx context.Context, fpath string, list [][]interface{}, forces ...bool) error {
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
func ReadXLSX(ctx context.Context, fpath string, sheets ...string) ([][]string, error) {
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
		f.SetSheetRow(sheet1, fmt.Sprintf("A%d", k+1), &v)
	}
	return f.SaveAs(fpath)
}

// ----------------------------------------------------------------

// WriteExcel ...
var WriteExcel = writeExcel

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
		val := reflect.ValueOf(list)
		if val.Kind() != reflect.Ptr ||
			val.Elem().Kind() != reflect.Slice ||
			val.Elem().Type().Elem().Elem().Kind() != reflect.Struct {
			return fmt.Errorf("only support []*struct")
		}
		var (
			elemType  = val.Elem().Type().Elem().Elem()
			elemValue = val.Elem()
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
		return fmt.Errorf("unsupported type: %s", kind)
	}
	return nil
}

// ----------------------------------------------------------------
