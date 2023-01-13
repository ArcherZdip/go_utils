// @Author zhanglingyu
package cast

import (
	"errors"
	"fmt"
	"html/template"
	"reflect"
	"strconv"
)

var errNegativeNotAllowed = errors.New("unable to cast negative value")

// ToBoolE
// @Description: cast i interfalce to bool type.
// @param i
// @return bool
// @return error
func ToBoolE(i interface{}, def ...bool) (bool, error) {
	i = indirect(i)

	switch b := i.(type) {
	case bool:
		return b, nil
	case nil:
		return false, nil
	case int:
		if i.(int) != 0 {
			return true, nil
		}
		return false, nil
	case string:
		return strconv.ParseBool(i.(string))
	default:
		if len(def) > 0 {
			return def[0], nil
		}
		return false, fmt.Errorf("unable to cast %#v of type %T to bool", i, i)
	}
}

// ToIntE
// @Description: interface to int.
// @param i
// @return int
// @return error
func ToIntE(i interface{}, def ...int) (int, error) {
	i = indirect(i)
	switch t := i.(type) {
	case int:
		return t, nil
	case int8:
		return int(t), nil
	case int16:
		return int(t), nil
	case int32:
		return int(t), nil
	case int64:
		return int(t), nil
	case uint:
		return int(t), nil
	case uint8:
		return int(t), nil
	case uint16:
		return int(t), nil
	case uint32:
		return int(t), nil
	case uint64:
		return int(t), nil
	case float32:
		return int(t), nil
	case float64:
		return int(t), nil
	case string:
		v, err := strconv.ParseInt(t, 0, 0)
		if err == nil {
			return int(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		if len(def) > 0 {
			return def[0], nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	}
}

// ToInt8E
// @Description: interface to int8
// @param i
// @return int8
// @return error
func ToInt8E(i interface{}, def ...int8) (int8, error) {
	i = indirect(i)
	switch t := i.(type) {
	case int:
		return int8(t), nil
	case int8:
		return t, nil
	case int16:
		return int8(t), nil
	case int32:
		return int8(t), nil
	case int64:
		return int8(t), nil
	case uint:
		return int8(t), nil
	case uint8:
		return int8(t), nil
	case uint16:
		return int8(t), nil
	case uint32:
		return int8(t), nil
	case uint64:
		return int8(t), nil
	case float32:
		return int8(t), nil
	case float64:
		return int8(t), nil
	case string:
		v, err := strconv.ParseInt(t, 0, 0)
		if err == nil {
			return int8(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		if len(def) > 0 {
			return def[0], nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
	}
}

// ToInt16E
// @Description: interface to int16
// @param i
// @return int16
// @return error
func ToInt16E(i interface{}, def ...int16) (int16, error) {
	i = indirect(i)
	switch t := i.(type) {
	case int:
		return int16(t), nil
	case int8:
		return int16(t), nil
	case int16:
		return t, nil
	case int32:
		return int16(t), nil
	case int64:
		return int16(t), nil
	case uint:
		return int16(t), nil
	case uint8:
		return int16(t), nil
	case uint16:
		return int16(t), nil
	case uint32:
		return int16(t), nil
	case uint64:
		return int16(t), nil
	case float32:
		return int16(t), nil
	case float64:
		return int16(t), nil
	case string:
		v, err := strconv.ParseInt(t, 0, 0)
		if err == nil {
			return int16(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		if len(def) > 0 {
			return def[0], nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
	}
}

// ToInt32E
// @Description: interface to int32
// @param i
// @return int32
// @return error
func ToInt32E(i interface{}, def ...int32) (int32, error) {
	i = indirect(i)
	switch t := i.(type) {
	case int:
		return int32(t), nil
	case int8:
		return int32(t), nil
	case int16:
		return int32(t), nil
	case int32:
		return t, nil
	case int64:
		return int32(t), nil
	case uint:
		return int32(t), nil
	case uint8:
		return int32(t), nil
	case uint16:
		return int32(t), nil
	case uint32:
		return int32(t), nil
	case uint64:
		return int32(t), nil
	case float32:
		return int32(t), nil
	case float64:
		return int32(t), nil
	case string:
		v, err := strconv.ParseInt(t, 0, 0)
		if err == nil {
			return int32(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		if len(def) > 0 {
			return def[0], nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
	}
}

// ToInt64E
// @Description: interface to int64.
// @param i
// @return int64
// @return error
func ToInt64E(i interface{}, def ...int64) (int64, error) {
	i = indirect(i)
	switch t := i.(type) {
	case int:
		return int64(t), nil
	case int8:
		return int64(t), nil
	case int16:
		return int64(t), nil
	case int32:
		return int64(t), nil
	case int64:
		return t, nil
	case uint:
		return int64(t), nil
	case uint8:
		return int64(t), nil
	case uint16:
		return int64(t), nil
	case uint32:
		return int64(t), nil
	case uint64:
		return int64(t), nil
	case float32:
		return int64(t), nil
	case float64:
		return int64(t), nil
	case string:
		v, err := strconv.ParseInt(t, 0, 0)
		if err == nil {
			return int64(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		if len(def) > 0 {
			return def[0], nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
	}
}

// ToUintE
// @Description: interface to uint.
// @param i
// @return uint
// @return error
func ToUintE(i interface{}, def ...uint) (uint, error) {
	i = indirect(i)
	switch t := i.(type) {
	case int:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(t), nil
	case int8:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(t), nil
	case int16:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(t), nil
	case int32:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(t), nil
	case int64:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(t), nil
	case uint:
		return t, nil
	case uint8:
		return uint(t), nil
	case uint16:
		return uint(t), nil
	case uint32:
		return uint(t), nil
	case uint64:
		return uint(t), nil
	case float32:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(t), nil
	case float64:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(t), nil
	case string:
		v, err := strconv.ParseUint(t, 0, 0)
		if err == nil {
			return uint(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint: %s", i, err)
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		if len(def) > 0 {
			return def[0], nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
	}
}

// ToUint8E
// @Description: interface to uint8.
// @param i
// @return uint8
// @return error
func ToUint8E(i interface{}, def ...uint8) (uint8, error) {
	i = indirect(i)
	switch t := i.(type) {
	case int:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(t), nil
	case int8:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(t), nil
	case int16:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(t), nil
	case int32:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(t), nil
	case int64:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(t), nil
	case uint:
		return uint8(t), nil
	case uint8:
		return t, nil
	case uint16:
		return uint8(t), nil
	case uint32:
		return uint8(t), nil
	case uint64:
		return uint8(t), nil
	case float32:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(t), nil
	case float64:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(t), nil
	case string:
		v, err := strconv.ParseUint(t, 0, 0)
		if err == nil {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint8: %s", i, err)
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		if len(def) > 0 {
			return def[0], nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
	}
}

// ToUint16E
// @Description: interface to uint16.
// @param i
// @return uint16
// @return error
func ToUint16E(i interface{}, def ...uint16) (uint16, error) {
	i = indirect(i)
	switch t := i.(type) {
	case int:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(t), nil
	case int8:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(t), nil
	case int16:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(t), nil
	case int32:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(t), nil
	case int64:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(t), nil
	case uint:
		return uint16(t), nil
	case uint8:
		return uint16(t), nil
	case uint16:
		return t, nil
	case uint32:
		return uint16(t), nil
	case uint64:
		return uint16(t), nil
	case float32:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(t), nil
	case float64:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(t), nil
	case string:
		v, err := strconv.ParseUint(t, 0, 0)
		if err == nil {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint16: %s", i, err)
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		if len(def) > 0 {
			return def[0], nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint16", i, i)
	}
}

// ToUint32E
// @Description: interface to uint32.
// @param i
// @return uint32
// @return error
func ToUint32E(i interface{}, def ...uint32) (uint32, error) {
	i = indirect(i)
	switch t := i.(type) {
	case int:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(t), nil
	case int8:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(t), nil
	case int16:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(t), nil
	case int32:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(t), nil
	case int64:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(t), nil
	case uint:
		return uint32(t), nil
	case uint8:
		return uint32(t), nil
	case uint16:
		return uint32(t), nil
	case uint32:
		return t, nil
	case uint64:
		return uint32(t), nil
	case float32:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(t), nil
	case float64:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(t), nil
	case string:
		v, err := strconv.ParseUint(t, 0, 0)
		if err == nil {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint32: %s", i, err)
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		if len(def) > 0 {
			return def[0], nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
	}
}

// ToUint64E
// @Description: interface to uint64.
// @param i
// @return uint64
// @return error
func ToUint64E(i interface{}, def ...uint64) (uint64, error) {
	i = indirect(i)
	switch t := i.(type) {
	case int:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(t), nil
	case int8:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(t), nil
	case int16:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(t), nil
	case int32:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(t), nil
	case int64:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(t), nil
	case uint:
		return uint64(t), nil
	case uint8:
		return uint64(t), nil
	case uint16:
		return uint64(t), nil
	case uint32:
		return uint64(t), nil
	case uint64:
		return t, nil
	case float32:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(t), nil
	case float64:
		if t < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(t), nil
	case string:
		v, err := strconv.ParseUint(t, 0, 0)
		if err == nil {
			return uint64(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint64: %s", i, err)
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		if len(def) > 0 {
			return def[0], nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
	}
}

// ToFloat32E
// @Description: interface to float32.
// @param i
// @return float32
// @return error
func Tofloat32E(i interface{}, def ...float32) (float32, error) {
	i = indirect(i)
	switch t := i.(type) {
	case int:
		return float32(t), nil
	case int8:
		return float32(t), nil
	case int16:
		return float32(t), nil
	case int32:
		return float32(t), nil
	case int64:
		return float32(t), nil
	case uint:
		return float32(t), nil
	case uint8:
		return float32(t), nil
	case uint16:
		return float32(t), nil
	case uint32:
		return float32(t), nil
	case uint64:
		return float32(t), nil
	case float32:
		return t, nil
	case float64:
		return float32(t), nil
	case string:
		v, err := strconv.ParseFloat(t, 32)
		if err == nil {
			return float32(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to float32", i, i)
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		if len(def) > 0 {
			return def[0], nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to float32", i, i)
	}
}

// Tofloat64E
// @Description: interface to float64.
// @param i
// @return float64
// @return error
func Tofloat64E(i interface{}, def ...float64) (float64, error) {
	i = indirect(i)
	switch t := i.(type) {
	case int:
		return float64(t), nil
	case int8:
		return float64(t), nil
	case int16:
		return float64(t), nil
	case int32:
		return float64(t), nil
	case int64:
		return float64(t), nil
	case uint:
		return float64(t), nil
	case uint8:
		return float64(t), nil
	case uint16:
		return float64(t), nil
	case uint32:
		return float64(t), nil
	case uint64:
		return float64(t), nil
	case float32:
		return float64(t), nil
	case float64:
		return t, nil
	case string:
		v, err := strconv.ParseFloat(t, 64)
		if err == nil {
			return float64(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", i, i)
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		if len(def) > 0 {
			return def[0], nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", i, i)
	}
}

// ToStringE
// @Description: interface to string.
// @param i
// @return string
// @return error
func ToStringE(i interface{}, def ...string) (string, error) {
	i = indirectToStringerOrError(i)

	switch s := i.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case int:
		return strconv.Itoa(s), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case int32:
		return strconv.Itoa(int(s)), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case uint:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(s), 10), nil
	case []byte:
		return string(s), nil
	case template.HTML:
		return string(s), nil
	case template.URL:
		return string(s), nil
	case template.JS:
		return string(s), nil
	case template.CSS:
		return string(s), nil
	case template.HTMLAttr:
		return string(s), nil
	case nil:
		return "", nil
	case fmt.Stringer:
		return s.String(), nil
	case error:
		return s.Error(), nil
	default:
		if len(def) > 0 {
			return def[0], nil
		}
		return "", fmt.Errorf("unable to cast %#v of type %T to string", i, i)
	}
}

// indirect
// @Description: returns the value, after dereferencing as many times
// @param i
// @return interface{}
func indirect(i interface{}) interface{} {
	if i == nil {
		return nil
	}

	if t := reflect.TypeOf(i); t.Kind() != reflect.Ptr {
		return i
	}

	v := reflect.ValueOf(i)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	return v.Interface()
}

// indirectToStringerOrError
// @Description:  returns the value, after dereferencing as many times as necessary to reach the base type (or nil) or an implementation of fmt.Stringer or error
// @param i
// @return interface{}
func indirectToStringerOrError(i interface{}) interface{} {
	if i == nil {
		return nil
	}
	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	v := reflect.ValueOf(i)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}
