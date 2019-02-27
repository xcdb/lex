package lex_test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/xcdb/lex"

	"github.com/stretchr/testify/assert"
)

type aliasedInt int

type testStruct struct {
	A int
	B string
	C float32
}

type invalidStruct struct {
	A map[string]int
	b int
}

func TestPutReflect(t *testing.T) {
	var a1 int16 = 42
	var a2 float32 = 9.2

	expected := make([]byte, 6)
	lex.PutInt16(expected, a1)
	lex.PutFloat32(expected[2:], a2)

	actual := make([]byte, 6)
	lex.PutReflect(actual, a1)
	lex.PutReflect(actual[2:], a2)

	assert.True(t, bytes.Equal(expected, actual))
}

func TestPutReflect_alias(t *testing.T) {
	var a1 aliasedInt = 42

	expected := make([]byte, 8)
	lex.PutInt(expected, int(a1))

	actual := make([]byte, 8)
	lex.PutReflect(actual, a1)

	assert.True(t, bytes.Equal(expected, actual))
}

func TestPutReflect_struct(t *testing.T) {
	var a1 testStruct = testStruct{42, "hello", 12.5}

	expected := make([]byte, 18)
	lex.PutInt(expected, a1.A)
	lex.PutString(expected[8:], a1.B)
	lex.PutFloat32(expected[14:], a1.C)

	actual := make([]byte, 18)
	lex.PutReflect(actual, a1)

	assert.True(t, bytes.Equal(expected, actual))
}

func TestPutReflect_struct_invalid(t *testing.T) {
	var tests = []interface{}{
		invalidStruct{},
		struct{}{},
		time.Time{},
	}
	b := make([]byte, 16)
	for _, tt := range tests {
		err := lex.PutReflect(b, tt)
		assert.NotNil(t, err)
	}
}

//

func TestReflect(t *testing.T) {
	var expected1, actual1 int16 = 42, 0
	var expected2, actual2 float32 = 9.2, 0

	b := make([]byte, 6)
	lex.PutInt16(b, expected1)
	lex.PutFloat32(b[2:], expected2)

	lex.Reflect(b, &actual1)
	lex.Reflect(b[2:], &actual2)

	assert.Equal(t, expected1, actual1)
	assert.Equal(t, expected2, actual2)
}

func TestReflect_string(t *testing.T) {
	var expected, actual string = "howdy", ""

	b := make([]byte, 6)
	lex.PutString(b, expected)

	lex.Reflect(b, &actual)

	assert.Equal(t, expected, actual)
}

func TestReflect_alias(t *testing.T) {
	var expected, actual aliasedInt = 42, 0

	b := make([]byte, 8)
	lex.PutInt(b, int(expected))

	lex.Reflect(b, &actual)

	assert.Equal(t, expected, actual)
}

func TestReflect_struct(t *testing.T) {
	var expected testStruct = testStruct{42, "hello", 12.5}
	var actual testStruct

	b := make([]byte, 18)
	lex.PutInt(b, expected.A)
	lex.PutString(b[8:], expected.B)
	lex.PutFloat32(b[14:], expected.C)

	lex.Reflect(b, &actual)

	assert.Equal(t, expected, actual)
}

func TestReflect_struct_invalid(t *testing.T) {
	var tests = []interface{}{
		&invalidStruct{},
		&struct{}{},
		&time.Time{},
	}
	b := make([]byte, 16)
	lex.PutInt(b, 42) //just something to decode
	for _, tt := range tests {
		err := lex.Reflect(b, tt)
		assert.NotNil(t, err)
	}
}

func TestReflect_notptr(t *testing.T) {
	var v, actual int = 42, 0

	b := make([]byte, 8)
	lex.PutInt(b, v)

	lex.Reflect(b, actual)

	assert.Equal(t, 0, actual)
}

//

func TestSize(t *testing.T) {
	i := lex.Size(int32(42))
	assert.Equal(t, 4, i)
}

func TestSize_alias(t *testing.T) {
	i := lex.Size(aliasedInt(42))
	assert.Equal(t, 8, i)
}

func TestSize_struct(t *testing.T) {
	st := testStruct{42, "hello", 12.5}
	i := lex.Size(st)
	assert.Equal(t, 18, i)
}

func TestSize_struct_invalid(t *testing.T) {
	var tests = []interface{}{
		invalidStruct{},
		struct{}{},
		time.Time{},
	}
	for _, tt := range tests {
		i := lex.Size(tt)
		assert.Equal(t, -1, i)
	}
}

func ExampleSize() {
	var a1 int16 = 42
	var a2 float32 = 9.2

	i := lex.Size(a1)
	j := lex.Size(a2)

	fmt.Printf("%v\n%v", i, j)

	// Output:
	// 2
	// 4
}

//

func TestKey(t *testing.T) {
	var a1 int16 = 42
	var a2 float32 = 9.2

	expected := make([]byte, 6)
	lex.PutInt16(expected, a1)
	lex.PutFloat32(expected[2:], a2)

	actual, err := lex.Key(a1, a2)
	assert.Nil(t, err)
	assert.True(t, bytes.Equal(expected, actual))
}

func TestKey_alias(t *testing.T) {
	var a1 aliasedInt = 42

	expected := make([]byte, 8)
	lex.PutInt(expected, int(a1))

	actual, err := lex.Key(a1)
	assert.Nil(t, err)
	assert.True(t, bytes.Equal(expected, actual))
}

func TestKey_noargs(t *testing.T) {
	b, err := lex.Key()
	assert.Nil(t, b)
	assert.NotNil(t, err)
}

func TestKey_nil(t *testing.T) {
	b, err := lex.Key(nil)
	assert.Nil(t, b)
	assert.NotNil(t, err)
}

func TestKey_nilptr(t *testing.T) {
	var p *int = nil
	b, err := lex.Key(p)
	assert.Nil(t, b)
	assert.NotNil(t, err)
}

func TestKey_invalid(t *testing.T) {
	var m map[string]int
	b, err := lex.Key(m)
	assert.Nil(t, b)
	assert.NotNil(t, err)
}

func TestKey_invalidnilptr(t *testing.T) {
	var m map[string]int
	b, err := lex.Key(&m)
	assert.Nil(t, b)
	assert.NotNil(t, err)
}

func ExampleKey() {
	var a1 int16 = 42
	var a2 float32 = 9.2

	x, _ := lex.Key(a1)
	y, _ := lex.Key(a2)
	z, _ := lex.Key(a1, a2)

	fmt.Printf("%v %v\n%v", x, y, z)

	// Output:
	// [128 42] [193 19 51 51]
	// [128 42 193 19 51 51]
}

//

func TestMustKey_noargs(t *testing.T) {
	assert.Panics(t, func() {
		b := lex.MustKey()
		assert.Nil(t, b)
	})
}

func TestMustKey_nil(t *testing.T) {
	assert.Panics(t, func() {
		b := lex.MustKey(nil)
		assert.Nil(t, b)
	})
}

func TestMustKey_nilptr(t *testing.T) {
	assert.Panics(t, func() {
		var p *int = nil
		b := lex.MustKey(p)
		assert.Nil(t, b)
	})
}

func TestMustKey_invalid(t *testing.T) {
	assert.Panics(t, func() {
		var m map[string]int
		b := lex.MustKey(m)
		assert.Nil(t, b)
	})
}

func TestMustKey_invalidnilptr(t *testing.T) {
	assert.Panics(t, func() {
		var m map[string]int
		b := lex.MustKey(&m)
		assert.Nil(t, b)
	})
}

//

type mystring string
type mybool bool
type myint8 int8
type myint16 int16
type myint32 int32
type myint64 int64
type myuint8 uint8
type myuint16 uint16
type myuint32 uint32
type myuint64 uint64
type myfloat32 float32
type myfloat64 float64
type mycomplex64 complex64
type mycomplex128 complex128
type mybyte byte
type myrune rune
type myint int
type myuint uint

type mystruct1 struct {
	X int32
	Y int32
}
type mystruct2 struct {
	A float32
	B float32
}
type mystruct3 struct {
	F1 mystruct1
	F2 mystruct2
}

func TestSizeReflectPutReflect(t *testing.T) {
	v01 := string("howdy")
	v02 := bool(true)
	v03 := int8(2)
	v04 := int16(3)
	v05 := int32(5)
	v06 := int64(8)
	v07 := uint8(2)
	v08 := uint16(3)
	v09 := uint32(5)
	v10 := uint64(8)
	v11 := float32(23)
	v12 := float64(46)
	v13 := complex(float32(1), float32(2))
	v14 := complex(float64(1), float64(2))
	v15 := byte(24)
	v16 := rune(24)
	v17 := int(42)
	v18 := uint(42)

	a01 := mystring("howdy")
	a02 := mybool(true)
	a03 := myint8(2)
	a04 := myint16(3)
	a05 := myint32(5)
	a06 := myint64(8)
	a07 := myuint8(2)
	a08 := myuint16(3)
	a09 := myuint32(5)
	a10 := myuint64(8)
	a11 := myfloat32(23)
	a12 := myfloat64(46)
	a13 := mycomplex64(complex(float32(1), float32(2)))
	a14 := mycomplex128(complex(float64(1), float64(2)))
	a15 := mybyte(24)
	a16 := myrune(24)
	a17 := myint(42)
	a18 := myuint(42)

	s01 := mystruct1{24, 42}
	s02 := mystruct2{1.5, 5.0}
	s03 := mystruct3{
		mystruct1{1, 2},
		mystruct2{3.0, 4.0},
	}

	var tests = []struct {
		val, ptr interface{}
		size     int
	}{
		//builtin
		{v01, &v01, len(v01) + 1},
		{v02, &v02, 1},
		{v03, &v03, 1},
		{v04, &v04, 2},
		{v05, &v05, 4},
		{v06, &v06, 8},
		{v07, &v07, 1},
		{v08, &v08, 2},
		{v09, &v09, 4},
		{v10, &v10, 8},
		{v11, &v11, 4},
		{v12, &v12, 8},
		{v13, &v13, 8},
		{v14, &v14, 16},
		{v15, &v15, 1},
		{v16, &v16, 4},
		{v17, &v17, 8},
		{v18, &v18, 8},
		//aliased
		{a01, &a01, len(a01) + 1},
		{a02, &a02, 1},
		{a03, &a03, 1},
		{a04, &a04, 2},
		{a05, &a05, 4},
		{a06, &a06, 8},
		{a07, &a07, 1},
		{a08, &a08, 2},
		{a09, &a09, 4},
		{a10, &a10, 8},
		{a11, &a11, 4},
		{a12, &a12, 8},
		{a13, &a13, 8},
		{a14, &a14, 16},
		{a15, &a15, 1},
		{a16, &a16, 4},
		{a17, &a17, 8},
		{a18, &a18, 8},
		//structs
		{s01, &s01, 8},
		{s02, &s02, 8},
		{s03, &s03, 16},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.size, lex.Size(tt.val), "(%v) %v", reflect.ValueOf(tt.val).Kind(), tt.val)
		assert.Equal(t, tt.size, lex.Size(tt.ptr), "(%v ptr)", reflect.ValueOf(tt.val).Kind())

		b1 := make([]byte, tt.size)
		b2 := make([]byte, tt.size)
		lex.PutReflect(b1, tt.val)

		lex.Reflect(b2, tt.ptr)              //underneath, val == *ptr, so we need to change the value of ptr
		lex.PutReflect(b2, tt.ptr)           //we're cheating a little and just using Reflect/PutReflect as a means
		assert.False(t, bytes.Equal(b1, b2)) //to change the underlying value, avoiding a bunch of type checks

		lex.Reflect(b1, tt.ptr)
		lex.PutReflect(b2, tt.ptr)
		assert.True(t, bytes.Equal(b1, b2))
	}
}

func TestSizeReflectPutReflect_nil(t *testing.T) {
	b := make([]byte, 8)
	assert.Equal(t, -1, lex.Size(nil))
	lex.Reflect(b, nil)
	lex.PutReflect(b, nil)
}

func TestSizeReflectPutReflect_nilptr(t *testing.T) {
	b := make([]byte, 8)
	var p *int = nil
	assert.Equal(t, -1, lex.Size(p))
	lex.Reflect(b, p)
	lex.PutReflect(b, p)
}

func TestSizeReflectPutReflect_invalid(t *testing.T) {
	b := make([]byte, 8)
	var m map[string]int
	assert.Equal(t, -1, lex.Size(m))
	lex.Reflect(b, m)
	lex.PutReflect(b, m)
}

func TestSizeReflectPutReflect_invalidnilptr(t *testing.T) {
	b := make([]byte, 8)
	var m map[string]int
	assert.Equal(t, -1, lex.Size(&m))
	lex.Reflect(b, &m)
	lex.PutReflect(b, &m)
}

//

func BenchmarkSizeString(b *testing.B) {
	s := "hello world"
	for n := 0; n < b.N; n++ {
		lex.Size(s)
	}
}

func BenchmarkSizeInt(b *testing.B) {
	v := 64
	for n := 0; n < b.N; n++ {
		lex.Size(v)
	}
}

func BenchmarkPutReflectString(b *testing.B) {
	s := "hello world"
	bs := make([]byte, 12)
	for n := 0; n < b.N; n++ {
		lex.PutReflect(bs, s)
	}
}

func BenchmarkPutString(b *testing.B) {
	s := "hello world"
	bs := make([]byte, 12)
	for n := 0; n < b.N; n++ {
		lex.PutString(bs, s)
	}
}

func BenchmarkPutReflectInt(b *testing.B) {
	v := 64
	bs := make([]byte, 8)
	for n := 0; n < b.N; n++ {
		lex.PutReflect(bs, v)
	}
}

func BenchmarkPutInt(b *testing.B) {
	v := 64
	bs := make([]byte, 8)
	for n := 0; n < b.N; n++ {
		lex.PutInt(bs, v)
	}
}

func BenchmarkStdLibWriteInt(b *testing.B) {
	v := 64
	buf := &bytes.Buffer{}
	for n := 0; n < b.N; n++ {
		binary.Write(buf, binary.BigEndian, v)
	}
}

func BenchmarkStdLibPutUint64(b *testing.B) {
	v := uint64(42)
	bs := make([]byte, 8)
	for n := 0; n < b.N; n++ {
		binary.BigEndian.PutUint64(bs, v)
	}
}
