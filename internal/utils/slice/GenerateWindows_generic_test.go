package slice

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleGenerateWindowsTValue_stepEqSize() {
	ds := []TValue{1, 2, 3, 4, 5, 6}
	fmt.Println(GenerateWindowsTValue(ds, 3, 3))
	// Output: [[1 2 3] [4 5 6]]
}

func ExampleGenerateWindowsTValue_stepNeSize() {
	ds := []TValue{1, 2, 3, 4, 5, 6}
	fmt.Println(GenerateWindowsTValue(ds, 3, 2))
	// Output: [[1 2 3] [3 4 5] [5 6]]
}

func TestGenerateWindowsTValue(t *testing.T) {
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
		got := GenerateWindowsTValue(row.src, row.size, row.step)
		assert.EqualValues(t, row.expect, got,
			"Parameters src: %v; size: %v; step: %v", row.src, row.size, row.step)
	}
}

func TestGenerateWindowsTValue_resultLinkedwithSrc(t *testing.T) {
	ds := []TValue{1, 2, 3, 4}
	got := GenerateWindowsTValue(ds, 3, 2)
	ds[2] = 999
	expect := [][]TValue{{1, 2, 999}, {999, 4}}
	assert.EqualValues(t, expect, got)
}

func BenchmarkGenerateWindowsTValue(b *testing.B) {
	const size = 10000
	ds := make([]TValue, 0, size)
	for i := 0; i < size; i++ {
		ds = append(ds, rand.Int())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GenerateWindowsTValue(ds, 3, 2)
	}
}
