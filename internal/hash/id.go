package hash

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// IntToMD5Checksum converts an integer to a MD5 checksum. It takes an optional
// prefix.
func IntToMD5Checksum(i int, prefix string) string {
	s := fmt.Sprintf("%s-%d", prefix, i)
	c := md5.Sum(
		[]byte(s),
	)

	return hex.EncodeToString(c[:])
}
