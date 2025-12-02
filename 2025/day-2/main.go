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

	input := strings.TrimSpace(string(dat)) // Removes trailing \n
	intervals := strings.Split(input, ",")

	password1, password2 := 0, 0

	for _, interval := range intervals {
		var min, max int
		fmt.Sscanf(interval, "%d-%d", &min, &max)
		for i := min; i < max+1; i++ {

			s := strconv.Itoa(i)
			// Part 1 Checks
			if len(s)%2 == 0 {
				middle := len(s) / 2
				numberOne := s[:middle]
				numberTwo := s[middle:]
				if numberOne == numberTwo {
					password1 += i
				}
			}
			// Part 2 Checks
			// Great stuff from https://stackoverflow.com/a/55840779
			if strings.Contains((s + s)[1:len(s+s)-1], s) {
				password2 += i
			}
		}
	}
	fmt.Println("The Part 1 password is: ", password1)
	fmt.Println("The Part 2 password is: ", password2)
}
