// @Author zhanglingyu
package cast

// ToBool
// @Description: interface to bool.
// @param i
// @return bool
func ToBool(i interface{}, def ...bool) bool {
	v, _ := ToBoolE(i, def...)
	return v
}

// ToInt
// @Description: interface to int.
// @param i
// @return int
func ToInt(i interface{}, def ...int) int {
	v, _ := ToIntE(i, def...)
	return v
}

// ToInt8
// @Description: interface to int8.
// @param i
// @return int8
func ToInt8(i interface{}, def ...int8) int8 {
	v, _ := ToInt8E(i, def...)
	return v
}

// ToInt16
// @Description: interface to int16.
// @param i
// @return int16
func ToInt16(i interface{}, def ...int16) int16 {
	v, _ := ToInt16E(i, def...)
	return v
}

// ToInt32
// @Description: interface to int32.
// @param i
// @return int32
func ToInt32(i interface{}, def ...int32) int32 {
	v, _ := ToInt32E(i, def...)
	return v
}

// ToInt64
// @Description: interface to int64.
// @param i
// @return int64
func ToInt64(i interface{}, def ...int64) int64 {
	v, _ := ToInt64E(i, def...)
	return v
}

// ToUint
// @Description: interface to uint.
// @param i
// @return uint
func ToUint(i interface{}, def ...uint) uint {
	v, _ := ToUintE(i, def...)
	return v
}

// ToUint8
// @Description: interface to uint8.
// @param i
// @return uint8
func ToUint8(i interface{}, def ...uint8) uint8 {
	v, _ := ToUint8E(i, def...)
	return v
}

// ToUint16
// @Description: interface to uint16.
// @param i
// @return uint16
func ToUint16(i interface{}, def ...uint16) uint16 {
	v, _ := ToUint16E(i, def...)
	return v
}

// ToUint32
// @Description: interface to uint32.
// @param i
// @return uint32
func ToUint32(i interface{}, def ...uint32) uint32 {
	v, _ := ToUint32E(i, def...)
	return v
}

// ToUint64
// @Description: interface to uint64.
// @param i
// @return uint64
func ToUint64(i interface{}, def ...uint64) uint64 {
	v, _ := ToUint64E(i, def...)
	return v
}

// ToFloat32
// @Description: interface to float32.
// @param i
// @return float32
func ToFloat32(i interface{}, def ...float32) float32 {
	v, _ := Tofloat32E(i, def...)
	return v
}

// ToFloat64
// @Description: interface to float64.
// @param i
// @return float64
func ToFloat64(i interface{}, def ...float64) float64 {
	v, _ := Tofloat64E(i, def...)
	return v
}

// ToString
// @Description: interface to string.
// @param i
// @return string
func ToString(i interface{}, def ...string) string {
	v, _ := ToStringE(i, def...)
	return v
}
