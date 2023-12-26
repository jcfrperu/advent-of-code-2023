package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"strings"
)

type Cell struct {
	value string
	row   int
	col   int
}

func getColValues(M [][]Cell, c int) string {
	maxRows, _ := findMatrixSizes(M)
	values := make([]string, 0)
	for r := 0; r < maxRows; r++ {
		values = append(values, M[r][c].value)
	}
	return strings.Join(values, "")
}

func getRowValues(M [][]Cell, r int) string {
	_, maxCols := findMatrixSizes(M)
	values := make([]string, 0)
	for c := 0; c < maxCols; c++ {
		values = append(values, M[r][c].value)
	}
	return strings.Join(values, "")
}

func findMatrixSizes(matrix [][]Cell) (int, int) {
	return len(matrix), len(matrix[0])
}

func readInputLines(lines []string) [][]Cell {
	var field = make([][]Cell, 0)
	for r, line := range lines {
		var fieldRow = make([]Cell, 0)
		for c := range line {
			var value = string(line[c])
			var cell = Cell{value, r, c}
			fieldRow = append(fieldRow, cell)
		}
		field = append(field, fieldRow)
	}

	return field
}

func readingAllMatrices(lines []string) [][][]Cell {
	// matrixList: even position = vertically, odd position = horizontally
	matrixList := make([][][]Cell, 0)
	startIndex := 0
	for i := 0; i < len(lines); i++ {
		if IsBlank(lines[i]) {
			matrix := readInputLines(lines[startIndex:i])
			matrixList = append(matrixList, matrix)
			startIndex = i + 1
		}
	}
	matrix := readInputLines(lines[startIndex:])
	matrixList = append(matrixList, matrix)

	return matrixList
}

func solutionPart01(lines []string) {
	total := 0
	matrixList := readingAllMatrices(lines)
	for i := range matrixList {
		verticalResult := processVertReflect(matrixList[i])
		horizontalResult := processHorizReflect(matrixList[i])
		total += verticalResult + 100*horizontalResult

		fmt.Printf("matrix %d VERT: %d\n", i, verticalResult)
		fmt.Printf("matrix %d HORI: %d\n", i, horizontalResult)
	}

	fmt.Printf("%d", total)
}

func processVertReflect(matrix [][]Cell) int {
	_, maxCols := findMatrixSizes(matrix)

	adjacentReflect := make([]int, 0)
	for c := 1; c < maxCols; c++ {
		if getColValues(matrix, c) == getColValues(matrix, c-1) {
			adjacentReflect = append(adjacentReflect, c)
		}
	}

	if len(adjacentReflect) > 0 {
		for _, col := range adjacentReflect {
			actualMatches := 0
			expectedMatches := min(col, maxCols-col)

			left := col - 1
			right := col
			for left >= 0 && right < maxCols && getColValues(matrix, left) == getColValues(matrix, right) {
				left--
				right++
				actualMatches++
			}

			if actualMatches == expectedMatches {
				return col
			}
		}
	}

	return 0
}

func processHorizReflect(matrix [][]Cell) int {
	maxRows, _ := findMatrixSizes(matrix)

	adjacentReflect := make([]int, 0)
	for r := 1; r < maxRows; r++ {
		if getRowValues(matrix, r) == getRowValues(matrix, r-1) {
			adjacentReflect = append(adjacentReflect, r)
		}
	}

	if len(adjacentReflect) > 0 {
		for _, row := range adjacentReflect {
			actualMatches := 0
			expectedMatches := min(row, maxRows-row)

			up := row - 1
			down := row
			for up >= 0 && down < maxRows && getRowValues(matrix, up) == getRowValues(matrix, down) {
				up--
				down++
				actualMatches++
			}

			if actualMatches == expectedMatches {
				return row
			}
		}
	}

	return 0
}

func solutionPart02(lines []string) {

	fmt.Printf("%d", 0)
}

// https://adventofcode.com/2023/day/13
func main() {
	// part 01: using string or input file
	RunAdventOfCodeWithString(solutionPart01, "#.##..##.\n..#.##.#.\n##......#\n##......#\n..#.##.#.\n..##..##.\n#.#.##.#.\n\n#...##..#\n#....#..#\n..##..###\n#####.##.\n#####.##.\n..##..###\n#....#..#")
	//RunAdventOfCodeWithString(solutionPart01, ".##...#..#...##..\n##..#.####.#..###\n#......#.......##\n##..#..##..#..###\n..###..##..###...\n#..#..####..#..##\n#.##.#....#.##.##\n#.##.#....#.##.##\n.###.#....#.###..")
	//RunAdventOfCodeWithString(solutionPart01, "#.######.\n#.#....##\n.#..#.###\n.#..#.###\n#.#....##\n#.######.\n......#.#\n.#....#.#\n#.######.\n#.#....##\n.#..#.###")
	//RunAdventOfCodeWithString(solutionPart01, "####...#.#.......\n#.#.####.####...#\n..#.##.#...##....\n..#.##.#...##....\n#.#.####.####...#\n####...#.#.......\n#..#.#..#.#.#.##.\n.#....#....#..###\n..#.##.#...#.#.#.\n..#.##.#...#.#.#.\n.##...#....#..###\n#..#.#..#.#.#.##.\n####...#.#.......")
	//RunAdventOfCodeWithString(solutionPart01, "..##.##.##...\n##.#......###\n..########...\n#####..######\n..#......#...\n.####..####..\n#....##....##\n....#..#.....\n#...####...##\n##.#....#.###\n#...#..#...##\n#####..######\n##..#..#..###\n.#...##...#..\n..#.#..#.#...")
	//RunAdventOfCodeWithFile(solutionPart01, "day_13/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "")
	//RunAdventOfCodeWithFile(solutionPart02, "day_13/testcases/input-part-02.txt")
}
