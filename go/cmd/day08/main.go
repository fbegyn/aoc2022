package main

import (
	"fmt"
	"os"

	"github.com/fbegyn/aoc2022/go/helpers"
)

type Grid [][]int

func main() {
	file := os.Args[1]
	input := helpers.InputToLines(file)

	xSize, ySize := len(input[0]), len(input)
	grid := make(Grid, ySize)
	for y := range input {
		grid[y] = make([]int, xSize)
		for x := range input[y] {
			grid[y][x] = helpers.Atoi(string(input[y][x]))
		}
	}

	visbleTrees := (xSize+ySize)*2 - 4

	maxScore := 0
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			if grid.IsVisible(x, y) {
				visbleTrees++
			}
			score := grid.ScenicScore(x, y)
			if maxScore < score {
				maxScore = score
			}
		}
	}

	fmt.Printf("There are %d trees visible.\n", visbleTrees)
	fmt.Printf("The highest scenic score is: %d\n", maxScore)
}

func (g Grid) ColumnSlice(yStart, yEnd, x int) []int {
	column := make([]int, len(g[yStart:yEnd]))
	for y := 0; y < yEnd-yStart; y++ {
		column[y] = g[y+yStart][x]
	}
	return column
}

func (g Grid) IsVisible(x, y int) bool {
	val := g[y][x]
	lSlice := g[y][:x]
	rSlice := g[y][x+1:]
	uSlice := g.ColumnSlice(0, y, x)
	dSlice := g.ColumnSlice(y+1, len(g), x)

	lVisible, rVisible := checkLOS(val, lSlice), checkLOS(val, rSlice)
	uVisible, dVisible := checkLOS(val, uSlice), checkLOS(val, dSlice)
	return lVisible || rVisible || uVisible || dVisible
}

func checkLOS(v int, slice []int) bool {
	for _, t := range slice {
		if v <= t {
			return false
		}
	}
	return true
}

func (g Grid) ScenicScore(x, y int) int {
	val := g[y][x]
	lSlice := g[y][:x]
	rSlice := g[y][x+1:]
	uSlice := g.ColumnSlice(0, y, x)
	dSlice := g.ColumnSlice(y+1, len(g), x)
	lView, rView := checkView(val, lSlice, true), checkView(val, rSlice, false)
	uView, dView := checkView(val, uSlice, true), checkView(val, dSlice, false)
	return lView * rView * uView * dView
}

func checkView(v int, slice []int, reverse bool) (view int) {
	c := make([]int, len(slice))
	copy(c, slice)
	if reverse {
		helpers.ReverseSlice(c)
	}
	for i, t := range c {
		view = i + 1
		if v <= t {
			break
		}
	}
	return
}

func (g Grid) Print() {
	for y, a := range g {
		for x := range a {
			fmt.Printf("%d", g[y][x])
		}
		fmt.Println()
	}
}
