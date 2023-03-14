package xfilter

import (
	"fmt"
	"strings"

	"github.com/evercyan/brick/xjson"
	"github.com/evercyan/brick/xtype"
)

// Condition 断言条件接口
type Condition interface {
	Name() string          // 条件描述
	Assert(*Context) error // 条件断言
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------

// ConditionSingle 单个条件
type ConditionSingle struct {
	Variable  Variable    // 变量
	Operation Operation   // 操作符
	Expect    interface{} // 预期结果
}

// Name ...
func (t *ConditionSingle) Name() string {
	return fmt.Sprintf("%s %s %v", t.Variable.Name(), t.Operation.Name(), t.Expect)
}

// Assert ...
func (t *ConditionSingle) Assert(ctx *Context) error {
	ok := t.Operation.Assert(ctx, t.Variable, t.Expect)
	if !ok {
		return fmt.Errorf("condition [%s] fail", t.Name())
	}
	return nil
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------

// ConditionGroup 群组条件, 包含 Logic 关系
type ConditionGroup struct {
	Logic      Logic
	Conditions []Condition
}

// Name ...
func (t *ConditionGroup) Name() string {
	res := make([]string, 0)
	for _, condition := range t.Conditions {
		res = append(res, condition.Name())
	}
	return fmt.Sprintf(`["%s", "%s"]`, strings.Join(res, `", "`), t.Logic.String())
}

// Assert ...
func (t *ConditionGroup) Assert(ctx *Context) (err error) {
	// 群组条件需关注条件之间的逻辑关联
	defer func() {
		// 返回错误带上条件组名称描述
		if err != nil {
			err = fmt.Errorf("group %s, %s", t.Name(), err.Error())
		}
	}()
	for _, condition := range t.Conditions {
		err = condition.Assert(ctx)
		if err != nil {
			if t.Logic == LogicAnd {
				return
			} else if t.Logic == LogicOr {
				continue
			}
		} else {
			if t.Logic == LogicAnd {
				continue
			} else if t.Logic == LogicOr {
				return
			}
		}
	}
	return
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------

// NewConditionSingle 初始化单个条件实例
func NewConditionSingle(filter []interface{}) (*ConditionSingle, error) {
	prefix := xjson.Encode(filter)
	if len(filter) != 3 {
		return nil, fmt.Errorf("%s: 条件必须有 3 个元素", prefix)
	}
	// 解析变量
	variableName, ok := filter[0].(string)
	if !ok {
		return nil, fmt.Errorf("%s: 条件的第 1 个元素必须是字符串: %v", prefix, filter[0])
	}
	variable, err := NewVariable(variableName)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", prefix, err.Error())
	}
	// 解析操作符
	operationName, ok := filter[1].(string)
	if !ok {
		return nil, fmt.Errorf("%s: 条件的第 2 个元素必须是字符串: %v", prefix, filter[1])
	}
	operation, err := NewOperation(operationName)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", prefix, err.Error())
	}
	// 解析预期值
	expect, err := operation.Expect(filter[2])
	if err != nil {
		return nil, fmt.Errorf("%s: %s", prefix, err.Error())
	}
	return &ConditionSingle{
		Variable:  variable,
		Operation: operation,
		Expect:    expect,
	}, nil
}

// NewConditionGroup 初始化群组条件实例
func NewConditionGroup(filters []interface{}) (Condition, error) {
	prefix := xjson.Encode(filters)
	if len(filters) == 0 {
		return nil, fmt.Errorf("%s: 条件组至少要有 1 个条件", prefix)
	}

	// 条件组间逻辑默认为 and
	logic := LogicAnd
	if s, ok := filters[len(filters)-1].(string); ok {
		// 如果最后一位元素是字符串
		logic = ToLogic(s)
		filters = filters[:len(filters)-1]
		if len(filters) == 0 {
			return nil, fmt.Errorf("%s: 条件组至少要有 1 个条件", prefix)
		}
	}

	conditionGroup := &ConditionGroup{
		Logic:      logic,
		Conditions: make([]Condition, 0),
	}
	for _, filter := range filters {
		if !xtype.IsSlice(filter) {
			return nil, fmt.Errorf(
				"条件组的元素必须是一个数组: %v", xjson.Encode(filter),
			)
		}
		subCondition, err := NewCondition(filter.([]interface{}))
		if err != nil {
			return nil, err
		}
		conditionGroup.Conditions = append(conditionGroup.Conditions, subCondition)
	}
	return conditionGroup, nil
}

// NewCondition 初始化复合条件
func NewCondition(filters []interface{}) (Condition, error) {
	conditionSingle, errSingle := NewConditionSingle(filters)
	if errSingle == nil {
		return conditionSingle, nil
	}
	conditionGroup, errGroup := NewConditionGroup(filters)
	if errGroup == nil {
		return conditionGroup, nil
	}
	return nil, fmt.Errorf(
		"解析失败: (条件: %v) (条件组: %v)",
		errSingle.Error(),
		errGroup.Error(),
	)
}
