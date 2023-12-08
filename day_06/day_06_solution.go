package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"strings"
)

func solutionPart01(lines []string) {
	var times = SplitInts(Split(lines[0], ":")[1], " ")
	var distances = SplitInts(Split(lines[1], ":")[1], " ")

	var sumProd = 1
	for i, time := range times {
		var counter = 0
		for startTime := 1; startTime < time; startTime++ {
			var movingTime = time - startTime
			var velocity = startTime
			var distance = movingTime * velocity
			if distance > distances[i] {
				counter++
			}
		}
		sumProd = sumProd * counter
	}

	fmt.Printf("%d", sumProd)
}

func solutionPart02(lines []string) {
	var time = ParseLong(strings.ReplaceAll(Split(lines[0], ":")[1], " ", ""))
	var distance = ParseLong(strings.ReplaceAll(Split(lines[1], ":")[1], " ", ""))

	var counter = int64(0)
	for startTime := int64(1); startTime < time; startTime++ {
		var movingTime = time - startTime
		var velocity = startTime
		if movingTime*velocity > distance {
			counter++
		}
	}

	fmt.Printf("%d", counter)
}

// https://adventofcode.com/2023/day/6
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "Time:      7  15   30\nDistance:  9  40  200")
	//RunAdventOfCodeWithFile(solutionPart01, "day_06/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "Time:      7  15   30\nDistance:  9  40  200")
	RunAdventOfCodeWithFile(solutionPart02, "day_06/testcases/input-part-02.txt")
}
