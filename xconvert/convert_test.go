package xconvert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTitle(t *testing.T) {
	assert.Equal(t, "FirstName", Title("firstName"))
}

func TestToCamelCase(t *testing.T) {
	assert.Equal(t, "userName", ToCamelCase("user_name"))
	assert.Equal(t, "userNAme", ToCamelCase("user_nAme"))
}

func TestToSnakeCase(t *testing.T) {
	assert.Equal(t, "user_name", ToSnakeCase("userName"))
	assert.Equal(t, "user_name", ToSnakeCase("UserName"))
}

func TestCopyStruct(t *testing.T) {
	type A1 struct {
		Name string
		age  int
	}
	type A2 struct {
		Name string
		age  int
	}

	a1 := &A1{
		Name: "hello",
		age:  10,
	}

	// CopyStructByJSON
	a21 := &A2{}
	assert.Nil(t, CopyStructByJSON(a1, a21))
	assert.Equal(t, "hello", a21.Name)
	assert.NotNil(t, CopyStructByJSON(make(chan int), nil))

	// CopyStructByReflect
	a22 := &A2{}
	CopyStructByReflect(a1, a22)
	assert.Equal(t, "hello", a22.Name)
	assert.Equal(t, 0, a22.age)

}
