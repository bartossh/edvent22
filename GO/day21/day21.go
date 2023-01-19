package day21

import (
	"fmt"
	"strings"
)

const (
	multiplication = "*"
	division       = "/"
	reduction      = "-"
	sum            = "+"
	equal          = "="
)

type puzzle struct {
	operation string
	a, b      string
}

type solver struct {
	puzzles map[string]puzzle
	known   map[string]int
}

// func PredictRootMonkeyNumberHumanGuess(d string) int {
// 	rows := strings.Split(d, "\n")

// 	slv := solver{
// 		puzzles: make(map[string]puzzle),
// 		known:   make(map[string]int),
// 	}

// 	for _, row := range rows {
// 		if row == "" {
// 			continue
// 		}

// 		var monkeyMain, monkeyA, operation, monkeyB string

// 		_, err := fmt.Sscanf(row, "%s %s %s %s", &monkeyMain, &monkeyA, &operation, &monkeyB)

// 		if err == nil {
// 			monkeyMain = monkeyMain[:len(monkeyMain)-1]
// 			var opr string
// 			if monkeyMain == "root" {
// 				operation = "="
// 			}

// 			switch operation {
// 			case "*":
// 				opr = multiplication
// 				slv.puzzles[monkeyMain] = puzzle{
// 					operation: opr,
// 					a:         monkeyA,
// 					b:         monkeyB,
// 				}
// 				continue
// 			case "/":
// 				opr = division
// 				slv.puzzles[monkeyMain] = puzzle{
// 					operation: opr,
// 					a:         monkeyA,
// 					b:         monkeyB,
// 				}
// 				continue
// 			case "-":
// 				opr = reduction
// 				slv.puzzles[monkeyMain] = puzzle{
// 					operation: opr,
// 					a:         monkeyA,
// 					b:         monkeyB,
// 				}
// 				continue
// 			case "+":
// 				opr = sum
// 				slv.puzzles[monkeyMain] = puzzle{
// 					operation: opr,
// 					a:         monkeyA,
// 					b:         monkeyB,
// 				}
// 				continue
// 			case "=":
// 				opr = sum
// 				slv.puzzles[monkeyMain] = puzzle{
// 					operation: equal,
// 					a:         monkeyA,
// 					b:         monkeyB,
// 				}
// 				continue
// 			default:
// 				panic(fmt.Errorf("not parsed, %s, %w", row, err))
// 			}

// 		}

// 		var number int

// 		_, err = fmt.Sscanf(row, "%s %d", &monkeyMain, &number)
// 		monkeyMain = monkeyMain[:len(monkeyMain)-1]
// 		if err != nil {
// 			panic(fmt.Errorf("at least one operation must pass: %s, %w", row, err))
// 		}
// 		slv.known[monkeyMain] = number

// 	}

// 	leafs := []string{slv.puzzles["root"].a, slv.puzzles["root"].b}

// 	humn := "humn"
// 	switch {
// 	case slv.isHumanPath(leafs[0]):
// 		slv.puzzles[leafs[0]] = puzzle{
// 			operation: equal,
// 			a:         leafs[1],
// 			b:         leafs[1],
// 		}
// 		delete(slv.puzzles, "root")
// 	case slv.isHumanPath(leafs[1]):
// 		slv.puzzles[leafs[1]] = puzzle{
// 			operation: equal,
// 			a:         leafs[0],
// 			b:         leafs[0],
// 		}
// 		delete(slv.puzzles, "root")
// 	}
// 	delete(slv.known, humn)
// 	var next string
// 	toReverse := []string{humn}

// 	for len(toReverse) > 0 {
// 		next, toReverse = getNext(toReverse)
// 	Inner:
// 		for k, p := range slv.puzzles {

// 			if p.a == next {
// 				if p.operation == equal {
// 					break Inner
// 				}
// 				delete(slv.puzzles, k)
// 				a, pa, b, _ := revertOperation(next, p)
// 				slv.puzzles[a] = pa

// 				toReverse = append(toReverse, k)

// 				if _, ok := slv.known[b]; !ok {
// 					toReverse = append(toReverse, b)
// 				}
// 				break Inner
// 			}

// 			if p.b == next {
// 				if p.operation == equal {
// 					break Inner
// 				}
// 				delete(slv.puzzles, k)
// 				a, _, b, pb := revertOperation(next, p)
// 				slv.puzzles[b] = pb

// 				toReverse = append(toReverse, k)

// 				if _, ok := slv.known[a]; !ok {
// 					toReverse = append(toReverse, a)
// 				}
// 				break Inner
// 			}
// 		}

// 	}

// 	return slv.solve(humn)
// }

// func (s *solver) isHumanPath(b string) bool {
// 	var n string
// 	validate := []string{b}
// 	for len(validate) > 0 {
// 		n, validate = getNext(validate)
// 		if n == "humn" {

// 			return true
// 		}
// 		if v, ok := s.puzzles[n]; ok {
// 			validate = append(validate, v.a)
// 			validate = append(validate, v.b)

// 		}
// 	}
// 	return false
// }

// func revertOperation(n string, p puzzle) (string, puzzle, string, puzzle) {

// 	switch p.operation {
// 	case sum:
// 		return p.a, puzzle{operation: reduction, a: n, b: p.b}, p.b, puzzle{operation: reduction, a: n, b: p.a}
// 	case reduction:
// 		return p.a, puzzle{operation: sum, a: n, b: p.b}, p.b, puzzle{operation: reduction, a: p.a, b: n}
// 	case division:
// 		return p.a, puzzle{operation: multiplication, a: n, b: p.b}, p.b, puzzle{operation: division, a: p.a, b: n}
// 	case multiplication:
// 		return p.a, puzzle{operation: division, a: n, b: p.b}, p.b, puzzle{operation: division, a: n, b: p.a}
// 	default:
// 		return p.a, puzzle{operation: equal, a: n, b: p.b}, p.b, puzzle{operation: equal, a: n, b: p.a}
// 	}

// }

func PredictRootMonkeyNumber(d, r string) int {
	rows := strings.Split(d, "\n")

	slv := solver{
		puzzles: make(map[string]puzzle),
		known:   make(map[string]int),
	}

	for _, row := range rows {
		if row == "" {
			continue
		}

		var monkeyMain, monkeyA, operation, monkeyB string

		_, err := fmt.Sscanf(row, "%s %s %s %s", &monkeyMain, &monkeyA, &operation, &monkeyB)
		monkeyMain = monkeyMain[:len(monkeyMain)-1]
		if err == nil {
			var opr string
			switch operation {
			case "*":
				opr = multiplication
				slv.puzzles[monkeyMain] = puzzle{
					operation: opr,
					a:         monkeyA,
					b:         monkeyB,
				}
				continue
			case "/":
				opr = division
				slv.puzzles[monkeyMain] = puzzle{
					operation: opr,
					a:         monkeyA,
					b:         monkeyB,
				}
				continue
			case "-":
				opr = reduction
				slv.puzzles[monkeyMain] = puzzle{
					operation: opr,
					a:         monkeyA,
					b:         monkeyB,
				}
				continue
			case "+":
				opr = sum
				slv.puzzles[monkeyMain] = puzzle{
					operation: opr,
					a:         monkeyA,
					b:         monkeyB,
				}
				continue
			default:
				panic(fmt.Errorf("not parsed, %s, %w", row, err))
			}

		}

		var number int

		_, err = fmt.Sscanf(row, "%s %d", &monkeyMain, &number)
		monkeyMain = monkeyMain[:len(monkeyMain)-1]
		if err != nil {
			panic(fmt.Errorf("at least one operation must pass: %s, %w", row, err))
		}
		slv.known[monkeyMain] = number

	}

	return slv.solve(r)

}

func (s *solver) solve(r string) int {
	unknown := make([]string, 0, len(s.puzzles))
	for n := range s.puzzles {
		unknown = append(unknown, n)
	}
	var next string
	for len(unknown) > 0 {
		next, unknown = getNext(unknown)
		p := s.puzzles[next]
		a, okA := s.known[p.a]
		b, okB := s.known[p.b]
		if okA && okB {
			var result int
			switch p.operation {
			case multiplication:
				result = a * b
			case division:
				result = a / b
			case reduction:
				result = a - b
			case sum:
				result = a + b
			case equal:
				result = a
			default:
				panic("it is required to know the operation")
			}
			s.known[next] = result
			continue
		}
		unknown = append(unknown, next)
	}

	return s.known[r]
}

func getNext(arr []string) (string, []string) {
	return arr[0], arr[1:]
}

func contains(m string, arr []string) bool {
	for _, mm := range arr {
		if mm == m {
			return true
		}
	}
	return false
}
