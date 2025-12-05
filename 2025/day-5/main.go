package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(strings.TrimSpace(string(dat)), "\n\n")
	freshDB := strings.Split(input[0], "\n")
	ingredients := strings.Split(input[1], "\n")

	freshRanges := [][2]int{}
	for _, row := range freshDB {
		var low, high int
		fmt.Sscanf(row, "%d-%d", &low, &high)
		freshRanges = append(freshRanges, [2]int{low, high})
	}
	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i][0] < freshRanges[j][0]
	})

	// Part 1
	password1 := 0
	for _, ingredient := range ingredients {
		i, _ := strconv.Atoi(ingredient)
		for _, r := range freshRanges {
			if r[0] <= i && i <= r[1] {
				password1++
				break
			}
		}
	}

	// Part 2
	mergedRanges := [][2]int{freshRanges[0]}
	for _, r := range freshRanges {
		last := &mergedRanges[len(mergedRanges)-1]
		lastHigh := last[1]
		currentLow := r[0]
		currentHigh := r[1]

		if currentLow <= lastHigh {
			last[1] = max(currentHigh, lastHigh)
		} else {
			mergedRanges = append(mergedRanges, r)
		}
	}

	password2 := 0
	for _, r := range mergedRanges {
		password2 += r[1] - r[0] + 1
	}

	fmt.Println("The part1 password is:", password1)
	fmt.Println("The part2 password is:", password2)
}
