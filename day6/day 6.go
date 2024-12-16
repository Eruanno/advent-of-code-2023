package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func main() {
	filePath := "day 6.input"
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

func prepareData(data []string) []Race {
	var races []Race

	times := strings.Fields(data[0])
	distances := strings.Fields(data[1])

	for i := 1; i < len(times); i++ {
		time, err := strconv.Atoi(times[i])
		distance, err := strconv.Atoi(distances[i])
		races = append(races, Race{time, distance})
		if err != nil {
			println("Błąd podczas konwersji danych.")
		}
	}

	return races
}

func calculateFirstStar(races []Race) int {
	result := 1
	for _, race := range races {
		acc := 0
		for i := 0; i < race.Time; i++ {
			speed := i
			remainingTime := race.Time - speed
			distance := speed * remainingTime
			if race.Distance < distance {
				acc += 1
			}
		}
		result *= acc
	}
	return result
}

func calculateSecondStar(races []Race) int {
	timeString := ""
	distanceString := ""
	for _, race := range races {
		timeStr := strconv.Itoa(race.Time)
		distanceStr := strconv.Itoa(race.Distance)
		timeString = timeString + timeStr
		distanceString = distanceString + distanceStr
	}
	time, err := strconv.Atoi(timeString)
	distance, err := strconv.Atoi(distanceString)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0
	}
	acc := 0
	for i := 0; i < time; i++ {
		speed := i
		remainingTime := time - speed
		calcDistance := speed * remainingTime
		if distance < calcDistance {
			acc += 1
		}
	}
	return acc
}
