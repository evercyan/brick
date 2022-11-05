package xfilter

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/evercyan/brick/xtype"
	"github.com/evercyan/brick/xutil"
)

// Operation 操作接口
type Operation interface {
	Name() string                                                   // 操作描述
	Expect(value interface{}) (interface{}, error)                  // 操作预期值
	Assert(ctx *Context, variable Variable, value interface{}) bool // 操作断言
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------

// OperationFactory 操作实例工厂
type OperationFactory struct {
	operations map[string]Operation
}

// Register 注册操作实例
func (t *OperationFactory) Register(operation Operation) {
	t.operations[operation.Name()] = operation
}

// Discovery 发现操作实例
func (t *OperationFactory) Discovery(name string) Operation {
	if operation, ok := t.operations[name]; ok {
		return operation
	}
	return nil
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------

// OperationEqual ...
type OperationEqual struct{}

// Name ...
func (t *OperationEqual) Name() string {
	return "="
}

// Expect ...
func (t *OperationEqual) Expect(value interface{}) (interface{}, error) {
	return value, nil
}

// Assert ...
func (t *OperationEqual) Assert(ctx *Context, variable Variable, expect interface{}) bool {
	return xutil.Compare(variable.Value(ctx), expect) == 0
}

// ----------------------------------------------------------------

// OperationEqualNot ...
type OperationEqualNot struct {
	OperationEqual
}

// Name ...
func (t *OperationEqualNot) Name() string {
	return "!="
}

// Assert ...
func (t *OperationEqualNot) Assert(ctx *Context, variable Variable, expect interface{}) bool {
	return xutil.Compare(variable.Value(ctx), expect) != 0
}

// ----------------------------------------------------------------

// OperationEqualGT ...
type OperationEqualGT struct {
	OperationEqual
}

// Name ...
func (t *OperationEqualGT) Name() string {
	return ">"
}

// Assert ...
func (t *OperationEqualGT) Assert(ctx *Context, variable Variable, expect interface{}) bool {
	return xutil.Compare(variable.Value(ctx), expect) > 0
}

// ----------------------------------------------------------------

// OperationEqualGTE ...
type OperationEqualGTE struct {
	OperationEqual
}

// Name ...
func (t *OperationEqualGTE) Name() string {
	return ">="
}

// Assert ...
func (t *OperationEqualGTE) Assert(ctx *Context, variable Variable, expect interface{}) bool {
	return xutil.Compare(variable.Value(ctx), expect) >= 0
}

// ----------------------------------------------------------------

// OperationEqualLT ...
type OperationEqualLT struct {
	OperationEqual
}

// Name ...
func (t *OperationEqualLT) Name() string {
	return "<"
}

// Assert ...
func (t *OperationEqualLT) Assert(ctx *Context, variable Variable, expect interface{}) bool {
	return xutil.Compare(variable.Value(ctx), expect) < 0
}

// ----------------------------------------------------------------

// OperationEqualLTE ...
type OperationEqualLTE struct {
	OperationEqual
}

// Name ...
func (t *OperationEqualLTE) Name() string {
	return "<="
}

// Assert ...
func (t *OperationEqualLTE) Assert(ctx *Context, variable Variable, expect interface{}) bool {
	return xutil.Compare(variable.Value(ctx), expect) <= 0
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------

// OperationBetween ...
type OperationBetween struct{}

// Name ...
func (t *OperationBetween) Name() string {
	return "between"
}

// Expect ...
func (t *OperationBetween) Expect(value interface{}) (interface{}, error) {
	section := xtype.ToSlice(value)
	if len(section) != 2 {
		return nil, fmt.Errorf("operation [%s] value must be a list with 2 elements", t.Name())
	}
	return section, nil
}

// Assert ...
func (t *OperationBetween) Assert(ctx *Context, variable Variable, expect interface{}) bool {
	value := variable.Value(ctx)
	list := expect.([]interface{})
	return xutil.Compare(value, list[0]) >= 0 && xutil.Compare(value, list[1]) <= 0
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------

// OperationHas ...
type OperationHas struct{}

// Name ...
func (t *OperationHas) Name() string {
	return "has"
}

// Expect ...
func (t *OperationHas) Expect(value interface{}) (interface{}, error) {
	list := xtype.ToSlice(value)
	if len(list) == 0 {
		return nil, fmt.Errorf("operation [%s] value must be a list", t.Name())
	}
	return list, nil
}

// Assert ...
func (t *OperationHas) Assert(ctx *Context, variable Variable, expect interface{}) bool {
	vList := xtype.ToSlice(variable.Value(ctx))
	eList := xtype.ToSlice(expect)
	if len(vList) == 0 || len(eList) == 0 {
		return false
	}
	for _, ev := range eList {
		ok := false
		for _, vv := range vList {
			if xutil.Compare(ev, vv) == 0 {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}
	return true
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------

// OperationIn ...
type OperationIn struct{}

// Name ...
func (t *OperationIn) Name() string {
	return "in"
}

// Expect ...
func (t *OperationIn) Expect(value interface{}) (interface{}, error) {
	section := xtype.ToSlice(value)
	if len(section) == 0 {
		return nil, fmt.Errorf("operation [%s] value must be a list", t.Name())
	}
	return section, nil
}

// Assert ...
func (t *OperationIn) Assert(ctx *Context, variable Variable, expect interface{}) bool {
	value := variable.Value(ctx)
	for _, v := range expect.([]interface{}) {
		if xutil.Compare(v, value) == 0 {
			return true
		}
	}
	return false
}

// ----------------------------------------------------------------

// OperationNotIn ...
type OperationNotIn struct {
	OperationIn
}

// Name ...
func (t *OperationNotIn) Name() string {
	return "not in"
}

// Assert ...
func (t *OperationNotIn) Assert(ctx *Context, variable Variable, expect interface{}) bool {
	return !t.OperationIn.Assert(ctx, variable, expect)
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------

// OperationMatch ...
type OperationMatch struct{}

// Name ...
func (t *OperationMatch) Name() string {
	return "match"
}

// Expect ...
func (t *OperationMatch) Expect(value interface{}) (interface{}, error) {
	s, ok := value.(string)
	if !ok || s == "" {
		return nil, fmt.Errorf("operation [%s] value must be a string", t.Name())
	}
	// 非 /xxx/ 格式的, 只计算是否 contains
	if !strings.HasPrefix(s, "/") || !strings.HasSuffix(s, "/") {
		return strings.ToLower(s), nil
	}
	s = strings.Trim(s, "/")
	if s == "" {
		return nil, fmt.Errorf("operation [%s] value is not a valid regexp expression", t.Name())
	}
	reg, err := regexp.Compile("(?i)" + s)
	if err != nil {
		return nil, fmt.Errorf("operation [%s] value is not a valid regexp expression: %s", t.Name(), err.Error())
	}
	return reg, nil
}

// Assert ...
func (t *OperationMatch) Assert(ctx *Context, variable Variable, expect interface{}) bool {
	value := variable.Value(ctx)
	valueStr, ok := value.(string)
	if !ok {
		return false
	}
	if reg, ok1 := expect.(*regexp.Regexp); ok1 {
		return reg.MatchString(valueStr)
	} else if str, ok2 := expect.(string); ok2 {
		return strings.Contains(strings.ToLower(valueStr), str)
	}
	return false
}

// ----------------------------------------------------------------

// OperationNotMatch ...
type OperationNotMatch struct {
	OperationMatch
}

// Name ...
func (t *OperationNotMatch) Name() string {
	return "not match"
}

// Assert ...
func (t *OperationNotMatch) Assert(ctx *Context, variable Variable, expect interface{}) bool {
	return !t.OperationMatch.Assert(ctx, variable, expect)
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------

// NewOperation ...
func NewOperation(operationName string) (Operation, error) {
	operation := _operationFactory.Discovery(operationName)
	if operation == nil {
		return nil, fmt.Errorf("operation unknown [%s]", operationName)
	}
	return operation, nil
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------

var _operationFactory *OperationFactory

func init() {
	_operationFactory = &OperationFactory{
		operations: make(map[string]Operation),
	}
	_operationFactory.Register(&OperationEqual{})    // =
	_operationFactory.Register(&OperationEqualNot{}) // !=
	_operationFactory.Register(&OperationEqualGT{})  // >
	_operationFactory.Register(&OperationEqualGTE{}) // >=
	_operationFactory.Register(&OperationEqualLT{})  // <
	_operationFactory.Register(&OperationEqualLTE{}) // <=
	_operationFactory.Register(&OperationBetween{})  // between
	_operationFactory.Register(&OperationIn{})       // in
	_operationFactory.Register(&OperationNotIn{})    // not in
	_operationFactory.Register(&OperationMatch{})    // match
	_operationFactory.Register(&OperationNotMatch{}) // not match
	_operationFactory.Register(&OperationHas{})      // has
}
