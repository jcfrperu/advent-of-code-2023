package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"unicode"
)

type Symbol struct {
	line  int
	index int
	value string
}

type Number struct {
	startIndex int
	endIndex   int
	value      int
	lineIndex  int
}

func findAllNumbers(lines []string) []Number {
	var numbers = make([]Number, 0)
	for i, line := range lines {
		var numbersInLine = extractNumbersInLine(line, i)
		for _, number := range numbersInLine {
			numbers = append(numbers, number)
		}
	}
	return numbers
}

func findAllSymbols(lines []string) []Symbol {
	var symbols = make([]Symbol, 0)
	for i, line := range lines {
		for index, c := range line {
			if !unicode.IsDigit(c) && c != '.' {
				symbols = append(symbols, Symbol{i, index, string(c)})
			}
		}
	}
	return symbols
}

func filterSymbolsByValue(symbols []Symbol, value string) []Symbol {
	var results = make([]Symbol, 0)
	for _, symbol := range symbols {
		if symbol.value == value {
			results = append(results, symbol)
		}
	}
	return symbols
}

func isNumberAdjacentToSymbol(number Number, symbol Symbol) bool {
	// check symbols in same line
	if symbol.line == number.lineIndex {
		if symbol.index == number.startIndex-1 || symbol.index == number.endIndex+1 {
			return true
		}
	}
	// check symbols in line before or after
	if symbol.line == number.lineIndex-1 || symbol.line == number.lineIndex+1 {
		if symbol.index >= number.startIndex-1 && symbol.index <= number.endIndex+1 {
			return true
		}
	}
	return false
}

func isNumberAdjacentToAnySymbol(number Number, symbols []Symbol) bool {
	for _, symbol := range symbols {
		if isNumberAdjacentToSymbol(number, symbol) {
			return true
		}
	}
	return false
}

func extractNumbersInLine(line string, lineIndex int) []Number {
	var numbers = make([]Number, 0)

	var runes = []rune(line)
	var startIndex = -1
	var endIndex = -1

	var number = ""
	for i := 0; i < len(runes); {
		if unicode.IsDigit(runes[i]) {
			startIndex = i
			for i < len(runes) && unicode.IsDigit(runes[i]) {
				number += string(runes[i])
				i++
			}
			endIndex = startIndex + len(number) - 1
			numbers = append(numbers, Number{startIndex, endIndex, ParseInt(number), lineIndex})
			number = ""
		}
		i++
	}
	return numbers
}

func findNumbersAdjacentToSymbol(numbers []Number, symbol Symbol) []Number {
	var matches = make([]Number, 0)
	for _, number := range numbers {
		if isNumberAdjacentToSymbol(number, symbol) {
			matches = append(matches, number)
		}
	}
	return matches
}

func solutionPart01(lines []string) {
	var symbols = findAllSymbols(lines)
	var numbers = findAllNumbers(lines)

	var sum = 0
	for _, number := range numbers {
		fmt.Printf("line: %d, number: %d, isAdjacent: %t\n", number.lineIndex, number.value, isNumberAdjacentToAnySymbol(number, symbols))
		if isNumberAdjacentToAnySymbol(number, symbols) {
			sum += number.value
		}
	}

	fmt.Printf("%d", sum)
}

func solutionPart02(lines []string) {
	var numbers = findAllNumbers(lines)
	var symbols = findAllSymbols(lines)
	var gears = filterSymbolsByValue(symbols, "*")

	var sum = 0
	for _, gear := range gears {
		var matches = findNumbersAdjacentToSymbol(numbers, gear)
		if len(matches) == 2 {
			sum += matches[0].value * matches[1].value
		} else if len(matches) > 2 {
			fmt.Printf("ERROR: %d\n", len(matches))
		}
	}

	fmt.Printf("%d", sum)
}

// https://adventofcode.com/2023/day/3
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..")
	RunAdventOfCodeWithFile(solutionPart01, "day_03/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..")
	//RunAdventOfCodeWithFile(solutionPart02, "day_03/testcases/input-part-02.txt")
}
