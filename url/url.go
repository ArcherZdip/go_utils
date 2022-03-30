// @Author zhanglingyu
package url

import (
	"encoding/base64"
	"net/url"
	"strings"
)

// HttpBuildQueryByMap
// @Description: Map to get query.
// @param querys
// @return string
func HttpBuildQueryByMap(querys map[string]interface{}) string {
	var uri url.URL
	q := uri.Query()
	for k, v := range querys {
		q.Add(k, v.(string))
	}

	return q.Encode()
}

// Base64Encode
// @Description: base64 encode.
// @param str
// @return string
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64Decode
// @Description: base64 decode.
// @param str
// @return string
// @return error
func Base64Decode(str string) (string, error) {
	switch len(str) % 4 {
	case 2:
		str += "=="
	case 3:
		str += "="
	}
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// URLEncode
// @Description: urlencode
// @param str
// @return string
func URLEncode(str string) string {
	return url.QueryEscape(str)
}

// URLDecode
// @Description: urldecode
// @param str
// @return string
// @return error
func URLDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}

// Rawurlencode
// @Description: rawurlencode
// @param str
// @return string
func Rawurlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

// Rawurldecode
// @Description: rawurldecode
// @param str
// @return string
// @return error
func Rawurldecode(str string) (string, error) {
	return url.QueryUnescape(strings.Replace(str, "%20", "+", -1))
}
