package day16

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

func TestReleaseMaxValvePressureTestData(t *testing.T) {
	result := CalcMaxPressureToRelease(data.ValveMapTest, 30)
	fmt.Printf("Max pressure to be released for test data is: %v\n", result)
}

func TestReleaseMaxValvePressurePuzzleData(t *testing.T) {
	result := CalcMaxPressureToRelease(data.ValveMapPuzzle, 30)
	fmt.Printf("Max pressure to be released for puzzle data is: %v\n", result)
}

func TestReleaseMaxValvePressureTestDataWithElephant(t *testing.T) {
	result := CalcMaxPressureToReleaseWithElephant(data.ValveMapTest, 26)
	fmt.Printf("Max pressure to be released for test data is: %v\n", result)
}

func TestReleaseMaxValvePressurePuzzleDataWithElephant(t *testing.T) {
	result := CalcMaxPressureToReleaseWithElephant(data.ValveMapPuzzle, 26)
	fmt.Printf("Max pressure to be released for puzzle data is: %v\n", result)
}
