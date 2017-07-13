package lex

import (
	"bytes"
	"math"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	b1 := make([]byte, 1)
	PutBool(b1, true)
	assert.Equal(t, b1, []byte{1})
	assert.True(t, Bool(b1))

	b2 := make([]byte, 1)
	PutBool(b2, false)
	assert.Equal(t, b2, []byte{0})
	assert.False(t, Bool(b2))

	assert.Equal(t, -1, bytes.Compare(b2, b1))
	assert.Equal(t, 1, bytes.Compare(b1, b2))
}

func TestBool_ZeroAllocs(t *testing.T) {
	b := make([]byte, 1)
	assert.Zero(t, testing.AllocsPerRun(1, func() { PutBool(b, true) }))
	assert.Zero(t, testing.AllocsPerRun(1, func() { Bool(b) }))
}

//

func TestUint8(t *testing.T) {
	r := []uint8{0, 1, 42, math.MaxUint8}
	var prev []byte
	for _, v := range r {
		b1 := make([]byte, 1)
		PutUint8(b1, v)

		v1 := Uint8(b1)
		assert.Equal(t, v, v1)

		if prev != nil {
			assert.Equal(t, -1, bytes.Compare(prev, b1))
		}
		prev = b1
	}
}

func TestUint8_Range(t *testing.T) {
	var prev []byte
	for i := 0; i <= math.MaxUint8; i++ {
		v := uint8(i)

		b1 := make([]byte, 1)
		PutUint8(b1, v)

		v1 := Uint8(b1)
		assert.Equal(t, v, v1)

		if prev != nil {
			assert.Equal(t, -1, bytes.Compare(prev, b1))
		}
		prev = b1
	}
}

func TestUint8_ZeroAllocs(t *testing.T) {
	b := make([]byte, 1)
	assert.Zero(t, testing.AllocsPerRun(1, func() { PutUint8(b, 42) }))
	assert.Zero(t, testing.AllocsPerRun(1, func() { Uint8(b) }))
}

func TestUint16(t *testing.T) {
	r := []uint16{0, 1, 42, math.MaxUint16}
	var prev []byte
	for _, v := range r {
		b := make([]byte, 2)
		PutUint16(b, v)

		v1 := Uint16(b)
		assert.Equal(t, v, v1)

		if prev != nil {
			assert.Equal(t, -1, bytes.Compare(prev, b))
		}
		prev = b
	}
}

func TestUint16_Random(t *testing.T) {
	f := func(a1 uint16) bool {
		b1 := make([]byte, 2)
		PutUint16(b1, a1)
		v1 := Uint16(b1)
		return v1 == a1
	}
	assert.Nil(t, quick.Check(f, nil))
}

func TestUint16_RandomCompare(t *testing.T) {
	f := func(a1, a2 uint16) bool {
		b1 := make([]byte, 2)
		PutUint16(b1, a1)

		b2 := make([]byte, 2)
		PutUint16(b2, a2)

		var expected int
		switch {
		case a1 < a2:
			expected = -1
		case a1 > a2:
			expected = +1
		}
		return bytes.Compare(b1, b2) == expected
	}
	assert.Nil(t, quick.Check(f, nil))
}

func TestUint16_ZeroAllocs(t *testing.T) {
	b := make([]byte, 2)
	assert.Zero(t, testing.AllocsPerRun(1, func() { PutUint16(b, 42) }))
	assert.Zero(t, testing.AllocsPerRun(1, func() { Uint16(b) }))
}

func TestUint32(t *testing.T) {
	r := []uint32{0, 1, 42, math.MaxUint32}
	var prev []byte
	for _, v := range r {
		b := make([]byte, 4)
		PutUint32(b, v)

		v1 := Uint32(b)
		assert.Equal(t, v, v1)

		if prev != nil {
			assert.Equal(t, -1, bytes.Compare(prev, b))
		}
		prev = b
	}
}

func TestUint32_Random(t *testing.T) {
	f := func(a1 uint32) bool {
		b1 := make([]byte, 4)
		PutUint32(b1, a1)
		v1 := Uint32(b1)
		return v1 == a1
	}
	assert.Nil(t, quick.Check(f, nil))
}

func TestUint32_RandomCompare(t *testing.T) {
	f := func(a1, a2 uint32) bool {
		b1 := make([]byte, 4)
		PutUint32(b1, a1)

		b2 := make([]byte, 4)
		PutUint32(b2, a2)

		var expected int
		switch {
		case a1 < a2:
			expected = -1
		case a1 > a2:
			expected = +1
		}
		return bytes.Compare(b1, b2) == expected
	}
	assert.Nil(t, quick.Check(f, nil))
}

func TestUint32_ZeroAllocs(t *testing.T) {
	b := make([]byte, 4)
	assert.Zero(t, testing.AllocsPerRun(1, func() { PutUint32(b, 42) }))
	assert.Zero(t, testing.AllocsPerRun(1, func() { Uint32(b) }))
}

func TestUint64(t *testing.T) {
	r := []uint64{0, 1, 42, math.MaxUint64}
	var prev []byte
	for _, v := range r {
		b := make([]byte, 8)
		PutUint64(b, v)

		v1 := Uint64(b)
		assert.Equal(t, v, v1)

		if prev != nil {
			assert.Equal(t, -1, bytes.Compare(prev, b))
		}
		prev = b
	}
}

func TestUint64_Random(t *testing.T) {
	f := func(a1 uint64) bool {
		b1 := make([]byte, 8)
		PutUint64(b1, a1)
		v1 := Uint64(b1)
		return v1 == a1
	}
	assert.Nil(t, quick.Check(f, nil))
}

func TestUint64_RandomCompare(t *testing.T) {
	f := func(a1, a2 uint64) bool {
		b1 := make([]byte, 8)
		PutUint64(b1, a1)

		b2 := make([]byte, 8)
		PutUint64(b2, a2)

		var expected int
		switch {
		case a1 < a2:
			expected = -1
		case a1 > a2:
			expected = +1
		}
		return bytes.Compare(b1, b2) == expected
	}
	assert.Nil(t, quick.Check(f, nil))
}

func TestUint64_ZeroAllocs(t *testing.T) {
	b := make([]byte, 8)
	assert.Zero(t, testing.AllocsPerRun(1, func() { PutUint64(b, 42) }))
	assert.Zero(t, testing.AllocsPerRun(1, func() { Uint64(b) }))
}

//

func TestInt8(t *testing.T) {
	r := []int8{math.MinInt8, -42, -1, 0, 1, 42, math.MaxInt8}
	var prev []byte
	for _, v := range r {
		b1 := make([]byte, 1)
		PutInt8(b1, v)

		v1 := Int8(b1)
		assert.Equal(t, v, v1)

		if prev != nil {
			assert.Equal(t, -1, bytes.Compare(prev, b1))
		}
		prev = b1
	}
}

func TestInt8_Range(t *testing.T) {
	var prev []byte
	for i := math.MinInt8; i <= math.MaxInt8; i++ {
		v := int8(i)

		b1 := make([]byte, 1)
		PutInt8(b1, v)

		v1 := Int8(b1)
		assert.Equal(t, v, v1)

		if prev != nil {
			assert.Equal(t, -1, bytes.Compare(prev, b1))
		}
		prev = b1
	}
}

func TestInt8_ZeroAllocs(t *testing.T) {
	b := make([]byte, 1)
	assert.Zero(t, testing.AllocsPerRun(1, func() { PutInt8(b, 42) }))
	assert.Zero(t, testing.AllocsPerRun(1, func() { Int8(b) }))
}

func TestInt16(t *testing.T) {
	r := []int16{math.MinInt16, -42, -1, 0, 1, 42, math.MaxInt16}
	var prev []byte
	for _, v := range r {
		b := make([]byte, 2)
		PutInt16(b, v)

		v1 := Int16(b)
		assert.Equal(t, v, v1)

		if prev != nil {
			assert.Equal(t, -1, bytes.Compare(prev, b))
		}
		prev = b
	}
}

func TestInt16_Random(t *testing.T) {
	f := func(a1 int16) bool {
		b1 := make([]byte, 2)
		PutInt16(b1, a1)
		v1 := Int16(b1)
		return v1 == a1
	}
	assert.Nil(t, quick.Check(f, nil))
}

func TestInt16_RandomCompare(t *testing.T) {
	f := func(a1, a2 int16) bool {
		b1 := make([]byte, 2)
		PutInt16(b1, a1)

		b2 := make([]byte, 2)
		PutInt16(b2, a2)

		var expected int
		switch {
		case a1 < a2:
			expected = -1
		case a1 > a2:
			expected = +1
		}
		return bytes.Compare(b1, b2) == expected
	}
	assert.Nil(t, quick.Check(f, nil))
}

func TestInt16_ZeroAllocs(t *testing.T) {
	b := make([]byte, 2)
	assert.Zero(t, testing.AllocsPerRun(1, func() { PutInt16(b, 42) }))
	assert.Zero(t, testing.AllocsPerRun(1, func() { Int16(b) }))
}

func TestInt32(t *testing.T) {
	r := []int32{math.MinInt32, -42, -1, 0, 1, 42, math.MaxInt32}
	var prev []byte
	for _, v := range r {
		b := make([]byte, 4)
		PutInt32(b, v)

		v1 := Int32(b)
		assert.Equal(t, v, v1)

		if prev != nil {
			assert.Equal(t, -1, bytes.Compare(prev, b))
		}
		prev = b
	}
}

func TestInt32_Random(t *testing.T) {
	f := func(a1 int32) bool {
		b1 := make([]byte, 4)
		PutInt32(b1, a1)
		v1 := Int32(b1)
		return v1 == a1
	}
	assert.Nil(t, quick.Check(f, nil))
}

func TestInt32_RandomCompare(t *testing.T) {
	f := func(a1, a2 int32) bool {
		b1 := make([]byte, 4)
		PutInt32(b1, a1)

		b2 := make([]byte, 4)
		PutInt32(b2, a2)

		var expected int
		switch {
		case a1 < a2:
			expected = -1
		case a1 > a2:
			expected = +1
		}
		return bytes.Compare(b1, b2) == expected
	}
	assert.Nil(t, quick.Check(f, nil))
}

func TestInt32_ZeroAllocs(t *testing.T) {
	b := make([]byte, 4)
	assert.Zero(t, testing.AllocsPerRun(1, func() { PutInt32(b, 42) }))
	assert.Zero(t, testing.AllocsPerRun(1, func() { Int32(b) }))
}

func TestInt64(t *testing.T) {
	r := []int64{math.MinInt64, -42, -1, 0, 1, 42, math.MaxInt64}
	var prev []byte
	for _, v := range r {
		b := make([]byte, 8)
		PutInt64(b, v)

		v1 := Int64(b)
		assert.Equal(t, v, v1)

		if prev != nil {
			assert.Equal(t, -1, bytes.Compare(prev, b))
		}
		prev = b
	}
}

func TestInt64_Random(t *testing.T) {
	f := func(a1 int64) bool {
		b1 := make([]byte, 8)
		PutInt64(b1, a1)
		v1 := Int64(b1)
		return v1 == a1
	}
	assert.Nil(t, quick.Check(f, nil))
}

func TestInt64_RandomCompare(t *testing.T) {
	f := func(a1, a2 int64) bool {
		b1 := make([]byte, 8)
		PutInt64(b1, a1)

		b2 := make([]byte, 8)
		PutInt64(b2, a2)

		var expected int
		switch {
		case a1 < a2:
			expected = -1
		case a1 > a2:
			expected = +1
		}
		return bytes.Compare(b1, b2) == expected
	}
	assert.Nil(t, quick.Check(f, nil))
}

func TestInt64_ZeroAllocs(t *testing.T) {
	b := make([]byte, 8)
	assert.Zero(t, testing.AllocsPerRun(1, func() { PutInt64(b, 42) }))
	assert.Zero(t, testing.AllocsPerRun(1, func() { Int64(b) }))
}

//

func TestFloat32(t *testing.T) {
	r := []float32{-math.MaxFloat32, -42.999, -1.5, -math.SmallestNonzeroFloat32, 0, math.SmallestNonzeroFloat32, 1.5, 42.999, math.MaxFloat32, float32(math.Inf(1))}

	var prev []byte
	for _, v := range r {
		b := make([]byte, 4)
		PutFloat32(b, v)

		v1 := Float32(b)
		assert.Equal(t, v, v1)

		if prev != nil {
			assert.Equal(t, -1, bytes.Compare(prev, b))
		}
		prev = b
	}
}

func TestFloat32_Zero(t *testing.T) {
	nzero := make([]byte, 4)
	PutFloat32(nzero, -0.0)

	pzero := make([]byte, 4)
	PutFloat32(pzero, +0.0)

	assert.True(t, -0.0 == +0.0)
	assert.True(t, bytes.Equal(nzero, pzero)) //Positive and negative zero are treated as equal
}

func TestFloat32_Inf(t *testing.T) {
	ninf := make([]byte, 4)
	PutFloat32(ninf, float32(math.Inf(0)))

	pinf := make([]byte, 4)
	PutFloat32(pinf, float32(math.Inf(1)))

	assert.True(t, math.Inf(0) == math.Inf(1))
	assert.True(t, bytes.Equal(ninf, pinf)) //Positive and negative infinity are treated as equal
}

func TestFloat32_NaN(t *testing.T) {
	nan := make([]byte, 4)
	PutFloat32(nan, float32(math.NaN()))

	pinf := make([]byte, 4)
	PutFloat32(pinf, float32(math.Inf(1)))

	assert.Equal(t, -1, bytes.Compare(pinf, nan)) //NaN sorts after positive infinity
}

func TestFloat32_Random(t *testing.T) {
	f := func(a1 float32) bool {
		b1 := make([]byte, 4)
		PutFloat32(b1, a1)
		v1 := Float32(b1)
		return v1 == a1
	}
	assert.Nil(t, quick.Check(f, nil))
}

func TestFloat32_RandomCompare(t *testing.T) {
	f := func(a1, a2 float32) bool {
		if math.IsNaN(float64(a1)) || math.IsNaN(float64(a2)) {
			return true //skip if NaN
		}

		b1 := make([]byte, 4)
		PutFloat32(b1, a1)

		b2 := make([]byte, 4)
		PutFloat32(b2, a2)

		var expected int
		switch {
		case a1 < a2:
			expected = -1
		case a1 > a2:
			expected = +1
		}
		return bytes.Compare(b1, b2) == expected
	}
	assert.Nil(t, quick.Check(f, nil))
}

func TestFloat32_ZeroAllocs(t *testing.T) {
	b := make([]byte, 4)
	assert.Zero(t, testing.AllocsPerRun(1, func() { PutFloat32(b, 42) }))
	assert.Zero(t, testing.AllocsPerRun(1, func() { Float32(b) }))
}

func TestFloat64(t *testing.T) {
	r := []float64{-math.MaxFloat64, -42.999, -1.5, -math.SmallestNonzeroFloat64, 0, math.SmallestNonzeroFloat64, 1.5, 42.999, math.MaxFloat64, math.Inf(1)}
	var prev []byte
	for _, v := range r {
		b := make([]byte, 8)
		PutFloat64(b, v)

		v1 := Float64(b)
		assert.Equal(t, v, v1)

		if prev != nil {
			assert.Equal(t, -1, bytes.Compare(prev, b))
		}
		prev = b
	}
}

func TestFloat64_Zero(t *testing.T) {
	nzero := make([]byte, 8)
	PutFloat64(nzero, -0.0)

	pzero := make([]byte, 8)
	PutFloat64(pzero, +0.0)

	assert.True(t, -0.0 == +0.0)
	assert.True(t, bytes.Equal(nzero, pzero)) //Positive and negative zero are treated as equal
}

func TestFloat64_Inf(t *testing.T) {
	ninf := make([]byte, 8)
	PutFloat64(ninf, math.Inf(0))

	pinf := make([]byte, 8)
	PutFloat64(pinf, math.Inf(1))

	assert.True(t, math.Inf(0) == math.Inf(1))
	assert.True(t, bytes.Equal(ninf, pinf)) //Positive and negative infinity are treated as equal
}

func TestFloat64_NaN(t *testing.T) {
	nan := make([]byte, 8)
	PutFloat64(nan, math.NaN())

	pinf := make([]byte, 8)
	PutFloat64(pinf, math.Inf(1))

	assert.Equal(t, -1, bytes.Compare(pinf, nan)) //NaN sorts after positive infinity
}

func TestFloat64_Random(t *testing.T) {
	f := func(a1 float64) bool {
		b1 := make([]byte, 8)
		PutFloat64(b1, a1)
		v1 := Float64(b1)
		return v1 == a1
	}
	assert.Nil(t, quick.Check(f, nil))
}

func TestFloat64_RandomCompare(t *testing.T) {
	f := func(a1, a2 float64) bool {
		if math.IsNaN(a1) || math.IsNaN(a2) {
			return true //skip if NaN
		}

		b1 := make([]byte, 8)
		PutFloat64(b1, a1)

		b2 := make([]byte, 8)
		PutFloat64(b2, a2)

		var expected int
		switch {
		case a1 < a2:
			expected = -1
		case a1 > a2:
			expected = +1
		}
		return bytes.Compare(b1, b2) == expected
	}
	assert.Nil(t, quick.Check(f, nil))
}

func TestFloat64_ZeroAllocs(t *testing.T) {
	b := make([]byte, 8)
	assert.Zero(t, testing.AllocsPerRun(1, func() { PutFloat64(b, 42) }))
	assert.Zero(t, testing.AllocsPerRun(1, func() { Float64(b) }))
}

//

func TestComplex64(t *testing.T) {
	v := complex(float32(1.2), float32(2.4))

	b := make([]byte, 8)
	PutComplex64(b, v)

	v1 := Complex64(b)
	assert.Equal(t, v, v1)
}

func TestComplex64_ZeroAllocs(t *testing.T) {
	v := complex(float32(1.2), float32(2.4))
	b := make([]byte, 8)
	assert.Zero(t, testing.AllocsPerRun(1, func() { PutComplex64(b, v) }))
	assert.Zero(t, testing.AllocsPerRun(1, func() { Complex64(b) }))
}

func TestComplex128(t *testing.T) {
	v := complex(float64(1.2), float64(2.4))

	b := make([]byte, 16)
	PutComplex128(b, v)

	v1 := Complex128(b)
	assert.Equal(t, v, v1)
}

func TestComplex128_ZeroAllocs(t *testing.T) {
	v := complex(float64(1.2), float64(2.4))
	b := make([]byte, 16)
	assert.Zero(t, testing.AllocsPerRun(1, func() { PutComplex128(b, v) }))
	assert.Zero(t, testing.AllocsPerRun(1, func() { Complex128(b) }))
}

//

func TestByte(t *testing.T) {
	r := []byte{0, 1, 42, math.MaxUint8}
	var prev []byte
	for _, v := range r {
		b := make([]byte, 1)
		PutByte(b, v)

		v1 := Byte(b)
		assert.Equal(t, v, v1)

		if prev != nil {
			assert.Equal(t, -1, bytes.Compare(prev, b))
		}
		prev = b
	}
}

func TestByte_Range(t *testing.T) {
	var prev []byte
	for i := 0; i <= math.MaxUint8; i++ {
		v := byte(i)

		b1 := make([]byte, 1)
		PutByte(b1, v)

		v1 := Byte(b1)
		assert.Equal(t, v, v1)

		if prev != nil {
			assert.Equal(t, -1, bytes.Compare(prev, b1))
		}
		prev = b1
	}
}

func TestByte_ZeroAllocs(t *testing.T) {
	b := make([]byte, 1)
	assert.Zero(t, testing.AllocsPerRun(1, func() { PutByte(b, 42) }))
	assert.Zero(t, testing.AllocsPerRun(1, func() { Byte(b) }))
}

//

func TestRune(t *testing.T) {
	r := []rune{math.MinInt32, -42, -1, 0, 1, 42, math.MaxInt32}
	var prev []byte
	for _, v := range r {
		b := make([]byte, 4)
		PutRune(b, v)
		v1 := Rune(b)
		assert.Equal(t, v, v1)

		if prev != nil {
			assert.Equal(t, -1, bytes.Compare(prev, b))
		}
		prev = b
	}
}

func TestRune_Random(t *testing.T) {
	f := func(a1 rune) bool {
		b1 := make([]byte, 4)
		PutRune(b1, a1)
		v1 := Rune(b1)
		return v1 == a1
	}
	assert.Nil(t, quick.Check(f, nil))
}

func TestRune_RandomCompare(t *testing.T) {
	f := func(a1, a2 rune) bool {
		b1 := make([]byte, 4)
		PutRune(b1, a1)

		b2 := make([]byte, 4)
		PutRune(b2, a2)

		var expected int
		switch {
		case a1 < a2:
			expected = -1
		case a1 > a2:
			expected = +1
		}
		return bytes.Compare(b1, b2) == expected
	}
	assert.Nil(t, quick.Check(f, nil))
}

func TestRune_ZeroAllocs(t *testing.T) {
	b := make([]byte, 4)
	assert.Zero(t, testing.AllocsPerRun(1, func() { PutRune(b, 42) }))
	assert.Zero(t, testing.AllocsPerRun(1, func() { Rune(b) }))
}

//

func TestInt(t *testing.T) {
	i := int(1)

	b := make([]byte, 8)
	PutInt(b, i)
	assert.Equal(t, byte(1<<7), b[0])
	assert.Equal(t, byte(1<<0), b[7])

	v := Int(b)
	assert.Equal(t, i, v)
}

func TestInt_ZeroAllocs(t *testing.T) {
	b := make([]byte, 8)
	assert.Zero(t, testing.AllocsPerRun(1, func() { PutInt(b, 42) }))
	assert.Zero(t, testing.AllocsPerRun(1, func() { Int(b) }))
}

func TestUint(t *testing.T) {
	i := uint(1)

	b := make([]byte, 8)
	PutUint(b, i)
	assert.Equal(t, byte(1<<0), b[7])

	v := Uint(b)
	assert.Equal(t, i, v)
}

func TestUint_ZeroAllocs(t *testing.T) {
	b := make([]byte, 8)
	assert.Zero(t, testing.AllocsPerRun(1, func() { PutUint(b, 42) }))
	assert.Zero(t, testing.AllocsPerRun(1, func() { Uint(b) }))
}

//

func TestString(t *testing.T) {
	a1 := "the quick brown fox jumped over the lazy dog"
	b1 := make([]byte, len(a1)+1)
	PutString(b1, a1)
	assert.Equal(t, []byte(a1), b1[:len(b1)-1])
	assert.Equal(t, byte(0), byte(b1[len(b1)-1]))

	v1 := String(b1)
	assert.Equal(t, a1, v1)
}

func TestString_ZeroAllocs(t *testing.T) {
	v := "jumped over the lazy dog"
	b := make([]byte, len(v)+1)
	assert.Zero(t, testing.AllocsPerRun(1, func() { PutString(b, v) }))
	assert.Zero(t, testing.AllocsPerRun(1, func() { String(b) }))

	//Note that while the String function requires zero allocations,
	//the decoded string can still escape...
	var x string
	assert.Equal(t, 1.0, testing.AllocsPerRun(1, func() {
		s := String(b)
		x = s
	}))
	assert.Equal(t, v, x)
}

//

func TestScanString(t *testing.T) {
	v1, v2, v3 := "jumped over the lazy dog", 42, float32(12.1)
	slen := len(v1) + 1
	b := make([]byte, slen+8+4)

	PutString(b, v1)
	PutInt(b[slen:], v2)
	PutFloat32(b[slen+8:], v3)

	s := ScanString(b)
	assert.Equal(t, v1, s)
}

func TestScanString_empty(t *testing.T) {
	v1, v2, v3 := "", 42, float32(12.1)
	slen := len(v1) + 1
	b := make([]byte, slen+8+4)

	PutString(b, v1)
	PutInt(b[slen:], v2)
	PutFloat32(b[slen+8:], v3)

	s := ScanString(b)
	assert.Equal(t, v1, s)
}

func TestScanString_badformat(t *testing.T) {
	var b []byte
	s := ScanString(b) //no bytes at all
	assert.Equal(t, "", s)

	b = []byte("howdy")
	s = ScanString(b) //string but no trailing nul
	assert.Equal(t, "", s)
}
