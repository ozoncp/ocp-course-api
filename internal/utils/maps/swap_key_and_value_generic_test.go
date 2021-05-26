package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwapKeyAndValueTKeyTValue_nonEmpty(t *testing.T) {
	xs := map[TKey]TValue{
		1: "one",
		2: "two",
		3: "three",
	}
	expect := map[TValue]TKey{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	assert.EqualValues(t, expect, SwapKeyAndValueTKeyTValue(xs))
}

func TestSwapKeyAndValueTKeyTValue_empty(t *testing.T) {
	xs := make(map[TKey]TValue)
	expect := make(map[TValue]TKey)
	assert.EqualValues(t, expect, SwapKeyAndValueTKeyTValue(xs))
}

func TestSwapKeyAndValueTKeyTValue_nil(t *testing.T) {
	expect := make(map[TValue]TKey)
	assert.EqualValues(t, expect, SwapKeyAndValueTKeyTValue(nil))
}

func TestSwapKeyAndValueTKeyTValue_hasDuplicates(t *testing.T) {
	xs := map[TKey]TValue{
		1: "one",
		2: "two",
		3: "one",
	}
	assert.Panics(t, func() { SwapKeyAndValueTKeyTValue(xs) })
}
