package saver

type OverflowStrategy struct {
	value int8
	new   bool
}

const (
	dropFirst int8 = iota
	dropAll
	block
)

func OverflowStrategyDropFirst() OverflowStrategy {
	return OverflowStrategy{value: dropFirst, new: true}
}

func OverflowStrategyDropAll() OverflowStrategy {
	return OverflowStrategy{value: dropAll, new: true}
}

func OverflowStrategyBlock() OverflowStrategy {
	return OverflowStrategy{value: block, new: true}
}
