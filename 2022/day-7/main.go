package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

type directory struct {
	name   string
	files  []file
	parent *directory
	child  map[string]*directory
}

type browser struct {
	dirSizes1   []int
	dirSizes2   []int
	spaceToFree int
}

func createTree(lines []string) *directory {
	root := &directory{"root", []file{}, nil, map[string]*directory{}}
	var current *directory
	for _, l := range lines {
		input := strings.Fields(l)
		switch input[0] {
		case "$":
			switch input[1] {
			case "cd":
				switch input[2] {
				case "/":
					current = root
				case "..":
					current = current.parent
				default:
					current = current.child[input[2]]
				}
			case "ls":
			}
		case "dir":
			current.child[input[1]] = &directory{input[1], []file{}, current, map[string]*directory{}}
		default:
			// file
			size, err := strconv.Atoi(input[0])
			if err != nil {
				fmt.Println("Error during str conversion")
			}
			current.files = append(current.files, file{input[1], size})
		}
	}
	return root
}

func (b *browser) getSize(dir *directory) int {
	size := 0
	if len(dir.child) != 0 {
		for _, d := range dir.child {
			size += b.getSize(d)
		}
	}
	if len(dir.files) != 0 {
		for _, f := range dir.files {
			size += f.size
		}
	}
	if size < 100000 {
		b.dirSizes1 = append(b.dirSizes1, size)
	}
	if b.spaceToFree > 0 {
		if size > b.spaceToFree {
			b.dirSizes2 = append(b.dirSizes2, size)
		}
	}
	return size
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSuffix(string(dat), "\n"), "\n")
	fs := createTree(lines)
	browser := &browser{}
	totalSize := browser.getSize(fs)
	// Part 1
	part1Size := 0
	for _, s := range browser.dirSizes1 {
		part1Size += s
	}
	fmt.Println("Sum of of folder < 100000: ", part1Size)
	// Part 2
	requiredSpace := 30000000
	availableSpace := 70000000 - totalSize
	browser.spaceToFree = requiredSpace - availableSpace
	browser.getSize(fs)
	sort.Ints(browser.dirSizes2)
	fmt.Println("Smallest directory to delete:", browser.dirSizes2[0])
}
