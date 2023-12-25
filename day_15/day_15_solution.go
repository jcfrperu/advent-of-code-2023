package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solutionPart01(lines []string) {
	blocks := Split(lines[0], ",")

	sum := int64(0)
	for _, block := range blocks {
		sum += getHash(block)
	}
	fmt.Printf("%d", sum)
}

func getHash(value string) int64 {
	hash := int64(0)
	for _, ch := range value {
		hash += int64(ch)
		hash *= 17
		hash = hash % 256
	}
	return hash
}

func solutionPart02(lines []string) {

	fmt.Printf("%d", 0)
}

// https://adventofcode.com/2023/day/15
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7")
	RunAdventOfCodeWithFile(solutionPart01, "day_15/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "")
	//RunAdventOfCodeWithFile(solutionPart02, "day_15/testcases/input-part-02.txt")
}
