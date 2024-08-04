package lib

import (
	"log"
	"sync"
)

type BitSetMap struct {
	data map[uint32]bool
	rwMu sync.RWMutex
}

func NewBitSetMap(size int) *BitSetMap {
	bitSetMap := make(map[uint32]bool)
	for i := 1; i <= size; i++ {
		bitSetMap[uint32(i)] = false
	}
	return &BitSetMap{
		data: bitSetMap,
	}
}

func (b *BitSetMap) Set(position uint32, value bool) {
	b.rwMu.Lock()
	log.Printf("Setting position: %d to isChecked: %t", position, value)
	(*b).data[uint32(position)] = value
	b.rwMu.Unlock()
}

func (b *BitSetMap) Get(position uint32) bool {
	b.rwMu.RLock()
	defer b.rwMu.RUnlock()
	return b.data[position]
}

func (b *BitSetMap) GetMap() map[uint32]bool {
	b.rwMu.RLock()
	defer b.rwMu.RUnlock()

	copyMap := make(map[uint32]bool, len(b.data))
	for k, v := range b.data {
		copyMap[k] = v
	}
	return copyMap
}
