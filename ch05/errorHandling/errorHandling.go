package errorHandling

import (
	"fmt"
	"os"
)

func CheckError(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v", e)
	}
}
