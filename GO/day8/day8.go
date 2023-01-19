package day8

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func CountVisibleTreesFormOutsideTheGrid(data string) int {
	grid := stringToGridArr(data)
	count := 0
	for x := range grid {
		for y := range grid[x] {
			if findIsVisible(x, y, grid) {
				count++
			}
		}
	}

	return count
}

func CountsScenicScoreAnyTree(data string) int {
	grid := stringToGridArr(data)
	maxScenicView := 0
	for x := range grid {
		for y := range grid[x] {
			scenicView := calcScenicView(x, y, grid)
			if maxScenicView < scenicView {
				maxScenicView = scenicView
			}

		}
	}
	return maxScenicView
}

func calcScenicView(x, y int, grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		panic(errors.New("unexpected grid of zero length"))
	}
	if x == 0 || x == len(grid)-1 {
		return 0 // outside the grid
	}
	if y == 0 || y == len(grid[0])-1 {
		return 0 // outside the grid
	}
	visibleTop := 0
	visibleBottom := 0
	visibleLeft := 0
	visibleRight := 0

	for i := len(grid[:x]) - 1; i >= 0; i-- {
		if grid[i][y] < grid[x][y] {
			visibleTop++
			continue
		}
		visibleTop++
		break
	}

	for i := x + 1; i < len(grid); i++ {
		if grid[i][y] < grid[x][y] {
			visibleBottom++
			continue
		}
		visibleBottom++
		break
	}

	for i := len(grid[:][:y]) - 1; i >= 0; i-- {
		if grid[x][i] < grid[x][y] {
			visibleLeft++
			continue
		}
		visibleLeft++
		break
	}

	for i := y + 1; i < len(grid[x]); i++ {
		if grid[x][i] < grid[x][y] {
			visibleRight++
			continue
		}
		visibleRight++
		break
	}

	return visibleTop * visibleBottom * visibleLeft * visibleRight
}

func findIsVisible(x, y int, grid [][]int) bool {
	if len(grid) == 0 || len(grid[0]) == 0 {
		panic(errors.New("unexpected grid of zero length"))
	}
	if x == 0 || x == len(grid)-1 {
		return true // outside the grid
	}
	if y == 0 || y == len(grid[0])-1 {
		return true // outside the grid
	}

	visibleTop := true
	visibleBottom := true
	visibleLeft := true
	visibleRight := true

	for ix := 0; ix < len(grid); ix++ {
		if ix == x {
			continue
		}
		if ix < x && grid[ix][y] >= grid[x][y] {
			visibleTop = false
			ix = x
			continue
		}
		if ix > x && grid[ix][y] >= grid[x][y] {
			visibleBottom = false
			break
		}

	}

	for iy := 0; iy < len(grid[x]); iy++ {
		if iy == y {
			continue
		}
		if iy < y && grid[x][iy] >= grid[x][y] {
			visibleLeft = false
			iy = y
			continue
		}
		if iy > y && grid[x][iy] >= grid[x][y] {
			visibleRight = false
			break
		}

	}

	return visibleTop || visibleBottom || visibleLeft || visibleRight
}

func stringToGridArr(s string) [][]int {
	arr := make([][]int, 0)

	for _, row := range strings.Split(s, "\n") {
		if row == "" {
			continue
		}
		r := make([]int, 0, len(row))
		for _, s := range strings.Split(row, "") {
			v, err := strconv.Atoi(s)
			if err != nil {
				panic(fmt.Errorf("unexpected err with ascii to integer conversion %w", err))
			}
			r = append(r, v)
		}
		arr = append(arr, r)
	}
	return arr
}
