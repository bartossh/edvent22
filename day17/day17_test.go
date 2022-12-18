package day17

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

// func TestFallingRocksTetrisTestData(t *testing.T) {
// 	res := CalcHighOfTetrisTowel(data.JetFlowTest, 2022)
// 	fmt.Printf("High of test tetris power is %v \n", res)
// }

func TestFallingRocksTetrisPuzzleData(t *testing.T) {
	res := CalcHighOfTetrisTowel(data.JetFlowPuzzle, 2022)
	fmt.Printf("High of puzzle tetris power is %v \n", res)
}

// func TestFallingRocksTetrisTestData1000000000000(t *testing.T) {
// 	res := CalcHighOfTetrisTowel(data.JetFlowTest, 1000000000000)
// 	fmt.Printf("High of test tetris power is %v \n", res)
// }

func TestFallingRocksTetrisPuzzleData1000000000000(t *testing.T) {
	res := CalcHighOfTetrisTowel(data.JetFlowPuzzle, 100000000)
	fmt.Printf("High of test tetris power is %v \n", res)
}

// func TestFallingRocksTetrisPuzzleData1000000000000(t *testing.T) {
// 	res := CalcHighOfTetrisTowel(data.JetFlowPuzzle, 1000000000000)
// 	fmt.Printf("High of test tetris power is %v \n", res)
// }
