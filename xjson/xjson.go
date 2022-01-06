package xjson

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type js struct {
	value interface{}
}

// New ...
func New(s string) *js {
	j := new(js)
	var v interface{}
	if err := json.Unmarshal([]byte(s), &v); err != nil {
		return j
	}
	j.value = v
	return j
}

// Key ...
func (j *js) Key(key string) *js {
	m, ok := (j.value).(map[string]interface{})
	if !ok {
		j.value = nil
		return j
	}
	if _, ok := m[key]; !ok {
		j.value = nil
		return j
	}
	j.value = m[key]
	return j
}

// Index ...
func (j *js) Index(index int) *js {
	l, ok := (j.value).([]interface{})
	if !ok {
		j.value = nil
		return j
	}

	if index > len(l)-1 {
		j.value = nil
		return j
	}

	j.value = l[index]
	return j
}

// Value ...
func (j *js) Value() interface{} {
	return j.value
}

// ToString ...
func (j *js) ToString() string {
	if j.value == nil {
		return ""
	}
	return fmt.Sprint(j.value)
}

// ToUint ...
func (j *js) ToUint() uint64 {
	v, err := strconv.ParseUint(j.ToString(), 10, 64)
	if err != nil {
		return 0
	}
	return v
}

// ToJSON ...
func (j *js) ToJSON() string {
	if j.value == nil {
		return ""
	}
	b, _ := json.Marshal(j.value)
	return string(b)
}

// ToArray ...
func (j *js) ToArray() interface{} {
	switch (j.value).(type) {
	case []interface{}:
		return (j.value).([]interface{})
	case map[string]interface{}:
		return (j.value).(map[string]interface{})
	default:
		return nil
	}
}
