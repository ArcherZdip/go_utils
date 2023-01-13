// @Author zhanglingyu
package string

import (
	"testing"

	"github.com/archerzdip/go_utils/cast"

	"github.com/archerzdip/go_utils/test"
)

// TestStudly
// @Description: TestStudly test.
// @param t
func TestStudly(t *testing.T) {
	test.Equal(t, "FooBarZAR", Studly("foo_bar_ZAR"))
}

//
// TestMd5
//  @Description: Md5 testing
//  @param t
//
func TestMd5(t *testing.T) {
	test.Equal(t, "e10adc3949ba59abbe56e057f20f883e", Md5("123456"))
}

//
// TestSha1
//  @Description: Sha1 testing
//  @param t
//
func TestSha1(t *testing.T) {
	test.Equal(t, "7c4a8d09ca3762af61e59520943dc26494f8941b", Sha1("123456"))
}

//
// TestCrc32
//  @Description: Crc32 testing
//  @param t
//
func TestCrc32(t *testing.T) {
	test.Equal(t, cast.ToUint32(158520161), Crc32("123456"))
}
