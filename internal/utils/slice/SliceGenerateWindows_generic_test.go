package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleSliceGenerateWindowsTValue_stepEqSize() {
	ds := []TValue{1, 2, 3, 4, 5, 6}
	fmt.Println(SliceGenerateWindowsTValue(ds, 3, 3))
	// Output: [[1 2 3] [4 5 6]]
}

func ExampleSliceGenerateWindowsTValue_stepNeSize() {
	ds := []TValue{1, 2, 3, 4, 5, 6}
	fmt.Println(SliceGenerateWindowsTValue(ds, 3, 2))
	// Output: [[1 2 3] [3 4 5] [5 6]]
}

func TestSliceGenerateWindowsTValue(t *testing.T) {
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

	for _, row := range table {
		got := SliceGenerateWindowsTValue(row.src, row.size, row.step)
		assert.EqualValues(t, row.expect, got,
			"Parameters src: %v; size: %v; step: %v", row.src, row.size, row.step)
	}
}

func TestGenerateWindowsTValue_resultLinkedwithSrc(t *testing.T) {
	ds := []TValue{1, 2, 3, 4}
	got := SliceGenerateWindowsTValue(ds, 3, 2)
	ds[2] = 999
	expect := [][]TValue{{1, 2, 999}, {999, 4}}
	assert.EqualValues(t, expect, got)
}
