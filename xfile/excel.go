package xfile

import (
	"context"
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
func ToXLSX(ctx context.Context, fpath string, list [][]string) error {
	if len(list) == 0 {
		return nil
	}
	if !strings.HasSuffix(fpath, ".xlsx") {
		fpath += ".xlsx"
	}
	if IsExist(fpath) {
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
