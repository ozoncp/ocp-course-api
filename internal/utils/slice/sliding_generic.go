//go:generate genny -in=$GOFILE -out=$GOFILE.gen.go gen "TValue=int,model.Course,model.Lesson"

package slice

import "github.com/ozoncp/ocp-course-api/internal/utils/commons"

func SlidingTValue(xs []TValue, size int, step int, f func(pos int, window []TValue) bool) {
	srcLen := len(xs)
	for i := 0; i < srcLen; i += step {
		end := commons.MinInt(i+size, srcLen)
		needContinue := f(i, xs[i:end:end])
		if !needContinue || (i+size >= srcLen) {
			break
		}
	}
}
