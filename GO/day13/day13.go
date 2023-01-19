package day13

import (
	"fmt"
	"strconv"
	"strings"
)

func CalcSumOfIndicesInOrder(d string) int {
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
		if compareTwoPackets(packets[i], packets[i+1]) {
			sum += idx
		}
		idx++
	}

	return sum
}

func compareTwoPackets(a, b string) bool {

	nsA := createNesting(a)
	nsB := createNesting(b)

	iA, iB := 0, 0

	for {
		if iA == len(nsA) || iB == len(nsB) {
			if iA > iB {
				return false
			}
			break
		}
		numbersA := getNestedNumbers(nsA[iA])
		numbersB := getNestedNumbers(nsB[iB])
		// fmt.Printf("%v | %v\n", numbersA, numbersB)

		if len(numbersB) < len(numbersA) {
			return false
		}

		for i := range numbersA {
			if numbersA[i] <= numbersB[i] {
				continue
			}
			return false
		}
		iA++
		iB++
	}

	return true
}

func createNesting(s string) []string {
	fmt.Printf(" == %v ==\n", s)
	openB := make([]int, 0)
	pairs := make([][2]int, 0)
	for i := range s {
		if s[i] == '[' {
			openB = append(openB, i)
		}
		if s[i] == ']' {
			oi := openB[len(openB)-1]
			openB = openB[:len(openB)-1]
			pairs = append(pairs, [2]int{oi, i})
		}
	}

	if len(openB) != 0 {
		panic("brackets are not matching")
	}

	// for i, j := 0, len(pairs)-1; i < j; i, j = i+1, j-1 {
	// 	pairs[i], pairs[j] = pairs[j], pairs[i]
	// }

	nestings := make([]string, 0)

	for _, p := range pairs {
		sl := s[p[0]+1 : p[1]]
		nestings = append(nestings, sl)
		fmt.Printf("%v\n", sl)

	}
	return nestings
}

func getNestedNumbers(s string) []int {
	iob := -1
	for i, l := range s {
		if l == '[' {
			iob = i
			break
		}
	}
	icb := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ']' {
			icb = i
			break
		}
	}

	var ns string

	if icb == -1 || iob == -1 {
		ns = s
	} else {
		ns = s[:iob] + s[icb+1:]
	}

	nums := strings.Split(ns, ",")
	arr := make([]int, 0)
	for _, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			continue
		}
		arr = append(arr, n)
	}
	return arr

}
