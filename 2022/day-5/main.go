package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type stack struct {
	crates []string
}

// https://github.com/golang/go/wiki/SliceTricks#push-frontunshift
func (s *stack) prepend(char string) {
	s.crates = append([]string{char}, s.crates...)
}

func (s *stack) push(crate string) {
	s.crates = append(s.crates, crate)
}

func (s *stack) multiPush(crate []string) {
	s.crates = append(s.crates, crate...)
}

func (s *stack) pop() string {
	r := s.crates[len(s.crates)-1]
	s.crates = s.crates[:len(s.crates)-1]
	return r
}

func (s *stack) multiPop(i int) []string {
	r := s.crates[len(s.crates)-i : len(s.crates)]
	s.crates = s.crates[:len(s.crates)-i]
	return r
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	stacksData := strings.Split(string(dat), "\n\n")[0]
	stacksLines := strings.Split(stacksData, "\n")
	stacks := make([]stack, 9)
	for _, sl := range stacksLines {
		for i := 1; i < len(sl); i += 4 {
			r := rune(sl[i])
			if unicode.IsLetter(r) {
				stacks[i/4].prepend(string(r))
			}
		}
	}
	stacks2 := make([]stack, 9)
	copy(stacks2, stacks)

	movesData := strings.Split(string(dat), "\n\n")[1]
	movesLines := strings.Split(movesData, "\n")
	var quantity, from, to int
	for _, ml := range movesLines {
		if ml != "" {
			fmt.Sscanf(ml, "move %d from %d to %d", &quantity, &from, &to)
			// Part 1
			for q := 0; q < quantity; q++ {
				cr := stacks[from-1].pop()
				stacks[to-1].push(cr)
			}
			// Part 2
			cr := stacks2[from-1].multiPop(quantity)
			stacks2[to-1].multiPush(cr)

		}
	}

	fmt.Println("Part 1:")
	for _, s := range stacks {
		fmt.Println(s.pop())
	}

	fmt.Println("Part 2:")
	for _, s := range stacks2 {
		fmt.Println(s.pop())
	}
}
