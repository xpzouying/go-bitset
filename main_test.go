package bitset

import (
	"testing"
)

func assertNil(t *testing.T, v interface{}) {
	t.Helper()

	if v != nil {
		t.Errorf("should be nil, but not: %v", v)
	}
}

func assertNotNil(t *testing.T, v interface{}) {
	t.Helper()

	if v == nil {
		t.Error("should not be nil, but is nil")
	}
}

func assertTrue(t *testing.T, v bool) {
	t.Helper()

	if v != true {
		t.Errorf("should be true, but not: %v", v)
	}
}

func assertFalse(t *testing.T, v bool) {
	t.Helper()

	if v != false {
		t.Errorf("should be false, but not: %v", v)
	}
}

func assertError(t *testing.T, err, expErr error) {
	t.Helper()

	if err != expErr {
		t.Errorf("error should be %v, but be %v", expErr, err)
	}
}

func TestGetAndSetArray1(t *testing.T) {
	bs := NewBitSet(64)

	err := bs.Set(4)
	assertNil(t, err)
	err = bs.Set(0)
	assertNil(t, err)
	err = bs.Set(63)
	assertNil(t, err)

	b, err := bs.Get(4)
	assertNil(t, err)
	assertTrue(t, b)
	b, err = bs.Get(0)
	assertNil(t, err)
	assertTrue(t, b)
	b, err = bs.Get(63)
	assertNil(t, err)
	assertTrue(t, b)
}

func TestGetAndSetArray2(t *testing.T) {
	bs := NewBitSet(65)

	err := bs.Set(64)
	assertNil(t, err)

	b, err := bs.Get(64)
	assertNil(t, err)
	assertTrue(t, b)
}

func TestGetNil(t *testing.T) {
	bs := NewBitSet(10)

	b, err := bs.Get(2)
	assertNil(t, err)
	assertFalse(t, b)
}

func TestSetOutOfRange(t *testing.T) {
	bs := NewBitSet(64)
	err := bs.Set(64)
	assertNotNil(t, err)
	assertError(t, err, ErrOutOfRange)
}
