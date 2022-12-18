package day18

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

func TestLaveCubesTestData(t *testing.T) {
	res := CalcSurfaceArea(data.TestLavaCubes)
	fmt.Printf("Total surface area for test data is %v \n", res)
}

func TestLaveCubesPuzzleData(t *testing.T) {
	res := CalcSurfaceArea(data.PuzzleLavaCube)
	fmt.Printf("Total surface area for puzzle data is %v \n", res)
}

func TestLaveCubesTouchingTestData(t *testing.T) {
	res := CalcTouching(data.TestLavaCubes)
	fmt.Printf("Total surface area for test data is %v \n", res)
}

func TestLaveCubesTouchingPuzzleData(t *testing.T) {
	res := CalcTouching(data.PuzzleLavaCube)
	fmt.Printf("Total surface area for puzzle data is %v \n", res)
}
