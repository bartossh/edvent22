package day11

import (
	"sort"

	"advent.com/2022/data"
)

// 45135754

func CalcKeepAwayWorryManaged(monkeys []data.Monkey) int {
	return calcKeepAway(monkeys, 20, func(v, d int) int {
		v = v / 3
		if v == 0 {
			return 0
		}
		return v % d
	})
}

func CalcKeepAwayWorryNotManaged(monkeys []data.Monkey) int {
	return calcKeepAway(monkeys, 10000, func(v, d int) int { return v % d })

}

func calcKeepAway(monkeys []data.Monkey, rounds int, managingWorryLevel func(int, int) int) int {
	monkeyIndex := 0
	inspections := make([]int, len(monkeys))

	denominator := monkeys[0].Prime

	for _, v := range monkeys[1:] {
		denominator = (v.Prime * denominator)
	}

R:
	for {
		if monkeyIndex == len(monkeys) {
			monkeyIndex = 0
			rounds--
		}
		if rounds == 0 {
			break R
		}
		m := &monkeys[monkeyIndex]
	M:
		for {
			if v, ok := removeFromMonkey(m); ok {
				nv := m.Operation(v)
				nv = managingWorryLevel(nv, denominator)
				t := m.Test(nv)
				passTo := m.Throw(t)
				appendToMonkey(&monkeys[passTo], nv)
				inspections[monkeyIndex]++
				continue
			}
			break M
		}

		monkeyIndex++
	}

	sort.Ints(inspections)
	return inspections[len(inspections)-2] * inspections[len(inspections)-1]
}

func appendToMonkey(m *data.Monkey, v int) {
	m.Items = append(m.Items, v)
}

func removeFromMonkey(m *data.Monkey) (int, bool) {
	if len(m.Items) == 0 {
		return 0, false
	}
	v := m.Items[0]
	m.Items = m.Items[1:]
	return v, true
}
