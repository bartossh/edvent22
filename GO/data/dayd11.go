package data

const TestMonkeySchema = `
Monkey 0:
Starting items: 79, 98
Operation: new = old * 19
Test: divisible by 23
  If true: throw to monkey 2
  If false: throw to monkey 3

Monkey 1:
Starting items: 54, 65, 75, 74
Operation: new = old + 6
Test: divisible by 19
  If true: throw to monkey 2
  If false: throw to monkey 0

Monkey 2:
Starting items: 79, 60, 97
Operation: new = old * old
Test: divisible by 13
  If true: throw to monkey 1
  If false: throw to monkey 3

Monkey 3:
Starting items: 74
Operation: new = old + 3
Test: divisible by 17
  If true: throw to monkey 0
  If false: throw to monkey 1
`

type Monkey struct {
	Num       int
	Items     []int
	Operation func(old int) int
	Test      func(new int) bool
	Throw     func(t bool) int
	Prime     int
}

var TestMonkeysSchema = []Monkey{
	{
		Num:   0,
		Items: []int{79, 98},
		Operation: func(old int) int {
			return old * 19
		},
		Test: func(new int) bool {
			return new%23 == 0
		},
		Throw: func(t bool) int {
			if t {
				return 2
			}
			return 3
		},
		Prime: 23,
	},
	{
		Num:   1,
		Items: []int{54, 65, 75, 74},
		Operation: func(old int) int {
			return old + 6
		},
		Test: func(new int) bool {
			return new%19 == 0
		},
		Throw: func(t bool) int {
			if t {
				return 2
			}
			return 0
		},
		Prime: 19,
	},
	{
		Num:   2,
		Items: []int{79, 60, 97},
		Operation: func(old int) int {
			return old * old
		},
		Test: func(new int) bool {
			return new%13 == 0
		},
		Throw: func(t bool) int {
			if t {
				return 1
			}
			return 3
		},
		Prime: 13,
	},
	{
		Num:   3,
		Items: []int{74},
		Operation: func(old int) int {
			return old + 3
		},
		Test: func(new int) bool {
			return new%17 == 0
		},
		Throw: func(t bool) int {
			if t {
				return 0
			}
			return 1
		},
		Prime: 17,
	},
}

const PuzzleMonkeyThrows = `
Monkey 0:
  Starting items: 75, 75, 98, 97, 79, 97, 64
  Operation: new = old * 13
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 7

Monkey 1:
  Starting items: 50, 99, 80, 84, 65, 95
  Operation: new = old + 2
  Test: divisible by 3
    If true: throw to monkey 4
    If false: throw to monkey 5

Monkey 2:
  Starting items: 96, 74, 68, 96, 56, 71, 75, 53
  Operation: new = old + 1
  Test: divisible by 11
    If true: throw to monkey 7
    If false: throw to monkey 3

Monkey 3:
  Starting items: 83, 96, 86, 58, 92
  Operation: new = old + 8
  Test: divisible by 17
    If true: throw to monkey 6
    If false: throw to monkey 1

Monkey 4:
  Starting items: 99
  Operation: new = old * old
  Test: divisible by 5
    If true: throw to monkey 0
    If false: throw to monkey 5

Monkey 5:
  Starting items: 60, 54, 83
  Operation: new = old + 4
  Test: divisible by 2
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 6:
  Starting items: 77, 67
  Operation: new = old * 17
  Test: divisible by 13
    If true: throw to monkey 4
    If false: throw to monkey 1

Monkey 7:
  Starting items: 95, 65, 58, 76
  Operation: new = old + 5
  Test: divisible by 7
    If true: throw to monkey 3
    If false: throw to monkey 6
`

var PuzzleMonkeysSchema = []Monkey{
	{
		Num:   0,
		Items: []int{75, 75, 98, 97, 79, 97, 64},
		Operation: func(old int) int {
			return old * 13
		},
		Test: func(new int) bool {
			return new%19 == 0
		},
		Throw: func(t bool) int {
			if t {
				return 2
			}
			return 7
		},
		Prime: 19,
	},
	{
		Num:   1,
		Items: []int{50, 99, 80, 84, 65, 95},
		Operation: func(old int) int {
			return old + 2
		},
		Test: func(new int) bool {
			return new%3 == 0
		},
		Throw: func(t bool) int {
			if t {
				return 4
			}
			return 5
		},
		Prime: 3,
	},
	{
		Num:   2,
		Items: []int{96, 74, 68, 96, 56, 71, 75, 53},
		Operation: func(old int) int {
			return old + 1
		},
		Test: func(new int) bool {
			return new%11 == 0
		},
		Throw: func(t bool) int {
			if t {
				return 7
			}
			return 3
		},
		Prime: 11,
	},
	{
		Num:   3,
		Items: []int{83, 96, 86, 58, 92},
		Operation: func(old int) int {
			return old + 8
		},
		Test: func(new int) bool {
			return new%17 == 0
		},
		Throw: func(t bool) int {
			if t {
				return 6
			}
			return 1
		},
		Prime: 17,
	},
	{
		Num:   4,
		Items: []int{99},
		Operation: func(old int) int {
			return old * old
		},
		Test: func(new int) bool {
			return new%5 == 0
		},
		Throw: func(t bool) int {
			if t {
				return 0
			}
			return 5
		},
		Prime: 5,
	},
	{
		Num:   5,
		Items: []int{60, 54, 83},
		Operation: func(old int) int {
			return old + 4
		},
		Test: func(new int) bool {
			return new%2 == 0
		},
		Throw: func(t bool) int {
			if t {
				return 2
			}
			return 0
		},
		Prime: 2,
	},
	{
		Num:   6,
		Items: []int{77, 67},
		Operation: func(old int) int {
			return old * 17
		},
		Test: func(new int) bool {
			return new%13 == 0
		},
		Throw: func(t bool) int {
			if t {
				return 4
			}
			return 1
		},
		Prime: 13,
	},
	{
		Num:   7,
		Items: []int{95, 65, 58, 76},
		Operation: func(old int) int {
			return old + 5
		},
		Test: func(new int) bool {
			return new%7 == 0
		},
		Throw: func(t bool) int {
			if t {
				return 3
			}
			return 6
		},
		Prime: 7,
	},
}
