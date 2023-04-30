package xjson

import (
	"encoding/json"

	"github.com/evercyan/brick/xtype"
)

// JSON ...
type JSON struct {
	value interface{}
}

// New ...
func New(s string) *JSON {
	j := &JSON{}
	if err := json.Unmarshal([]byte(s), &j.value); err != nil {
		j.value = nil
	}
	return j
}

// Key ...
func (j *JSON) Key(key string) *JSON {
	if m, ok := j.value.(map[string]interface{}); ok {
		if v, ok := m[key]; ok {
			j.value = v
			return j
		}
	}
	j.value = nil
	return j
}

// Index ...
func (j *JSON) Index(index int) *JSON {
	if l, ok := j.value.([]interface{}); ok && index >= 0 && index < len(l) {
		j.value = l[index]
		return j
	}
	j.value = nil
	return j
}

// Value ...
func (j *JSON) Value() interface{} {
	return j.value
}

// ToString ...
func (j *JSON) ToString() string {
	return xtype.ToString(j.value)
}

// ToInt ...
func (j *JSON) ToInt64() int64 {
	return xtype.ToInt64(j.value)
}

// ToJSON ...
func (j *JSON) ToJSON() string {
	if j.value == nil {
		return ""
	}
	b, _ := json.Marshal(j.value)
	return string(b)
}

// ToSlice ...
func (j *JSON) ToSlice() []interface{} {
	switch (j.value).(type) {
	case []interface{}:
		return (j.value).([]interface{})
	default:
		return nil
	}
}

// ToMap ...
func (j *JSON) ToMap() map[string]interface{} {
	switch (j.value).(type) {
	case map[string]interface{}:
		return (j.value).(map[string]interface{})
	default:
		return nil
	}
}
