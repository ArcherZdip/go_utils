// @Author zhanglingyu
package time

import (
	"go_utils/cast"
	"go_utils/test"
	"testing"
)

// TestTime
// @Description: test Time.
// @param t
func TestTime(t *testing.T) {
	test.Gt(t, cast.ToFloat64(Time()), 1648519083)
}

// TestStrtotime
// @Description: test Strtotime.
// @param t
func TestStrtotime(t *testing.T) {
	v, _ := Strtotime("2006-01-02 15:04:05", "2012-01-01 10:10:10")
	test.Equal(t, int64(1325383810), v)
}

// TestDate
// @Description: test Date.
// @param t
func TestDate(t *testing.T) {
	test.Equal(t, "2022-03-29 10:10:10", Date("2006-01-02 15:04:05", 1648519810))
}

// TestCheckdate
// @Description: test Checkdate
// @param t
func TestCheckdate(t *testing.T) {
	test.Equal(t, false, Checkdate(2, 29, 2018))
	test.Equal(t, true, Checkdate(2, 29, 2020))
}
