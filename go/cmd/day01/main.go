package main

import (
	"fmt"
	"os"

	"github.com/fbegyn/aoc2022/go/helpers"
	"golang.org/x/exp/slices"
)

func main() {
	file := os.Args[1]
	input := helpers.InputToLines(file)
	var elvesCal []int
	elvesCount, maxCal := 0, 0
	sumCal := 0

	for _, l := range input {
		if l == "" {
			elvesCal = append(elvesCal, sumCal)
			if sumCal > maxCal {
				maxCal = sumCal
			}
			elvesCount++
			sumCal = 0
		} else {
			sumCal += helpers.Atoi(l)
		}
	}
	elvesCal = append(elvesCal, sumCal)
	if sumCal > maxCal {
		maxCal = sumCal
	}
	elvesCount++

	fmt.Printf("The elf carrying the most calories carries: %d.\n", maxCal)

	slices.Sort(elvesCal)
	sum := helpers.Sum(elvesCal[elvesCount-3:])
	fmt.Printf("The top 3 elves carry %d calories.\n", sum)
}
