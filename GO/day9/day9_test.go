package day9

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

func TestDataToMoves(t *testing.T) {
	res := dataToMoves(data.TestRopeHeadMoves1)
	fmt.Printf("Test moves structure: %#v\n", res)
}

func TestVisitedByTailTestData(t *testing.T) {
	res := CountVisitedByTheTailAtLeastOnce(data.TestRopeHeadMoves1)
	fmt.Printf("Visited number of nodes by tail for test data: %v\n", res)
}

func TestVisitedByTailPuzzleData(t *testing.T) {
	res := CountVisitedByTheTailAtLeastOnce(data.PuzzleRopeHeadMoves)
	fmt.Printf("Visited number of nodes by tail for puzzle data: %v\n", res)
}

func TestVisitedByTailTestDataMultiKnots(t *testing.T) {
	res1 := CountVisitedByTheTailAtLeastOnceMultiKnots(data.TestRopeHeadMoves1, 10)
	fmt.Printf("Visited number of nodes by tail for test 1 data with 10 knots: %v\n", res1)
	res2 := CountVisitedByTheTailAtLeastOnceMultiKnots(data.TestRopeHeadMoves2, 10)
	fmt.Printf("Visited number of nodes by tail for test 2 data with 10 knots: %v\n", res2)
}

func TestVisitedByTailPuzzleDataMultiKnots(t *testing.T) {
	res := CountVisitedByTheTailAtLeastOnceMultiKnots(data.PuzzleRopeHeadMoves, 10)
	fmt.Printf("Visited number of nodes by tail for puzzle data with 10 knots: %v\n", res)
}
