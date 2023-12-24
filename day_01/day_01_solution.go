package main

import (
	. "github.com/jcfrperu/go-competitive-programming"
	"sort"
	"strings"
	"unicode"
)

func getTwoDigitsNumber(digit01 int, digit02 int) int {
	// example: 3, 5 -> "3" + "5" -> "35" -> 35
	return ParseInt(FormatInt(digit01) + FormatInt(digit02))
}

func solutionPart01(lines []string) {
	var sum = 0
	for _, line := range lines {
		var digits = make([]int, 0)
		for _, c := range line {
			if unicode.IsDigit(c) {
				digits = append(digits, ParseInt(string(c)))
			}
		}

		if len(digits) == 1 {
			sum += getTwoDigitsNumber(digits[0], digits[0])
		} else {
			var lastIndex = len(digits) - 1
			sum += getTwoDigitsNumber(digits[0], digits[lastIndex])
		}
	}

	print(sum)
}

func solutionPart02(lines []string) {
	var digitsCatalog = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	var sum = 0
	for _, line := range lines {

		var matches = make(map[int]string)  // map: key -> position of digit, value -> "five" or "5" (digitsCatalog)
		var matchesIndexes = make([]int, 0) // list of indexes of matches map

		for k, _ := range digitsCatalog {
			// finding first matches with digits in catalog
			var firstMatch = strings.Index(line, k)
			if firstMatch >= 0 {
				matches[firstMatch] = k
				matchesIndexes = append(matchesIndexes, firstMatch)
			}

			// finding last matches with digits in catalog
			var lastMatch = strings.LastIndex(line, k)
			if lastMatch >= 0 {
				matches[lastMatch] = k
				matchesIndexes = append(matchesIndexes, lastMatch)
			}
		}
		sort.Ints(matchesIndexes) // we need to sort them to know positions were found

		var digits = make([]int, 0)
		for _, v := range matchesIndexes {
			// getting the values of digitsCatalog map using the string key
			digits = append(digits, digitsCatalog[matches[v]])
		}

		// same as part 01
		if len(digits) == 1 {
			sum += getTwoDigitsNumber(digits[0], digits[0])
		} else {
			var lastIndex = len(digits) - 1
			sum += getTwoDigitsNumber(digits[0], digits[lastIndex])
		}
	}

	print(sum)
}

// https://adventofcode.com/2023/day/1
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet")
	RunAdventOfCodeWithFile(solutionPart01, "day_01/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen")
	//RunAdventOfCodeWithFile(solutionPart02, "day_01/testcases/input-part-02.txt")
}
