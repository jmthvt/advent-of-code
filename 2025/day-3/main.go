package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func seekJoltage(bank string, index int, digits int, currentJoltage string) (joltage int) {
	for i := 9; i > 0; i-- {
		battery := strconv.Itoa(i)
		maxIndex := len(bank) - (digits - len(currentJoltage) - 1)
		if strings.Contains(bank[index:maxIndex], battery) {
			index += strings.Index(bank[index:maxIndex], battery)
			currentJoltage += battery
			if len(currentJoltage) == digits {
				joltage, _ = strconv.Atoi(currentJoltage)
				break
			} else {
				joltage = seekJoltage(bank, index+1, digits, currentJoltage)
				break
			}
		}
	}
	return joltage
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(dat))
	banks := strings.Split(input, "\n")

	password1, password11, password2 := 0, 0, 0

	// Part 1
	battery1, battery2 := "", ""
OUTER:
	for _, bank := range banks {
		for i := 9; i > 0; i-- {
			battery1 = strconv.Itoa(i)
			if strings.Contains(bank[:len(bank)-1], battery1) {
				index := strings.Index(bank[:len(bank)-1], battery1)
				for j := 9; j > 0; j-- {
					battery2 = strconv.Itoa(j)
					if strings.Contains(bank[index+1:], battery2) {
						battery, _ := strconv.Atoi(battery1 + battery2)
						password1 += battery
						continue OUTER
					}
				}
			}
		}
	}

	// Part 2
	for _, bank := range banks {
		password11 += seekJoltage(bank, 0, 2, "")
		password2 += seekJoltage(bank, 0, 12, "")
	}

	fmt.Println("The part1 password is:", password1)
	fmt.Println("The part1 password is: (using part2 code)", password11)
	fmt.Println("The part2 password is:", password2)
}
