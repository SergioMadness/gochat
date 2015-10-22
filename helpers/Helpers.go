package helpers

/**
* Different helpers
 */

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

/**
* Get md5 string
 */
func GetMD5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

/**
* Join integers
 */
func JoinI(items []int, separator string) string {
	result := ""
	for _, val := range items {
		if result != "" {
			result += separator
		}
		result += strconv.Itoa(val)
	}

	return result
}
