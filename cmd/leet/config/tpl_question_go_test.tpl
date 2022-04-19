package solution

import (
	"fmt"
	"testing"

	"github.com/evercyan/brick/cmd/leet/xleet"
	"github.com/stretchr/testify/assert"
)

func cases() []xleet.Case {
	return []xleet.Case{
		{
			[]interface{}{
				"input",
			},
			[]interface{}{
				"output",
			},
		},
	}
}

func Test_FuncToReplace(t *testing.T) {
	for k, v := range cases() {
		t.Run(fmt.Sprintf("FuncToReplace_%d", k+1), func(t *testing.T) {
			res := FuncToReplace(v.Args[0].(string))
			assert.Equal(t, v.Wants[0].(string), res)
		})
	}
}