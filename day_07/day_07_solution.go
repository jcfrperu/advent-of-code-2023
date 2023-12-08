package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"sort"
)

type Card struct {
	hand             string
	handListUnsorted []string
	handListSorted   []string
	frequenciesMap   map[string]int
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

func createGameTypeSet() map[string]GameType {
	var gameTypesList = []GameType{
		{"five of a kind", 100},
		{"four of a kind", 90},
		{"full house", 80},
		{"three of a kind", 70},
		{"two pair", 60},
		{"one pair", 50},
		{"high card", 40},
	}

	var gameTypesMap = make(map[string]GameType)
	for _, gameType := range gameTypesList {
		gameTypesMap[gameType.name] = gameType
	}
	return gameTypesMap
}

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

func readGames(lines []string) []Game {
	var gamesTypesSet = createGameTypeSet()
	var games = make([]Game, 0)
	for _, line := range lines {

		var hand = Trim(Split(line, " ")[0])
		var handAsList = Split(hand, "")
		var bind = ParseInt(Split(line, " ")[1])
		var frequenciesList, frequenciesMap = Frequencies(handAsList)

		games = append(games, Game{
			card: Card{
				hand:             hand,
				handListUnsorted: handAsList,
				handListSorted:   frequenciesList,
				frequenciesMap:   frequenciesMap,
			},
			bind:     bind,
			gameType: gamesTypesSet["high card"], // by default the lowest
		})
	}

	return games
}

func sortFrequencies(game *Game) {
	// update list of hands in order of frequencies
	var handListSorted = game.card.handListSorted
	var freqMap = game.card.frequenciesMap

	// we need the card set because "K" is better than "2"
	// and we want to consider that in the sorting when cards has the same frequency
	var cardSet = createCardSet()

	sort.SliceStable(handListSorted, func(i, j int) bool {
		if freqMap[handListSorted[i]] == freqMap[handListSorted[j]] {
			return cardSet[handListSorted[i]] > cardSet[handListSorted[j]]
		} else {
			return freqMap[handListSorted[i]] > freqMap[handListSorted[j]]
		}
	})
}

func hasFrequencyOf(game *Game, n int) bool {
	return countMatchesFrequencyOf(game, n) > 0
}

func countMatchesFrequencyOf(game *Game, n int) int {
	var count = 0
	var freqMap = game.card.frequenciesMap
	for _, hand := range game.card.handListSorted {
		if freqMap[hand] == n {
			count++
		}
	}
	return count
}

func calculateGameType(game *Game) {
	var gamesTypesSet = createGameTypeSet()

	if hasFrequencyOf(game, 5) {
		game.gameType = gamesTypesSet["five of a kind"]
		return
	}

	if hasFrequencyOf(game, 4) {
		game.gameType = gamesTypesSet["four of a kind"]
		return
	}

	if hasFrequencyOf(game, 3) && hasFrequencyOf(game, 2) {
		game.gameType = gamesTypesSet["full house"]
		return
	}

	if hasFrequencyOf(game, 3) {
		game.gameType = gamesTypesSet["three of a kind"]
		return
	}

	if countMatchesFrequencyOf(game, 2) >= 2 {
		game.gameType = gamesTypesSet["two pair"]
		return
	}

	if hasFrequencyOf(game, 2) {
		game.gameType = gamesTypesSet["one pair"]
		return
	}
	// it is already "high card"  by default
}

func sortByRank(games []Game) {
	var cardSet = createCardSet()

	// rank is the order of the list of games: first consider the type of hand, then the highest label
	sort.SliceStable(games, func(i, j int) bool {
		if games[i].gameType.points == games[j].gameType.points {
			for k, _ := range games[i].card.handListUnsorted {
				if cardSet[games[i].card.handListUnsorted[k]] == cardSet[games[j].card.handListUnsorted[k]] {
					continue
				}
				return cardSet[games[i].card.handListUnsorted[k]] < cardSet[games[j].card.handListUnsorted[k]]
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
		fmt.Printf("RANK %d, HAND: %s, TYPE: %s\n", i+1, game.card.hand, game.gameType.name)
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
