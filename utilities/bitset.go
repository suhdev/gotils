package utilities

import "fmt"

type BitSet struct {
	value    []byte
	noOfBits int
}

// NewBitSet creates a new BitSet
func NewBitSet() *BitSet {
	return &BitSet{
		value:    make([]byte, 0),
		noOfBits: 0,
	}
}

func (set *BitSet) setValueOfBit(idx, value int) {
	byteIndex := idx / 8
	bitIndex := uint8(idx % 8)
	val := (value << bitIndex)
	if idx >= set.noOfBits {
		set.noOfBits = idx + 1
	}
	if byteIndex >= len(set.value) {
		count := byteIndex - len(set.value) + 1
		for i := 0; i < count; i++ {
			set.value = append(set.value, 0)
		}
	}

	if value == 1 {
		set.value[byteIndex] = set.value[byteIndex] | byte(val)
	} else {
		set.value[byteIndex] = set.value[byteIndex] &^ (byte(val))
	}

}

func (set *BitSet) setRangeValue(startIdx, endIndex, value int) {
	if startIdx == endIndex {
		set.setValueOfBit(startIdx, value)
	} else if startIdx > endIndex {
		set.setRangeValue(endIndex, startIdx, value)
	} else {
		for i := startIdx; i < endIndex; i++ {
			set.setValueOfBit(i, value)
		}
	}
}

// Set sets the bit at the given index
func (set *BitSet) Set(idx int) {
	set.setValueOfBit(idx, 1)
}

// Clear clears the bit at the given index
func (set *BitSet) Clear(idx int) {
	set.setValueOfBit(idx, 0)
}

// SetRange sets a range of bits given a start and end indices
func (set *BitSet) SetRange(startIdx, endIndex int) {
	set.setRangeValue(startIdx, endIndex, 1)
}

// ClearRange clears a range of bits given a start and end indices
func (set *BitSet) ClearRange(startIdx, endIndex int) {
	set.setRangeValue(startIdx, endIndex, 0)
}

// Length returns the no. of bits that have been set
func (set *BitSet) Length() int {
	return set.noOfBits
}

// Size returns the no. of bytes used to store the bits within the set
func (set *BitSet) Size() int {
	return len(set.value)
}

// Get returns whether the bit at the given index is set or not
func (set *BitSet) Get(idx int) bool {
	byteIndex := idx / 8
	bitIndex := uint8(idx % 8)
	x := set.value[byteIndex]
	return ((x >> bitIndex) & 1) == 1
}

func (set *BitSet) String() string {
	str := ""
	l := set.noOfBits
	for i := 0; i < l; i++ {
		if set.Get(i) {
			str = "1" + str
		} else {
			str = "0" + str
		}
	}
	return fmt.Sprintf("{%s %d}", str, set.Size())
}
