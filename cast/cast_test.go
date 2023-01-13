// @Author zhanglingyu
package cast

import (
	"html/template"
	"testing"

	"github.com/archerzdip/go_utils/test"
)

// TestToBoolE
// @Description: Test ToBoolE
// @param t
func TestToBoolE(t *testing.T) {
	tests := []struct {
		actual interface{}
		expect bool
		is_err bool
	}{
		{0, false, false},
		{nil, false, false},
		{"false", false, false},
		{"FALSE", false, false},
		{"False", false, false},
		{"f", false, false},
		{"F", false, false},
		{false, false, false},
		{"true", true, false},
		{"TRUE", true, false},
		{"True", true, false},
		{"t", true, false},
		{"T", true, false},
		{1, true, false},
		{true, true, false},
		{-1, true, false},

		// errors
		{"test", false, true},
		{testing.T{}, false, true},
	}

	// def
	rv, _ := ToBoolE(testing.T{}, false)
	test.Equal(t, rv, false)

	for _, vv := range tests {
		v, err := ToBoolE(vv.actual)
		if vv.is_err {
			t.Skip(err)
			continue
		}

		test.Equal(t, vv.expect, v)
	}
}

// TestToIntE
// @Description: test ToIntE
// @param t
func TestToIntE(t *testing.T) {
	tests := []struct {
		actual interface{}
		expect int
		is_err bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	// def
	rv, _ := ToIntE(testing.T{}, 123)
	test.Equal(t, rv, 123)

	for _, vv := range tests {
		v, err := ToIntE(vv.actual)
		if vv.is_err {
			t.Skip(err)
			continue
		}

		test.Equal(t, vv.expect, v)
	}
}

// TestToInt8E
// @Description: test ToInt8E
// @param t
func TestToInt8E(t *testing.T) {
	tests := []struct {
		actual interface{}
		expect int8
		is_err bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	// def
	rv, _ := ToInt8E(testing.T{}, 8)
	var actual int8 = 8
	test.Equal(t, rv, actual)

	for _, vv := range tests {
		v, err := ToInt8E(vv.actual)
		if vv.is_err {
			t.Skip(err)
			continue
		}

		test.Equal(t, vv.expect, v)
	}
}

// TestToInt16E
// @Description: test ToInt16E
// @param t
func TestToInt16E(t *testing.T) {
	tests := []struct {
		actual interface{}
		expect int16
		is_err bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	// def
	rv, _ := ToInt16E(testing.T{}, 8)
	var actual int16 = 8
	test.Equal(t, rv, actual)

	for _, vv := range tests {
		v, err := ToInt16E(vv.actual)
		if vv.is_err {
			t.Skip(err)
			continue
		}

		test.Equal(t, vv.expect, v)
	}
}

// TestToInt32E
// @Description: test ToInt32
// @param t
func TestToInt32E(t *testing.T) {
	tests := []struct {
		actual interface{}
		expect int32
		is_err bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for _, vv := range tests {
		v, err := ToInt32E(vv.actual)
		if vv.is_err {
			t.Skip(err)
			continue
		}

		test.Equal(t, vv.expect, v)
	}
}

// TestToInt64E
// @Description: test ToInt64E
// @param t
func TestToInt64E(t *testing.T) {
	tests := []struct {
		actual interface{}
		expect int64
		is_err bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for _, vv := range tests {
		v, err := ToInt64E(vv.actual)
		if vv.is_err {
			t.Skip(err)
			continue
		}

		test.Equal(t, vv.expect, v)
	}
}

// TestToUintE
// @Description: test ToUintE
// @param t
func TestToUintE(t *testing.T) {
	tests := []struct {
		actual interface{}
		expect uint
		is_err bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for _, vv := range tests {
		v, err := ToUintE(vv.actual)
		if vv.is_err {
			t.Skip(err)
			continue
		}

		test.Equal(t, vv.expect, v)
	}
}

// TestToUint8E
// @Description: test ToUint8E
// @param t
func TestToUint8E(t *testing.T) {
	tests := []struct {
		actual interface{}
		expect uint8
		is_err bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for _, vv := range tests {
		v, err := ToUint8E(vv.actual)
		if vv.is_err {
			t.Skip(err)
			continue
		}

		test.Equal(t, vv.expect, v)
	}
}

// TestToUint16E
// @Description: test ToUint16E
// @param t
func TestToUint16E(t *testing.T) {
	tests := []struct {
		actual interface{}
		expect uint16
		is_err bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for _, vv := range tests {
		v, err := ToUint16E(vv.actual)
		if vv.is_err {
			t.Skip(err)
			continue
		}

		test.Equal(t, vv.expect, v)
	}
}

// TestToUint32E
// @Description: test ToUint32E
// @param t
func TestToUint32E(t *testing.T) {
	tests := []struct {
		actual interface{}
		expect uint32
		is_err bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for _, vv := range tests {
		v, err := ToUint32E(vv.actual)
		if vv.is_err {
			t.Skip(err)
			continue
		}

		test.Equal(t, vv.expect, v)
	}
}

// TestToUint64E
// @Description: Test ToUint64E
// @param t
func TestToUint64E(t *testing.T) {
	tests := []struct {
		actual interface{}
		expect uint64
		is_err bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for _, vv := range tests {
		v, err := ToUint64E(vv.actual)
		if vv.is_err {
			t.Skip(err)
			continue
		}

		test.Equal(t, vv.expect, v)
	}
}

// TestTofloat32E
// @Description: test ToFloat32E.
// @param t
func TestTofloat32E(t *testing.T) {
	tests := []struct {
		actual interface{}
		expect float32
		is_err bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8.31, false},
		{float64(8.31), 8.31, false},
		{"8", 8, false},
		{true, 1, false},
		{false, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for _, vv := range tests {
		v, err := Tofloat32E(vv.actual)
		if vv.is_err {
			t.Skip(err)
			continue
		}

		test.Equal(t, vv.expect, v)
	}
}

// TestToFloat64E
// @Description: test ToFloat64E.
// @param t
func TestToFloat64E(t *testing.T) {
	tests := []struct {
		actual interface{}
		expect float64
		is_err bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8), 8, false},
		{float64(8.31), 8.31, false},
		{"8", 8, false},
		{true, 1, false},
		{false, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for _, vv := range tests {
		v, err := Tofloat64E(vv.actual)
		if vv.is_err {
			t.Skip(err)
			continue
		}

		test.Equal(t, vv.expect, v)
	}
}

// TestToStringE
// @Description: test ToStringE.
// @param t
func TestToStringE(t *testing.T) {
	type Key struct {
		k string
	}
	key := &Key{"foo"}

	tests := []struct {
		actual interface{}
		expect string
		is_err bool
	}{
		{int(8), "8", false},
		{int8(8), "8", false},
		{int16(8), "8", false},
		{int32(8), "8", false},
		{int64(8), "8", false},
		{uint(8), "8", false},
		{uint8(8), "8", false},
		{uint16(8), "8", false},
		{uint32(8), "8", false},
		{uint64(8), "8", false},
		{float32(8.31), "8.31", false},
		{float64(8.31), "8.31", false},
		{true, "true", false},
		{false, "false", false},
		{nil, "", false},
		{[]byte("one time"), "one time", false},
		{"one more time", "one more time", false},
		{template.HTML("one time"), "one time", false},
		{template.URL("http://somehost.foo"), "http://somehost.foo", false},
		{template.JS("(1+2)"), "(1+2)", false},
		{template.CSS("a"), "a", false},
		{template.HTMLAttr("a"), "a", false},
		// errors
		{testing.T{}, "", true},
		{key, "", true},
	}

	for _, vv := range tests {
		v, err := ToStringE(vv.actual)
		if vv.is_err {
			t.Skip(err)
			continue
		}

		test.Equal(t, vv.expect, v)
	}
}
