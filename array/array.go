// @Author zhanglingyu
package array

import "bytes"

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
