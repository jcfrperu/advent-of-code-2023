package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"sort"
	"strings"
)

// types of hands
const (
	FiveKind  = "five of a kind "
	FourKind  = "four of a kind "
	FullHouse = "full house     "
	ThreeKind = "three of a kind"
	TwoPairs  = "two pair       "
	OnePair   = "one pair       "
	HighCard  = "high card      "
)

// GLOBAL gCardScore is the list of possible label with its respective score/value
var gCardScore = map[string]int{
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

// GLOBAL gHandScore is the list of possible types of hands with its respective score/value
var gHandScore = map[string]int{
	FiveKind:  100,
	FourKind:  90,
	FullHouse: 80,
	ThreeKind: 60,
	TwoPairs:  40,
	OnePair:   20,
	HighCard:  10,
}

// Game abstracts a hand with its bind (one line of the input file), additional field were added due calculations
type Game struct {
	hand        []string // list of 5 labels of the hand
	bind        int
	frequencies map[string]int
	labels      []string // non-duplicate and sorted labels in frequencies map
	handType    string
}

func newGame(handStr string, bind int) Game {
	var hand = Split(handStr, "")
	var frequencies, labels = calculateFrequencies(hand)
	return Game{
		hand:        hand,
		bind:        bind,
		frequencies: frequencies,
		labels:      labels,
		handType:    calculateGameType(frequencies, labels),
	}
}

func readGames(lines []string) []Game {
	var games = make([]Game, 0)

	for _, line := range lines {
		var bind = ParseInt(Split(line, " ")[1])
		var handStr = Trim(Split(line, " ")[0])
		games = append(games, newGame(handStr, bind))
	}

	return games
}

func calculateFrequencies(list []string) (map[string]int, []string) {
	var frequencies = make(map[string]int, len(list))
	for _, label := range list {
		frequencies[label]++
	}

	var labels = make([]string, 0, len(list)) // no duplicates labels
	for label := range frequencies {
		labels = append(labels, label)
	}

	// sort labels by frequency first, then by card value
	sort.Slice(labels, func(i, j int) bool {
		if frequencies[labels[i]] == frequencies[labels[j]] {
			return gCardScore[labels[i]] > gCardScore[labels[j]]
		}
		return frequencies[labels[i]] > frequencies[labels[j]]
	})

	return frequencies, labels
}

func hasFrequencyOf(frequencies map[string]int, labels []string, n int) bool {
	return countLabelsWithFrequency(frequencies, labels, n) > 0
}

func countLabelsWithFrequency(frequencies map[string]int, labels []string, frequency int) int {
	var count = 0
	for _, label := range labels {
		if frequencies[label] == frequency {
			count++
		}
	}
	return count
}

func calculateGameType(frequencies map[string]int, labels []string) string {
	switch {
	case hasFrequencyOf(frequencies, labels, 5):
		return FiveKind
	case hasFrequencyOf(frequencies, labels, 4):
		return FourKind
	case hasFrequencyOf(frequencies, labels, 3) && hasFrequencyOf(frequencies, labels, 2):
		return FullHouse
	case hasFrequencyOf(frequencies, labels, 3):
		return ThreeKind
	case countLabelsWithFrequency(frequencies, labels, 2) >= 2:
		return TwoPairs
	case hasFrequencyOf(frequencies, labels, 2):
		return OnePair
	default:
		return HighCard
	}
}

func sortByRank(games []Game) {
	var HS = gHandScore
	var CS = gCardScore

	sort.Slice(games, func(i, j int) bool {
		// first by hand type
		if HS[games[i].handType] == HS[games[j].handType] {
			// then by label
			for k := 0; k < len(games[i].hand); k++ {
				if CS[games[i].hand[k]] == CS[games[j].hand[k]] {
					continue
				}
				return CS[games[i].hand[k]] < CS[games[j].hand[k]]
			}
		}
		return HS[games[i].handType] < HS[games[j].handType]
	})
}

func improveHandUsingJoker(game Game) string {
	if game.frequencies["J"] == 0 {
		return strings.Join(game.hand, "")
	}

	if game.frequencies["J"] == 5 {
		return strings.Repeat("A", 5) // create a FiveKind with the best highest ("A")
	}

	if game.frequencies["J"] == 4 {
		// check the one remaining card
		for _, label := range game.labels {
			if label != "J" {
				return strings.Repeat(label, 5) // create a FiveKind with the first non-J label
			}
		}
		return strings.Join(game.hand, "")
	}

	if game.frequencies["J"] == 3 {
		// check the two remaining cards to find a pair
		for _, label := range game.labels {
			if label != "J" && game.frequencies[label] == 2 {
				// create a FiveKind with the existing pair
				return strings.Repeat(label, 5)
			}
		}
		// "J" has 3 matches and the remaining two are not a pair, we can create a FourKind
		// in card.labels label "J" will be at beginning (label[0]) since labels are sorted by frequencies.
		// the highest label and second-highest label will be the next two labels after "J"
		var highest = game.labels[1]
		var secondHighest = game.labels[2]
		return strings.Repeat(highest, 4) + secondHighest
	}

	var allCombinations = make([]Game, 0, 120)
	if game.frequencies["J"] == 2 {
		for label01 := range gCardScore {
			for label02 := range gCardScore {
				if label01 != "J" && label02 != "J" {
					var handStr = strings.ReplaceAll(strings.Join(game.hand, ""), "J", "")
					var combination = newGame(handStr+label01+label02, 0)
					allCombinations = append(allCombinations, combination)
				}
			}
		}
	} else {
		if game.frequencies["J"] == 1 {
			for label01 := range gCardScore {
				if label01 != "J" {
					var handStr = strings.ReplaceAll(strings.Join(game.hand, ""), "J", "")
					var combination = newGame(handStr+label01, 0)
					allCombinations = append(allCombinations, combination)
				}
			}
		}
	}

	var bestGame Game
	var highest = 0
	for _, possibleGame := range allCombinations {
		if gHandScore[possibleGame.handType] > highest {
			highest = gHandScore[possibleGame.handType]
			bestGame = possibleGame
		}
	}

	return strings.Join(bestGame.hand, "")
}

func solutionPart01(lines []string) {
	var games = readGames(lines)

	sortByRank(games)

	var sum = 0
	for i := range games {
		// fmt.Printf("hand=%v, type=%s, labels=%v, bind=%d\n", games[i].hand, games[i].handType, games[i].labels, games[i].bind)
		sum += (i + 1) * games[i].bind
	}

	fmt.Printf("%d", sum)
}

func solutionPart02(lines []string) {
	gCardScore["J"] = 1

	var games = readGames(lines)

	for i := range games {
		var bestHandWithJoker = improveHandUsingJoker(games[i])
		var bestGameWithJoker = newGame(bestHandWithJoker, 0)
		games[i].handType = calculateGameType(bestGameWithJoker.frequencies, bestGameWithJoker.labels)
	}

	sortByRank(games)

	var sum = 0
	for i := range games {
		// fmt.Printf("hand=%v, type=%s, labels=%v, bind=%d\n", games[i].hand, games[i].handType, games[i].labels, games[i].bind)
		sum += (i + 1) * games[i].bind
	}

	fmt.Printf("%d", sum)
}

// https://adventofcode.com/2023/day/7
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483")
	RunAdventOfCodeWithFile(solutionPart01, "day_07/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483")
	//RunAdventOfCodeWithFile(solutionPart02, "day_07/testcases/input-part-02.txt") // 255086951 too high  254837398
}
