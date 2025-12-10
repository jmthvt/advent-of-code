package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	points := []Point{}
	input := strings.TrimSpace(string(dat))
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, Point{x, y})
	}

	maxArea := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]

			width := p1.x - p2.x + 1
			height := p1.y - p2.y + 1

			area := width * height
			if area > maxArea {
				maxArea = area
			}
		}
	}

	fmt.Println("The part1 password is:", maxArea)
}
