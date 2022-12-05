package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Fields(string(dat))
	containedPairs, overlappedPairs := 0, 0

	for _, l := range lines {
		pairs := strings.Split(l, ",")
		pair1 := strings.Split(pairs[0], "-")
		pair2 := strings.Split(pairs[1], "-")

		startElf1, _ := strconv.Atoi(pair1[0])
		endElf1, _ := strconv.Atoi(pair1[1])
		startElf2, _ := strconv.Atoi(pair2[0])
		endElf2, _ := strconv.Atoi(pair2[1])

		if (startElf1 <= startElf2 && endElf1 >= endElf2) ||
			(startElf2 <= startElf1 && endElf2 >= endElf1) {
			containedPairs++
			overlappedPairs++
		} else if endElf1 >= startElf2 && endElf2 >= startElf1 {
			overlappedPairs++
		}
	}
	fmt.Println("Contained pairs: ", containedPairs)
	fmt.Println("Overlapped pairs: ", overlappedPairs)
}
