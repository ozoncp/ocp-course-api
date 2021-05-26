//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "TKey=int,uint TValue=string,model.Course,model.Lesson"

package maps

import (
	"fmt"
)

func SwapKeyAndValueTKeyTValue(xs map[TKey]TValue) map[TValue]TKey {
	res := make(map[TValue]TKey, len(xs))
	if xs == nil {
		return res
	}
	for k, v := range xs {
		if pk, contains := res[v]; contains {
			panic(
				fmt.Sprintf("The '%v' is not unique. Keys '%v' and '%v' have it",
					v, k, pk))
		}
		res[v] = k
	}
	return res
}
