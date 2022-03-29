// @Author zhanglingyu
package url

import (
	"go_utils/test"
	"testing"
)

// TestHttpBuildQueryByMap
// @Description: Test TestHttpBuildQueryByMap
// @param t
func TestHttpBuildQueryByMap(t *testing.T) {
	test.Equal(t, "foo1=bar1&foo2=bar2", HttpBuildQueryByMap(map[string]interface{}{
		"foo1": "bar1",
		"foo2": "bar2",
	}))
}
