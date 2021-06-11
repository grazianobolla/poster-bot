package hasher

import (
	"crypto/sha256"
	"encoding/base64"
	"shitposter-bot/shared"
)

var sha256_hasher = sha256.New()

//hashes byte array to sha256
func Byte2Sha256(data []byte) string {
	sha256_hasher.Reset()
	_, err := sha256_hasher.Write(data)
	shared.CheckError(err)
	return base64.URLEncoding.EncodeToString(sha256_hasher.Sum(nil))
}

func String2Sha256(data string) string {
	return Byte2Sha256([]byte(data))
}
