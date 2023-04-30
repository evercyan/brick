package xfilter

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	content := `[
	{
		"id": "1",
		"filter": [
        	["ctx.uid", "=", "1"],
			"and"
		]
	},
	{
		"id": "2",
		"filter": [
        	["ctx.uid", "=", "2"],
			"and"
		]
	},
	{
		"id": "3"
	}
]`
	list := make([]map[string]interface{}, 0)
	json.Unmarshal([]byte(content), &list)

	ctx := NewContext()
	ctx.Set("uid", "1")

	res, err := Filter(ctx, list)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(res))
	assert.Equal(t, "1", res[0]["id"])
	assert.Equal(t, "3", res[1]["id"])
}

func TestAssert(t *testing.T) {
	rules := `[
	[
		["year", "=", 2022],
		["year", "=", 2023],
		"or"
	],
	[
		["year", ">=", 0],
		["month", ">=", 0],
		["day", ">=", 0],
		["hour", ">=", 0],
		["minute", ">=", 0],
		["second", ">=", 0],
		["wday", ">=", 0],
		["date", "!=", ""],
		["time", "!=", ""],
		["unixtime", "!=", ""],
		["datetime", "!=", ""],
		"and"
	],
	["ctx.trace_id", "=", "123456"],
	[
		["ctx.uid", "=", "7"],
		["ctx.uid", "!=", "8"],
		["ctx.uid", ">", "0"],
		["ctx.uid", ">=", "0"],
		["ctx.uid", "<", "10"],
		["ctx.uid", "<=", "10"],
		["ctx.uid", "between", "5, 10"],
		["ctx.uid", "in", "6, 7, 8"],
		["ctx.uid", "not in", "6, 8"],
		["ctx.user.name", "=", "hello"],
		["ctx.user.detail.nums.1", "=", "2"],
		["ctx.user.detail.nums", "has", "4"],
		["ctx.user.name", "match", "/l{2}/"],
		["ctx.user.name", "match", "ell"],
		["ctx.user.name", "not match", "/l{3}/"],
		["ctx.user.name", "not match", "ellll"],
		"and"
	],
	"and"
]`

	c := context.Background()
	c = context.WithValue(c, "trace_id", "123456")

	ctx := NewContext(c)
	ctx.Set("uid", 7)
	ctx.Set("user", map[string]interface{}{
		"name": "hello",
		"detail": map[string]interface{}{
			"sex": "male",
			"nums": []int{
				1, 2, 3, 4,
			},
		},
	})

	err := Assert(ctx, rules)
	assert.Nil(t, err)
}

func TestAssertCoverage(t *testing.T) {
	ctx := NewContext()

	assert.NotNil(t, Assert(ctx, `[
	[
		["year", "=", 2020],
		["year", "=", 2021],
		"or"
	],
	"and"
]`))
	assert.NotNil(t, Assert(ctx, `[
	[
		["ddd", "=", 2020],
		"or"
	],
	"and"
]`))
	assert.NotNil(t, Assert(ctx, `[
	[
		["year", 0, 2020],
		["year", "haha", 2020],
		"or"
	],
	"and"
]`))
	assert.NotNil(t, Assert(ctx, `[
	[
		["year", "haha", 2020],
		"or"
	],
	"and"
]`))
	assert.NotNil(t, Assert(ctx, `[
	[
		["year", "between", "2020, 2021, 2022"],
		"or"
	],
	"and"
]`))
	assert.NotNil(t, Assert(ctx, `[
	[
		["year", "=", 2020]
	],
	"and"
]`))

	_, err1 := ParseFilter("")
	assert.NotNil(t, err1)

	_, err2 := ParseFilter("[]")
	assert.NotNil(t, err2)

	assert.NotNil(t, Assert(ctx, "[]"))

	_, err3 := Filter(ctx, nil)
	assert.NotNil(t, err3)
}
