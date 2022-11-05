package xfilter

import (
	"fmt"
	"strings"
	"time"

	"github.com/evercyan/brick/xtype"
)

// Variable 变量接口
type Variable interface {
	Name() string               // 变量名称
	Value(*Context) interface{} // 变量值
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------

// VariableCreator 变量生成器
type VariableCreator = func(string) Variable

// VariableFactory 变量实例工厂
type VariableFactory struct {
	creators map[string]VariableCreator
}

// Register 注册变量实例
func (t *VariableFactory) Register(name string, creator VariableCreator) {
	t.creators[name] = creator
}

// Discovery 发现变量实例
func (t *VariableFactory) Discovery(name string) Variable {
	if creator, ok := t.creators[name]; ok {
		return creator(name)
	}
	segments := strings.Split(name, ".")
	if len(segments) > 1 {
		if creator, ok := t.creators[segments[0]+"."]; ok {
			return creator(name)
		}
	}
	return nil
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------

// VariableCtx ...
type VariableCtx struct {
	name string
	key  string
}

// Name ...
func (t *VariableCtx) Name() string {
	return t.name
}

// Value ...
func (t *VariableCtx) Value(ctx *Context) interface{} {
	// 读取 ctx 数据
	if v, ok := ctx.Get(t.key); ok {
		return v
	}
	// 读取 ctx.ctx 数据
	if v := ctx.Ctx(t.key); v != nil {
		return v
	}
	// key 存在 . 分割, 取第一级进行匹配
	segments := strings.Split(t.key, ".")
	if len(segments) > 1 {
		if v, ok := ctx.Get(segments[0]); ok {
			// 解析数据结构
			path := strings.TrimPrefix(t.key, segments[0]+".")
			return xtype.Parse(v, path)
		}
	}
	return nil
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------

// VariableTime ...
type VariableTime struct {
	name string
	key  string
}

// Name ...
func (t *VariableTime) Name() string {
	return t.name
}

// Value ...
func (t *VariableTime) Value(ctx *Context) interface{} {
	now := time.Now()
	switch t.key {
	case "year":
		return now.Year()
	case "month":
		return int(now.Month())
	case "day":
		return now.Day()
	case "hour":
		return now.Hour()
	case "minute":
		return now.Minute()
	case "second":
		return now.Second()
	case "wday":
		wday := int(now.Weekday())
		if wday == 0 {
			wday = 7
		}
		return wday
	case "date":
		return now.Format("2006-01-02")
	case "time":
		return now.Format("15:04:05")
	case "unixtime":
		return now.Unix()
	case "datetime":
		fallthrough
	default:
		return now.Format("2006-01-02 15:04:05")
	}
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------

// NewVariable ...
func NewVariable(name string) (Variable, error) {
	variable := _variableFactory.Discovery(name)
	if variable == nil {
		return nil, fmt.Errorf("variable unknown [%s]", name)
	}
	return variable, nil
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------

var _variableFactory *VariableFactory

func init() {
	_variableFactory = &VariableFactory{
		creators: make(map[string]VariableCreator),
	}
	_variableFactory.Register("ctx.", func(name string) Variable {
		key := strings.TrimPrefix(name, "ctx.")
		if key == "" {
			return nil
		}
		return &VariableCtx{
			name: name,
			key:  key,
		}
	})
	variableTimes := []string{
		"year",
		"month",
		"day",
		"hour",
		"minute",
		"second",
		"wday",
		"date",
		"time",
		"unixtime",
		"datetime",
	}
	for _, name := range variableTimes {
		_variableFactory.Register(name, func(name string) Variable {
			return &VariableTime{
				name: name,
				key:  name,
			}
		})
	}
}
