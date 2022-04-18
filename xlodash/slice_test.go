package xlodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	assert.ElementsMatch(t, []string{"a", "b"}, Unique([]string{"a", "b", "a"}))
	assert.ElementsMatch(t, []int{1, 2}, Unique([]int{1, 2, 1}))
}

func TestMap(t *testing.T) {
	a1 := []int{
		1, 2,
	}
	w1 := []int{2, 3}
	r1 := Map(a1, func(k int, v int) int {
		return v + 1
	})
	assert.ElementsMatch(t, w1, r1)

	// ----------------------------------------------------------------

	a2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	w2 := []int{1, 4}
	r2 := Map(a2, func(k int, v []int) int {
		return v[0]
	})
	assert.ElementsMatch(t, w2, r2)

	// ----------------------------------------------------------------

	type Person struct {
		Name string
		Age  int
	}
	a3 := []*Person{
		{
			Name: "a",
			Age:  1,
		},
		{
			Name: "b",
			Age:  2,
		},
	}
	w3 := []string{"a", "b"}
	r3 := Map(a3, func(k int, v *Person) string {
		return v.Name
	})
	assert.ElementsMatch(t, w3, r3)
}

func TestFilter(t *testing.T) {
	assert.ElementsMatch(t, []string{"c"}, Filter([]string{"a", "b", "c"}, func(k int, v string) bool {
		return v == "c"
	}))
	assert.ElementsMatch(t, []int{3, 4, 5}, Filter([]int{1, 2, 3, 4, 5}, func(k int, v int) bool {
		return v >= 3
	}))
}

func TestGroupBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	a := []*Person{
		{
			Name: "a",
			Age:  1,
		},
		{
			Name: "a",
			Age:  2,
		},
		{
			Name: "b",
			Age:  3,
		},
	}
	w := map[string][]*Person{
		"a": {
			{
				Name: "a",
				Age:  1,
			},
			{
				Name: "a",
				Age:  2,
			},
		},
		"b": {
			{
				Name: "b",
				Age:  3,
			},
		},
	}
	assert.Equal(t, w, GroupBy(a, func(v *Person) string {
		return v.Name
	}))
}

func TestChunk(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	w := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7},
	}
	assert.Equal(t, w, Chunk(a, 2))
	assert.Equal(t, [][]int{}, Chunk(a, 0))
}

func TestUnion(t *testing.T) {
	assert.ElementsMatch(t, []int{2, 3}, Intersect(
		[]int{1, 2, 3},
		[]int{2, 3, 4},
	))
	assert.ElementsMatch(t, []int{1}, Diff(
		[]int{1, 2, 3},
		[]int{2, 3, 4},
	))
	assert.ElementsMatch(t, []int{1, 2, 3, 4}, Union(
		[]int{1, 2, 3},
		[]int{2, 3, 4},
	))
}
