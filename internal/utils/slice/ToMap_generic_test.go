package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleToMapTValueTKey() {
	type Value struct {
		id int
		v  string
	}

	xs := []TValue{
		Value{1, "one"},
		Value{2, "second"},
		Value{1, "first"},
	}
	res := ToMapTValueTKey(xs, func(x TValue) TKey {
		return x.(Value).id
	})

	for k, v := range res {
		fmt.Printf("(%v,%v)\n", k, v)
	}
	//Unordered output:
	//(1,{1 first})
	//(2,{2 second})
}

func TestToMapTValueTKey(t *testing.T) {
	type Value struct {
		id int
		v  string
	}

	type Row struct {
		in     []TValue
		expect map[TKey]TValue
	}

	op := func(x TValue) TKey {
		return x.(Value).id
	}

	table := []Row{
		{nil, make(map[TKey]TValue)},
		{[]TValue{}, make(map[TKey]TValue)},
		{[]TValue{Value{5, "b"}, Value{3, "a"}},
			map[TKey]TValue{5: Value{5, "b"}, 3: Value{3, "a"}}},
		{[]TValue{Value{3, "w"}, Value{5, "b"}, Value{3, "a"}},
			map[TKey]TValue{5: Value{5, "b"}, 3: Value{3, "a"}}},
	}

	for _, row := range table {
		got := ToMapTValueTKey(row.in, op)
		assert.EqualValues(t, row.expect, got)
	}
}
