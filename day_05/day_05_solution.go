package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"sort"
	"strings"
)

type Data struct {
	destRangeStart int64
	srcRangeStart  int64
	rangeDestSrc   int64
}

func fillMap(startIndex int, lines []string) []Data {
	var dataList = make([]Data, 0)
	var i = startIndex + 1

	for i < len(lines) && Trim(lines[i]) != "" {
		var destRangeStart = SplitGetLongAt(lines[i], " ", 0) // example 50
		var srcRangeStart = SplitGetLongAt(lines[i], " ", 1)  // example 98
		var rangeDestSrc = SplitGetLongAt(lines[i], " ", 2)   // example 2
		dataList = append(dataList, Data{destRangeStart, srcRangeStart, rangeDestSrc})
		i++
	}

	// sorting by srcRangeStart (need it for binary search in part 02 of problem)
	sort.Slice(dataList, func(i, j int) bool {
		return dataList[i].srcRangeStart < dataList[j].srcRangeStart
	})

	return dataList
}

func readCatalog(lines []string) ([]int64, map[string][]Data) {
	var seeds = make([]int64, 0)
	var catalog = make(map[string][]Data)

	for i, line := range lines {
		if i == 0 {
			seeds = SplitLongs(Split(line, ":")[1], " ")
		}
		if strings.Contains(line, "seed-to-soil") {
			catalog["seed-to-soil"] = fillMap(i, lines)
			continue
		}
		if strings.Contains(line, "soil-to-fertilizer") {
			catalog["soil-to-fertilizer"] = fillMap(i, lines)
			continue
		}
		if strings.Contains(line, "fertilizer-to-water") {
			catalog["fertilizer-to-water"] = fillMap(i, lines)
			continue
		}
		if strings.Contains(line, "water-to-light") {
			catalog["water-to-light"] = fillMap(i, lines)
			continue
		}
		if strings.Contains(line, "light-to-temperature") {
			catalog["light-to-temperature"] = fillMap(i, lines)
			continue
		}
		if strings.Contains(line, "temperature-to-humidity") {
			catalog["temperature-to-humidity"] = fillMap(i, lines)
			continue
		}
		if strings.Contains(line, "humidity-to-location") {
			catalog["humidity-to-location"] = fillMap(i, lines)
			continue
		}
	}

	return seeds, catalog
}

func findDataIndexBS(dataList []Data, value int64) int {
	var left = 0
	var right = len(dataList) - 1

	for left <= right {
		var mid = (left + right) / 2

		if dataList[mid].srcRangeStart <= value && value < dataList[mid].srcRangeStart+dataList[mid].rangeDestSrc {
			return mid
		} else if dataList[mid].srcRangeStart < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func getNextValue(dataList []Data, value int64) int64 {
	var index = findDataIndexBS(dataList, value)

	if index >= 0 && index < len(dataList) {
		return dataList[index].destRangeStart + (value - dataList[index].srcRangeStart)
	}

	return value
}

func findLocation(catalog map[string][]Data, seed int64) int64 {
	var soil = getNextValue(catalog["seed-to-soil"], seed)
	var fertilizer = getNextValue(catalog["soil-to-fertilizer"], soil)
	var water = getNextValue(catalog["fertilizer-to-water"], fertilizer)
	var light = getNextValue(catalog["water-to-light"], water)
	var temperature = getNextValue(catalog["light-to-temperature"], light)
	var humidity = getNextValue(catalog["temperature-to-humidity"], temperature)
	var location = getNextValue(catalog["humidity-to-location"], humidity)

	return location
}

func solutionPart01(lines []string) {
	var seeds, catalog = readCatalog(lines)

	var minLoc = int64(0)
	var firstTime = true
	for _, seed := range seeds {
		var location = findLocation(catalog, seed)

		if firstTime {
			minLoc = location
			firstTime = false
		} else if location < minLoc {
			minLoc = location
		}
	}

	fmt.Printf("%d", minLoc)
}

func solutionPart02(lines []string) {
	var seeds, catalog = readCatalog(lines)

	var i int64 = 0
	var minLoc = int64(0)
	var firstTime = true

	for i = 0; i < int64(len(seeds))-1; i = i + 2 {
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			var location = findLocation(catalog, seed)

			if firstTime {
				minLoc = location
				firstTime = false
			} else if location < minLoc {
				minLoc = location
			}
		}
	}

	fmt.Printf("%d", minLoc)
}

// https://adventofcode.com/2023/day/5
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4")
	RunAdventOfCodeWithFile(solutionPart01, "day_05/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4")
	//RunAdventOfCodeWithFile(solutionPart02, "day_05/testcases/input-part-02.txt")
}
