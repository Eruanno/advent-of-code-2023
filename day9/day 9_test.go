package main

import "testing"

func TestFirstStarTestData(t *testing.T) {
	fileContent := readFile("day 9 test.input")
	data := prepareData(fileContent)
	result := calculateFirstStar(data)
	expected := 114

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestFirstStarRealData(t *testing.T) {
	fileContent := readFile("day 9.input")
	data := prepareData(fileContent)
	result := calculateFirstStar(data)
	expected := 14429

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSecondStarTestData(t *testing.T) {
	fileContent := readFile("day 9 test 2.input")
	data := prepareData(fileContent)
	result := calculateSecondStar(data)
	expected := 6

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSecondStarRealData(t *testing.T) {
	fileContent := readFile("day 9.input")
	data := prepareData(fileContent)
	result := calculateSecondStar(data)
	expected := 10921547990923
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
