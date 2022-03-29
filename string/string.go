// @Author zhanglingyu
package string

import "strings"

// Studly
// @Description: 下划线转驼峰
// @param str
// @return string
func Studly(str string) string {
	if str != "" {
		c := strings.Split(str, "_")
		var s string
		for _, v := range c {
			vv := []rune(v)
			if len(vv) > 0 {
				if vv[0] >= 'a' && vv[0] <= 'z' {
					vv[0] -= 32
				}
				s += string(vv)
			}
		}

		return s
	}

	return str
}
