package bitset

import (
	"errors"
	"math"
)

var ErrOutOfRange = errors.New("out of offset range")

// BitSet define interface which have Set & Get
// NOTE: offset start from 0
type BitSet interface {
	Set(offset uint64) error
	Get(offset uint64) (bool, error)
}

// BitSet using uint64 for store value
type bitset struct {
	// arrays to store bit status
	arrays []uint64

	// numBit is how many bits could store
	numBit uint64
}

// NewBitSet to get BitSet object which implemented by bitset
func NewBitSet(numBit uint64) BitSet {
	// func NewBitSet(numBit uint64) *bitset {
	// the length of arrays is the ceil of numBit/sizeof(uint64)
	len := uint64(math.Ceil(float64(numBit) / 64))

	return &bitset{
		arrays: make([]uint64, len),
		numBit: numBit,
	}
}

// Set the offset bit in BitSet
func (bs *bitset) Set(offset uint64) error {
	// get the right offset in the arrays array
	idx, err := bs.idxFromOffset(offset)
	if err != nil {
		return err
	}

	// the right index of the arrays[idx]
	offsetInArray := offset - idx*64

	bs.arrays[idx] |= 1 << offsetInArray

	return nil
}

// Get the offset bit in BitSet
func (bs *bitset) Get(offset uint64) (bool, error) {
	idx, err := bs.idxFromOffset(offset)
	if err != nil {
		return false, err
	}

	// the right index of the arrays[idx]
	offsetInArray := offset - idx*64

	return ((bs.arrays[idx] >> offsetInArray) & 1) == 1, nil
}

// get index of arrays from offset
// return the index of array to store the offset bit data
func (bs *bitset) idxFromOffset(offset uint64) (idx uint64, err error) {
	if offset >= bs.numBit {
		// offset start from 0, so equal is out of range
		return 0, ErrOutOfRange
	}

	// get the setting idx of arrays
	return uint64(math.Floor(float64(offset) / 64)), nil
}
