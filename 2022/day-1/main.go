package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sum(calories []string) int {
	caloriesTotal := 0
	for _, v := range calories {
		intCal, _ := strconv.Atoi(v)
		caloriesTotal += intCal
	}
	return caloriesTotal
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	elves := strings.Split(string(dat), "\n\n")
	var totalCaloriesPerElf []int

	for _, v := range elves {
		totalCalories := sum(strings.Split(v, "\n"))
		totalCaloriesPerElf = append(totalCaloriesPerElf, totalCalories)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(totalCaloriesPerElf)))
	topElf := totalCaloriesPerElf[0]
	topThreeElves := totalCaloriesPerElf[0] + totalCaloriesPerElf[1] + totalCaloriesPerElf[2]

	fmt.Printf("Top Elf carries %v calories.\n", topElf)
	fmt.Printf("Top 3 Elves carry a total of %v calories.\n", topThreeElves)
}
