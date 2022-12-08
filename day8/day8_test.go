package day8

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

func TestForestGridVisibilityStringToData(t *testing.T) {
	res := stringToGridArr(data.TestForrestGrid)
	fmt.Printf("Converted data looks like %v\n", res)
}

func TestForestGridVisibilityTestData(t *testing.T) {
	res := CountVisibleTreesFormOutsideTheGrid(data.TestForrestGrid)
	fmt.Printf("Number of trees visible from outside the grid for test data is: %v\n", res)
}

func TestForestGridVisibilityPuzzleData(t *testing.T) {
	res := CountVisibleTreesFormOutsideTheGrid(data.PuzzleForrestGrid)
	fmt.Printf("Number of trees visible from outside the grid for puzzle data is: %v\n", res)
}

func TestForestCountsScenicScoreAnyTreeTestData(t *testing.T) {
	res := CountsScenicScoreAnyTree(data.TestForrestGrid)
	fmt.Printf("Scenic score for the grid for test data is: %v\n", res)
}

func TestForestCountsScenicScoreAnyTreePuzzleData(t *testing.T) {
	res := CountsScenicScoreAnyTree(data.PuzzleForrestGrid)
	fmt.Printf("Scenic score for the grid for puzzle data is: %v\n", res)
}
