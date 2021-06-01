package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/ozoncp/ocp-course-api/internal/utils/commons"
)

func TestSlidingTValue(t *testing.T) {
	table := []struct {
		src    []TValue
		size   int
		step   int
		expect [][]TValue
	}{
		{nil, 1, 1, [][]TValue{}},
		{[]TValue{}, 1, 1, [][]TValue{}},
		{[]TValue{1, 2, 3}, 2, 2, [][]TValue{{1, 2}, {3}}},
		{[]TValue{1, 2, 3}, 3, 3, [][]TValue{{1, 2, 3}}},
		{[]TValue{1, 2, 3}, 2, 1, [][]TValue{{1, 2}, {2, 3}}},
		{[]TValue{1, 2, 3}, 1, 1, [][]TValue{{1}, {2}, {3}}},
		{[]TValue{1, 2, 3}, 4, 1, [][]TValue{{1, 2, 3}}},
	}

	testFunc := func(out *[][]TValue) func(int, []TValue) bool {
		return func(pos int, xs []TValue) bool {
			*out = append(*out, xs)
			return true
		}
	}

	for _, row := range table {
		got := make([][]TValue, 0)
		SlidingTValue(row.src, NewNaturalIntUnsafe(row.size), NewNaturalIntUnsafe(row.step), testFunc(&got))
		assert.EqualValues(t, row.expect, got,
			"Parameters src: %v; size: %v; step: %v", row.src, row.size, row.step)
	}
}
