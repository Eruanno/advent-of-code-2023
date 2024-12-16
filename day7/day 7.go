package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Figure struct {
	Name     string
	Strength int
}

var figures = []Figure{
	{Name: "A", Strength: 14},
	{Name: "K", Strength: 13},
	{Name: "Q", Strength: 12},
	{Name: "J", Strength: 1},
	{Name: "T", Strength: 10},
	{Name: "9", Strength: 9},
	{Name: "8", Strength: 8},
	{Name: "7", Strength: 7},
	{Name: "6", Strength: 6},
	{Name: "5", Strength: 5},
	{Name: "4", Strength: 4},
	{Name: "3", Strength: 3},
	{Name: "2", Strength: 2},
}

type HandType struct {
	Name     string
	Strength int
}

var (
	Fives    = HandType{Name: "FIVES", Strength: 6}
	Fours    = HandType{Name: "FOURS", Strength: 5}
	Full     = HandType{Name: "FULL", Strength: 4}
	Threes   = HandType{Name: "THREES", Strength: 3}
	TwoPairs = HandType{Name: "TWO_PAIRS", Strength: 2}
	Pair     = HandType{Name: "PAIR", Strength: 1}
	High     = HandType{Name: "HIGH", Strength: 0}
)

type Hand struct {
	Elements [5]Figure
	HandType HandType
	Bid      int
}

type Game []Hand

func (h Game) Len() int { return len(h) }
func (h Game) Less(i, j int) bool {
	if h[i].HandType.Strength != h[j].HandType.Strength {
		return h[i].HandType.Strength < h[j].HandType.Strength
	}
	for k := 0; k < 5; k++ {
		if h[i].Elements[k] != h[j].Elements[k] {
			return h[i].Elements[k].Strength < h[j].Elements[k].Strength
		}
	}
	return false
}
func (h Game) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func main() {
	filePath := "day 7.input"
	fileContent := readFile(filePath)
	data := prepareData(fileContent)
	fmt.Println(calculateFirstStar(data))
	//fmt.Println(calculateSecondStar(games))
}

func readFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func prepareData(data []string) Game {
	var hands Game
	for _, line := range data {
		parts := strings.Fields(line)

		var elements [5]Figure
		for i, char := range parts[0] {
			elements[i] = findFigureByName(string(char))
		}

		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil
		}

		handType := resolveTypeJoker(countEnumOccurrences(elements))

		hands = append(hands, Hand{Elements: elements, Bid: bid, HandType: handType})
	}

	return hands
}

func findFigureByName(name string) Figure {
	for _, figure := range figures {
		if figure.Name == name {
			return figure
		}
	}
	return figures[0]
}

func countEnumOccurrences(hands [5]Figure) map[Figure]int {
	enumCounts := make(map[Figure]int)

	for _, hand := range hands {
		enumCounts[hand]++
	}

	return enumCounts
}

func resolveType(occurances map[Figure]int) HandType {
	for _, count := range occurances {
		if count == 5 {
			return Fives
		}
		if count == 4 {
			return Fours
		}
	}
	for _, count := range occurances {
		if count == 3 {
			for _, count2 := range occurances {
				if count2 == 2 {
					return Full
				}
			}
			return Threes
		}
	}
	for enum, count := range occurances {
		if count == 2 {
			for enum2, count2 := range occurances {
				if count2 == 2 && enum != enum2 {
					return TwoPairs
				}
			}
			return Pair
		}
	}
	return High
}

func calculateFirstStar(game Game) int {
	acc := 0
	sort.Sort(Game(game))
	for i := 0; i < len(game); i++ {
		acc += (i + 1) * game[i].Bid
	}
	return acc
}

func calculateSecondStar(game Game) int64 {
	var acc int64 = 0
	sort.Sort(Game(game))
	for i := 0; i < len(game); i++ {
		acc = acc + int64((i+1)*game[i].Bid)
	}
	return acc
}

func resolveTypeJoker(occurances map[Figure]int) HandType {
	jokers := 0
	for enum, count := range occurances {
		if enum.Name == "J" {
			jokers = count
		}
	}
	for enum, count := range occurances {
		if enum.Name != "J" {
			count += jokers
		}
		if count == 5 {
			return Fives
		}
		if count == 4 {
			if enum.Name == "J" {
				return Fives
			}
			return Fours
		}
	}
	for enum, count := range occurances {
		if enum.Name != "J" {
			count += jokers
		}
		if count == 3 {
			for enum2, count2 := range occurances {
				if count2 == 2 && enum != enum2 && enum2.Name != "J" {
					return Full
				}
			}
			return Threes
		}
	}
	for enum, count := range occurances {
		if enum.Name != "J" {
			count += jokers
		}
		if count == 2 {
			for enum2, count2 := range occurances {
				if count2 == 2 && enum != enum2 {
					return TwoPairs
				}
			}
			return Pair
		}
	}
	return High
}
