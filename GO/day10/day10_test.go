package day10

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

func TestRegisterSumTestData(t *testing.T) {
	sum := CalculateRegisterSum(data.TestCathodeRayTubeData)
	fmt.Printf("CathodeRay register sum for test data is: %v\n", sum)
}

func TestRegisterSumPuzzle(t *testing.T) {
	sum := CalculateRegisterSum(data.PuzzleCathodeRayTubeData)
	fmt.Printf("CathodeRay register sum for puzzle data is: %v\n", sum)
}

func TestCTRSumTestData(t *testing.T) {
	result := DrawCTR(data.TestCathodeRayTubeData)
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
}

func TestCTRSumPuzzleData(t *testing.T) {
	result := DrawCTR(data.PuzzleCathodeRayTubeData)
	for _, r := range result {
		fmt.Printf("%v\n", r)
	}
}
