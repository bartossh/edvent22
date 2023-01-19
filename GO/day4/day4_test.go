package day4

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
	"github.com/stretchr/testify/assert"
)

func TestCountOverlapsTestData(t *testing.T) {
	res := CountOverlaps(data.TestCampCleanupAssignments)
	assert.Equal(t, data.CampCleanupOverlapAssignmentsTestResult, res)
	fmt.Printf("Overlapping test result: %v\n", res)
}

func TestCountOverlapsTaskData(t *testing.T) {
	res := CountOverlaps(data.DataCampCleanupAssignments)
	fmt.Printf("Overlapping task result: %v\n", res)
}

func TestCountPartlyOverlapsTestData(t *testing.T) {
	res := CountOverlapPartly(data.TestCampCleanupAssignments)
	assert.Equal(t, data.CampCleanupOverlapPartlyAssignmentsTestResult, res)
	fmt.Printf("Overlapping test result: %v\n", res)
}

func TestCountPartlyOverlapsTaskData(t *testing.T) {
	res := CountOverlapPartly(data.DataCampCleanupAssignments)
	fmt.Printf("Overlapping task result: %v\n", res)
}
