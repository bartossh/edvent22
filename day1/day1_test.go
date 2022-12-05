package day1

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

func TestDay1_1(t *testing.T) {
	result := CalcElvesMaxCalories(data.InputCalories)
	fmt.Printf("max calories carried by elf is: %v\n", result)
}

func TestDay1_2(t *testing.T) {
	largestThree := CalcElvesMaxCaloriesLast(data.InputCalories, 3)
	fmt.Printf("last three elves calories count: %v\n", largestThree)
}
