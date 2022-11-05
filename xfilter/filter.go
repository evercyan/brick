package xfilter

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/evercyan/brick/xencoding"
)

// ParseFilter 解析 filter 规则
func ParseFilter(rule string) ([]interface{}, error) {
	res := make([]interface{}, 0)
	if err := json.Unmarshal([]byte(rule), &res); err != nil {
		return nil, fmt.Errorf("filter rule invalid")
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("filter rule empty")
	}
	return res, nil
}

// Assert filter 规则断言
func Assert(ctx *Context, rule string) error {
	filters, err := ParseFilter(rule)
	if err != nil {
		return err
	}
	condition, err := NewCondition(filters)
	if err != nil {
		return err
	}
	return condition.Assert(ctx)
}

// Filter 列表数据过滤
func Filter(ctx *Context, list []map[string]interface{}) ([]map[string]interface{}, error) {
	if len(list) == 0 {
		return nil, errors.New("empty filter list")
	}
	res := make([]map[string]interface{}, 0)
	for _, v := range list {
		if _, ok := v["filter"]; !ok {
			res = append(res, v)
			continue
		}
		err := Assert(ctx, xencoding.JSONEncode(v["filter"]))
		if err != nil {
			continue
		}
		res = append(res, v)
	}
	return res, nil
}
