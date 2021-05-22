package utils

import (
	"fmt"
	"testing"
)
import . "github.com/ozoncp/ocp-course-api/testtool/matchers"

func ExampleGenerateWindowsInt_stepEqSize() {
	ds := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(GenerateWindowsInt(ds, 3, 3))
	// Output: [[1 2 3] [4 5 6]]
}

func ExampleGenerateWindowsInt_stepNeSize() {
	ds := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(GenerateWindowsInt(ds, 3, 2))
	// Output: [[1 2 3] [3 4 5] [5 6]]
}

func TestGenerateWindowsInt(t *testing.T) {
	table := []struct {
		src    []int
		size   int
		step   int
		expect [][]int
	}{
		{nil, 1, 1, [][]int{}},
		{[]int{}, 1, 1, [][]int{}},
		{[]int{1, 2, 3}, 2, 2, [][]int{{1, 2}, {3}}},
		{[]int{1, 2, 3}, 3, 3, [][]int{{1, 2, 3}}},
		{[]int{1, 2, 3}, 2, 1, [][]int{{1, 2}, {2, 3}}},
		{[]int{1, 2, 3}, 1, 1, [][]int{{1}, {2}, {3}}},
		{[]int{1, 2, 3}, 4, 1, [][]int{{1, 2, 3}}},
	}

	for _, row := range table {
		ShouldBeE_with(t,
			GenerateWindowsInt(row.src, row.size, row.step),
			row.expect,
			fmt.Sprintf("\n\tsrc: %v; size: %v; step: %v", row.src, row.size, row.step))
	}
}

func TestGenerateWindowsInt_resultLinkedwithSrc(t *testing.T) {
	ds := []int{1, 2, 3, 4}
	got := GenerateWindowsInt(ds, 3, 2)
	ds[2] = 999
	expect := [][]int{{1, 2, 999}, {999, 4}}
	ShouldBe(t, got, expect)
}

func TestSwapKeyValueIntString_nonEmpty(t *testing.T) {
	src := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	expect := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	ShouldBe(t, SwapKeyValueIntString(src), expect)
}

func TestSwapKeyValueIntString_empty(t *testing.T) {
	src := make(map[int]string)
	expect := make(map[string]int)
	ShouldBe(t, SwapKeyValueIntString(src), expect)
}

func TestSwapKeyValueIntString_nil(t *testing.T) {
	expect := make(map[string]int)
	ShouldBe(t, SwapKeyValueIntString(nil), expect)
}

func TestSwapKeyValueIntString_hasDuplicates(t *testing.T) {
	src := map[int]string{
		1: "one",
		2: "two",
		3: "one",
	}
	ShouldPanic(t, func() { SwapKeyValueIntString(src) })
}

func TestFilterSliceOfInt(t *testing.T) {
	table := []struct {
		in     []int
		skip   []int
		expect []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{1, 1}, nil, []int{1, 1}},
		{[]int{1, 2}, []int{}, []int{1, 2}},
		{[]int{1, 2}, []int{1, 2}, []int{}},
		{[]int{1, 2, 1}, []int{1}, []int{2}},
	}

	for _, row := range table {
		got := FilterSliceOfInt(row.in, row.skip)
		ShouldBeE_with(t,
			got, row.expect,
			fmt.Sprintf("\n\tin: %v; skip: %v", row.in, row.skip))
	}
}
