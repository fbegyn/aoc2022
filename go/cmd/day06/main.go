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

	for _, l := range input {
		fmt.Printf("The first packet can be found after: %9d\n", checkMarker(l, 4))
		fmt.Printf("The first message can be found after: %8d\n", checkMarker(l, 14))
		fmt.Println()
	}
}

func checkMarker(inp string, size int) int {
	searching := true
	var i int
	for i = 0; searching; i++ {
		searching = false
		buffer := inp[i : size+i]
		for _, r := range buffer {
			if strings.Count(buffer, string(r)) > 1 {
				searching = true
				break
			}
		}
	}
	return i + size - 1
}
