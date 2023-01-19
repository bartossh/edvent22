package day19

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

// func TestBestBlueprintStrategyTestData(t *testing.T) {
// 	res := CalcBestStrategy(data.TestRobotFactoryOptions, 24)
// 	fmt.Printf("Best quality level for test data is: %v\n", res)
// }

// func TestBestBlueprintStrategyPuzzleData(t *testing.T) {
// 	res := CalcBestStrategy(data.PuzzleRobotFactoryOptions, 24)
// 	fmt.Printf("Best quality level for puzzle data is: %v\n", res)
// }

// func TestBestBlueprintGeodesOpenedTestData(t *testing.T) {
// 	res := CalcMaxGeodesOpened(data.TestRobotFactoryOptions, 32)
// 	fmt.Printf("Max geodes opened for test data is: %v\n", res)
// }

func TestBestBlueprintGeodesOpenedPuzzleData(t *testing.T) {
	res := CalcMaxGeodesOpened(data.PuzzleRobotFactoryOptions, 32)
	fmt.Printf("Max geodes opened for puzzle data is: %v\n", res)
}
