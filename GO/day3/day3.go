package day3

import (
	"strings"
	"unicode"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

const (
	lowercaseStart = 1
	uppercaseStart = 27
)

func CalcElvesGroupPriority(data string) int {
	rucksacks := strings.Split(data, "\n")
	sum := 0
	group := make([]string, 0, 3)
	counter := 0
	for _, r := range rucksacks {
		if r == "" {
			continue
		}
		counter++
		group = append(group, r)
		if counter%3 == 0 {
			sum += calcSingleGroupPriority(group)
			group = make([]string, 0, 3)
		}
	}
	return sum
}

func calcSingleGroupPriority(g []string) int {
	if len(g) != 3 {
		return 0
	}
	l := findRepetitions(g[0], g[1])
	r := findRepetitions(g[1], g[2])
	repetition, ok := findRepetition(l, r)
	if !ok {
		return 0
	}
	return calcRepValue(repetition)

}

func findRepetitions(l, r string) string {
	set := make(map[rune]struct{}, len(l))
	var s strings.Builder
	for _, v := range l {
		set[v] = struct{}{}
	}
	for _, v := range r {
		if _, ok := set[v]; ok {
			s.WriteRune(v)
		}
	}
	return s.String()
}

func CalcSumPriorities(data string) int {
	rucksacks := strings.Split(data, "\n")
	sum := 0
	for _, r := range rucksacks {
		if r == "" {
			continue
		}
		sum += singleRucksackRepeatedPriority(r)
	}
	return sum
}

func singleRucksackRepeatedPriority(r string) int {
	half := len(r) / 2
	repetition, ok := findRepetition(r[:half], r[half:])
	if !ok {
		return 0
	}
	return calcRepValue(repetition)
}

func findRepetition(l, r string) (rune, bool) {
	set := make(map[rune]struct{}, len(l))
	for _, v := range l {
		set[v] = struct{}{}
	}
	for _, v := range r {
		if _, ok := set[v]; ok {
			return v, true
		}
	}
	return 0, false
}

func calcRepValue(r rune) int {
	sum := 0

	for i, a := range alphabet {
		switch {
		case r == a:
			sum += i + lowercaseStart
		case r == unicode.ToUpper(a):
			sum += i + uppercaseStart
		}
	}

	return sum
}
