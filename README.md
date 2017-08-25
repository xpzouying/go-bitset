# go-bitset

[![Build Status](https://travis-ci.org/xpzouying/go-bitset.svg?branch=master)](https://travis-ci.org/xpzouying/go-bitset)

Bitset in go way.


## Test

NOTE: Testing is only support `go1.9`. If you want testing this pkg in the lower go env, please remove t.Helper() first.

More detail information for `T.Helper()` (release in Go1.9), [T.Helper](https://golang.org/pkg/testing/#T.Helper)

Run testcase: `go test -v`


## Usage

```go
// setting 100 bits length
bs := NewBitSet(100)

// NOTE: the offset for Set() is from 0 to 99(NOT 100)
err := bs.Set(85)
if err != nil {
    panic(err)
}

b, err := bs.Get(64)
if err != nil {
    panic(err)
}

if b == true {
    // 64 bit is setting
} else {
    // 64 bit is not setting
}
```

