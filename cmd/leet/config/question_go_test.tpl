package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FuncToReplace(t *testing.T) {
	cases := []struct {
		inputs  []interface{}
		expects []interface{}
	}{
		{
			[]interface{}{
				"input",
			},
			[]interface{}{
				"output",
			},
		},
	}
	for _, c := range cases {
		t.Run("FuncToReplace", func(t *testing.T) {
			ret := FuncToReplace(c.inputs[0].(string))
			assert.Equal(t, c.expects[0].(string), ret)
		})
	}
}