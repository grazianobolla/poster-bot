//common utilities
package shared

import (
	"fmt"
)

func CheckError(err error) bool {
	if err != nil {
		fmt.Printf("Error Func: %s\n", err)
		return true
	}
	return false
}
