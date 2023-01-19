package day11

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

func TestMonkeyKeepAwayTestData(t *testing.T) {
	res := CalcKeepAwayWorryManaged(copyMonkey(data.TestMonkeysSchema))
	fmt.Printf("Monkey business for test data equals: %v\n", res)
}

func TestMonkeyKeepAwayPuzzleData(t *testing.T) {
	res := CalcKeepAwayWorryManaged(copyMonkey(data.PuzzleMonkeysSchema))
	fmt.Printf("Monkey business for puzzle data equals: %v\n", res)
}

func TestMonkeyKeepAwayNotManagedTestData(t *testing.T) {
	res := CalcKeepAwayWorryNotManaged(copyMonkey(data.TestMonkeysSchema))
	fmt.Printf("Monkey business for test data worry not manged equals: %v\n", res)
}

func TestMonkeyKeepAwayNotManagedPuzzleData(t *testing.T) {
	res := CalcKeepAwayWorryNotManaged(copyMonkey(data.PuzzleMonkeysSchema))
	fmt.Printf("Monkey business for test data worry not manged equals: %v\n", res)
}

func copyMonkey(monkeys []data.Monkey) []data.Monkey {
	cp := make([]data.Monkey, 0, len(monkeys))
	for _, m := range monkeys {
		items := make([]int, len(m.Items))
		copy(items, m.Items)
		cp = append(cp, data.Monkey{
			Num:       m.Num,
			Items:     items,
			Operation: m.Operation,
			Test:      m.Test,
			Throw:     m.Throw,
			Prime:     m.Prime,
		})
	}
	return cp
}
