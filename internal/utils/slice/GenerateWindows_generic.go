//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "TValue=int,model.Course,model.Lesson"

package slice

import "github.com/ozoncp/ocp-course-api/internal/utils/commons"

func GenerateWindowsTValue(xs []TValue, size int, step int) [][]TValue {
	srcLen := len(xs)
	dstLen := srcLen / step
	if (srcLen % step) > 0 {
		dstLen++
	}
	res := make([][]TValue, 0, dstLen)
	for i := 0; i < srcLen; i += step {
		end := commons.MinInt(i+size, srcLen)
		res = append(res, xs[i:end:end])
		if i+size >= srcLen {
			break
		}
	}
	return res
}
