// @Author zhanglingyu
package url

import "net/url"

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
