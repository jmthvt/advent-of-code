package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Part 2
// Replace digits words by digit encapsulated with the first and last letter
// to avoid breaking digits with overpalpping letters.
var digits = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func firstDigit(s string) string {
	var digit rune
	for _, r := range s {
		if unicode.IsDigit(r) {
			digit = r
			break
		}
	}
	return string(digit)
}

func lastDigit(s string) string {
	var digit byte
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			digit = s[i]
			break
		}
	}
	return string(digit)
}

// Part 2
func replaceLettersDigits(s string) string {
	for word, digit := range digits {
		s = strings.ReplaceAll(s, word, digit)
	}
	fmt.Println(s)
	return s
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	calibrationValuesSum := 0

	for _, v := range lines {
		v = replaceLettersDigits(v)
		calibrationValue, _ := strconv.Atoi(firstDigit(v) + lastDigit(v))
		calibrationValuesSum += calibrationValue
	}
	fmt.Println("Calibration Values Sum:", calibrationValuesSum)
}
