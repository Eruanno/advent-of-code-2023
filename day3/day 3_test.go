package main

import "testing"

func TestFirstStarTestData(t *testing.T) {
	fileContent := readFile("day 3 test.input")
	data := prepareData(fileContent)
	result := calculateFirstStar(data)
	expected := 4361

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestFirstStarRealData(t *testing.T) {
	fileContent := readFile("day 3.input")
	data := prepareData(fileContent)
	result := calculateFirstStar(data)
	expected := 521601

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSecondStarTestData(t *testing.T) {
	fileContent := readFile("day 3 test.input")
	data := prepareData(fileContent)
	result := calculateSecondStar(data)
	expected := 467835

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSecondStarRealData(t *testing.T) {
	fileContent := readFile("day 3.input")
	data := prepareData(fileContent)
	result := calculateSecondStar(data)
	expected := 80694070

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
