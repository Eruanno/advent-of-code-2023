package main

import "testing"

func TestFirstStarTestData(t *testing.T) {
	fileContent := readFile("day 4 test.input")
	data := prepareData(fileContent)
	result := calculateFirstStar(data)
	expected := 13

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestFirstStarRealData(t *testing.T) {
	fileContent := readFile("day 4.input")
	data := prepareData(fileContent)
	result := calculateFirstStar(data)
	expected := 25010

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSecondStarTestData(t *testing.T) {
	fileContent := readFile("day 4 test.input")
	data := prepareData(fileContent)
	result := calculateSecondStar(data)
	expected := 30

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSecondStarRealData(t *testing.T) {
	fileContent := readFile("day 4.input")
	data := prepareData(fileContent)
	result := calculateSecondStar(data)
	expected := 9924412

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
