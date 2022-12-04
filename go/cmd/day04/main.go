package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fbegyn/aoc2022/go/helpers"
)

func main() {
	file := os.Args[1]
	input := helpers.InputToLines(file)

	overlap, overlapPartially := 0, 0
	for _, l := range input {
		assignments := strings.Split(l, ",")
		firstRangeStr := strings.Split(assignments[0], "-")
		secondRangeStr := strings.Split(assignments[1], "-")
		firstRange := [2]int{helpers.Atoi(firstRangeStr[0]), helpers.Atoi(firstRangeStr[1])}
		secondRange := [2]int{helpers.Atoi(secondRangeStr[0]), helpers.Atoi(secondRangeStr[1])}

		if checkPairs(firstRange, secondRange) {
			overlap++
			overlapPartially++
		} else {
			if checkPairsPartially(firstRange, secondRange) {
				overlapPartially++
			}
		}
	}
	fmt.Printf("There are %d pairs with fully overlapping ranges.\n", overlap)
	fmt.Printf("There are %d pairs with overlapping ranges.\n", overlapPartially)
}

func checkPairs(x, y [2]int) bool {
	if x[0] <= y[0] && y[1] <= x[1] {
		return true
	}
	if y[0] <= x[0] && x[1] <= y[1] {
		return true
	}
	return false
}

func checkPairsPartially(x, y [2]int) bool {
	if y[0] <= x[0] && x[0] <= y[1] {
		return true
	}
	if y[0] <= x[1] && x[1] <= y[1] {
		return true
	}
	return false
}
