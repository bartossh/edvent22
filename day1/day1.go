package day1

import (
	"sort"
	"strconv"
	"strings"
)

func CalcElvesMaxCalories(in string) int {
	elves := strings.Split(in, "\n\n")
	maxCal := 0
	for _, e := range elves {
		s := strings.Split(e, "\n")
		ec := sumSlice(s)
		if ec > maxCal {
			maxCal = ec
		}
	}

	return maxCal
}

func sumSlice(sl []string) int {
	sum := 0
	for _, s := range sl {
		n, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		sum += n
	}
	return sum
}

func CalcElvesMaxCaloriesLast(in string, lastElves int) int {
	elves := strings.Split(in, "\n\n")
	caloriesPerElf := make([]int, 0, len(elves))
	for _, e := range elves {
		s := strings.Split(e, "\n")
		ec := sumSlice(s)
		caloriesPerElf = append(caloriesPerElf, ec)
	}

	sort.Ints(caloriesPerElf)

	return sumLast(caloriesPerElf, lastElves)

}

func sumLast(sl []int, i int) int {
	sum := 0
	for _, v := range sl[len(sl)-i:] {
		sum += v
	}
	return sum
}
