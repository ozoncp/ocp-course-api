//go:generate genny -in=$GOFILE -out=$GOFILE.gen.go gen "TValue=model.Course,model.Lesson"

package flusher

import (
	"github.com/cheekybits/genny/generic"

	"github.com/ozoncp/ocp-course-api/internal/utils/commons"
	"github.com/ozoncp/ocp-course-api/internal/utils/slice"
)

type TValue = generic.Type

type FlusherTValue interface {
	FlushTValue(vs []TValue) []TValue
}

type repoTValue interface {
	AddTValues(vs []TValue) error
}

type flusherTValue struct {
	repo      repoTValue
	batchSize commons.NaturalInt
}

func (this *flusherTValue) FlushTValue(vs []TValue) []TValue {
	var res []TValue
	slice.SlidingTValue(vs, this.batchSize, this.batchSize, func(pos int, window []TValue) bool {
		err := this.repo.AddTValues(window)
		if err != nil {
			res = vs[pos:]
			return false
		}
		return true
	})
	return res
}

func NewFlusherTValue(repo repoTValue, batchSize commons.NaturalInt) FlusherTValue {
	return &flusherTValue{repo, batchSize}
}
