package main

import (
	"fmt"

	"github.com/abhikvarma/advent-of-code-2025/utils"
)

func main() {
	banks := utils.ReadLinesInDir("input.txt", "")

	sum1 := int64(0)
	sum2 := int64(0)
	for _, bankStr := range banks {
		bank := utils.ToIntSlice(bankStr)
		sum1 += maxJoltage(bank, 2)
		sum2 += maxJoltage(bank, 12)
	}

	utils.PrintAnswer("part1", sum1)
	utils.PrintAnswer("part2", sum2)
}

func maxJoltage(bank []int, k int) int64 {
	result := int64(0)
	start := 0
	fmt.Printf("\n\n\n->%v", bank)

	for skip := k - 1; skip >= 0; skip-- {
		searchEnd := len(bank) - skip

		maxIdx := start
		fmt.Printf("\nsearching from %d to %d,\t%v", start, searchEnd, bank[start:searchEnd])
		for j := start + 1; j < searchEnd; j++ {
			if bank[j] > bank[maxIdx] {
				maxIdx = j
			}
		}
		fmt.Printf("\nfound -> %d", bank[maxIdx])

		result = result*10 + int64(bank[maxIdx])
		start = maxIdx + 1
	}

	return result
}
