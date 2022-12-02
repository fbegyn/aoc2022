package main

import (
	"fmt"
	"os"

	"github.com/fbegyn/aoc2022/go/helpers"
)

func main() {
	file := os.Args[1]
	input := helpers.InputToLines(file)

	// Parse the input and calculate the total score
	totalScore, totalScore2 := 0, 0
	for _, line := range input {
		totalScore += calculateRoundScore(line)
		totalScore2 += calculateRoundScorePart2(line)
	}

	// Print the result
	fmt.Printf("Total score part 1: %d\n", totalScore)
	fmt.Printf("Total score part 2: %d\n", totalScore2)
}

func calculateRoundScore(line string) int {
	// Parse the round info from the line
	opponentShape := line[0] - 'A'
	playerShape := line[2] - 'X'

	// Calculate the score for the player's shape
	var playerShapeScore int
	switch playerShape {
	case 0:
		playerShapeScore = 1
	case 1:
		playerShapeScore = 2
	case 2:
		playerShapeScore = 3
	}

	// Calculate the outcome score
	var outcomeScore int
	if opponentShape == playerShape {
		// Draw
		outcomeScore = 3
	} else {
		// Win or loss
		var winsAgainst byte
		switch playerShape {
		case 0:
			winsAgainst = 2
		case 1:
			winsAgainst = 0
		case 2:
			winsAgainst = 1
		}
		if opponentShape == winsAgainst {
			// Win
			outcomeScore = 6
		} else {
			// Loss
			outcomeScore = 0
		}
	}

	// Return the total round score
	return playerShapeScore + outcomeScore
}

func determineHand(outcome, opponent byte) (byte, int) {
	switch outcome {
        // lose
	case 0:
		switch opponent {
		case 0:
			return 2, 0
		case 1:
			return 0, 0
		case 2:
			return 1, 0
		}
        // draw
	case 1:
		switch opponent {
		case 0:
			return 0, 3
		case 1:
			return 1, 3
		case 2:
			return 2, 3
		}
        // win
	case 2:
		switch opponent {
		case 0:
			return 1, 6
		case 1:
			return 2, 6
		case 2:
			return 0, 6
		}
	}
	return 5, 0
}

func calculateRoundScorePart2(line string) int {
	// Parse the round info from the line
	opponentShape := line[0] - 'A'
	outcome := line[2] - 'X'


	// Calculate the outcome score
	playerShape, outcomeScore := determineHand(outcome, opponentShape)

	// Calculate the score for the player's shape
	var playerShapeScore int
	switch playerShape {
	case 0:
		playerShapeScore = 1
	case 1:
		playerShapeScore = 2
	case 2:
		playerShapeScore = 3
	}

	// Return the total round score
	return playerShapeScore + outcomeScore
}
