package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/fbegyn/aoc2022/go/helpers"
)

func main() {
	file := os.Args[1]
	input := helpers.InputToLines(file)

	totalPrio, groupPrio := 0, 0
	var group [3]string
	counter := 0
	for _, l := range input {
		middle := len(l) / 2
		comp1, comp2 := l[:middle], l[middle:]
		var overlap rune
		for _, r := range comp1 {
			if strings.ContainsRune(comp2, r) {
				overlap = r
				break
			}
		}

		var priority int
		if unicode.IsUpper(overlap) {
			priority = int(overlap - 38)
		} else {
			priority = int(overlap - 96)
		}

		totalPrio += priority

		group[counter] = l
		counter++
		if counter > 2 {
			for _, r := range group[0] {
				if strings.ContainsRune(group[1], r) && strings.ContainsRune(group[2], r) {
					overlap = r
					break
				}
			}
			if unicode.IsUpper(overlap) {
				priority = int(overlap - 38)
			} else {
				priority = int(overlap - 96)
			}
			groupPrio += priority
			counter = 0
		}
	}

	fmt.Printf("The total priority is %d.\n", totalPrio)
	fmt.Printf("The total group priority is %d.\n", groupPrio)
}
