package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	Id      int
	Subsets []Subset
}

type Subset struct {
	Red   int
	Green int
	Blue  int
}

type Color int

const (
	RED Color = iota
	GREEN
	BLUE
)

func main() {
	filePath := "day 2 first test.input"
	fileContent := readFile(filePath)
	games := prepareData(fileContent)
	fmt.Println(calculateFirstStar(games))
	fmt.Println(calculateSecondStar(games))
}

func readFile(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Błąd wczytywania pliku:", err)
		return ""
	}

	return string(content)
}

func prepareData(fileContent string) []Game {
	lines := strings.Split(fileContent, "\n")
	games := []Game{}
	for _, line := range lines {
		basicParts := strings.Split(line, ":")
		id := findNumberAtEnds(basicParts[0])
		supersets := strings.Split(basicParts[1], ";")
		subsets := []Subset{}
		for _, superset := range supersets {
			subset := parseSubsetString(superset)
			subsets = append(subsets, subset)
		}
		games = append(games, Game{Id: id, Subsets: subsets})
	}
	return games
}

func findNumberAtEnds(input string) int {
	re := regexp.MustCompile(`(\d+)$`)

	match := re.FindStringSubmatch(input)
	if match == nil {
		return 0
	}

	number := match[1]
	result, err := strconv.Atoi(number)
	if err != nil {
		return 0
	}

	return result
}

func parseSubsetString(input string) Subset {
	re := regexp.MustCompile(`(\d+)\s+(red|green|blue)`)

	matches := re.FindAllStringSubmatch(input, -1)

	subset := Subset{}

	for _, match := range matches {
		number, err := strconv.Atoi(match[1])
		if err != nil {
			return Subset{}
		}

		color := match[2]
		switch color {
		case "red":
			subset.Red = number
		case "green":
			subset.Green = number
		case "blue":
			subset.Blue = number
		}
	}
	return subset
}

func calculateFirstStar(games []Game) int {
	acc := 0
	for _, game := range games {
		if valid(game.Subsets) {
			acc += game.Id
		}
	}
	return acc
}

func valid(subsets []Subset) bool {
	for _, subset := range subsets {
		if subset.Red > 12 || subset.Green > 13 || subset.Blue > 14 {
			return false
		}
	}
	return true
}

func calculateSecondStar(games []Game) int {
	return calculate(games)
}

func calculate(games []Game) int {
	acc := 0
	for _, game := range games {
		acc += powerOfSet(game.Subsets)
	}
	return acc
}

func powerOfSet(subsets []Subset) int {
	red := 0
	green := 0
	blue := 0
	for _, subset := range subsets {
		if subset.Red > red && subset.Red > 0 {
			red = subset.Red
		}
		if subset.Green > green && subset.Green > 0 {
			green = subset.Green
		}
		if subset.Blue > blue && subset.Blue > 0 {
			blue = subset.Blue
		}
	}
	return red * green * blue
}
