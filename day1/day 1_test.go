package main

import "testing"

func TestFirstStarTestData(t *testing.T) {
	fileContent := readFile("day 1 first test.input")
	data := prepareData(fileContent)
	result := calculateFirstStar(data)
	expected := 142

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestFirstStarRealData(t *testing.T) {
	fileContent := readFile("day 1.input")
	data := prepareData(fileContent)
	result := calculateFirstStar(data)
	expected := 54877

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSecondStarTestData(t *testing.T) {
	fileContent := readFile("day 1 second test.input")
	data := prepareData(fileContent)
	result := calculateSecondStar(data)
	expected := 281

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSecondStarRealData(t *testing.T) {
	fileContent := readFile("day 1.input")
	data := prepareData(fileContent)
	result := calculateSecondStar(data)
	expected := 54100

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
