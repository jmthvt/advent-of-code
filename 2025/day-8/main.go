package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z int
}

// Connection between 2 points with a calculated distance
type Edge struct {
	dist, u, v int
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
		z, _ := strconv.Atoi(parts[2])
		points = append(points, Point{x, y, z})
	}

	// Calculate distances & sort
	var edges []Edge
	n := len(points)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			p1 := points[i]
			p2 := points[j]

			// Calculate squared Euclidean distance
			dist := (p1.x-p2.x)*(p1.x-p2.x) +
				(p1.y-p2.y)*(p1.y-p2.y) +
				(p1.z-p2.z)*(p1.z-p2.z)

			if i > j {
				edges = append(edges, Edge{dist: dist, u: i, v: j})
			}
		}
	}
	sort.Slice(edges, func(a, b int) bool {
		if edges[a].dist != edges[b].dist {
			return edges[a].dist < edges[b].dist
		}
		if edges[a].u != edges[b].u {
			return edges[a].u < edges[b].u
		}
		return edges[a].v < edges[b].v
	})

	// Union-Find
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) {
		rootX := find(x)
		rootY := find(y)
		if rootX != rootY {
			parent[rootX] = rootY
		}
	}

	password1, password2 := 0, 0
	connections := 0
	for pairs, edge := range edges {
		i, j := edge.u, edge.v

		// Part 1 (snapshot at pairs=1000):
		if pairs == 1000 {
			sz := map[int]int{}
			for x := 0; x < n; x++ {
				sz[find(x)]++
			}

			// Extract values and sort
			var sizes []int
			for _, count := range sz {
				sizes = append(sizes, count)
			}
			sort.Ints(sizes)

			slen := len(sizes)
			if slen >= 3 {
				password1 = sizes[slen-1] * sizes[slen-2] * sizes[slen-3]
			}
		}

		// Construct MST
		if find(i) != find(j) {
			connections++
			if connections == n-1 {
				// Last edge needed to connect the graph
				password2 = points[i].x * points[j].x
			}
			union(i, j)
		}
	}

	fmt.Println("The part1 password is:", password1)
	fmt.Println("The part2 password is:", password2)
}
