package command

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/evercyan/brick/cmd/leet/config"
	"github.com/evercyan/brick/cmd/leet/internal"
	"github.com/evercyan/brick/xcli/xcolor"
	"github.com/evercyan/brick/xconvert"
	"github.com/evercyan/brick/xutil"
	"github.com/peterh/liner"
	"github.com/spf13/cobra"
)

var (
	// QuestionCommand ...
	QuestionCommand = &cobra.Command{
		Use:     "question",
		Aliases: []string{"q"},
		Short:   "生成答题文件, e.g. leet question two-sum",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				xcolor.Fail(config.SymbolError, "请输入题目标识或ID")
				return
			}
			defer func(begin time.Time) {
				xcolor.Success(
					config.SymbolTime,
					fmt.Sprintf("耗时: %s", time.Now().Sub(begin).String()),
				)
			}(time.Now())
			app := internal.NewApp()
			// 校验题目
			list, err := app.GetQuestionList()
			if err != nil {
				xcolor.Fail(config.SymbolError, err.Error())
				return
			}
			question := &config.Question{}
			text, num := args[0], int64(xconvert.ToUint(args[0]))
			for _, v := range list {
				if v.Slug == text || v.Qid == num {
					question = v
					break
				}
			}
			if question.Qid == 0 {
				xcolor.Fail(config.SymbolError, "未匹配到题目")
				return
			}
			xcolor.Success(
				config.SymbolNotice,
				fmt.Sprintf("匹配题目: %d. %s (%s)", question.Qid, question.Title, question.Link),
			)
			// 题目详情
			detail, err := app.GetQuestionDetail(question.Slug)
			if err != nil {
				xcolor.Fail(config.SymbolError, err.Error())
				return
			}
			question.Detail = detail
			// 校验语言
			lang := ""
			if app.Lang != "" && xutil.IsContains(detail.LangList, app.Lang) {
				lang = app.Lang
			} else {
				xcolor.Success(
					config.SymbolNotice,
					fmt.Sprintf("支持的编程语言是: [%s]", strings.Join(detail.LangList, ", ")),
				)
				// 监听输入, 自动补全语言
				line := liner.NewLiner()
				defer line.Close()
				line.SetCtrlCAborts(true)
				line.SetCompleter(func(line string) []string {
					list := make([]string, 0)
					for _, v := range detail.LangList {
						if strings.HasPrefix(v, strings.ToLower(line)) {
							list = append(list, v)
						}
					}
					return list
				})
				text, err := line.Prompt(config.SymbolNotice + " 输入编程语言: ")
				if err == liner.ErrPromptAborted {
					os.Exit(0)
				}
				if !xutil.IsContains(detail.LangList, text) {
					xcolor.Fail(config.SymbolError, "无效的编程语言")
					return
				}
				lang = text
			}
			xcolor.Success(config.SymbolNotice, fmt.Sprintf("使用的编程语言是: %s", lang))
			// 生成答题文件相关
			if err := app.GenerateQuestion(question, app.Path, lang); err != nil {
				xcolor.Fail(config.SymbolError, err.Error())
				return
			}
			xcolor.Success(config.SymbolSuccess, "生成答题文件成功")
		},
	}
)
