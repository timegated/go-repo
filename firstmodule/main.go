package firstmodule

import (
	"fmt"
)

func Firstmodule(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}