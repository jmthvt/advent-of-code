package main

import (
	"fmt"
	"maps"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

var adjacents = []Point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(dat))
	lines := strings.Split(input, "\n")

	grid := map[Point]int{}
	for y, l := range lines {
		for x, slot := range l {
			if slot == '@' {
				grid[Point{x, y}] = 1
			}
		}
	}

	// Part 1
	forkliftableRolls := 0
	for point := range grid {
		rolls := 0
		for _, a := range adjacents {
			rolls += grid[point.Add(a)]
		}
		if rolls < 4 {
			forkliftableRolls++
		}
	}

	// Part 2
	removed := 0
	for {
		nextGrid := maps.Clone(grid)
		for point := range grid {
			rolls := 0
			for _, a := range adjacents {
				rolls += grid[point.Add(a)]
			}
			if rolls < 4 {
				delete(nextGrid, point)
				removed++
			}
		}
		if maps.Equal(nextGrid, grid) {
			break
		}
		grid = nextGrid
	}

	fmt.Println("The part1 password is:", forkliftableRolls)
	fmt.Println("The part2 password is:", removed)
}
