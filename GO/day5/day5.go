package day5

import (
	"strconv"
	"strings"
)

func MoveCrates9000(crates, moves string) string {
	cratesArr := cratesStringToArr(crates)
	movesArr := movesStringToArr(moves)
	cratesArr = crateStack(cratesArr)

	for _, m := range movesArr {
		cratesArr = applyMove9000(m[0], m[1], m[2], cratesArr)
	}

	return readTopCrates(cratesArr)
}

func MoveCrates9001(crates, moves string) string {
	cratesArr := cratesStringToArr(crates)
	movesArr := movesStringToArr(moves)
	cratesArr = crateStack(cratesArr)

	for _, m := range movesArr {
		cratesArr = applyMove9001(m[0], m[1], m[2], cratesArr)
	}

	return readTopCrates(cratesArr)
}

func cratesStringToArr(data string) [][]rune {
	rows := strings.Split(data, "\n")
	crates := make([][]rune, 0)
	for j, row := range rows {
		if row == "" {
			continue
		}
		crate := make([]rune, 0)
		if j == len(rows)-2 {
			// row numbers not included
			continue
		}
		for i, r := range row {
			if i == 1 || (i-1)%4 == 0 {
				crate = append(crate, r)
			}
		}
		crates = append(crates, crate)
	}
	return crates
}

func movesStringToArr(data string) [][]int {
	data = strings.ReplaceAll(data, "	", "")
	rows := strings.Split(data, "\n")
	moves := make([][]int, 0)

	for _, row := range rows {
		row = strings.ReplaceAll(row, "move ", "")
		row = strings.ReplaceAll(row, " from ", ",")
		row = strings.ReplaceAll(row, " to ", ",")
		nums := strings.Split(row, ",")
		if len(nums) != 3 {
			continue
		}
		move := make([]int, 0, 3)
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				continue
			}
			move = append(move, n)
		}
		moves = append(moves, move)
	}
	return moves
}

func crateStack(arr [][]rune) [][]rune {
	if len(arr) == 0 {
		return nil
	}
	stack := make([][]rune, len(arr[0]))
	for i := range arr {
		for j := range arr[i] {
			r := arr[i][j]
			if r == ' ' {
				continue
			}
			stack[j] = append(stack[j], r)
		}
	}
	return stack
}

func applyMove9000(moves, from, to int, arr [][]rune) [][]rune {
	for m := 0; m < moves; m++ {
		if len(arr[from-1]) == 0 {
			break
		}
		v := arr[from-1][0]
		arr[from-1] = arr[from-1][1:]
		arr[to-1] = append([]rune{v}, arr[to-1]...)
	}
	return arr
}

func applyMove9001(moves, from, to int, arr [][]rune) [][]rune {

	if len(arr[from-1]) == 0 {
		return arr
	}
	if len(arr[from-1]) <= moves {
		arr[to-1] = append(arr[from-1], arr[to-1]...)
		arr[from-1] = make([]rune, 0)
		return arr
	}
	cs := make([]rune, moves)
	copy(cs, arr[from-1][:moves])
	rs := make([]rune, len(arr[from-1])-moves)
	copy(rs, arr[from-1][moves:])
	arr[to-1] = append(cs, arr[to-1]...)
	arr[from-1] = rs

	return arr
}

func readTopCrates(arr [][]rune) string {
	var res strings.Builder
	for _, stack := range arr {
		if len(stack) == 0 {
			continue
		}
		r := stack[0]
		res.WriteRune(r)
	}
	return res.String()
}
