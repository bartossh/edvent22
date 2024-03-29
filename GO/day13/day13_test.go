package day13

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"advent.com/2022/data"
)

func TestSumOfIndicesInRightOrderTestData(t *testing.T) {
	sum := CalcSumOfIndicesInOrder(data.TestReorderedDistressPackets)
	fmt.Printf("Sum of indices in order for test data is: %v\n", sum)
}

func TestSumOfIndicesInRightOrderTestData2(t *testing.T) {
	sum := aCalcSumOfIndicesInOrder(data.TestReorderedDistressPackets)
	fmt.Printf("Sum of indices in order for test data is: %v\n", sum)
}

func TestSumOfIndicesInRightOrderPuzzleData2(t *testing.T) {
	sum := aCalcSumOfIndicesInOrder(data.PuzzleReorderedDistressPackets)
	fmt.Printf("Sum of indices in order for puzzle data is: %v\n", sum)
}

func TestDecoderKeyOrderTestData2(t *testing.T) {
	sum := aDecodeKeyOfIndices(data.TestReorderedDistressPackets)
	fmt.Printf("Decoder key for test data is: %v\n", sum)
}

func TestDecoderKeyOrderPuzzleData2(t *testing.T) {
	sum := aDecodeKeyOfIndices(data.PuzzleReorderedDistressPackets)
	fmt.Printf("Decoder key  for puzzle data is: %v\n", sum)
}

func aDecodeKeyOfIndices(d string) int {

	pck := make([]string, 0)
	rows := strings.Split(d, "\n")

	for _, row := range rows {
		if row == "" {
			continue
		}
		pck = append(pck, row)
	}

	var pockets []tree
	for _, pac := range pck {
		pockets = append(pockets, readTree(pac))
	}
	pockets = append(pockets, readTree("[[2]]"))
	pockets = append(pockets, readTree("[[6]]"))

	sort.Slice(pockets, func(i, j int) bool {
		return areOrdered(pockets[i], pockets[j]) == 1
	})

	decoderKey := 1
	for i, p := range pockets {
		if areOrdered(p, readTree("[[2]]")) == 0 || areOrdered(p, readTree("[[6]]")) == 0 {
			decoderKey *= i + 1
		}
	}
	return decoderKey
}

func aCalcSumOfIndicesInOrder(d string) int {
	sum := 0

	packets := make([]string, 0)
	rows := strings.Split(d, "\n")

	for _, row := range rows {
		if row == "" {
			continue
		}
		packets = append(packets, row)
	}

	idx := 1

	for i := 0; i < len(packets)-1; i += 2 {
		if areOrdered(readTree(packets[i]), readTree(packets[i+1])) == 1 {
			sum += idx
		}
		idx++
	}

	return sum
}

type tree struct {
	valueLeaf int
	elements  []*tree
	father    *tree
}

func readTree(input string) tree {
	root := tree{-1, []*tree{}, nil}
	temp := &root

	var currentNumber string
	for _, r := range input {
		switch r {
		case '[':
			newTree := tree{-1, []*tree{}, temp}
			temp.elements = append(temp.elements, &newTree)
			temp = &newTree
		case ']':
			if len(currentNumber) > 0 {
				number, _ := strconv.Atoi(currentNumber)
				temp.valueLeaf = number
				currentNumber = ""
			}
			temp = temp.father
		case ',':
			if len(currentNumber) > 0 {
				number, _ := strconv.Atoi(currentNumber)
				temp.valueLeaf = number
				currentNumber = ""
			}
			temp = temp.father
			newTree := tree{-1, []*tree{}, temp}
			temp.elements = append(temp.elements, &newTree)
			temp = &newTree
		default:
			currentNumber += string(r)
		}
	}
	return root
}

func areOrdered(first, second tree) int {
	switch {
	case len(first.elements) == 0 && len(second.elements) == 0:
		if first.valueLeaf > second.valueLeaf {
			return -1
		} else if first.valueLeaf == second.valueLeaf {
			return 0
		}
		return 1

	case first.valueLeaf >= 0:
		return areOrdered(tree{-1, []*tree{&first}, nil}, second)

	case second.valueLeaf >= 0:
		return areOrdered(first, tree{-1, []*tree{&second}, nil})
	default:
		var i int
		for i = 0; i < len(first.elements) && i < len(second.elements); i++ {
			ordered := areOrdered(*first.elements[i], *second.elements[i])
			if ordered != 0 {
				return ordered
			}
		}
		if i < len(first.elements) {
			return -1
		} else if i < len(second.elements) {
			return 1
		}
	}
	return 0
}
