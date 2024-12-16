package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SeedMap struct {
	Seeds                    []int
	SeedToSoilMap            [][]int
	SoilToFertilizerMap      [][]int
	FertilizerToWaterMap     [][]int
	WaterToLightMap          [][]int
	LightToTemperatureMap    [][]int
	TemperatureToHumidityMap [][]int
	HumidityToLocationMap    [][]int
}

func main() {
	filePath := "day 5.input"
	fileContent := readFile(filePath)
	games := prepareData(fileContent)
	//fmt.Println(calculateFirstStar(games))
	fmt.Println(calculateSecondStar(games))
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

func prepareData(data []string) SeedMap {
	var seedMap SeedMap
	var currentMap [][]int

	for _, line := range data {
		if strings.Contains(line, ":") && !containsNumericValues(line) {
			continue
		}

		if line == "" {
			storeMap(&seedMap, currentMap)
			currentMap = nil
		} else {
			fields := strings.Fields(line)
			var row []int
			for _, field := range fields {
				value, err := strconv.Atoi(field)
				if err == nil {
					row = append(row, value)
				}
			}
			currentMap = append(currentMap, row)
		}
	}

	if len(currentMap) > 0 {
		storeMap(&seedMap, currentMap)
	}

	return seedMap
}

func storeMap(seedMap *SeedMap, currentMap [][]int) {
	switch {
	case seedMap.Seeds == nil:
		seedMap.Seeds = currentMap[0]
	case seedMap.SeedToSoilMap == nil:
		seedMap.SeedToSoilMap = currentMap
	case seedMap.SoilToFertilizerMap == nil:
		seedMap.SoilToFertilizerMap = currentMap
	case seedMap.FertilizerToWaterMap == nil:
		seedMap.FertilizerToWaterMap = currentMap
	case seedMap.WaterToLightMap == nil:
		seedMap.WaterToLightMap = currentMap
	case seedMap.LightToTemperatureMap == nil:
		seedMap.LightToTemperatureMap = currentMap
	case seedMap.TemperatureToHumidityMap == nil:
		seedMap.TemperatureToHumidityMap = currentMap
	case seedMap.HumidityToLocationMap == nil:
		seedMap.HumidityToLocationMap = currentMap
	}
}

func containsNumericValues(s string) bool {
	fields := strings.Fields(s)
	for _, field := range fields {
		if _, err := strconv.Atoi(field); err == nil {
			return true
		}
	}
	return false
}

func calculateFirstStar(seedMap SeedMap) int {
	lowestLocation := 1<<63 - 1
	for _, seed := range seedMap.Seeds {
		soil := findNextThing(seed, seedMap.SeedToSoilMap)
		fertilizer := findNextThing(soil, seedMap.SoilToFertilizerMap)
		water := findNextThing(fertilizer, seedMap.FertilizerToWaterMap)
		light := findNextThing(water, seedMap.WaterToLightMap)
		temperature := findNextThing(light, seedMap.LightToTemperatureMap)
		humidity := findNextThing(temperature, seedMap.TemperatureToHumidityMap)
		location := findNextThing(humidity, seedMap.HumidityToLocationMap)
		if location < lowestLocation {
			lowestLocation = location
		}
	}
	return lowestLocation
}

func findNextThing(input int, currentMap [][]int) int {
	output := -1
	for i := 0; i < len(currentMap); i++ {
		destination := currentMap[i][0]
		source := currentMap[i][1]
		rangeLength := currentMap[i][2]
		if input >= source && input < source+rangeLength {
			output = destination + (input - source)
		}
	}
	if output >= 0 {
		return output
	}
	return input
}

func calculateSecondStar(seedMap SeedMap) int {
	lowestLocation := 1<<63 - 1
	for i := 0; i < len(seedMap.Seeds); i += 2 {
		start := seedMap.Seeds[i]
		seedRange := seedMap.Seeds[i+1]
		for seed := start; seed < start+seedRange; seed++ {
			soil := findNextThing(seed, seedMap.SeedToSoilMap)
			fertilizer := findNextThing(soil, seedMap.SoilToFertilizerMap)
			water := findNextThing(fertilizer, seedMap.FertilizerToWaterMap)
			light := findNextThing(water, seedMap.WaterToLightMap)
			temperature := findNextThing(light, seedMap.LightToTemperatureMap)
			humidity := findNextThing(temperature, seedMap.TemperatureToHumidityMap)
			location := findNextThing(humidity, seedMap.HumidityToLocationMap)
			//fmt.Printf("Seed %d soil %d location %d\n", seed, soil, location)
			if location < lowestLocation {
				lowestLocation = location
			}
		}
	}
	return lowestLocation
}
