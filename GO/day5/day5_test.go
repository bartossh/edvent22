package day5

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

func TestStringDataToArr(t *testing.T) {
	res := cratesStringToArr(data.CrateArrangement)
	fmt.Printf("crates: %v\n", res)
}

func TestStringDataToMoves(t *testing.T) {
	res := movesStringToArr(data.Moves)
	fmt.Printf("moves: %v\n", res)
}

func TestCreateStack(t *testing.T) {
	rows := cratesStringToArr(data.CrateArrangement)
	res := crateStack(rows)
	for _, v := range res {
		fmt.Printf("%s\n", string(v))
	}
}

func TestCreateStackTest(t *testing.T) {
	rows := cratesStringToArr(data.TestCrateArrangement)
	res := crateStack(rows)
	for _, v := range res {
		fmt.Printf("%s\n", string(v))
	}
}

func TestCraneMovingCratesTestData9000(t *testing.T) {
	res := MoveCrates9000(data.TestCrateArrangement, data.TestMoves)
	fmt.Printf("Test data top crates are: %s\n", res)
}

func TestCraneMovingCrates9000(t *testing.T) {
	res := MoveCrates9000(data.CrateArrangement, data.Moves)
	fmt.Printf("Task data top crates are: %s\n", res)
}

func TestCraneMovingCratesTestData9001(t *testing.T) {
	res := MoveCrates9001(data.TestCrateArrangement, data.TestMoves)
	fmt.Printf("Task data top crates are: %s\n", res)
}

func TestCraneMovingCrates9001(t *testing.T) {
	res := MoveCrates9001(data.CrateArrangement, data.Moves)
	fmt.Printf("Task data top crates are: %s\n", res)
}
