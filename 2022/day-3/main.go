package main

import (
	"fmt"
	"os"
	"strings"
)

const contents = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func findSharedItem(source string, target string) string {
	for _, c := range source {
		i := strings.Index(target, string(c))
		if i > -1 {
			return string(c)
		}
	}
	return ""
}

func itemMatches(source string, target string) map[rune]bool {
	matches := map[rune]bool{}
	for _, c := range source {
		i := strings.Index(target, string(c))
		if i > -1 {
			matches[c] = true
		}
	}
	return matches
}

func main() {

	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Fields(string(dat))
	fmt.Println("P1. The sum of priorities is: ", p1(lines))
	fmt.Println("P2. The sum of priorities is: ", p2(lines))
}

func p1(lines []string) int {
	prioritySum := 0
	for _, sack := range lines {
		items := len(sack) + 1
		rucksackCompartment1 := sack[:items/2]
		rucksackCompartment2 := sack[items/2:]
		sharedItem := findSharedItem(rucksackCompartment1, rucksackCompartment2)
		sharedItemPriority := strings.Index(contents, sharedItem) + 1
		prioritySum += sharedItemPriority
	}
	return prioritySum
}

func p2(lines []string) int {
	prioritySum := 0
	for i := 0; i < len(lines); i += 3 {
		elf1 := lines[i]
		elf2 := lines[i+1]
		elf3 := lines[i+2]
		sharedItemsElf1and2 := itemMatches(elf1, elf2)
		sharedItemsElf2and3 := itemMatches(elf2, elf3)
		for item := range sharedItemsElf1and2 {
			if sharedItemsElf2and3[item] {
				sharedItemPriority := strings.Index(contents, string(item)) + 1
				prioritySum += sharedItemPriority
				break
			}
		}
	}
	return prioritySum
}
