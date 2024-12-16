package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Scratchpad struct {
	ID      int
	checked []int
	winning []int
}

type Scratchpads []Scratchpad

func (o Scratchpads) Len() int           { return len(o) }
func (o Scratchpads) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
func (o Scratchpads) Less(i, j int) bool { return o[i].ID < o[j].ID }

func main() {
	filePath := "day 4 test.input"
	fileContent := readFile(filePath)
	games := prepareData(fileContent)
	//fmt.Println(calculateFirstStar(games))
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

func prepareData(fileContent string) []Scratchpad {
	lines := strings.Split(fileContent, "\n")
	var scratchpads []Scratchpad
	for _, line := range lines {
		scratchpads = append(scratchpads, convertLineToScratchpad(line))
	}
	return scratchpads
}

func convertLineToScratchpad(line string) Scratchpad {
	parts := strings.Split(line, "|")

	var id int
	fmt.Sscanf(parts[0], "Card %d:", &id)

	checkedStrings := strings.Fields(strings.Split(parts[0], ":")[1])
	check := make([]int, len(checkedStrings))
	for i, str := range checkedStrings {
		fmt.Sscanf(str, "%d", &check[i])
	}

	winningStrings := strings.Fields(parts[1])
	win := make([]int, len(winningStrings))
	for i, str := range winningStrings {
		fmt.Sscanf(str, "%d", &win[i])
	}

	return Scratchpad{
		ID:      id,
		checked: check,
		winning: win,
	}
}

func calculateFirstStar(scratchpads Scratchpads) int {
	acc := 0
	for _, scratchpad := range scratchpads {
		power := 0
		for _, checkedNumber := range scratchpad.checked {
			for _, winningNumber := range scratchpad.winning {
				if checkedNumber == winningNumber {
					if power == 0 {
						power = 1
					} else {
						power *= 2
					}
					break
				}
			}
		}
		acc += power
	}
	return acc
}

func calculateSecondStar(scratchpads Scratchpads) int {
	someNumbers := make([]int, len(scratchpads))
	for i := range someNumbers {
		someNumbers[i] = 1
	}
	sort.Sort(scratchpads)
	for _, scratchpad := range scratchpads {
		id := scratchpad.ID
		nextId := id
		for _, checkedNumber := range scratchpad.checked {
			for _, winningNumber := range scratchpad.winning {
				if checkedNumber == winningNumber {
					someNumbers[nextId] += someNumbers[id-1]
					nextId += 1
				}
			}
		}
	}
	acc := 0
	for i := 0; i < len(someNumbers); i++ {
		acc += someNumbers[i]
	}
	return acc
}
