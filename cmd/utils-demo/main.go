package main

import (
	"fmt"
	"github.com/ozoncp/ocp-course-api/internal/utils/maps"
	"github.com/ozoncp/ocp-course-api/internal/utils/slice"
)

func main() {
	{
		in := []int{1, 2, 3, 4, 5, 6}
		size := 3
		step := 2
		res := slice.GenerateWindowsInt(in, size, step)
		fmt.Printf("GenerateWindowsInt(%v, %v, %v): %v\n", in, size, step, res)
	}

	{
		in := []int{1, 2, 3, 1, 2, 3}
		skip := []int{1, 2}
		res := slice.FilterOutInt(in, skip)
		fmt.Printf("FilterSliceOfInt(%v, %v): %v\n", in, skip, res)
	}

	{
		in := map[int]string{1: "first", 2: "second", 3: "third"}
		res := maps.SwapKeyAndValueIntString(in)
		fmt.Printf("SwapKeyValueIntString(%v): %v\n", in, res)
	}
}
