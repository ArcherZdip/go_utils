// @Author zhanglingyu
package array

import (
	"bytes"
	"reflect"
)

// Implode
// @Description: 以指定字符分隔数组
// @param pieces 数组
// @param separator 分隔符
// @return string
func Implode(pieces []string, separator string) string {
	var buf bytes.Buffer
	l := len(pieces)
	for _, str := range pieces {
		buf.WriteString(str)
		if l--; l > 0 {
			buf.WriteString(separator)
		}
	}

	return buf.String()
}

//
// InArray
//  @Description: in_array haystack supported types: slice, array or map
//  @param needle
//  @param haystack
//  @return bool
//
func InArray(needle interface{}, haystack interface{}) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	default:
		panic("haystack: haystack type muset be slice, array or map")
	}

	return false
}
