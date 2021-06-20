package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

func AsSHA256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func AsMD5(o interface{}) string {
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))
	return fmt.Sprintf("%x", h.Sum(nil))
}
