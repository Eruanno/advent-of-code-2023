package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := "day 8.input"
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

func prepareData(data []string) []string {
	var reports [][]int
	for _, line := range data {
		fields := strings.Fields(line)
		var report []int
		for _, field := range fields {
			bid, err := strconv.Atoi(field)
			if err != nil {
				return nil
			}
			report = append(report, bid)
		}
		reports = append(reports, report)
		var secondReport []int
		for i := 0; i < len(report)-1; i++ {
			secondReport = append(secondReport, (report[i+1] - report[i]))
		}
		for i := 0; i < len(secondReport); i++ {
			if secondReport[i] != 0 {
				continue
			}
		}
	}
	return data
}

func calculateFirstStar(network []string) int {
	return 0
}

func calculateSecondStar(network []string) int {
	return 0
}
