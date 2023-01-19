package day22

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

func TestAlgorithmMazeTraversal(t *testing.T) {
	res := MazeTraversal(data.TestPasswordMaze)
	fmt.Printf("Result password for test data is: %v\n", res)
}
