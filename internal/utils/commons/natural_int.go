package commons

import (
	"fmt"
)

type NaturalInt struct {
	value int
	new   bool
}

func (this NaturalInt) ToInt() int {
	if !this.new {
		panic("the instance of NaturalInt was wrong created")
	}
	return this.value
}

func NewNaturalIntPanic(v int) NaturalInt {
	res, err := NewNaturalInt(v)
	if err != nil {
		panic(err)
	}
	return res
}

func NewNaturalInt(v int) (NaturalInt, error) {
	if v < 1 {
		return NaturalInt{1, false}, fmt.Errorf("a natural int must be greater then 0, got %v", v)
	}
	return NaturalInt{v, true}, nil
}
