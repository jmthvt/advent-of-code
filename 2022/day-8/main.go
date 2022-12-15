package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverse(s []int) []int {
	// Making a copy - don't touch the original.
	c := make([]int, len(s))
	copy(c, s)
	for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]
	}
	return c
}

func check(trees []int, h int) (bool, int) {
	visibleTrees := 0
	for _, tree := range trees {
		visibleTrees++
		if tree >= h {
			return false, visibleTrees
		}
	}
	return true, visibleTrees
}

func isVisible(y int, x int, matrix [][]int) (bool, int) {
	height := matrix[y][x]
	column := make([]int, len(matrix))
	for y := range matrix {
		column[y] = matrix[y][x]
	}
	// slice a[low : high]
	// low is included, high is excluded.
	left, leftScore := check(reverse(matrix[y][:x]), height)
	right, rightScore := check(matrix[y][x+1:], height)
	up, upScore := check(reverse(column[:y]), height)
	down, downScore := check(column[y+1:], height)
	isVisible := left || right || up || down
	scenicScore := leftScore * rightScore * upScore * downScore
	return isVisible, scenicScore
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSuffix(string(dat), "\n"), "\n")
	matrix := make([][]int, len(lines))
	for y, l := range lines {
		matrix[y] = make([]int, len(l))
		for x, height := range l {
			matrix[y][x], _ = strconv.Atoi(string(height))
		}
	}
	visibleCount := len(matrix)*2 + len(matrix[0])*2 - 4
	bestScenicScore := 0
	for y := 1; y < len(matrix)-1; y++ {
		for x := 1; x < len(matrix[0])-1; x++ {
			visible, scenicScore := isVisible(y, x, matrix)
			if visible {
				visibleCount++
			}
			if scenicScore > bestScenicScore {
				bestScenicScore = scenicScore
			}
		}
	}
	fmt.Println("Visible trees: ", visibleCount)
	fmt.Println("Best scenic score: ", bestScenicScore)
}
