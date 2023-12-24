package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

type Cell struct {
	value  string
	row    int
	column int
}

type Stack []Cell

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(cell Cell) {
	*s = append(*s, cell)
}

func (s *Stack) Pop() (Cell, bool) {
	if s.IsEmpty() {
		return Cell{}, false
	} else {
		var index = len(*s) - 1 // remove last element -> LIFO
		var element = (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func makeBooleanMatrix(maxRows int, maxCols int) [][]bool {
	var matrix = make([][]bool, maxRows)
	for i := 0; i < maxRows; i++ {
		matrix[i] = make([]bool, maxCols)
	}
	return matrix
}

func isStartPoint(cell Cell) bool {
	return cell.value == "S"
}

func findNeighbors(field [][]Cell, cell Cell, visited [][]bool) []Cell {
	var neighbors = make([]Cell, 0)

	// finding horizontal neighbors
	if isValidColumn(field, cell.column+1) {
		var nextCell = field[cell.row][cell.column+1]
		if nextCell.value != "." && !visited[nextCell.row][nextCell.column] && isNextCellValid(field, cell, nextCell) {
			neighbors = append(neighbors, nextCell)
		}
	}
	if isValidColumn(field, cell.column-1) {
		var nextCell = field[cell.row][cell.column-1]
		if nextCell.value != "." && !visited[nextCell.row][nextCell.column] && isNextCellValid(field, cell, nextCell) {
			neighbors = append(neighbors, nextCell)
		}
	}

	// finding vertical neighbors
	if isValidRow(field, cell.row+1) {
		var nextCell = field[cell.row+1][cell.column]
		if nextCell.value != "." && !visited[nextCell.row][nextCell.column] && isNextCellValid(field, cell, nextCell) {
			neighbors = append(neighbors, nextCell)
		}
	}
	if isValidRow(field, cell.row-1) {
		var nextCell = field[cell.row-1][cell.column]
		if nextCell.value != "." && !visited[nextCell.row][nextCell.column] && isNextCellValid(field, cell, nextCell) {
			neighbors = append(neighbors, nextCell)
		}
	}

	return neighbors
}

func isValidRow(field [][]Cell, row int) bool {
	var maxRows, _ = findMaxSizes(field)
	return row >= 0 && row < maxRows
}

func isValidColumn(field [][]Cell, col int) bool {
	var _, maxCols = findMaxSizes(field)
	return col >= 0 && col < maxCols
}

func findMaxSizes(field [][]Cell) (int, int) {
	// return (maxRows, maxCols)
	return len(field), len(field[0])
}

func isNextCellValid(field [][]Cell, current Cell, next Cell) bool {
	var valid, _ = nextCellByDir(field, current, next.row-current.row, next.column-current.column)
	return valid
}

func nextCellByDir(field [][]Cell, cell Cell, rowDir int, colDir int) (bool, Cell) {
	if colDir != 0 {
		var newColumn = cell.column + colDir
		if isValidColumn(field, newColumn) {
			var nextCell = field[cell.row][newColumn]
			// left to right
			if colDir > 0 && (nextCell.value == "-" || nextCell.value == "J" || nextCell.value == "7") {
				if cell.value == "S" || cell.value == "-" || cell.value == "L" || cell.value == "F" {
					return true, nextCell
				}
			}
			// right to left
			if colDir < 0 && (nextCell.value == "-" || nextCell.value == "F" || nextCell.value == "L") {
				if cell.value == "S" || cell.value == "-" || cell.value == "7" || cell.value == "J" {
					return true, nextCell
				}
			}
		}
		return false, cell
	}

	if rowDir != 0 {
		var newRow = cell.row + rowDir
		if isValidRow(field, newRow) {
			var nextCell = field[newRow][cell.column]
			// top to bottom
			if rowDir > 0 && (nextCell.value == "|" || nextCell.value == "J" || nextCell.value == "L") {
				if cell.value == "S" || cell.value == "|" || cell.value == "7" || cell.value == "F" {
					return true, nextCell
				}
			}
			// bottom to top
			if rowDir < 0 && (nextCell.value == "|" || nextCell.value == "F" || nextCell.value == "7") {
				if cell.value == "S" || cell.value == "|" || cell.value == "L" || cell.value == "J" {
					return true, nextCell
				}
			}
		}
		return false, cell
	}

	return false, cell
}

func isNeighborOfStartPoint(field [][]Cell, cell Cell) bool {

	// finding horizontal neighbors
	if isValidColumn(field, cell.column+1) {
		if isStartPoint(field[cell.row][cell.column+1]) {
			return true
		}
	}
	if isValidColumn(field, cell.column-1) {
		if isStartPoint(field[cell.row][cell.column-1]) {
			return true
		}
	}

	// finding vertical neighbors
	if isValidRow(field, cell.row+1) {
		if isStartPoint(field[cell.row+1][cell.column]) {
			return true
		}
	}
	if isValidRow(field, cell.row-1) {
		if isStartPoint(field[cell.row-1][cell.column]) {
			return true
		}
	}

	return false
}

func readInputLines(lines []string) (int, int, [][]Cell) {
	var field = make([][]Cell, 0)
	var startRow = 0
	var startCol = 0
	var found = false
	for r, line := range lines {
		var fieldRow = make([]Cell, 0)
		for c := range line {
			var cell = string(line[c])
			fieldRow = append(fieldRow, Cell{cell, r, c})
			if !found && cell == "S" {
				startRow = r
				startCol = c
				found = true
			}
		}
		field = append(field, fieldRow)
	}

	return startRow, startCol, field
}

func findPaths(startRow int, startCol int, field [][]Cell) [][]Cell {

	var startCell = field[startRow][startCol]
	var cell = Cell{}
	var path = make([]Cell, 0)
	var paths = make([][]Cell, 0)
	var visited = makeBooleanMatrix(findMaxSizes(field))

	var stack = Stack{}
	stack.Push(startCell)

	for !stack.IsEmpty() {
		cell, _ = stack.Pop()
		path = append(path, cell)
		visited[cell.row][cell.column] = true
		var neighbors = findNeighbors(field, cell, visited)
		if len(neighbors) != 0 {
			for _, neighbor := range neighbors {
				stack.Push(neighbor)
			}
		} else {
			// you finish your path and start over again from startCell
			if isNeighborOfStartPoint(field, cell) {
				path = append(path, startCell) // mark is a cycle
			}
			paths = append(paths, path)
			//fmt.Printf("%v\n", path)
			path = path[:0]
			path = append(path, startCell)

			visited = makeBooleanMatrix(findMaxSizes(field))
			visited[startCell.row][startCell.column] = true
			continue
		}
	}

	return paths
}

func solutionPart01(lines []string) {
	var startRow, startCol, field = readInputLines(lines)
	var paths = findPaths(startRow, startCol, field)

	var max = 0
	for i := range paths {
		var length = len(paths[i]) - 1
		if isStartPoint(paths[i][len(paths[i])-1]) {
			length = (len(paths[i]) - 1) / 2
		}

		if length > max {
			max = length
		}
	}

	fmt.Printf("%d\n", max)
}

func printMatrix(visited [][]bool, visitedSymbol string) {
	fmt.Printf("\n")
	for i := 0; i < len(visited); i++ {
		for j := 0; j < len(visited[i]); j++ {
			if visited[i][j] == true {
				fmt.Printf(visitedSymbol)
			} else {
				fmt.Printf("0")
			}
		}
		fmt.Printf("\n")
	}
}

func solutionPart02(lines []string) {
	var startRow, startCol, field = readInputLines(lines)
	var paths = findPaths(startRow, startCol, field)

	var visited = makeBooleanMatrix(findMaxSizes(field))

	for i := range paths {
		for _, cell := range paths[i] {
			visited[cell.row][cell.column] = true
		}
	}

	var matrix = make([]string, 0)
	for i := 0; i < len(visited); i++ {
		var line = ""
		for j := 0; j < len(visited[i]); j++ {
			if visited[i][j] {
				line += "1"
			} else {
				line += "0"
			}
		}
		matrix = append(matrix, line)
	}

	printMatrix(visited, "1")

	var count = 0
	var matrix2 = make([]string, 0)
	for _, matrixLine := range matrix {
		var line = ""
		for j := 0; j < len(matrixLine); j++ {
			if string(matrixLine[j]) == "0" {
				var index01 = findItemForward(matrixLine, "1", j+1)
				var index02 = findItemBackward(matrixLine, "1", j-1)
				if index01 >= 0 && index02 >= 0 {
					line += "0"
					count++
				} else {
					line += "1"
				}
			} else {
				line += "1"
			}
		}
		matrix2 = append(matrix2, line)
	}

	matrix = matrix2

	fmt.Printf("\n")
	for i := 0; i < len(matrix); i++ {
		fmt.Printf("%s\n", matrix[i])
	}

	fmt.Printf("%d", count)
}

func findItemForward(line string, search string, fromIndex int) int {
	for i := fromIndex; i < len(line); i++ {
		if i >= 0 && i < len(line) && string(line[i]) == search {
			return i
		}
	}
	return -1
}

func findItemBackward(line string, search string, fromIndex int) int {
	for i := fromIndex; i >= 0; i-- {
		if i >= 0 && i < len(line) && string(line[i]) == search {
			return i
		}
	}
	return -1
}

// https://adventofcode.com/2023/day/10
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, ".....\n.S-7.\n.|.|.\n.L-J.\n.....")
	//RunAdventOfCodeWithString(solutionPart01, "..F7.\n.FJ|.\nSJ.L7\n|F--J\nLJ...")
	RunAdventOfCodeWithFile(solutionPart01, "day_10/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, ".F----7F7F7F7F-7....\n.|F--7||||||||FJ....\n.||.FJ||||||||L7....\nFJL7L7LJLJ||LJ.L-7..\nL--J.L7...LJS7F-7L7.\n....F-J..F7FJ|L7L7L7\n....L7.F7||L7|.L7L7|\n.....|FJLJ|FJ|F7|.LJ\n....FJL-7.||.||||...\n....L---J.LJ.LJLJ...")
	//RunAdventOfCodeWithString(solutionPart02, "...........\n.S-------7.\n.|F-----7|.\n.||OOOOO||.\n.||OOOOO||.\n.|L-7OF-J|.\n.|II|O|II|.\n.L--JOL--J.\n.....O.....")
	//RunAdventOfCodeWithString(solutionPart02, ".F----7F7F7F7F-7....\n.|F--7||||||||FJ....\n.||.FJ||||||||L7....\nFJL7L7LJLJ||LJ.L-7..\nL--J.L7...LJS7F-7L7.\n....F-J..F7FJ|L7L7L7\n....L7.F7||L7|.L7L7|\n.....|FJLJ|FJ|F7|.LJ\n....FJL-7.||.||||...\n....L---J.LJ.LJLJ...")
	//RunAdventOfCodeWithString(solutionPart02, "FF7FSF7F7F7F7F7F---7\nL|LJ||||||||||||F--J\nFL-7LJLJ||||||LJL-77\nF--JF--7||LJLJIF7FJ-\nL---JF-JLJIIIIFJLJJ7\n|F|F-JF---7IIIL7L|7|\n|FFJF7L7F-JF7IIL---7\n7-L-JL7||F7|L7F-7F7|\nL.L7LFJ|||||FJL7||LJ\nL7JLJL-JLJLJL--JLJ.L")
	//RunAdventOfCodeWithFile(solutionPart02, "day_10/testcases/input-part-02.txt")
}
