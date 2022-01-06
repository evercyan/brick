package internal

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/evercyan/brick/xcli/xcolor"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// Log ...
func Log(symbol Symbol, args ...string) {
	prefix := time.Now().Format("2006-01-02 15:04:05")
	if symbol == SymbolSuccess {
		xcolor.Success(symbol.String(), prefix, strings.Join(args, " "))
	} else if symbol == SymbolFail {
		xcolor.Fail(symbol.String(), prefix, strings.Join(args, " "))
	} else {
		fmt.Println(symbol.String(), prefix, strings.Join(args, " "))
	}
}

// Exec ...
func Exec(cmd string) string {
	b, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return ""
	}
	return string(b)
}

// Charter ...
func Charter(info map[string]string, data []map[string]string) components.Charter {
	globalOptions := []charts.GlobalOpts{
		charts.WithTitleOpts(opts.Title{
			Title: info["title"],
			Left:  "0%",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "1200px",
			Height: "600px",
		}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:  true,
			Right: "0%",
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show:  true,
					Type:  "png",
					Title: "Save png",
				},
				DataView: &opts.ToolBoxFeatureDataView{
					Show:  true,
					Title: "Data view",
					Lang:  []string{"data view", "turn off", "refresh"},
				},
			}},
		),
	}

	keys := make([]string, 0)
	switch info["type"] {
	case "bar":
		barVals := make([]opts.BarData, 0)
		for _, v := range data {
			keys = append(keys, v["k"])
			barVals = append(barVals, opts.BarData{
				Value: v["v"],
			})
		}
		globalOptions = append(globalOptions, []charts.GlobalOpts{
			charts.WithXAxisOpts(opts.XAxis{
				Name: info["x"],
			}),
			charts.WithYAxisOpts(opts.YAxis{
				Name: info["y"],
			}),
		}...)
		bar := charts.NewBar()
		bar.SetGlobalOptions(globalOptions...)
		bar.SetXAxis(keys).
			AddSeries("commit", barVals).
			SetSeriesOptions(
				charts.WithLabelOpts(opts.Label{
					Show:     true,
					Position: "top",
				}),
			)
		return bar
	case "line":
		lineVals := make([]opts.LineData, 0)
		for _, v := range data {
			keys = append(keys, v["k"])
			lineVals = append(lineVals, opts.LineData{
				Value: v["v"],
			})
		}
		globalOptions = append(globalOptions, []charts.GlobalOpts{
			charts.WithXAxisOpts(opts.XAxis{
				Name: info["x"],
			}),
			charts.WithYAxisOpts(opts.YAxis{
				Name: info["y"],
			}),
		}...)
		line := charts.NewLine()
		line.SetGlobalOptions(globalOptions...)
		line.SetXAxis(keys).
			AddSeries("commit", lineVals).
			SetSeriesOptions(
				charts.WithLabelOpts(opts.Label{
					Show:     true,
					Position: "top",
				}),
				charts.WithLineChartOpts(
					opts.LineChart{
						Smooth: true,
					},
				),
			)
		return line
	case "pie":
		pieVals := make([]opts.PieData, 0)
		for _, v := range data {
			keys = append(keys, v["k"])
			pieVals = append(pieVals, opts.PieData{
				Name:  v["k"],
				Value: v["v"],
			})
		}
		pie := charts.NewPie()
		pie.SetGlobalOptions(globalOptions...)
		pie.AddSeries("commit", pieVals).
			SetSeriesOptions(
				charts.WithLabelOpts(opts.Label{
					Show:      true,
					Formatter: "{b}: {c}",
				}),
				charts.WithPieChartOpts(opts.PieChart{
					Radius: []string{"40%", "75%"},
				}),
			)
		return pie
	case "wordcloud":
		wcVals := make([]opts.WordCloudData, 0)
		for _, v := range data {
			keys = append(keys, v["k"])
			wcVals = append(wcVals, opts.WordCloudData{
				Name:  v["k"],
				Value: v["v"],
			})
		}
		wc := charts.NewWordCloud()
		wc.SetGlobalOptions(globalOptions...)
		wc.AddSeries("commit", wcVals).
			SetSeriesOptions(
				charts.WithWorldCloudChartOpts(opts.WordCloudChart{
					SizeRange: []float32{14, 80},
					Shape:     "cardioid",
				}),
			)
		return wc
	default:
		return nil
	}
}
