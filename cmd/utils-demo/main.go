package main

import (
	"fmt"
	um "github.com/ozoncp/ocp-course-api/internal/utils/maps"
	us "github.com/ozoncp/ocp-course-api/internal/utils/slice"
)

func main() {
	{
		in := []int{1, 2, 3, 4, 5, 6}
		size := 3
		step := 2
		res := us.SliceGenerateWindowsInt(in, size, step)
		fmt.Printf("GenerateWindowsInt(%v, %v, %v): %v\n", in, size, step, res)
	}

	{
		in := []int{1, 2, 3, 1, 2, 3}
		skip := []int{1, 2}
		res := us.SliceFilterOutInt(in, skip)
		fmt.Printf("FilterSliceOfInt(%v, %v): %v\n", in, skip, res)
	}

	{
		in := map[int]string{1: "first", 2: "second", 3: "third"}
		res := um.MapSwapKeyAndValueIntString(in)
		fmt.Printf("SwapKeyValueIntString(%v): %v\n", in, res)
	}
}
