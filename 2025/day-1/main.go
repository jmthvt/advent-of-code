package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	data int
	next *Number
	prev *Number
}

type Dial struct {
	head    *Number
	tail    *Number
	pointer *Number
	click1  int
	click2  int
}

func (d *Dial) Append(data int) {
	newNumber := &Number{data: data}
	if d.head == nil {
		d.head = newNumber
		d.tail = newNumber
	} else {
		newNumber.prev = d.tail
		d.tail.next = newNumber
		d.tail = newNumber

		d.tail.next = d.head
		d.head.prev = d.tail

		if data == 50 {
			d.pointer = newNumber
		}
	}
}

func (d *Dial) MoveForward(distance int) {
	for i := 0; i < distance; i++ {
		d.pointer = d.pointer.next
		if d.pointer.data == 0 {
			d.click2++
		}
	}
}

func (d *Dial) MoveBackward(distance int) {
	for i := 0; i < distance; i++ {
		d.pointer = d.pointer.prev
		if d.pointer.data == 0 {
			d.click2++
		}
	}
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(dat)) // Removes trailing \n
	lines := strings.Split(input, "\n")

	dial := &Dial{}
	for i := 0; i < 100; i++ {
		dial.Append(i)
	}
	dial.click1 = 0
	dial.click2 = 0
	for _, rotation := range lines {
		distance, _ := strconv.Atoi(rotation[1:])
		if strings.Contains(rotation, "R") {
			dial.MoveForward(distance)
		} else if strings.Contains(rotation, "L") {
			dial.MoveBackward(distance)
		}
		if dial.pointer.data == 0 {
			dial.click1++
		}
	}
	fmt.Println("Part 1 Password: ", dial.click1)
	fmt.Println("Part 2 Password: ", dial.click2)
}
