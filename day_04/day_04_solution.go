package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"strings"
)

type Card struct {
	cardID       int
	winners      []int
	myNumbers    []int
	matchesCount int
}

func createCard(line string) Card {
	var gameLine = Split(line, ":")

	var cardID = ParseInt(strings.ReplaceAll(gameLine[0], "Card ", ""))

	var listsLine = Split(gameLine[1], "|")
	var winners = SplitInts(listsLine[0], " ")
	var myNumbers = SplitInts(listsLine[1], " ")

	return Card{cardID, winners, myNumbers, 0}
}

func findAllCards(lines []string, findMatches bool) []Card {
	var cards = make([]Card, 0)
	for _, line := range lines {
		var card = createCard(line)
		if findMatches {
			var matchesCount = 0
			for _, winner := range card.winners {
				var index = Find[int](card.myNumbers, winner)
				if index >= 0 {
					matchesCount++
				}
			}
			card.matchesCount = matchesCount
		}
		cards = append(cards, card)
	}
	return cards
}

func solutionPart01(lines []string) {
	var sum = 0
	var cards = findAllCards(lines, false)
	for _, card := range cards {
		var count = 0
		var firstMatch = true
		for _, winner := range card.winners {
			if Find[int](card.myNumbers, winner) >= 0 {
				if firstMatch {
					count = 1
					firstMatch = false
				} else {
					count = 2 * count
				}
			}
		}
		sum += count
	}
	fmt.Printf("%d", sum)
}

func solutionPart02(lines []string) {
	var stack = make(map[int]int)
	var cards = findAllCards(lines, true)

	for _, card := range cards {
		var count = 1
		for count <= card.matchesCount {
			stack[card.cardID+count]++
			count++
		}

		if stack[card.cardID] > 0 {
			var stackCount = 1
			for stackCount <= card.matchesCount {
				stack[card.cardID+stackCount] += stack[card.cardID]
				stackCount++
			}
		}
	}

	var sum = len(cards)
	for _, value := range stack {
		sum += value
	}

	fmt.Printf("%d", sum)
}

// https://adventofcode.com/2023/day/4
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11")
	RunAdventOfCodeWithFile(solutionPart01, "day_04/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11")
	//RunAdventOfCodeWithFile(solutionPart02, "day_04/testcases/input-part-02.txt")
}
