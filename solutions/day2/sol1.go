package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/abhikvarma/advent-of-code-2025/utils"
)

func main() {
	idRanges := utils.ReadLinesInDir("input1.txt", ",")[0]

	sum1 := 0
	sum2 := 0
	for _, idRange := range idRanges {
		parts := strings.Split(idRange, "-")
		fmt.Printf("\n%v", parts)
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("couldn't convert %s to int", parts[0])
			panic(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("couldn't convert %s to int", parts[1])
			panic(err)
		}
		fmt.Printf("\nchecking from %d to %d", start, end)

		for i := start; i <= end; i++ {
			str := strconv.Itoa(i)
			length := len(str)
			if length&1 == 0 {
				if str[0:length/2] == str[length/2:] {
					fmt.Printf("\nadding %d", i)
					sum1 += i
				}
			}

			// "abcabc" -> "bc|abcabc|ab" contains "abcabc"
			// "somesome" -> "ome|somesome|som" contains "somesome"
			// "abcabcabd" -> "bcabcabdabcabcab" doesn't contain "abcabcabd"
			doubleStr := str + str
			if strings.Index(doubleStr[1:2*length-1], str) != -1 {
				sum2 += i
			}

		}
	}

	utils.PrintAnswer("part1", sum1)
	utils.PrintAnswer("part2", sum2)
}
