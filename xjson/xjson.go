package xjson

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// JSON ...
type JSON struct {
	value interface{}
}

// New ...
func New(s string) *JSON {
	j := new(JSON)
	var v interface{}
	if err := json.Unmarshal([]byte(s), &v); err == nil {
		j.value = v
	}
	return j
}

// Key ...
func (j *JSON) Key(key string) *JSON {
	m, ok := (j.value).(map[string]interface{})
	if ok {
		j.value = m[key]
	} else {
		j.value = nil
	}
	return j
}

// Index ...
func (j *JSON) Index(index int) *JSON {
	l, ok := (j.value).([]interface{})
	if ok && index >= 0 && index < len(l) {
		j.value = l[index]
	} else {
		j.value = nil
	}
	return j
}

// Value ...
func (j *JSON) Value() interface{} {
	return j.value
}

// ToString ...
func (j *JSON) ToString() string {
	if j.value == nil {
		return ""
	}
	return fmt.Sprint(j.value)
}

// ToInt ...
func (j *JSON) ToInt() int64 {
	v, err := strconv.ParseInt(j.ToString(), 10, 64)
	if err != nil {
		return 0
	}
	return v
}

// ToJSON ...
func (j *JSON) ToJSON() string {
	if j.value == nil {
		return ""
	}
	b, _ := json.Marshal(j.value)
	return string(b)
}

// ToArray ...
func (j *JSON) ToArray() interface{} {
	switch (j.value).(type) {
	case []interface{}:
		return (j.value).([]interface{})
	case map[string]interface{}:
		return (j.value).(map[string]interface{})
	default:
		return nil
	}
}
