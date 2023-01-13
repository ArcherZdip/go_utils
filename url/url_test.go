// @Author zhanglingyu
package url

import (
	"testing"

	"github.com/archerzdip/go_utils/test"
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

// TestBase64Encode
// @Description: test Base64Encode
// @param t
func TestBase64Encode(t *testing.T) {
	test.Equal(t, "VGhpcyBpcyBhbiBlbmNvZGUgZGVtby4=", Base64Encode("This is an encode demo."))
}

// TestBase64Decode
// @Description: test Base64Decode
// @param t
func TestBase64Decode(t *testing.T) {
	rv, _ := Base64Decode("VGhpcyBpcyBhbiBlbmNvZGUgZGVtby4")
	test.Equal(t, "This is an encode demo.", rv)
}

// TestURLDecode
// @Description: test URLDecode
// @param t
func TestURLDecode(t *testing.T) {
	rv, _ := URLDecode("https%3A%2F%2Farcherzdip.github.io%2F%3Ffoo%3Dbar%26baz%3D20+22")
	test.Equal(t, "https://archerzdip.github.io/?foo=bar&baz=20 22", rv)
}

// TestURLEncode
// @Description: test URLEncode
// @param t
func TestURLEncode(t *testing.T) {
	test.Equal(t, "https%3A%2F%2Farcherzdip.github.io%2F%3Ffoo%3Dbar%26baz%3D20+22", URLEncode("https://archerzdip.github.io/?foo=bar&baz=20 22"))
}

// TestRawurlencode
// @Description: test Rawurlencode
// @param t
func TestRawurlencode(t *testing.T) {
	test.Equal(t, "https%3A%2F%2Farcherzdip.github.io%2F%3Ffoo%3Dbar%26baz%3D20%2022", Rawurlencode("https://archerzdip.github.io/?foo=bar&baz=20 22"))
}

// TestRawurldecode
// @Description: test Rawurldecode
// @param t
func TestRawurldecode(t *testing.T) {
	rv, _ := Rawurldecode("https%3A%2F%2Farcherzdip.github.io%2F%3Ffoo%3Dbar%26baz%3D20%2022")
	test.Equal(t, "https://archerzdip.github.io/?foo=bar&baz=20 22", rv)
}
