package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/evercyan/brick/xcli/xcolor"
	"github.com/evercyan/brick/xfile"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/yanyiwu/gojieba"
)

var g = newGla()

type glat struct {
	storageDir  string
	projectName string
	projectPath string
}

// newGla ...
func newGla() *glat {
	homeDir := xfile.GetHomeDir()
	storageDir := fmt.Sprintf("%s/.config/.glat", homeDir)
	if !xfile.IsExist(storageDir) {
		os.MkdirAll(storageDir, os.ModePerm)
	}
	return &glat{
		storageDir: storageDir,
	}
}

// init ...
func (g *glat) init() error {
	if g.projectName != "" {
		return nil
	}
	res := Exec(cmdGetProject)
	if res == "" {
		return errors.New("invalid git project: " + xfile.GetCurrentDir())
	}
	g.projectName = strings.TrimSpace(res)
	Log(SymbolSuccess, "projectName:", g.projectName)
	g.projectPath = fmt.Sprintf("%s/%s", g.storageDir, g.projectName)
	if !xfile.IsExist(g.projectPath) {
		os.Mkdir(g.projectPath, os.ModePerm)
	}
	return nil
}

// Html ...
func (g *glat) Html() {
	if err := g.init(); err != nil {
		Log(SymbolFail, err.Error())
		os.Exit(0)
	}
	logContent := Exec(cmdGetLog)
	if logContent == "" {
		Log(SymbolFail, "fetch git log fail")
		os.Exit(0)
	}
	logFile := fmt.Sprintf("%s/git.log", g.projectPath)
	if err := xfile.Write(logFile, logContent); err != nil {
		Log(SymbolFail, "save git log fail:", logFile)
		os.Exit(0)
	}
	Log(SymbolSuccess, "fetch git log success\n")

	x := gojieba.NewJieba()
	defer x.Free()
	page := components.NewPage()
	for _, info := range analysisItems {
		Log(SymbolBegin, "generate charter begin -", info["key"])
		res := Exec(fmt.Sprintf(info["cmd"], logFile))
		if res == "" {
			Log(SymbolFail, "parse log fail -", info["key"])
			continue
		}
		Log(SymbolSuccess, "parse log success -", info["key"])

		res = fmt.Sprintf("[\n%s]", res)
		if info["key"] == "word" {
			lines := make([]string, 0)
			if err := json.Unmarshal([]byte(res), &lines); err != nil {
				Log(SymbolFail, "parse word fail -", info["key"])
				continue
			}
			wordMap := make(map[string]int)
			for _, line := range lines {
				words := x.Cut(line, true)
				if len(words) == 0 {
					continue
				}
				for _, word := range words {
					if len(word) <= 1 {
						continue
					}
					if _, ok := wordMap[word]; !ok {
						wordMap[word] = 1
					} else {
						wordMap[word]++
					}
				}
			}
			wordList := make([]map[string]string, 0)
			for kk, vv := range wordMap {
				wordList = append(wordList, map[string]string{
					"k": kk,
					"v": fmt.Sprintf("%d", vv),
				})
			}
			b, _ := json.Marshal(wordList)
			res = string(b)
		}
		xfile.Write(fmt.Sprintf("%s/%s.json", g.projectPath, info["key"]), res)
		data := make([]map[string]string, 0)
		if err := json.Unmarshal([]byte(res), &data); err != nil {
			Log(SymbolFail, "parse data fail -", info["key"])
			continue
		}
		page.AddCharts(Charter(info, data))
		Log(SymbolSuccess, "generate charter success -", info["key"], "\n")
	}

	f := filepath.Join(os.TempDir(), g.projectName+".html")
	Log(SymbolSuccess, "generate all charter success\n\n")
	xcolor.Success("🍺🍺🍺 Click To Open:", f)
	fd, _ := os.Create(f)
	page.Render(io.MultiWriter(fd))
}
