package day21

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

func TestRootMonkeyNumberTestData(t *testing.T) {
	res := PredictRootMonkeyNumber(data.TestRootMonkeyRiddle, "root")
	fmt.Printf("Monkey test riddle is: %v\n", res)
}

func TestRootMonkeyNumberPuzzleData(t *testing.T) {
	res := PredictRootMonkeyNumber(data.PuzzleRootMonkeyRiddle, "root")
	fmt.Printf("Monkey puzzle riddle is: %v\n", res)
}
