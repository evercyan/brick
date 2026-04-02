package xfile

import (
	"context"
	"encoding/csv"
	"os"
	"strings"

	"github.com/evercyan/brick/xtype"
)

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
func WriteCSV(ctx context.Context, fpath string, list [][]interface{}) error {
	if !strings.HasSuffix(fpath, ".csv") {
		fpath += ".csv"
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
