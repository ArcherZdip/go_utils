// @Author zhanglingyu
package array

import (
	"testing"

	"github.com/archerzdip/go_utils/test"
)

// TestImplode
// @Description: Implode testing
// @param t
func TestImplode(t *testing.T) {
	test.Equal(t, "1,a,张", Implode([]string{"1", "a", "张"}, ","))
}

//
// TestInArray
//  @Description: InArray testing
//  @param t
//
func TestInArray(t *testing.T) {
	test.Equal(t, true, InArray("a", []string{"a", "aa", "aaa"}))
	test.Equal(t, true, InArray(1, []int{1, 2, 3}))
	test.Equal(t, false, InArray(4, []int{1, 2, 3}))
}
