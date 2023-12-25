package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"strings"
)

var SEEN = make(map[string]bool)

type Cell struct {
	value string
	row   int
	col   int
}

func (c Cell) posAsStr() string {
	return FormatInt(c.row) + "-" + FormatInt(c.col)
}

func (c Cell) isSeen() bool {
	return SEEN[c.posAsStr()]
}

func findMaxSizes(field [][]Cell) (int, int) {
	// return (maxRows, maxCols)
	return len(field), len(field[0])
}

func readInputLines(lines []string) [][]Cell {
	var field = make([][]Cell, 0)
	for r, line := range lines {
		var fieldRow = make([]Cell, 0)
		for c := range line {
			var value = string(line[c])
			var cell = Cell{value, r, c}
			SEEN[cell.posAsStr()] = false
			fieldRow = append(fieldRow, cell)
		}
		field = append(field, fieldRow)
	}

	return field
}

func addNewLineRowWith(lines []string, val string) []string {
	line := strings.Repeat(val, len(lines[0]))
	lines = append([]string{line}, lines[:]...)
	return lines
}

func solutionPart01(lines []string) {
	var field = readInputLines(addNewLineRowWith(lines, "#"))
	var _, maxCols = findMaxSizes(field)

	var amount = 0
	for col := 0; col < maxCols; col++ {
		cubeShapeRow := findNextCubeShape(col, field)
		for cubeShapeRow >= 0 {
			allRocks := findNotSeenRocksDown(field[cubeShapeRow][col], field)
			for i, _ := range allRocks {
				amount += maxCols - (cubeShapeRow + i)
			}
			cubeShapeRow = findNextCubeShape(col, field)
		}
	}

	fmt.Printf("%d", amount)
}

func findNextCubeShape(col int, field [][]Cell) int {
	var maxRows, _ = findMaxSizes(field)

	var row = maxRows - 1
	for row >= 0 {
		if field[row][col].value == "#" && !SEEN[field[row][col].posAsStr()] {
			return row
		}
		row--
	}

	return -1
}

func findNotSeenRocksDown(init Cell, field [][]Cell) []Cell {
	var maxRows, _ = findMaxSizes(field)

	var col = init.col
	SEEN[init.posAsStr()] = true

	var matches = make([]Cell, 0)

	var row = init.row + 1
	for row >= 0 && row < maxRows {
		if field[row][col].value == "O" && !SEEN[field[row][col].posAsStr()] {
			matches = append(matches, field[row][col])
			SEEN[field[row][col].posAsStr()] = true
		}
		row++
	}

	return matches
}

func solutionPart02(lines []string) {
	var field = readInputLines(lines)

	fmt.Printf("%v", field)
}

// https://adventofcode.com/2023/day/14
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "O....#....\nO.OO#....#\n.....##...\nOO.#O....O\n.O.....O#.\nO.#..O.#.#\n..O..#O..O\n.......O..\n#....###..\n#OO..#....")
	RunAdventOfCodeWithFile(solutionPart01, "day_14/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "")
	//RunAdventOfCodeWithFile(solutionPart02, "day_14/testcases/input-part-02.txt")
}
