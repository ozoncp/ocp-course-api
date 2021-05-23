//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "TValue=int,model.Course,model.Lesson"

package slice

import "github.com/ozoncp/ocp-course-api/internal/utils/commons"

func SliceGenerateWindowsTValue(xs []TValue, size int, step int) [][]TValue {
	res := [][]TValue{}
	srcLen := len(xs)
	for i := 0; i < srcLen; i += commons.IfOrElse(i+size < srcLen, step, size).(int) {
		end := commons.MinInt(i+size, srcLen)
		res = append(res, xs[i:end])
	}
	return res
}
