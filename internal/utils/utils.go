package utils

import "fmt"

func GenerateWindowsInt(src []int, size int, step int) [][]int {
	res := [][]int{}
	srcLen := len(src)
	for i := 0; i < srcLen; i += IfOrElse(i+size < srcLen, step, size).(int) {
		end := MinInt(i+size, srcLen)
		res = append(res, src[i:end])
	}
	return res
}

func SplitOnBunchesInt(src []int, size int) [][]int {
	return GenerateWindowsInt(src, size, size)
}

func SwapKeyValueIntString(src map[int]string) map[string]int {
	res := make(map[string]int)

	if src == nil {
		return res
	}

	for k, v := range src {
		if pk, contains := res[v]; contains {
			panic(
				fmt.Sprintf("The '%v' is not unique. Keys '%v' and '%v' have it",
					v, k, pk))
		}
		res[v] = k
	}

	return res
}

func FilterSliceOfInt(in []int, skip []int) []int {
	filter := make(map[int]struct{})
	for _, v := range skip {
		filter[v] = struct{}{}
	}
	res := []int{}
	for _, v := range in {
		if _, shouldSkip := filter[v]; !shouldSkip {
			res = append(res, v)
		}
	}
	return res
}

func IfOrElse(cond bool, ifTrue, ifFalse interface{}) interface{} {
	if cond {
		return ifTrue
	} else {
		return ifFalse
	}
}

func MinInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
