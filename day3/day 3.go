package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Part struct {
	number int
	endX   int
	gearX  int
	gearY  int
}

type Gear struct {
	x             int
	y             int
	numberOfParts int
	value         int
}

func main() {
	filePath := "day 3 test.input"
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

func prepareData(fileContent string) [][]byte {
	lines := strings.Split(fileContent, "\n")
	var engine [][]byte
	for _, line := range lines {
		charArray := []byte(strings.TrimRight(line, "\r"))
		engine = append(engine, charArray)
	}
	return engine
}

func calculateFirstStar(engine [][]byte) int {
	acc := 0
	for y := 0; y < len(engine); y++ {
		for x := 0; x < len(engine[y]); x++ {
			if unicode.IsDigit(rune(engine[y][x])) {
				//fmt.Printf("Line nr %d column nr %d\n", y, x)
				part := checkPart(y, x, engine)
				acc += part.number
				x = part.endX
			}
		}
	}
	return acc
}

func checkPart(i int, j int, engine [][]byte) Part {
	startY := i
	startX := j
	endX := j
	value := ""
	for x := startX; x < len(engine[startY]); x++ {
		if unicode.IsDigit(rune(engine[startY][x])) {
			endX = x
			value = value + string(engine[startY][x])
		} else {
			break
		}
	}
	for x := startX - 1; x <= endX+1; x++ {
		for y := startY - 1; y <= startY+1; y++ {
			if y >= 0 && y < len(engine) && x >= 0 && x < len(engine[y]) {
				//fmt.Printf("Y: %d X: %d len %d\t", y, x, len(engine[y]))
				if checkField(engine[y][x]) {
					return Part{stringToInteger(value), endX, x, y}
				}
			}
		}
	}
	return Part{0, endX, -1, -1}
}

func checkField(field byte) bool {
	return !unicode.IsDigit(rune(field)) && field != '.'
}

func calculateSecondStar(engine [][]byte) int {
	var gears []Gear
	acc := 0
	for y := 0; y < len(engine); y++ {
		for x := 0; x < len(engine[y]); x++ {
			if unicode.IsDigit(rune(engine[y][x])) {
				//fmt.Printf("Line nr %d column nr %d\n", y, x)
				part := checkGear(y, x, engine)
				if part.number > 0 {
					found := false
					for i, gear := range gears {
						if gear.x == part.gearX && gear.y == part.gearY {
							gear.numberOfParts += 1
							gear.value *= part.number
							gears[i] = gear
							found = true
						}
					}
					if found == false {
						gears = append(gears, Gear{part.gearX, part.gearY, 1, part.number})
					}
				}
				x = part.endX
			}
		}
	}
	for _, gear := range gears {
		if gear.numberOfParts > 1 {
			acc += gear.value
		}
	}
	return acc
}

func checkGear(i int, j int, engine [][]byte) Part {
	startY := i
	startX := j
	endX := j
	value := ""
	for x := startX; x < len(engine[startY]); x++ {
		if unicode.IsDigit(rune(engine[startY][x])) {
			endX = x
			value = value + string(engine[startY][x])
		} else {
			break
		}
	}
	for x := startX - 1; x <= endX+1; x++ {
		for y := startY - 1; y <= startY+1; y++ {
			if y >= 0 && y < len(engine) && x >= 0 && x < len(engine[y]) {
				//fmt.Printf("Y: %d X: %d len %d\t", y, x, len(engine[y]))
				if engine[y][x] == '*' {
					return Part{stringToInteger(value), endX, x, y}
				}
			}
		}
	}
	return Part{0, endX, -1, -1}
}

func stringToInteger(number string) int {
	intNumber, err := strconv.Atoi(number)
	if err != nil {
		return 0
	}

	return intNumber
}
