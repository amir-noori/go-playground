package fp

import (
	"fmt"
)

func closureScope() func() {

	count := 0
	return func() {
		fmt.Println("count: ", count)
		count++
	}

}

func tryClosure() {
	fn := closureScope()
	fn()
	fn()
}
