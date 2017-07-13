lex [![GoDoc](https://godoc.org/github.com/xcdb/lex?status.svg)](https://godoc.org/github.com/xcdb/lex) [![Go Report Card](https://goreportcard.com/badge/github.com/xcdb/lex)](https://goreportcard.com/report/github.com/xcdb/lex)
====

Implements an encoder that preserves lexicographical order for [Boolean, Numeric and String](https://golang.org/ref/spec#Types) types in Go.

Sorted key/value stores such as [LMDB](http://symas.com/mdb), [Bolt](https://github.com/boltdb) and [LevelDB](http://LevelDB.org) support sequential iteration over the data. While custom [comparison](http://lmdb.tech/doc/group__mdb.html#gaa8e6e7a6f99bd7142947c48f0c4b970f) [functions](https://github.com/google/leveldb/blob/master/include/leveldb/comparator.h) are possible, by default a bytewise comparison of keys is used and an appropriate encoding must be selected so that comparisons remain consistent.

Strings are often used, as selecting formats compatible with bytewise comparison is trivial. Big-endian unsigned integers are similarly easy to use. Signed integers are a little harder to get right as, for example, -1 will normally sort after +1.

Lex provides functions that allow the safe usage of many more types with the default bytewise comparison. Efficient implementations are provided for many core types, with structs and aliased types also supported via a reflection-based approach.

Boolean and Numeric types are encoded as appropriate fixed-size values, while Strings are encoded simply as their underlying bytes with a single `NUL` character appended. Note that type information is *not* serialized with the value, and needs to be maintained separately.


