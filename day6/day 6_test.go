package main

import "testing"

func TestFirstStarTestData(t *testing.T) {
	fileContent := readFile("day 6 test.input")
	data := prepareData(fileContent)
	result := calculateFirstStar(data)
	expected := 288

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestFirstStarRealData(t *testing.T) {
	fileContent := readFile("day 6.input")
	data := prepareData(fileContent)
	result := calculateFirstStar(data)
	expected := 1624896

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSecondStarTestData(t *testing.T) {
	fileContent := readFile("day 6 test.input")
	data := prepareData(fileContent)
	result := calculateSecondStar(data)
	expected := 71503

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSecondStarRealData(t *testing.T) {
	fileContent := readFile("day 6.input")
	data := prepareData(fileContent)
	result := calculateSecondStar(data)
	expected := 32583852

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
