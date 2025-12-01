package utils

import "fmt"

func PrintAnswer(message string, answer any) {
	fmt.Printf("\n\nAnswer...  (%s)\n%v\n", message, answer)
}
