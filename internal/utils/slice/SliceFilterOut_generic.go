//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "TValue=int,model.Course,model.Lesson"

package slice

func SliceFilterOutTValue(xs []TValue, skip []TValue) []TValue {
	filter := make(map[TValue]struct{})
	for _, v := range skip {
		filter[v] = struct{}{}
	}
	res := []TValue{}
	for _, v := range xs {
		if _, shouldSkip := filter[v]; !shouldSkip {
			res = append(res, v)
		}
	}
	return res
}
