//go:generate genny -in=$GOFILE -out=$GOFILE.gen.go gen "TValue=model.Course,model.Lesson"

package repo

import (
	"github.com/cheekybits/genny/generic"
)

type TValue = generic.Type

type RepoTValue interface {
	DescribeTValue(id uint64) (TValue, error)
	ListTValues(limit uint64, offset uint64) ([]TValue, error)
	AddTValue(v TValue) (uint64, error)
	AddTValues(vs []TValue) error
	RemoveTValue(id uint64) error
}
