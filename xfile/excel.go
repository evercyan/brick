package xfile

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/evercyan/brick/xtype"
)

// readExcel ...
func readExcel(ctx context.Context, fpath string) ([][]string, error) {
	if strings.HasSuffix(fpath, ".csv") {
		return ReadCSV(ctx, fpath)
	} else if strings.HasSuffix(fpath, ".xlsx") {
		list, err := ReadXLSX(ctx, fpath)
		if err != nil {
			return nil, err
		}
		return xtype.Interface2string(list), nil
	}
	return nil, fmt.Errorf("无效的文件类型")
}

// writeExcel ...
func writeExcel(
	ctx context.Context,
	fpath string,
	list [][]interface{},
	rowColors ...map[int]string,
) error {
	if strings.HasSuffix(fpath, ".csv") {
		return WriteCSV(ctx, fpath, list)
	} else if strings.HasSuffix(fpath, ".xlsx") {
		return WriteXLSX(ctx, fpath, list, rowColors...)
	}
	return fmt.Errorf("无效的文件类型")
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
			return fmt.Errorf("未找到 excel 标签")
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

// WriteExcel ...
func WriteExcel(
	ctx context.Context,
	fpath string,
	list interface{},
	colors ...map[int]string,
) error {
	switch list.(type) {
	case [][]interface{}:
		return writeExcel(ctx, fpath, list.([][]interface{}), colors...)
	default:
		listValue := reflect.ValueOf(list)
		if listValue.Kind() != reflect.Slice {
			return fmt.Errorf("只支持 []interface{} 类型")
		}
		lines := make([][]interface{}, 0)
		header := make([]interface{}, 0)
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
					fieldName := fieldType.Name
					if tag := fieldType.Tag.Get("excel"); tag != "" {
						fieldName = tag
					}
					fieldColIdx[fieldType.Name] = len(header)
					header = append(header, fieldName)
				}
				line = append(line, item.Field(j).Interface())
			}
			lines = append(lines, line)
			rowColors = append(rowColors, rowColor)
		}
		lines = append([][]interface{}{header}, lines...)
		return writeExcel(ctx, fpath, lines, rowColors...)
	}
}
