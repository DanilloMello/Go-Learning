package greetings

import (
	"fmt"
)

func Hello(name string) string {
	message := name
	return fmt.Sprintf("PREFIX: %v", message)
}
