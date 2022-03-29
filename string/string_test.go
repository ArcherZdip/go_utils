// @Author zhanglingyu
package string

import (
	"go_utils/test"
	"testing"
)

// TestStudly
// @Description: TestStudly test.
// @param t
func TestStudly(t *testing.T) {
	test.Equal(t, "FooBarZAR", Studly("foo_bar_ZAR"))
}
