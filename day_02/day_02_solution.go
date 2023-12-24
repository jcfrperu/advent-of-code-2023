package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solutionPart01(lines []string) {
	var result = 0
	var limitColorMap = map[string]int{"red": 12, "green": 13, "blue": 14}

	for _, line := range lines {

		var cubeSetIsValid = true
		var gameID = ParseInt(Split(Split(line, ":")[0], " ")[1])

		var cubeSets = Split(Split(line, ":")[1], ";") // cubeSets (; separated):  cubeSet01 ; cubeSet02 ; ...
		for _, cubeSet := range cubeSets {

			var gameMap = make(map[string]int)

			var cubes = Split(cubeSet, ",") // cubeSet (, separated): cube01, cube02, cube 02 = 3 green, 10 blue, 2 red
			for _, cube := range cubes {
				cube = Trim(cube)
				var number = ParseInt(Split(cube, " ")[0])
				var color = Trim(Split(cube, " ")[1])

				gameMap[color] += number
			}

			// checking if cube set is valid
			for color, _ := range limitColorMap {
				if gameMap[color] > limitColorMap[color] {
					cubeSetIsValid = false
					break
				}
			}

			if !cubeSetIsValid {
				break
			}
		}

		if cubeSetIsValid {
			result += gameID
		}
	}

	fmt.Printf("%d", result)
}

func solutionPart02(lines []string) {
	var result = 0
	for _, line := range lines {

		var maxColorMap = map[string]int{"red": 0, "green": 0, "blue": 0}

		var cubeSets = Split(Split(line, ":")[1], ";") // cubeSets (; separated):  cubeSet01 ; cubeSet02 ; ...
		for _, cubeSet := range cubeSets {

			var cubes = Split(cubeSet, ",") // cubeSet (, separated): cube01, cube02, cube 02 = 3 green, 10 blue, 2 red
			for _, cube := range cubes {
				cube = Trim(cube)
				var number = ParseInt(Trim(Split(cube, " ")[0]))
				var color = Trim(Split(cube, " ")[1])

				// we need to find the max number per color
				if number > maxColorMap[color] {
					maxColorMap[color] = number
				}
			}
		}

		// multiplying number of each color
		var product = 1
		for color, _ := range maxColorMap {
			product *= maxColorMap[color]
		}

		result += product
	}

	fmt.Printf("%d", result)
}

// https://adventofcode.com/2023/day/2
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green")
	RunAdventOfCodeWithFile(solutionPart01, "day_02/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green")
	//RunAdventOfCodeWithFile(solutionPart02, "day_02/testcases/input-part-02.txt")
}
