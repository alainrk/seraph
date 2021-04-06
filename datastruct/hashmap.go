package datastruct

import (
	// "fmt"
	"bytes"
	"strconv"
)

// See: http://www.cse.yorku.ca/~oz/hash.html
func sdbmHash(text string) uint64 {
	textBytes := []byte(text)
	var hash uint64 = 0
	for _, c := range textBytes {
		hash = uint64(c) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}

func (hm *HashMap) calcIndex (key string) uint64 {
	// return sdbmHash(key) % hm.Len()
	return sdbmHash(key) & (hm.Len() - 1)
}


func (hm HashMap) Len() uint64 {
	return 1024
}
type HashMap struct {
	items [1024] *LinkedList
}

func (hm HashMap) Print() string {
	var visit bytes.Buffer
	for i, item := range hm.items {
		if item != nil {
			visit.WriteString(strconv.Itoa(i))
			visit.WriteString(": ")
			visit.WriteString(item.Visit())
			visit.WriteString("\n")
			} else {
				visit.WriteString(strconv.Itoa(i))
				visit.WriteString(": [EMPTY]\n")
		}
	}
	return visit.String()
}

func (hm *HashMap) Add(key string, value int) {
	index := hm.calcIndex(key)
	if hm.items[index] == nil {
		hm.items[index] = &LinkedList{}
	}
	hm.items[index].PushBack(key, value)
}

func (hm *HashMap) Get(key string) int {
	index := hm.calcIndex(key)
	return hm.items[index].Get(key)
}