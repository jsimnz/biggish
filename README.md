A big(ish) math package for Go
==============================

A math package for go that works with special Int types backed arbitrary and definable byte arrays.

### Why

The current Go math/big package allows you to use arbitrary sized byte arrays for integer math, but theres no overflowing. This package is desigend to work with intergers with a defined byte array backend.

So you can have normal int32 of 4 byte size, or int64 of 8 bytes. But you can also have ints backed by 6 bytes, or 256 bytes, and so on.