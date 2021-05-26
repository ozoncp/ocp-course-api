package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterOutTValue(t *testing.T) {
	table := []struct {
		xs     []TValue
		skip   []TValue
		expect []TValue
	}{
		{[]TValue{}, []TValue{}, []TValue{}},
		{[]TValue{1, 1}, nil, []TValue{1, 1}},
		{[]TValue{1, 2}, []TValue{}, []TValue{1, 2}},
		{[]TValue{1, 2}, []TValue{1, 2}, []TValue{}},
		{[]TValue{1, 2, 1}, []TValue{1}, []TValue{2}},
	}

	for _, row := range table {
		got := FilterOutTValue(row.xs, row.skip)
		assert.EqualValues(t,
			got, row.expect,
			fmt.Sprintf("Input: xs: %v; skip: %v", row.xs, row.skip))
	}
}
