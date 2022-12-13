package day12

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

func TestTravelGraphTestData(t *testing.T) {
	d := CalcDistance(data.TestElevationMap)
	fmt.Printf("distance: %v\n", d)
}

func TestTravelGraphPuzzleData(t *testing.T) {
	d := CalcDistance(data.PuzzleElevationMap)
	fmt.Printf("distance: %v\n", d)
}

func TestTravelGraphTestDataAnyElevation(t *testing.T) {
	d := CalcDistanceAnyLowerElevationStart(data.TestElevationMap)
	fmt.Printf("distance any lowest elevation: %v\n", d)
}

func TestTravelGraphPuzzleDataAnyElevation(t *testing.T) {
	d := CalcDistanceAnyLowerElevationStart(data.PuzzleElevationMap)
	fmt.Printf("distance any lowest elevation : %v\n", d)
}
