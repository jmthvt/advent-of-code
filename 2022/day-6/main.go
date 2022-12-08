package main

import (
	"fmt"
	"os"
)

func markerChecker(s string, distinctChars int) int {
	var ret int
	for i := 0; i < len(s)-distinctChars; i++ {
		keys := map[rune]bool{}
		for _, v := range s[i : i+distinctChars] {
			if !keys[v] {
				keys[v] = true
			} else {
				break
			}
		}
		if len(keys) == distinctChars {
			ret = i + distinctChars
			break
		}
	}
	return ret
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("First start-of-packet marker after character ", markerChecker(string(dat), 4))
	fmt.Println("First start-of-message marker after character ", markerChecker(string(dat), 14))
}
