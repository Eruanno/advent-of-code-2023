package main

import "testing"

func TestFirstStarTestData(t *testing.T) {
	fileContent := readFile("day 5 test.input")
	data := prepareData(fileContent)
	result := calculateFirstStar(data)
	expected := 35

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestFirstStarRealData(t *testing.T) {
	fileContent := readFile("day 5.input")
	data := prepareData(fileContent)
	result := calculateFirstStar(data)
	expected := 318728750

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSecondStarTestData(t *testing.T) {
	fileContent := readFile("day 5 test.input")
	data := prepareData(fileContent)
	result := calculateSecondStar(data)
	expected := 46

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSecondStarRealData(t *testing.T) {
	fileContent := readFile("day 5.input")
	data := prepareData(fileContent)
	result := calculateSecondStar(data)
	expected := 37384986

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
