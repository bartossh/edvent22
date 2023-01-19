package day10

import (
	"fmt"
	"strconv"
	"strings"
)

func DrawCTR(s string) []string {
	rows := strings.Split(s, "\n")

	tick := 0
	padding := 0
	register := 1

	pixels := make([]byte, 240)

	for _, r := range rows {
		if r == "" {
			continue
		}
		v, ok := addx(r)
		if ok {
			if isPixelInSprite(tick, register) {
				pixels[tick+padding] = 1
			}
			tick++
			if tick == 40 {
				padding += 40
				tick = 0
			}
			if isPixelInSprite(tick, register) {
				pixels[tick+padding] = 1
			}
			tick++
			if tick == 40 {
				padding += 40
				tick = 0
			}
			register += v

			continue
		}

		if isPixelInSprite(tick, register) {
			pixels[tick+padding] = 1
		}
		tick++
		if tick == 40 {
			padding += 40
			tick = 0
		}

	}
	fmt.Printf("%v\n", pixels)

	ctrLines := make([]string, 0, 6)
	for i := 0; i < 240; i += 40 {
		var buf strings.Builder
		for _, p := range pixels[i : i+40] {
			if p == 1 {
				buf.WriteString("#")
				continue
			}
			buf.WriteString(".")
			_ = 1
		}
		ctrLines = append(ctrLines, buf.String())
	}

	return ctrLines

}

func isPixelInSprite(p, s int) bool {
	if p == s-1 || p == s || p == s+1 {
		return true
	}
	return false
}

func CalculateRegisterSum(s string) int {
	rows := strings.Split(s, "\n")

	tick := 0
	register := 1

	sum := 0

	for _, r := range rows {
		if r == "" {
			continue
		}

		v, ok := addx(r)
		if ok {
			tick++
			if isSumCycle(tick) {
				sum = sum + (tick * register)
			}
			tick++
			if isSumCycle(tick) {
				sum = sum + (tick * register)
			}
			register += v
			continue
		}

		tick++
		if isSumCycle(tick) {
			sum = sum + (tick * register)
		}
	}
	return sum
}

func addx(s string) (int, bool) {
	if strings.Contains(s, "addx") {
		v, err := strconv.Atoi(s[5:])
		if err != nil {
			panic("conversion to number impossible")
		}
		return v, true
	}
	return 0, false
}

func isSumCycle(v int) bool {
	switch v {
	case 20, 60, 100, 140, 180, 220:
		return true
	default:
		return false
	}
}
