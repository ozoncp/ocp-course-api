package commons

func IfOrElse(cond bool, ifTrue, ifFalse interface{}) interface{} {
	if cond {
		return ifTrue
	} else {
		return ifFalse
	}
}

func MinInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
