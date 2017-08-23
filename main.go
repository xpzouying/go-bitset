package bitset

import "errors"

// TODO(zouying): only offset 64 now, implement more offset
var ErrOutOfRange = errors.New("out of offset range")

// BitSet using uint64 for store value
type BitSet struct {
	set uint64
}

// Set the offset bit in BitSet
func (bs *BitSet) Set(offset uint64) error {
	if offset >= 64 || offset < 0 {
		return ErrOutOfRange
	}

	bs.set |= 1 << offset

	return nil
}

// Get the offset bit in BitSet
func (bs *BitSet) Get(offset uint64) bool {
	return ((bs.set >> offset) & 1) == 1
}
