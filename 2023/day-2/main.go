package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var bag = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	r := regexp.MustCompile(`(\d+) (\w+)`)
	lines := strings.Split(string(dat), "\n")
	gameSum, powerSum := 0, 0
	for i, line := range lines {
		if line != "" {
			valid := true
			minCubes := map[string]int{}
			matches := r.FindAllString(line, -1)
			for _, match := range matches {
				var cubeNumber int
				var cubeColor string
				fmt.Sscanf(match, "%d %s", &cubeNumber, &cubeColor)
				minCubes[cubeColor] = Max(minCubes[cubeColor], cubeNumber)
				if cubeNumber > bag[cubeColor] {
					valid = false
				}
			}
			if valid {
				gameSum += i + 1
			}
			powerSum += minCubes["red"] * minCubes["green"] * minCubes["blue"]
		}
	}
	fmt.Println(gameSum)
	fmt.Println(powerSum)
}
