package main

import (
	"fmt"
	"os"
	"strings"

	"index/suffixarray"
)

const START = "S"
const SPLITTER = "^"

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(dat))
	lines := strings.Split(input, "\n")

	password1 := 0
	startIndex := strings.Index(lines[0], START)
	beamIndex := map[int]int{}
	beamIndex[startIndex] = 1

	for _, line := range lines[1:] {
		index := suffixarray.New([]byte(line))
		offsets := index.Lookup([]byte(SPLITTER), -1)
		delta := map[int]int{}

		for _, v := range offsets {
			if count, ok := beamIndex[v]; ok {
				delta[v-1] += count
				delta[v+1] += count
				delete(beamIndex, v)
				password1++
			}
		}
		for k, v := range delta {
			beamIndex[k] += v
		}
	}

	password2 := 0
	for _, count := range beamIndex {
		password2 += count
	}

	fmt.Println("The part1 password is:", password1)
	fmt.Println("The part2 password is:", password2)
}
