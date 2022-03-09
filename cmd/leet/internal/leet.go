package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/evercyan/brick/cmd/leet/config"
	"github.com/evercyan/brick/xcli/xtable"
	"github.com/evercyan/brick/xconvert"
	"github.com/evercyan/brick/xhttp"
	"github.com/evercyan/brick/xjson"
)

type Leet struct {
	Ctx  context.Context
	List []*config.Question
}

func newLeet() *Leet {
	return &Leet{
		Ctx:  context.Background(),
		List: make([]*config.Question, 0),
	}
}

// ----------------------------------------------------------------

func (t *Leet) GetList(keywords ...string) (string, error) {
	res := xhttp.Get(t.Ctx, config.LeetCodeAllURL)
	content := xjson.New(res).Key("stat_status_pairs").ToJSON()
	if content == "" {
		return "", fmt.Errorf("获取 LeetCode 问题列表失败")
	}
	originList := make([]struct {
		Stat struct {
			QuestionId          int64  `json:"question_id"`
			QuestionTitle       string `json:"question__title"`
			QuestionTitleSlug   string `json:"question__title_slug"`
			QuestionHide        bool   `json:"question__hide"`
			TotalAcs            int    `json:"total_acs"`
			TotalSubmitted      int    `json:"total_submitted"`
			TotalColumnArticles int    `json:"total_column_articles"`
			FrontendQuestionId  string `json:"frontend_question_id"`
			IsNewQuestion       bool   `json:"is_new_question"`
		} `json:"stat"`
		Status     interface{} `json:"status"`
		Difficulty struct {
			Level config.QuestionLevel `json:"level"`
		} `json:"difficulty"`
		PaidOnly  bool `json:"paid_only"`
		IsFavor   bool `json:"is_favor"`
		Frequency int  `json:"frequency"`
		Progress  int  `json:"progress"`
	}, 0)
	if err := json.Unmarshal([]byte(content), &originList); err != nil {
		return "", fmt.Errorf("解析问题信息失败")
	}

	list := make([]*config.Question, 0)
	for _, v := range originList {
		list = append([]*config.Question{{
			Fid:   v.Stat.FrontendQuestionId,
			Qid:   v.Stat.QuestionId,
			Title: v.Stat.QuestionTitle,
			Slug:  v.Stat.QuestionTitleSlug,
			Link:  GetQuestionLink(v.Stat.QuestionTitleSlug),
			Level: v.Difficulty.Level,
		}}, list...)
	}
	t.List = list

	matchs := make([][]interface{}, 0)
	text, num := "", int64(0)
	if len(keywords) > 0 {
		text, num = keywords[0], int64(xconvert.ToUint(keywords[0]))
	}
	for _, v := range list {
		if strings.Contains(v.Fid, text) ||
			strings.Contains(v.Title, text) ||
			strings.Contains(v.Slug, text) ||
			v.Qid == num {
			matchs = append(matchs, []interface{}{
				v.Fid,
				v.Title,
				v.Slug,
				v.Level.String(),
			})
		}
	}
	return xtable.New(matchs).Style(xtable.Dashed).Border(true).Header([]string{
		"ID", "标题", "标识", "难度",
	}).Text(), nil
}

// GetDetail ...
func (t *Leet) GetDetail(slug string) (*config.QuestionDetail, error) {
	res := xhttp.Post(t.Ctx, config.LeetCodeGraphqlURL, map[string]interface{}{
		"operationName": "questionData",
		"query":         "query questionData($titleSlug: String!) {question(titleSlug: $titleSlug) {questionId questionFrontendId title titleSlug content translatedTitle translatedContent topicTags {name slug translatedName} codeSnippets {lang langSlug code}}}",
		"variables": map[string]string{
			"titleSlug": slug,
		},
	})
	content := xjson.New(res).Key("data").Key("question").ToJSON()
	if content == "" {
		return nil, fmt.Errorf("获取 LeetCode 问题详情失败")
	}
	originDetail := struct {
		QuestionId         string `json:"questionId"`
		QuestionFrontendId string `json:"questionFrontendId"`
		Title              string `json:"title"`
		TitleSlug          string `json:"titleSlug"`
		Content            string `json:"content"`
		TranslatedTitle    string `json:"translatedTitle"`
		TranslatedContent  string `json:"translatedContent"`
		TopicTags          []struct {
			Name           string `json:"name"`
			Slug           string `json:"slug"`
			TranslatedName string `json:"translatedName"`
		} `json:"topicTags"`
		CodeSnippets []struct {
			Lang     string `json:"lang"`
			LangSlug string `json:"langSlug"`
			Code     string `json:"code"`
		} `json:"codeSnippets"`
	}{}
	if err := json.Unmarshal([]byte(content), &originDetail); err != nil {
		return nil, fmt.Errorf("解析问题详情失败")
	}
	detail := &config.QuestionDetail{
		Title:    originDetail.TranslatedTitle,
		Content:  FormatQuestionContent(originDetail.TranslatedContent),
		TagList:  make([]map[string]string, 0),
		LangList: make([]string, 0),
		CodeMap:  make(map[string]string),
	}
	for _, v := range originDetail.TopicTags {
		detail.TagList = append(detail.TagList, map[string]string{
			"slug": v.Slug,
			"name": v.TranslatedName,
		})
	}
	for _, v := range originDetail.CodeSnippets {
		lang := strings.ToLower(v.LangSlug)
		detail.LangList = append(detail.LangList, lang)
		detail.CodeMap[lang] = v.Code
	}
	return detail, nil
}
