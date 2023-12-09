package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"strings"
)

type ProblemMap struct {
	instructions []string
	places       map[string]Destiny // place["AAA"] map{ "L" -> "BBB", "R"-> "CCC" }
}

func readProblemMap(lines []string) ProblemMap {
	var instructions = Split(Trim(lines[0]), "")

	var places = make(map[string]Destiny)
	for i := 2; i < len(lines); i++ {
		var line = Trim(lines[i])

		var place = Trim(Split(line, "=")[0])

		var right = Trim(Split(line, " = ")[1])
		right = strings.ReplaceAll(right, ")", "")
		right = strings.ReplaceAll(right, "(", "")

		places[place] = Destiny{
			nextPlace: map[string]string{
				"L": Trim(Split(right, ",")[0]),
				"R": Trim(Split(right, ",")[1]),
			},
		}
	}

	return ProblemMap{instructions, places}
}

type Destiny struct {
	nextPlace map[string]string
}

func isEqualTo(currentPlace string, equalTo string) bool {
	return strings.HasSuffix(currentPlace, equalTo)
}

func findAllPointsEndsWith(problemMap ProblemMap, ends string) []string {
	var placesMap = make(map[string]string)

	for place, _ := range problemMap.places {
		if strings.HasSuffix(place, ends) {
			placesMap[place] = place
		}
	}

	var matches = make([]string, 0)
	for place, _ := range placesMap {
		matches = append(matches, place)
	}
	return matches
}

func solutionPart01(lines []string) {
	var problemMap = readProblemMap(lines)

	var count = 0
	var currentPlace = "AAA" // start point
	for !isEqualTo(currentPlace, "ZZZ") {
		for _, inst := range problemMap.instructions {
			count++
			var nextPlace = problemMap.places[currentPlace].nextPlace[inst]
			currentPlace = nextPlace
			if isEqualTo(currentPlace, "ZZZ") {
				break
			}
		}
	}
	fmt.Printf("%d", count)
}

func solutionPart02(lines []string) {
	var problemMap = readProblemMap(lines)
	var startPoints = findAllPointsEndsWith(problemMap, "A")
	var instSize = len(problemMap.instructions)

	var counts = make([]int, len(startPoints))
	var currentPlaces = make(map[int]string)
	for i, _ := range startPoints {
		currentPlaces[i] = startPoints[i]
		for !strings.HasSuffix(currentPlaces[i], "Z") {
			var instruction = problemMap.instructions[counts[i]%instSize]
			currentPlaces[i] = problemMap.places[currentPlaces[i]].nextPlace[instruction]
			counts[i]++
		}
	}

	var result = 1
	for _, count := range counts {
		result = LCMInt(result, count)
	}
	fmt.Printf("%d", result)
}

// https://adventofcode.com/2023/day/8
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "RL\n\nAAA = (BBB, CCC)\nBBB = (DDD, EEE)\nCCC = (ZZZ, GGG)\nDDD = (DDD, DDD)\nEEE = (EEE, EEE)\nGGG = (GGG, GGG)\nZZZ = (ZZZ, ZZZ")
	//RunAdventOfCodeWithString(solutionPart01, "LLR\n\nAAA = (BBB, BBB)\nBBB = (AAA, ZZZ)\nZZZ = (ZZZ, ZZZ)")
	RunAdventOfCodeWithFile(solutionPart01, "day_08/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "LR\n\n11A = (11B, XXX)\n11B = (XXX, 11Z)\n11Z = (11B, XXX)\n22A = (22B, XXX)\n22B = (22C, 22C)\n22C = (22Z, 22Z)\n22Z = (22B, 22B)\nXXX = (XXX, XXX)")
	//RunAdventOfCodeWithFile(solutionPart02, "day_08/testcases/input-part-02.txt")
}
