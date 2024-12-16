package main

import "testing"

func TestFirstStarTestData(t *testing.T) {
	fileContent := readFile("day 7 test.input")
	data := prepareData(fileContent)
	result := calculateFirstStar(data)
	expected := 6440

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestFirstStarRealData(t *testing.T) {
	fileContent := readFile("day 7.input")
	data := prepareData(fileContent)
	result := calculateFirstStar(data)
	expected := 253933213

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSecondStarTestData(t *testing.T) {
	fileContent := readFile("day 7 test.input")
	data := prepareData(fileContent)
	result := calculateSecondStar(data)
	expected := int64(6839)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSecondStarRealData(t *testing.T) {
	fileContent := readFile("day 7.input")
	data := prepareData(fileContent)
	result := calculateSecondStar(data)
	expected := int64(253473930)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
