package cryptoutils

import (
	"crypto/md5"
	"encoding/hex"
)

//GetMd5 function
func GetMd5(input string) string {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
