package main

import "testing"

func TestFirstStarTestData(t *testing.T) {
	fileContent := readFile("day 2 test.input")
	data := prepareData(fileContent)
	result := calculateFirstStar(data)
	expected := 8

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestFirstStarRealData(t *testing.T) {
	fileContent := readFile("day 2.input")
	data := prepareData(fileContent)
	result := calculateFirstStar(data)
	expected := 2512

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSecondStarTestData(t *testing.T) {
	fileContent := readFile("day 2 test.input")
	data := prepareData(fileContent)
	result := calculateSecondStar(data)
	expected := 2286

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSecondStarRealData(t *testing.T) {
	fileContent := readFile("day 2.input")
	data := prepareData(fileContent)
	result := calculateSecondStar(data)
	expected := 67335

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
