//Package lex implements an encoder that preserves lexicographical order for Boolean, Numeric and String types.
//
//Sorted key/value stores such as LMDB, Bolt and LevelDB support sequential iteration over the data. While custom comparison functions are possible, by default a bytewise comparison of keys is used and an appropriate encoding must be selected so that comparisons remain consistent.
//
//Strings are often used, as selecting formats compatible with bytewise comparison is trivial. Big-endian unsigned integers are similarly easy to use. Signed integers are a little harder to get right as, for example, -1 will normally sort after +1.
//
//Lex provides functions that allow the safe usage of many more types with the default bytewise comparison. Efficient implementations are provided for many core types, with structs and aliased types also supported via a reflection-based approach.
//
//Boolean and Numeric types are encoded as appropriate fixed-size values, while Strings are encoded simply as their underlying bytes with a single `NUL` character appended. Note that type information is *not* serialized with the value, and needs to be maintained separately.
package lex

import (
	"errors"
	"reflect"
)

//Size returns the number of bytes PutReflect would generate to encode the value d.
//Data must be of Boolean, Numeric or String based type, or a pointer to such data.
//If d is not of a supported type, Size returns -1.
func Size(d interface{}) int {
	return size(reflect.ValueOf(d))
}

func size(v reflect.Value) int {
	v = reflect.Indirect(v)
	switch v.Kind() {
	case reflect.String:
		return v.Len() + 1

	case reflect.Bool:
		return 1

	case reflect.Uint, reflect.Int:
		return 8

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128:
		return int(v.Type().Size())

	case reflect.Struct:
		sum := 0
		for i, n := 0, v.NumField(); i < n; i++ {
			s := size(v.Field(i))
			if s < 0 {
				return -1
			}
			sum += s
		}
		if sum == 0 {
			return -1
		}
		return sum
	}

	return -1
}

//PutReflect writes a lexicographically encoded representation of data into b.
//Data must be of Boolean, Numeric or String type, or a pointer to such data.
func PutReflect(b []byte, data interface{}) error {
	i := putReflect(b, reflect.ValueOf(data))
	if i < 0 {
		return errors.New("lex.PutReflect: invalid")
	}
	return nil
}

func putReflect(b []byte, v reflect.Value) int {
	v = reflect.Indirect(v)
	switch v.Kind() {
	case reflect.String:
		s := v.String()
		PutString(b, s)
		return len(s) + 1
	case reflect.Bool:
		PutBool(b, v.Bool())
		return 1
	case reflect.Int:
		PutInt(b, int(v.Int()))
		return 8
	case reflect.Uint:
		PutUint(b, uint(v.Uint()))
		return 8
	case reflect.Int8:
		PutInt8(b, int8(v.Int()))
	case reflect.Uint8:
		PutUint8(b, uint8(v.Uint()))
	case reflect.Int16:
		PutInt16(b, int16(v.Int()))
	case reflect.Uint16:
		PutUint16(b, uint16(v.Uint()))
	case reflect.Int32:
		PutInt32(b, int32(v.Int()))
	case reflect.Uint32:
		PutUint32(b, uint32(v.Uint()))
	case reflect.Int64:
		PutInt64(b, v.Int())
	case reflect.Uint64:
		PutUint64(b, v.Uint())
	case reflect.Float32:
		PutFloat32(b, float32(v.Float()))
	case reflect.Float64:
		PutFloat64(b, v.Float())
	case reflect.Complex64:
		PutComplex64(b, complex64(v.Complex()))
	case reflect.Complex128:
		PutComplex128(b, v.Complex())
	case reflect.Struct:
		sum := 0
		for i, n := 0, v.NumField(); i < n; i++ {
			s := putReflect(b[sum:], v.Field(i))
			if s < 0 {
				return -1
			}
			sum += s
		}
		if sum == 0 {
			return -1
		}
		return sum
	default:
		return -1
	}
	return int(v.Type().Size())
}

//Reflect reads lexicographically encoded data from b into data.
//Data must be a pointer to a Boolean, Numeric or String based type.
//When reading into a struct, all fields must be exported.
func Reflect(b []byte, data interface{}) error {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Ptr {
		return errors.New("lex.Reflect: invalid (data must be a pointer)")
	}

	v = v.Elem()

	//if data is string, then we can assume the whole slice is the string value
	//and avoid the much more expensive ScanString operation
	if v.Kind() == reflect.String {
		v.SetString(String(b))
		return nil
	}

	i := _reflect(b, v)
	if i < 0 {
		return errors.New("lex.Reflect: invalid")
	}
	return nil
}

func _reflect(b []byte, v reflect.Value) int {
	switch v.Kind() {
	case reflect.String:
		s := ScanString(b)
		v.SetString(s)
		return len(s) + 1
	case reflect.Bool:
		v.SetBool(Bool(b))
		return 1
	case reflect.Int:
		v.SetInt(int64(Int(b)))
		return 8
	case reflect.Uint:
		v.SetUint(uint64(Uint(b)))
		return 8
	case reflect.Int8:
		v.SetInt(int64(Int8(b)))
	case reflect.Uint8:
		v.SetUint(uint64(Uint8(b)))
	case reflect.Int16:
		v.SetInt(int64(Int16(b)))
	case reflect.Uint16:
		v.SetUint(uint64(Uint16(b)))
	case reflect.Int32:
		v.SetInt(int64(Int32(b)))
	case reflect.Uint32:
		v.SetUint(uint64(Uint32(b)))
	case reflect.Int64:
		v.SetInt(Int64(b))
	case reflect.Uint64:
		v.SetUint(Uint64(b))
	case reflect.Float32:
		v.SetFloat(float64(Float32(b)))
	case reflect.Float64:
		v.SetFloat(Float64(b))
	case reflect.Complex64:
		v.SetComplex(complex128(Complex64(b)))
	case reflect.Complex128:
		v.SetComplex(Complex128(b))
	case reflect.Struct:
		sum := 0
		for i, n := 0, v.NumField(); i < n; i++ {
			if f := v.Field(i); f.CanSet() {
				s := _reflect(b[sum:], f)
				if s < 0 {
					return -1
				}
				sum += s
			}
		}
		if sum == 0 {
			return -1
		}
		return sum
	default:
		return -1
	}
	return int(v.Type().Size())
}

//Key creates an appropriately-sized slice and writes passed data to it.
func Key(data ...interface{}) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("lex.Key: no data")
	}

	sum := 0
	for _, d := range data {
		n := Size(d)
		if n < 0 {
			return nil, errors.New("lex.Key: invalid")
		}
		sum += n
	}

	b := make([]byte, sum)
	offset := 0

	for _, d := range data {
		bs := b[offset:]
		PutReflect(bs, d)
		offset += Size(d)
	}

	return b, nil
}
