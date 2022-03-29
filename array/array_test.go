// @Author zhanglingyu
package array

import (
	"go_utils/test"
	"testing"
)

// TestImplode
// @Description: Implode testing
// @param t
func TestImplode(t *testing.T) {
	test.Equal(t, "1,a,张", Implode([]string{"1", "a", "张"}, ","))
}
