package commons

import (
	"fmt"
)

type NaturalInt struct {
	value int
}

func (this NaturalInt) ToInt() int {
	return this.value
}

func NewNaturalIntPanic(v int) NaturalInt {
	if v < 1 {
		panic(fmt.Sprintf("A natural int must be greater then 0, got %v", v))
	}
	return NaturalInt{v}
}

func NewNaturalInt(v int) (NaturalInt, error) {
	if v < 1 {
		return NaturalInt{v}, fmt.Errorf("a natural int must be greater then 0, got %v", v)
	}
	return NaturalInt{v}, nil
}
