package day14

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

func TestCalculateSandDropComingToRestTestData(t *testing.T) {
	res := CalcSandDropComingToRest(data.CavePlanTest, point2D{x: 500, y: 0})
	fmt.Printf("Total amount of sand coming to rest for test data is %v\n", res)
}

func TestCalculateSandDropComingToRestPuzzleData(t *testing.T) {
	res := CalcSandDropComingToRest(data.CavePlanPuzzle, point2D{x: 500, y: 0})
	fmt.Printf("Total amount of sand coming to rest for test data is %v\n", res)
}

func TestCalculateSandDropComingToDropPointTestData(t *testing.T) {
	res := CalcSandDropComingToReachDropPoint(data.CavePlanTest, point2D{x: 500, y: 0})
	fmt.Printf("Total amount of sand coming to rest for test data is %v\n", res)
}

func TestCalculateSandDropComingToDropPointPuzzleData(t *testing.T) {
	res := CalcSandDropComingToReachDropPoint(data.CavePlanPuzzle, point2D{x: 500, y: 0})
	fmt.Printf("Total amount of sand coming to rest for test data is %v\n", res)
}
