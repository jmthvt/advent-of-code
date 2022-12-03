package main

import (
	"fmt"
	"os"
	"strings"
)

type Score struct {
	part1 int
	part2 int
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	// pre-baked score for each move
	scoreMap := map[string]Score{
		"A X": {4, 3},
		"A Y": {8, 4},
		"A Z": {3, 8},

		"B X": {1, 1},
		"B Y": {5, 5},
		"B Z": {9, 9},

		"C X": {7, 2},
		"C Y": {2, 6},
		"C Z": {6, 7},
	}

	moves := strings.Split(string(dat), "\n")
	score1, score2 := 0, 0
	for _, move := range moves {
		score1 += scoreMap[move].part1
		score2 += scoreMap[move].part2
	}
	fmt.Println("Part1 score: ", score1)
	fmt.Println("Part2 score: ", score2)
}
