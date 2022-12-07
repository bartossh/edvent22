package day7

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

func TestCalcFolderBelowMaxFolderSize1TestData(t *testing.T) {
	res := CalcFoldersSize(data.MaxFolderSize1, data.TestTerminalMessages)
	fmt.Printf("Total size of test data: %v\n", res)
}

func TestGraphCreationTestData(t *testing.T) {
	res := splitTerminalMsgs(data.TestTerminalMessages)
	res.printStructure("")
}

func TestCalcFolderBelowMaxFolderSize1PuzzleData(t *testing.T) {
	res := CalcFoldersSize(data.MaxFolderSize1, data.PuzzleTerminalMessages)
	fmt.Printf("Total size of puzzle data: %v\n", res)
}

func TestCalcSmallestSizeDirectoryOverThresholdTestData(t *testing.T) {
	res := CalcSmallestSizeDirectoryOverThreshold(data.FreeSpaceNeeded, data.DiskSpace, data.TestTerminalMessages)
	fmt.Printf("Smallest directory size to free up the space of test data is: %v\n", res)
}

func TestCalcSmallestSizeDirectoryOverThresholdPuzzleData(t *testing.T) {
	res := CalcSmallestSizeDirectoryOverThreshold(data.FreeSpaceNeeded, data.DiskSpace, data.PuzzleTerminalMessages)
	fmt.Printf("Smallest directory size to free up the space of puzzle data is: %v\n", res)
}
