package lib

type BitSetMap map[uint32]bool

func NewBitSetMap() *BitSetMap {
	bitSetMap := make(BitSetMap)
	return &bitSetMap
}

func (b *BitSetMap) Set(position uint32, value bool) {
	if value {
		(*b)[position] = value
	} else {
		delete((*b), position)
	}
}

func (b BitSetMap) Get(position uint32) bool {
	return b[position]
}
