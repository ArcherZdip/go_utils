// @Author zhanglingyu
package time

import "time"

// Time
// @Description: timestamp
// @return int64
func Time() int64 {
	return time.Now().Unix()
}

// Strtotime
// @Description: time string to int64.
// @param format
// @param strtime
// @return int64
// @return error
func Strtotime(format, strtime string) (int64, error) {
	t, err := time.ParseInLocation(format, strtime, time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// Date
// Date("02/01/2006 15:04:05 PM", 1524799394)
// @Description: time int64 to string.
// @param format
// @param timestamp
// @return string
func Date(format string, timestamp int64) string {
	return time.Unix(timestamp, 0).Format(format)
}

// Checkdate
// @Description: Validate a Gregorian date
// @param month
// @param day
// @param year
// @return bool
func Checkdate(month, day, year int) bool {
	if month < 1 || month > 12 || day < 1 || day > 31 || year < 1 || year > 32767 {
		return false
	}
	switch month {
	case 4, 6, 9, 11:
		if day > 30 {
			return false
		}
	case 2:
		// leap year
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			if day > 29 {
				return false
			}
		} else if day > 28 {
			return false
		}
	}

	return true
}

// Sleep
// @Description: sleep(second)
// @param t
func Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}

// Usleep
// @Description: usleep(microsecond)
// @param t
func Usleep(t int64) {
	time.Sleep(time.Duration(t) * time.Microsecond)
}
