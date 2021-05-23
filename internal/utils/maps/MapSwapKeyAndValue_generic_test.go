package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapSwapKeyAndValueTKeyTValue_nonEmpty(t *testing.T) {
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
	assert.EqualValues(t, expect, MapSwapKeyAndValueTKeyTValue(xs))
}

func TestMapSwapKeyAndValueTKeyTValue_empty(t *testing.T) {
	xs := make(map[TKey]TValue)
	expect := make(map[TValue]TKey)
	assert.EqualValues(t, expect, MapSwapKeyAndValueTKeyTValue(xs))
}

func TestMapSwapKeyAndValueTKeyTValue_nil(t *testing.T) {
	expect := make(map[TValue]TKey)
	assert.EqualValues(t, expect, MapSwapKeyAndValueTKeyTValue(nil))
}

func TestMapSwapKeyAndValueTKeyTValue_hasDuplicates(t *testing.T) {
	xs := map[TKey]TValue{
		1: "one",
		2: "two",
		3: "one",
	}
	assert.Panics(t, func() { MapSwapKeyAndValueTKeyTValue(xs) })
}
