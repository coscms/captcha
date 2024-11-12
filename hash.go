package captcha

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(b []byte) string {
	m := md5.New()
	m.Write(b)
	return hex.EncodeToString(m.Sum(nil))
}
