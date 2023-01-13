// @Author zhanglingyu
package string

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash/crc32"
	"os"
	"strings"
)

// Studly
// @Description: 下划线转驼峰
// @param str
// @return string
func Studly(str string) string {
	if str != "" {
		c := strings.Split(str, "_")
		var s string
		for _, v := range c {
			vv := []rune(v)
			if len(vv) > 0 {
				if vv[0] >= 'a' && vv[0] <= 'z' {
					vv[0] -= 32
				}
				s += string(vv)
			}
		}

		return s
	}

	return str
}

//
// Md5
//  @Description: md5
//  @param str
//  @return string
//
func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

//
// Md5File
//  @Description: md5 file
//  @param path
//  @return string
//  @return error
//
func Md5File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return "", err
	}

	var size int64 = 1048576 // 1M
	hash := md5.New()

	if fi.Size() < size {
		data, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		hash.Write(data)
	} else {
		b := make([]byte, size)
		for {
			n, err := f.Read(b)
			if err != nil {
				break
			}

			hash.Write(b[:n])
		}
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

//
// Sha1
//  @Description: sha1
//  @param str
//  @return string
//
func Sha1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

//
// Sha1File
//  @Description: sha1 file
//  @param path
//  @return string
//  @return error
//
func Sha1File(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := sha1.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

//
// Crc32
//  @Description: crc32
//  @param str
//  @return uint32
//
func Crc32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}
