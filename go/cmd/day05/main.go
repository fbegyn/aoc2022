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

	stackLines := 0
	for i, l := range input {
		if strings.ContainsRune(l, '1') {
			stackLines = i
			break
		}
	}

	stackCount := len(input[0])/4 + 1
	stacks := make([][]byte, stackCount)
	stacks2 := make([][]byte, stackCount)

	for _, l := range input[:stackLines] {
		for i := 0; i < stackCount; i++ {
			runeIndex := 1 + i*4
			if l[runeIndex] != ' ' {
				stacks[i] = append([]byte{l[runeIndex]}, stacks[i]...)
			}
		}
	}

	// create a duplicate for part 2
	copy(stacks2, stacks)

	// solve part1
	for _, l := range input[stackLines+2:] {
		sections := strings.Split(l, " ")
		move, from, to := helpers.Atoi(sections[1]), helpers.Atoi(sections[3])-1, helpers.Atoi(sections[5])-1
		for i := 0; i < move; i++ {
			x, a := stacks[from][len(stacks[from])-1], stacks[from][:len(stacks[from])-1]
			stacks[from] = a
			stacks[to] = append(stacks[to], x)
		}
	}
	fmt.Printf("The first creates for part 1 are: ")
	for _, s := range stacks {
		fmt.Printf("%c", s[len(s)-1])
	}
	fmt.Println()

	// solve part2
	for _, l := range input[stackLines+2:] {
		fmt.Println(l)
		fmt.Println(stacks2)
		sections := strings.Split(l, " ")
		move, from, to := helpers.Atoi(sections[1]), helpers.Atoi(sections[3])-1, helpers.Atoi(sections[5])-1
		x, a := stacks2[from][len(stacks2[from])-move:], stacks2[from][:len(stacks2[from])-move]
		stacks2[from] = a
		stacks2[to] = append(stacks2[to], x...)
		fmt.Println(stacks2)
		fmt.Println()
	}
	fmt.Printf("The first creates for part 2 are: ")
	for _, s := range stacks2 {
		fmt.Printf("%c", s[len(s)-1])
	}
	fmt.Println()

}
