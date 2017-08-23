package bitset

import (
	"testing"
)

func assertNil(t *testing.T, v interface{}) {
	if v != nil {
		t.Errorf("should be nil, but not: %v", v)
	}
}

func assertNotNil(t *testing.T, v interface{}) {
	if v == nil {
		t.Error("should not be nil, but is nil")
	}
}

func assertTrue(t *testing.T, v bool) {
	if v != true {
		t.Errorf("should be true, but not: %v", v)
	}
}

func assertFalse(t *testing.T, v bool) {
	if v != false {
		t.Errorf("should be false, but not: %v", v)
	}
}

func assertError(t *testing.T, err, expErr error) {
	if err != expErr {
		t.Errorf("error should be %v, but be %v", expErr, err)
	}
}

func TestSetterAndGetter(t *testing.T) {
	bs := BitSet{}

	err := bs.Set(4)
	assertNil(t, err)

	err = bs.Set(30)
	assertNil(t, err)

	ok := bs.Get(2)
	assertFalse(t, ok)

	ok = bs.Get(4)
	assertTrue(t, ok)

	ok = bs.Get(30)
	assertTrue(t, ok)
}

func TestSetOutOfRange(t *testing.T) {
	bs := BitSet{}
	err := bs.Set(64)
	assertNotNil(t, err)
	assertError(t, err, ErrOutOfRange)
}
