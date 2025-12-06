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

	input := strings.TrimSpace(string(dat))
	lines := strings.Split(input, "\n")

	// Part 1
	password1 := 0
	problems := map[int][]int{}
	for _, line := range lines {
		cells := strings.Fields(line)
		for i, cell := range cells {
			switch cell {
			case "*":
				mult := 1
				for _, v := range problems[i] {
					mult *= v
				}
				password1 += mult
			case "+":
				sum := 0
				for _, v := range problems[i] {
					sum += v
				}
				password1 += sum
			default:
				n, _ := strconv.Atoi(cell)
				problems[i] = append(problems[i], n)
			}
		}
	}

	// Part 2
	ops := strings.Fields(lines[len(lines)-1])
	rows := lines[:len(lines)-1]
	shifted := ""
	for x := range rows[0] {
		vertical := ""
		for y := range rows {
			vertical += string(rows[y][x])
		}
		shifted += vertical + " "
		if strings.TrimSpace(vertical) == "" {
			shifted += "\n"
		}
	}

	password2 := 0
	problems2 := map[int][]int{}
	slines := strings.Split(shifted, "\n")
	for i, line := range slines {
		cells := strings.Fields(line)
		for _, cell := range cells {
			n, _ := strconv.Atoi(cell)
			problems2[i] = append(problems2[i], n)
		}
	}

	for i, op := range ops {
		switch op {
		case "*":
			mult := 1
			for _, v := range problems2[i] {
				mult *= v
			}
			password2 += mult
		case "+":
			sum := 0
			for _, v := range problems2[i] {
				sum += v
			}
			password2 += sum
		}
	}

	fmt.Println("The part1 password is:", password1)
	fmt.Println("The part2 password is:", password2)
}
