package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Network struct {
	path    string
	network []Node
}

type Node struct {
	start string
	left  string
	right string
}

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

func prepareData(data []string) Network {
	var network Network
	network.network = make([]Node, len(data)-2)
	for i, line := range data {
		fields := strings.Fields(line)
		if i == 0 {
			network.path = fields[0]
		}
		if len(fields) < 3 {
			continue
		}

		node := Node{
			start: fields[0],
			left:  fields[2][1 : len(fields[2])-1],
			right: fields[3][0 : len(fields[3])-1],
		}
		network.network[i-2] = node
	}
	return network
}

func calculateFirstStar(network Network) int {
	steps := 0
	pathStep := 0
	currentTunnel := "AAA"

	for steps < 99999999 {
		if currentTunnel == "ZZZ" {
			return steps
		}

		i := pathStep % len(network.path)
		var actualNode Node
		for _, node := range network.network {
			if node.start == currentTunnel {
				actualNode = node
			}
		}

		if network.path[i] == 'R' {
			currentTunnel = actualNode.right
		} else {
			currentTunnel = actualNode.left
		}

		pathStep = i + 1
		steps++
	}

	return -1
}

func calculateSecondStar(network Network) int {
	var numbers []int
	for _, node := range network.network {
		if strings.HasSuffix(node.start, "A") {
			steps := 0
			pathStep := 0
			currentTunnel := node.start
			for steps < 99999999 {
				if strings.HasSuffix(currentTunnel, "Z") {
					break
				}

				i := pathStep % len(network.path)
				var actualNode Node
				for _, node := range network.network {
					if node.start == currentTunnel {
						actualNode = node
					}
				}

				if network.path[i] == 'R' {
					currentTunnel = actualNode.right
				} else {
					currentTunnel = actualNode.left
				}

				pathStep = i + 1
				steps++
			}
			numbers = append(numbers, steps)
		}
	}
	return lcmOfArray(numbers)
}

// Calculate the greatest common divisor (GCD) using Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Calculate the least common multiple (LCM) of an array of integers
func lcmOfArray(numbers []int) int {
	result := numbers[0]

	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}

	return result
}

// Calculate the least common multiple (LCM) of two integers
func lcm(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}
