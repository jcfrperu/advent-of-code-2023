package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func createListOfDifferences(seq []int64) (bool, []int64) {
	var diffs = make([]int64, 0, 128)

	var hasNonZero = false
	for i := 0; i < len(seq)-1; i++ {
		var diff = seq[i+1] - seq[i]
		if diff != 0 {
			hasNonZero = true
		}
		diffs = append(diffs, diff)
	}

	return hasNonZero, diffs
}

func findNextValue(sequence []int64) int64 {
	var history = make([][]int64, 0, 64)
	history = append(history, sequence)

	var seq = sequence
	var diff []int64
	var hasNonZero = true
	for hasNonZero {
		hasNonZero, diff = createListOfDifferences(seq)
		history = append(history, diff)
		seq = diff
	}

	for i := len(history) - 1; i >= 1; i-- {
		var currentIndex = len(history[i]) - 1
		var beforeIndex = len(history[i-1]) - 1
		var newNumber = history[i-1][beforeIndex] + history[i][currentIndex]
		history[i-1] = append(history[i-1], newNumber)
	}

	return history[0][len(history[0])-1]
}

func findFirstValue(sequence []int64) int64 {
	var history = make([][]int64, 0, 64)
	history = append(history, sequence)

	var seq = sequence
	var diff []int64
	var hasNonZero = true
	for hasNonZero {
		hasNonZero, diff = createListOfDifferences(seq)
		history = append(history, diff)
		seq = diff
	}

	for i := len(history) - 1; i >= 1; i-- {
		var newNumber = history[i-1][0] - history[i][0]
		history[i-1] = append([]int64{newNumber}, history[i-1]...)
	}

	return history[0][0]
}

func solutionPart01(lines []string) {
	var sum = int64(0)
	for _, line := range lines {
		var seq = SplitLongs(line, " ")
		sum += findNextValue(seq)
	}

	fmt.Printf("%d", sum)
}

func solutionPart02(lines []string) {
	var sum = int64(0)
	for _, line := range lines {
		var seq = SplitLongs(line, " ")
		sum += findFirstValue(seq)
	}
	fmt.Printf("%d", sum)
}

// https://adventofcode.com/2023/day/9
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45")
	//RunAdventOfCodeWithFile(solutionPart01, "day_09/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "10  13  16  21  30  45")
	RunAdventOfCodeWithFile(solutionPart02, "day_09/testcases/input-part-02.txt")
}
