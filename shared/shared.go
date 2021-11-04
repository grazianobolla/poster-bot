//common utilities
package shared

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func CheckError(err error) bool {
	if err != nil {
		fmt.Printf("Error Func: %s\n", err)
		return true
	}
	return false
}

//byte[] to base64 string
func ToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func GetContentType(url string) string {
	res, err := http.Get(url)
	if CheckError(err) {
		return ""
	}

	t := res.Header.Get("Content-Type")
	return t
}
