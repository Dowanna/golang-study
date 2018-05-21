package errorHandling

import (
	"fmt"
)

func CheckError(e error) {
	if e != nil {
		fmt.Errorf("v", e)
	}
}
