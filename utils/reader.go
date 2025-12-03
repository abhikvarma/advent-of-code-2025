package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadLinesInDir(filename, separator string) [][]string {
	cwd, _ := os.Getwd()
	fmt.Printf("current working directory... %s", cwd)
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	rawLines := strings.Split(string(content), "\r\n")
	var lines []string
	for _, line := range rawLines {
		if line != "" {
			lines = append(lines, line)
		}
	}

	output := make([][]string, len(lines))

	for i, line := range lines {
		output[i] = strings.Split(line, separator)
	}

	return output
}

func ToIntSlice(s []string) []int {
	var output []int
	for _, str := range s {
		num, _ := strconv.Atoi(str)
		output = append(output, num)
	}
	return output
}
