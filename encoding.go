package lex

import (
	"encoding/binary"
	"math"
)

//PutBool serializes bool as a single byte.
//True is encoded as 1; false as 0.
func PutBool(b []byte, v bool) {
	if v {
		b[0] = 1
	} else {
		b[0] = 0
	}
}

//Bool deserializes bool from a single byte.
func Bool(b []byte) bool {
	return Byte(b) == 1
}

//PutUint8 serializes uint8 as a single byte.
func PutUint8(b []byte, v uint8) {
	b[0] = byte(v)
}

//Uint8 deserializes uint8 from a single byte.
func Uint8(b []byte) uint8 {
	return uint8(Byte(b))
}

//PutUint16 serializes uint16 as 2 bytes.
//Order is preserved by encoding in big-endian.
func PutUint16(b []byte, v uint16) {
	binary.BigEndian.PutUint16(b, v)
}

//Uint16 deserializes uint16 from 2 bytes.
func Uint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

//PutUint32 serializes uint32 as 4 bytes.
//Order is preserved by encoding in big-endian.
func PutUint32(b []byte, v uint32) {
	binary.BigEndian.PutUint32(b, v)
}

//Uint32 deserializes uint32 from 4 bytes.
func Uint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

//PutUint64 serializes uint64 as 8 bytes.
//Order is preserved by encoding in big-endian.
func PutUint64(b []byte, v uint64) {
	binary.BigEndian.PutUint64(b, v)
}

//Uint64 deserializes uint64 from 8 bytes.
func Uint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

//PutInt8 serializes int8 as 1 byte.
//Order is preserved by flipping the sign and encoding in big-endian.
func PutInt8(b []byte, v int8) {
	PutUint8(b, uint8(v^math.MinInt8))
}

//Int8 deserializes int8 from 1 byte.
func Int8(b []byte) int8 {
	v := int8(Uint8(b))
	return v ^ math.MinInt8
}

//PutInt16 serializes int16 as 2 bytes.
//Order is preserved by flipping the sign and encoding in big-endian.
func PutInt16(b []byte, v int16) {
	PutUint16(b, uint16(v^math.MinInt16))
}

//Int16 deserializes int16 from 2 bytes.
func Int16(b []byte) int16 {
	v := int16(Uint16(b))
	return v ^ math.MinInt16
}

//PutInt32 serializes int32 as 4 bytes.
//Order is preserved by flipping the sign and encoding in big-endian.
func PutInt32(b []byte, v int32) {
	PutUint32(b, uint32(v^math.MinInt32))
}

//Int32 deserializes int32 from 4 bytes.
func Int32(b []byte) int32 {
	v := int32(Uint32(b))
	return v ^ math.MinInt32
}

//PutInt64 serializes int64 as 8 bytes.
//Order is preserved by flipping the sign and encoding in big-endian.
func PutInt64(b []byte, v int64) {
	PutUint64(b, uint64(v^math.MinInt64))
}

//Int64 deserializes int64 from 8 bytes.
func Int64(b []byte) int64 {
	v := int64(Uint64(b))
	return v ^ math.MinInt64
}

//PutFloat32 serializes float32 as 4 bytes.
//Order is preserved by transforming to a comparable format and encoding in big-endian.
//
//Behaviour follows conventions as far as possible, specifically:
// -0.0 and +0.0 are treated as equal.
// Positive and negative infinity are treated as equal.
// Infinity sorts after max value.
// NaN sorts after infinity.
//
//Value transformation is as per Hacker's Delight 2nd Edition, 17-3.
func PutFloat32(b []byte, v float32) {
	n := int32(math.Float32bits(v))
	if n >= 0 {
		n = n + math.MinInt32
	} else {
		n = -n
	}
	PutUint32(b, uint32(n))
}

//Float32 deserializes float32 from 4 bytes.
func Float32(b []byte) float32 {
	//See Hacker's Delight 2nd Edition, 17-3
	n := int32(Uint32(b))
	if n >= 0 {
		n = -n
	} else {
		n = n + math.MinInt32
	}
	return math.Float32frombits(uint32(n))
}

//PutFloat64 serializes float64 as 8 bytes.
//Order is preserved by transforming to a comparable format and encoding in big-endian.
//
//Behaviour follows conventions as far as possible, specifically:
// -0.0 and +0.0 are treated as equal.
// Positive and negative infinity are treated as equal.
// Infinity sorts after max value.
// NaN sorts after infinity.
//
//Value transformation is as per Hacker's Delight 2nd Edition, 17-3.
func PutFloat64(b []byte, v float64) {
	n := int64(math.Float64bits(v))
	if n >= 0 {
		n = n + math.MinInt64
	} else {
		n = -n
	}
	PutUint64(b, uint64(n))
}

//Float64 deserializes float64 from 8 bytes.
func Float64(b []byte) float64 {
	//See Hacker's Delight 2nd Edition, 17-3
	n := int64(Uint64(b))
	if n >= 0 {
		n = -n
	} else {
		n = n + math.MinInt64
	}
	return math.Float64frombits(uint64(n))
}

//PutComplex64 serializes complex64 as 8 bytes.
//Behaviour is identical to PutFloat32(real) followed by PutFloat32(imag).
func PutComplex64(b []byte, v complex64) {
	r, i := real(v), imag(v)
	PutFloat32(b, r)
	PutFloat32(b[4:8], i)
}

//Complex64 deserializes complex64 from 8 bytes.
func Complex64(b []byte) complex64 {
	r := Float32(b)
	i := Float32(b[4:8])
	return complex(r, i)
}

//PutComplex128 serializes complex128 as 16 bytes.
//Behaviour is identical to PutFloat64(real) followed by PutFloat64(imag).
func PutComplex128(b []byte, v complex128) {
	r, i := real(v), imag(v)
	PutFloat64(b, r)
	PutFloat64(b[8:16], i)
}

//Complex128 deserializes complex128 from 16 bytes.
func Complex128(b []byte) complex128 {
	r := Float64(b)
	i := Float64(b[8:16])
	return complex(r, i)
}

//PutByte serializes a single byte.
func PutByte(b []byte, v byte) {
	b[0] = v
}

//Byte deserializes a single byte.
func Byte(b []byte) byte {
	return b[0]
}

//PutRune serializes rune as 4 bytes.
//Behaviour is identical to PutInt32.
func PutRune(b []byte, v rune) {
	PutInt32(b, int32(v))
}

//Rune deserializes rune from 4 bytes.
//Behaviour is identical to Int32.
func Rune(b []byte) rune {
	return rune(Int32(b))
}

//PutUint serializes uint as 8 bytes, regardless of architecture.
//Behaviour is identical to PutUint64.
func PutUint(b []byte, v uint) {
	PutUint64(b, uint64(v))
}

//Uint deserializes uint from 8 bytes, regardless of architecture.
//Behaviour is identical to Uint64.
func Uint(b []byte) uint {
	return uint(Uint64(b))
}

//PutInt serializes int as 8 bytes, regardless of architecture.
//Behaviour is identical to PutInt64.
func PutInt(b []byte, v int) {
	PutInt64(b, int64(v))
}

//Int deserializes int from 8 bytes, regardless of architecture.
//Behaviour is identical to Int64.
func Int(b []byte) int {
	return int(Int64(b))
}

//PutString serializes string as len(v) + 1 bytes.
//Behaviour is to append a single NUL character.
func PutString(b []byte, v string) {
	const nul = 0
	copy(b, []byte(v))
	b[len(v)] = nul
}

//String deserializes string from byte slice.
//Assumes the whole slice represents the value to deserialize.
func String(b []byte) string {
	l := len(b) - 1
	return string(b[:l])
}

//ScanString deserializes string from byte slice.
//Assumes that other values may be stored after the encoded string value.
//Prefer String if there are no other values stored in the slice.
func ScanString(b []byte) string {
	const nul = 0
	for i := 0; i < len(b); i++ {
		if b[i] == nul {
			return string(b[:i])
		}
	}
	return ""
}
