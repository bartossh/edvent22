package day3

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

const preTestData = `
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

func TestRucksacksPriorityPreTestData(t *testing.T) {
	result := CalcSumPriorities(preTestData)
	fmt.Printf("Pretest sum of rucksacks priorities is: %v\n", result)
}

func TestRucksacksPriority(t *testing.T) {
	result := CalcSumPriorities(data.Rucksacks)
	fmt.Printf("Input data sum of rucksacks priorities is: %v\n", result)
}

func TestCalcElvesGroupPriorityPretestData(t *testing.T) {
	result := CalcElvesGroupPriority(preTestData)
	fmt.Printf("Pretest sum of group priorities is: %v\n", result)
}

func TestCalcElvesGroupPriority(t *testing.T) {
	result := CalcElvesGroupPriority(data.Rucksacks)
	fmt.Printf("Input sum of group priorities is: %v\n", result)
}
