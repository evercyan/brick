package xlodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeysAndValues(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	assert.ElementsMatch(t, []string{"a", "b", "c"}, Keys(m))
	assert.ElementsMatch(t, []int{1, 2, 3}, Values(m))
}
