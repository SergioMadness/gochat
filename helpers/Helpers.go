package helpers

/**
* Different helpers
*/

import (
	"crypto/md5"
	"encoding/hex"
)

/**
* Get md5 string
*/
func GetMD5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}
