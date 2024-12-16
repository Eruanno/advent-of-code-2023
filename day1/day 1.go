package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Digit struct {
	Number int
	Name   string
}

var digits = []Digit{
	{0, "0"},
	{1, "1"},
	{2, "2"},
	{3, "3"},
	{4, "4"},
	{5, "5"},
	{6, "6"},
	{7, "7"},
	{8, "8"},
	{9, "9"},
}

var names = []Digit{
	{0, "zero"},
	{1, "one"},
	{2, "two"},
	{3, "three"},
	{4, "four"},
	{5, "five"},
	{6, "six"},
	{7, "seven"},
	{8, "eight"},
	{9, "nine"},
}

func main() {
	filePath := "day 1.input"
	fileContent := readFile(filePath)
	lines := prepareData(fileContent)
	fmt.Println(calculateFirstStar(lines))
	fmt.Println(calculateSecondStar(lines))
}

func readFile(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Błąd wczytywania pliku:", err)
		return ""
	}

	return string(content)
}

func prepareData(fileContent string) []string {
	return strings.Split(fileContent, "\n")
}

func calculateFirstStar(lines []string) int {
	return calculate(lines, digits)
}

func calculateSecondStar(lines []string) int {
	return calculate(lines, append(digits, names...))
}

func calculate(lines []string, digitsNames []Digit) int {
	acc := 0
	for _, line := range lines {
		positions := findDigitPositions(line, digitsNames)
		min := 1<<63 - 1
		max := -1
		minV := ""
		maxV := ""
		for name, pos := range positions {
			if pos[0] < min {
				min = pos[0]
				minV = findNumberByName(name, digitsNames)
			}
			if pos[1] > max {
				max = pos[1]
				maxV = findNumberByName(name, digitsNames)
			}
		}
		acc += stringToInt(minV + maxV)
	}
	return acc
}

func findDigitPositions(inputString string, digitArray []Digit) map[string][2]int {
	positions := make(map[string][2]int)
	for _, digit := range digitArray {
		startPos := strings.Index(inputString, digit.Name)
		endPos := strings.LastIndex(inputString, digit.Name)

		if startPos != -1 {
			positions[digit.Name] = [2]int{startPos, endPos + len(digit.Name) - 1}
		}
	}

	return positions
}

func findNumberByName(name string, digitArray []Digit) string {
	for _, digit := range digitArray {
		if strings.EqualFold(digit.Name, name) {
			return fmt.Sprint(digit.Number)
		}
	}
	return ""
}

func stringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}
