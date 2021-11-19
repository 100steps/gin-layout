package encryptor

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5(input string) string {
	writer := md5.New()
	io.WriteString(writer, input)
	res := writer.Sum(nil)
	return fmt.Sprintf("%x", res)
}
