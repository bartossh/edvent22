package day15

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

func TestLocateWhereBeaconNotPresentTestData(t *testing.T) {
	pos := CountPositionWhereBeaconNotPresent(data.TestBeaconsAndSensorsMap, 10)
	fmt.Printf("Test answer is: %v \n", pos)
}

func TestLocateWhereBeaconNotPresentPuzzleData(t *testing.T) {
	pos := CountPositionWhereBeaconNotPresent(data.PuzzleBeaconsAndSensorsMap, 2000000)
	fmt.Printf("Puzzle answer is : %v\n", pos)
}

func TestLocateWhereBeaconTuneFrequencyPuzzleData(t *testing.T) {
	pos := CountPositionWhereBeaconTuningFrequency(data.PuzzleBeaconsAndSensorsMap, 4000000)
	fmt.Printf("Puzzle answer is : %v\n", pos)
}
