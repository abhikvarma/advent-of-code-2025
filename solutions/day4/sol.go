package main

import "github.com/abhikvarma/advent-of-code-2025/utils"

func main() {
	lines := utils.ReadLinesInDir("input.txt", "")
	totalRows := len(lines)
	totalCols := len(lines[0])

	sum1 := 0
	for row, line := range lines {
		for col, char := range line {
			if char == "@" {
				toiletPaperCount := countToiletPaper(row, col, totalRows, totalCols, lines)
				if toiletPaperCount < 4 {
					sum1 += 1
				}
			}
		}
	}

	sum2 := 0
	rerun := true
	for rerun {
		rerun = false
		for row, line := range lines {
			for col, char := range line {
				if char == "@" {
					toiletPaperCount := countToiletPaper(row, col, totalRows, totalCols, lines)
					if toiletPaperCount < 4 {
						sum2 += 1
						lines[row][col] = "."
						rerun = true
					}
				}
			}
		}
	}

	utils.PrintAnswer("part1", sum1)
	utils.PrintAnswer("part2", sum2)
}

func countToiletPaper(row, col, totalRows, totalCols int, lines [][]string) int {
	cnt := 0

	if row-1 >= 0 && col-1 >= 0 && lines[row-1][col-1] == "@" {
		cnt++
	}
	if row-1 >= 0 && lines[row-1][col] == "@" {
		cnt++
	}
	if row-1 >= 0 && col+1 < totalCols && lines[row-1][col+1] == "@" {
		cnt++
	}

	if col-1 >= 0 && lines[row][col-1] == "@" {
		cnt++
	}
	if col+1 < totalCols && lines[row][col+1] == "@" {
		cnt++
	}

	if row+1 < totalRows && col-1 >= 0 && lines[row+1][col-1] == "@" {
		cnt++
	}
	if row+1 < totalRows && lines[row+1][col] == "@" {
		cnt++
	}
	if row+1 < totalRows && col+1 < totalCols && lines[row+1][col+1] == "@" {
		cnt++
	}

	return cnt
}
