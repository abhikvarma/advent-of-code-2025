package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadLinesInDir(filename, separator string) [][]string {
	cwd, _ := os.Getwd()
	fmt.Printf("current working directory... %s", cwd)
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\r\n")
	output := make([][]string, len(lines))

	for i, line := range lines {
		output[i] = strings.Split(line, separator)
	}

	return output
}
