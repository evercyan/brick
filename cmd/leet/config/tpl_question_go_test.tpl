package solution

import (
	"fmt"
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
	for k, v := range cases {
		t.Run(fmt.Sprintf("FuncToReplace_%d", k), func(t *testing.T) {
			res := FuncToReplace(v.inputs[0].(string))
			assert.Equal(t, v.expects[0].(string), res)
		})
	}
}