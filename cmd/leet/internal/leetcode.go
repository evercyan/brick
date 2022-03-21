package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/evercyan/brick/xencoding"

	"github.com/evercyan/brick/cmd/leet/config"
	"github.com/evercyan/brick/xfile"
	"github.com/evercyan/brick/xhttp"
	"github.com/evercyan/brick/xjson"
)

// LeetCode ...
type LeetCode struct {
	Ctx context.Context
}

func newLeetCode() *LeetCode {
	return &LeetCode{
		Ctx: context.Background(),
	}
}

// ----------------------------------------------------------------

// GetQuestionList ...
func (t *LeetCode) GetQuestionList() ([]*config.Question, error) {
	res := xhttp.Get(t.Ctx, config.LeetCodeAllURL)
	content := xjson.New(res).Key("stat_status_pairs").ToJSON()
	if content == "" {
		return nil, fmt.Errorf("获取 LeetCode 题目列表失败")
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
		return nil, fmt.Errorf("解析题目信息失败")
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
	return list, nil
}

// GetQuestionDetail ...
func (t *LeetCode) GetQuestionDetail(slug string) (*config.QuestionDetail, error) {
	res := xhttp.Post(t.Ctx, config.LeetCodeGraphqlURL, map[string]interface{}{
		"operationName": "questionData",
		"query":         "query questionData($titleSlug: String!) {question(titleSlug: $titleSlug) {questionId questionFrontendId title titleSlug content translatedTitle translatedContent topicTags {name slug translatedName} codeSnippets {lang langSlug code}}}",
		"variables": map[string]string{
			"titleSlug": slug,
		},
	})
	content := xjson.New(res).Key("data").Key("question").ToJSON()
	if content == "" {
		return nil, fmt.Errorf("获取 LeetCode 题目详情失败")
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
		return nil, fmt.Errorf("解析题目详情失败")
	}
	detail := &config.QuestionDetail{
		Title:       originDetail.TranslatedTitle,
		Content:     FormatQuestionContent(originDetail.TranslatedContent),
		TagList:     make([]config.Tag, 0),
		TagSlugList: make([]string, 0),
		LangList:    make([]string, 0),
		CodeMap:     make(map[string]string),
	}
	for _, v := range originDetail.TopicTags {
		tagName := v.TranslatedName
		if tagName == "" {
			tagName = v.Name
		}
		detail.TagList = append(detail.TagList, config.Tag{
			Name: tagName,
			Slug: v.Slug,
		})
		detail.TagSlugList = append(detail.TagSlugList, v.Slug)
	}
	for _, v := range originDetail.CodeSnippets {
		lang := strings.ToLower(v.LangSlug)
		detail.LangList = append(detail.LangList, lang)
		detail.CodeMap[lang] = v.Code
	}
	return detail, nil
}

// GetTagList ...
func (t *LeetCode) GetTagList() ([]*config.Tag, error) {
	res := xhttp.Post(t.Ctx, config.LeetCodeGraphqlURL, map[string]interface{}{
		"query":     "query questionTagTypeWithTags {questionTagTypeWithTags {name transName tagRelation {questionNum tag {name id nameTranslated slug}}}}",
		"variables": map[string]string{},
	})
	content := xjson.New(res).Key("data").Key("questionTagTypeWithTags").ToJSON()
	if content == "" {
		return nil, fmt.Errorf("获取 LeetCode 标签列表失败")
	}
	var tagList []struct {
		Name        string `json:"name"`
		TransName   string `json:"transName"`
		TagRelation []struct {
			QuestionNum int `json:"questionNum"`
			Tag         struct {
				Name           string `json:"name"`
				Id             string `json:"id"`
				NameTranslated string `json:"nameTranslated"`
				Slug           string `json:"slug"`
			} `json:"tag"`
		} `json:"tagRelation"`
	}
	if err := json.Unmarshal([]byte(content), &tagList); err != nil {
		return nil, fmt.Errorf("解析标签列表失败")
	}
	list := make([]*config.Tag, 0)
	for _, v := range tagList {
		for _, vv := range v.TagRelation {
			tagName := vv.Tag.NameTranslated
			if tagName == "" {
				tagName = vv.Tag.Name
			}
			list = append(list, &config.Tag{
				Name:  tagName,
				Slug:  vv.Tag.Slug,
				Count: vv.QuestionNum,
			})
		}
	}
	return list, nil
}

// GetTagQuestionMap ...
func (t *LeetCode) GetTagQuestionMap(app *App) (map[string][]*config.Question, error) {
	res := xhttp.Post(t.Ctx, config.LeetCodeGraphqlURL, map[string]interface{}{
		"operationName": "allQuestionUrls",
		"query":         "query allQuestionUrls{allQuestionUrls{questionUrl __typename}}",
		"variables":     map[string]string{},
	})
	content := xjson.New(res).Key("data").Key("allQuestionUrls").Key("questionUrl").ToString()
	if content == "" {
		return nil, fmt.Errorf("获取 LeetCode 问题列表失败")
	}
	questionContent := ""
	if content == app.QuestionFile {
		// 文件无变化
		questionContent = xfile.Read(GetQuestionFilepath())
	} else {
		// 获取文件
		jsonContent := xhttp.Get(t.Ctx, content)
		if jsonContent == "" {
			return nil, fmt.Errorf("获取 LeetCode 问题数据失败")
		}
		tmp := make([]config.QuestionTag, 0)
		if err := json.Unmarshal([]byte(jsonContent), &tmp); err != nil {
			return nil, fmt.Errorf("问题数据解析失败")
		}
		questionContent = xencoding.JSONEncode(tmp)
		// 更新配置文件 & 保存 JSON 文件
		app.QuestionFile = content
		app.Update()
		xfile.Write(GetQuestionFilepath(), questionContent)
	}
	qList := make([]config.QuestionTag, 0)
	if err := json.Unmarshal([]byte(questionContent), &qList); err != nil {
		return nil, fmt.Errorf("问题数据解析失败")
	}
	// 问题列表
	questionList, err := t.GetQuestionList()
	if err != nil {
		return nil, err
	}
	questionMap := make(map[string]*config.Question)
	for _, v := range questionList {
		questionMap[v.Slug] = v
	}
	// 拼组数据
	tagMap := make(map[string][]*config.Question)
	for _, v := range qList {
		if _, ok := questionMap[v.Slug]; !ok {
			continue
		}
		for _, vv := range v.TopicTags {
			if _, ok := tagMap[vv.Slug]; !ok {
				tagMap[vv.Slug] = make([]*config.Question, 0)
			}
			tagMap[vv.Slug] = append(tagMap[vv.Slug], questionMap[v.Slug])
		}
	}
	return tagMap, nil
}
