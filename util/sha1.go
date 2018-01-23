package util

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func EncryptSha1(data string) (enData string) {
	t := sha1.New()
	io.WriteString(t, data);
	return fmt.Sprintf("%x", t.Sum(nil))
}
