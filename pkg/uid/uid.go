package uid

import (
	"crypto/rand"
	"encoding/base32"
)

func Gen(l int) string {
	byt := make([]byte, 20)
	_, err := rand.Read(byt)
	if err != nil {
		return ""
	}
	return base32.HexEncoding.EncodeToString(byt)[:l]
}
