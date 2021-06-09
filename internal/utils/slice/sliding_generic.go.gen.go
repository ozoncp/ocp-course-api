// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package slice

import (
	"github.com/ozoncp/ocp-course-api/api/model"
	"github.com/ozoncp/ocp-course-api/internal/utils/commons"
)

func SlidingInt(xs []int, size commons.NaturalInt, step commons.NaturalInt, f func(pos int, window []int) bool) {
	srcLen := len(xs)
	for i := 0; i < srcLen; i += step.ToInt() {
		end := commons.MinInt(i+size.ToInt(), srcLen)
		needContinue := f(i, xs[i:end:end])
		if !needContinue || (i+size.ToInt() >= srcLen) {
			break
		}
	}
}

func SlidingModelCourse(xs []model.Course, size commons.NaturalInt, step commons.NaturalInt, f func(pos int, window []model.Course) bool) {
	srcLen := len(xs)
	for i := 0; i < srcLen; i += step.ToInt() {
		end := commons.MinInt(i+size.ToInt(), srcLen)
		needContinue := f(i, xs[i:end:end])
		if !needContinue || (i+size.ToInt() >= srcLen) {
			break
		}
	}
}

func SlidingModelLesson(xs []model.Lesson, size commons.NaturalInt, step commons.NaturalInt, f func(pos int, window []model.Lesson) bool) {
	srcLen := len(xs)
	for i := 0; i < srcLen; i += step.ToInt() {
		end := commons.MinInt(i+size.ToInt(), srcLen)
		needContinue := f(i, xs[i:end:end])
		if !needContinue || (i+size.ToInt() >= srcLen) {
			break
		}
	}
}
