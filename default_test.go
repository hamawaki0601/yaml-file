package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	以下は対象外
	Uintptr
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
*/
type testType struct {
	Bool    bool    `default:"true"`
	Int     int     `default:"-1"`
	Int8    int8    `default:"-128"`
	Int16   int16   `default:"-32768"`
	Int32   int32   `default:"-2147483648"`
	Int64   int64   `default:"-9223372036854775808"`
	Uint    uint    `default:"1"`
	Uint8   uint8   `default:"255"`
	Uint16  uint16  `default:"65535"`
	Uint32  uint32  `default:"4294967295"`
	Uint64  uint64  `default:"18446744073709551615"`
	Float32 float32 `default:"0.1"`
	Float64 float64 `default:"0.01"`
	String  string  `default:"abc"`
}

func TestSetDefaultValue(t *testing.T) {
	target := testType{}
	err := SetDefaultValue(&target)
	assert.Nil(t, err)
	assert.Equal(t, true, target.Bool)
	assert.Equal(t, -1, target.Int)
	assert.Equal(t, int8(-128), target.Int8)
	assert.Equal(t, int16(-32768), target.Int16)
	assert.Equal(t, int32(-2147483648), target.Int32)
	assert.Equal(t, int64(-9223372036854775808), target.Int64)
	assert.Equal(t, uint(1), target.Uint)
	assert.Equal(t, uint8(255), target.Uint8)
	assert.Equal(t, uint16(65535), target.Uint16)
	assert.Equal(t, uint32(4294967295), target.Uint32)
	assert.Equal(t, uint64(18446744073709551615), target.Uint64)
	assert.Equal(t, float32(0.1), target.Float32)
	assert.Equal(t, 0.01, target.Float64)
	assert.Equal(t, "abc", target.String)
}
