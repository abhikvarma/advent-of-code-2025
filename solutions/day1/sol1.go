package main

import (
	"fmt"
	"strconv"

	"github.com/abhikvarma/advent-of-code-2025/utils"
)

const Left = -1
const Right = 1

func main() {
	lines := utils.ReadLinesInDir("input1.txt", " ")

	password1 := 0
	password2 := 0

	pos := 50
	for _, line := range lines {
		direction := getDirection(line[0])
		amt, err := strconv.Atoi(line[0][1:])
		if err != nil {
			fmt.Printf("couldn't convert %s to integer\n", line[0])
			panic(err)
		}

		if direction == Right {
			password2 += countZeroes(pos+amt, pos+1)
		} else {
			password2 += countZeroes(pos-1, pos-amt)
		}

		pos += direction * amt
		pos %= 100
		if pos < 0 {
			pos += 100
		}

		if pos == 0 {
			password1 += 1
		}
	}
	utils.PrintAnswer("part1", password1)
	utils.PrintAnswer("part2", password2)
}

func getDirection(rotation string) int {
	if rotation[0] == 'R' {
		return Right
	}
	return Left
}

func countZeroes(hi, lo int) int {
	if lo > hi {
		return 0
	}
	return utils.FloorDiv(hi, 100) - utils.FloorDiv(lo-1, 100)
}
