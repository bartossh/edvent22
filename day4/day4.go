package day4

import (
	"strconv"
	"strings"
)

type boundary struct {
	low, high int
}

func CountOverlapPartly(data string) int {
	bounds := splitToArr(data)

	count := 0
	for _, b := range bounds {
		if contains(b[0], b[1]) {
			count++
			continue
		}
		if containsPartly(b[0], b[1]) {
			count++
		}
	}

	return count
}

func CountOverlaps(data string) int {
	bounds := splitToArr(data)

	count := 0
	for _, b := range bounds {
		if contains(b[0], b[1]) {
			count++
		}
	}

	return count
}

func containsPartly(a, b boundary) bool {
	if a.high < b.low {
		return false
	}
	if b.high < a.low {
		return false
	}
	return true
}

func contains(a, b boundary) bool {
	if a.low <= b.low && a.high >= b.high {
		return true
	}
	if b.low <= a.low && b.high >= a.high {
		return true
	}
	return false
}

func splitToArr(data string) [][2]boundary {
	lines := strings.Split(data, "\n")
	res := make([][2]boundary, 0, len(lines))

	for _, l := range lines {
		if l == "" {
			continue
		}
		pairs := strings.Split(l, ",")
		if len(pairs) != 2 {
			continue
		}

		left := strings.Split(pairs[0], "-")
		right := strings.Split(pairs[1], "-")
		if len(left) != 2 || len(right) != 2 {
			continue
		}
		ll, err := strconv.Atoi(left[0])
		if err != nil {
			continue
		}
		hl, err := strconv.Atoi(left[1])
		if err != nil {
			continue
		}

		lr, err := strconv.Atoi(right[0])
		if err != nil {
			continue
		}
		hr, err := strconv.Atoi(right[1])
		if err != nil {
			continue
		}
		res = append(res, [2]boundary{
			{low: ll, high: hl}, {low: lr, high: hr},
		})

	}
	return res
}
