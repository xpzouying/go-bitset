# go-bitset

Bitset in go way.


## Test

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

