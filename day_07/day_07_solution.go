package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"sort"
)

type Card struct {
	hand     string
	handList []string
	freqMap  map[string]int
	freqKeys []string
}

type Game struct {
	card     Card
	bind     int
	gameType GameType
}

type GameType struct {
	name   string
	points int // just a reference, strongest hand means more points
}

var gCardSet = createCardSet()
var gGameTypeSet = createGameTypeSet()

func createCardSet() map[string]int {
	return map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
}

func createGameTypeSet() map[string]GameType {
	var gameTypes = []GameType{
		{"five of a kind", 100},
		{"four of a kind", 90},
		{"full house", 80},
		{"three of a kind", 70},
		{"two pair", 60},
		{"one pair", 50},
		{"high card", 40},
	}

	var gameTypesMap = make(map[string]GameType)
	for _, gameType := range gameTypes {
		gameTypesMap[gameType.name] = gameType
	}
	return gameTypesMap
}

func readGames(lines []string) []Game {
	var games = make([]Game, 0)
	for _, line := range lines {
		var hand = Trim(Split(line, " ")[0])
		var bind = ParseInt(Split(line, " ")[1])
		var handList = Split(hand, "")
		var freqMap, freqKeys = Frequencies(handList, false)

		games = append(games, Game{
			card: Card{
				hand:     hand,
				handList: handList,
				freqMap:  freqMap,
				freqKeys: freqKeys,
			},
			bind:     bind,
			gameType: gGameTypeSet["high card"], // by default the lowest
		})
	}

	return games
}

func sortFrequencies(game *Game) {
	var freqMap = game.card.freqMap
	var freqKeys = game.card.freqKeys

	// update freqKeys order, using frequency matches first, then card label value
	// example: if we found 2 cards with same label, we need to validate the label since "K" is better than "2"
	sort.SliceStable(freqKeys, func(i, j int) bool {
		if freqMap[freqKeys[i]] == freqMap[freqKeys[j]] {
			return gCardSet[freqKeys[i]] > gCardSet[freqKeys[j]]
		} else {
			return freqMap[freqKeys[i]] > freqMap[freqKeys[j]]
		}
	})
}

func hasFrequencyOf(game *Game, n int) bool {
	return countMatchesFrequencyOf(game, n) > 0
}

func countMatchesFrequencyOf(game *Game, n int) int {
	var count = 0
	var freqMap = game.card.freqMap
	for _, hand := range game.card.freqKeys {
		if freqMap[hand] == n {
			count++
		}
	}
	return count
}

func calculateGameType(game *Game) {
	if hasFrequencyOf(game, 5) {
		game.gameType = gGameTypeSet["five of a kind"]
		return
	}

	if hasFrequencyOf(game, 4) {
		game.gameType = gGameTypeSet["four of a kind"]
		return
	}

	if hasFrequencyOf(game, 3) && hasFrequencyOf(game, 2) {
		game.gameType = gGameTypeSet["full house"]
		return
	}

	if hasFrequencyOf(game, 3) {
		game.gameType = gGameTypeSet["three of a kind"]
		return
	}

	if countMatchesFrequencyOf(game, 2) >= 2 {
		game.gameType = gGameTypeSet["two pair"]
		return
	}

	if hasFrequencyOf(game, 2) {
		game.gameType = gGameTypeSet["one pair"]
		return
	}
	// it is already "high card"  by default
}

func sortByRank(games []Game) {
	// rank is the order of the list of games: first consider the type of hand, then the highest label
	sort.SliceStable(games, func(i, j int) bool {
		if games[i].gameType.points == games[j].gameType.points {
			for k, _ := range games[i].card.handList {
				if gCardSet[games[i].card.handList[k]] == gCardSet[games[j].card.handList[k]] {
					continue
				}
				return gCardSet[games[i].card.handList[k]] < gCardSet[games[j].card.handList[k]]
			}
			return true // doesn't matter
		} else {
			return games[i].gameType.points < games[j].gameType.points
		}
	})
}

func calculateRanks(games []Game) {
	for i, _ := range games {
		var game = &games[i]
		sortFrequencies(game)
	}

	for i, _ := range games {
		var game = &games[i]
		calculateGameType(game)
	}

	sortByRank(games)
}

func solutionPart01(lines []string) {
	var games = readGames(lines)

	calculateRanks(games)

	// the rank is the index (is sorted by rank)
	var sum = 0
	for i, game := range games {
		// fmt.Printf("RANK %d, HAND: %s, TYPE: %s\n", i+1, game.card.hand, game.gameType.name)
		sum += (i + 1) * game.bind
	}

	fmt.Printf("%d", sum)
}

func solutionPart02(lines []string) {

	fmt.Printf("%d", 1)
}

// https://adventofcode.com/2023/day/6
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483")
	RunAdventOfCodeWithFile(solutionPart01, "day_07/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "")
	//RunAdventOfCodeWithFile(solutionPart02, "day_07/testcases/input-part-02.txt")
}
