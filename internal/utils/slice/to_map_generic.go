//go:generate genny -in=$GOFILE -out=$GOFILE.gen.go gen "TValue=model.Course,model.Lesson TKey=uint64"

package slice

// Converts the slice xs to map with the type of slice value as a value and the
// result of op as a key. If the op returns the same result for some values,
//the resulting map includes the latest value.
func ToMapTValueTKey(xs []TValue, op func(TValue) TKey) map[TKey]TValue {
	res := make(map[TKey]TValue, len(xs))
	for _, v := range xs {
		res[op(v)] = v
	}
	return res
}
