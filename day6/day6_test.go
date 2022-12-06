package day6

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
	"github.com/stretchr/testify/assert"
)

func TestCalcMarkerPositionTestData(t *testing.T) {
	for i, v := range data.MarkerPositions {
		res := CalcMarkerPosition(data.Subroutine[i])
		fmt.Printf("Tes Data Market position is: %v\n", res)
		assert.Equal(t, v, res)
	}
}

func TestCalcMarkerPositionPuzzleData(t *testing.T) {
	res := CalcMarkerPosition(data.PuzzleSubroutine)
	fmt.Printf("Puzzle Market position is: %v\n", res)
}

func TestCalcMessagePositionTestData(t *testing.T) {
	for i, v := range data.MessageMarkerPosition {
		res := CalcMessagePosition(data.Subroutine[i])
		fmt.Printf("Test Data Message Market position is: %v\n", res)
		assert.Equal(t, v, res)
	}
}

func TestCalcMessagePositionPuzzleData(t *testing.T) {
	res := CalcMessagePosition(data.PuzzleSubroutine)
	fmt.Printf("Puzzle Data Message Market position is: %v\n", res)
}
