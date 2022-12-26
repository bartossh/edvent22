package day22

import (
	"strconv"
	"strings"
)

const (
	left  = 'L'
	right = 'R'
)

var progress = [4]point{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}

type step struct {
	turn  rune
	tiles int
}

type point struct {
	x, y int
}

func MazeTraversal(d string) int {
	rows := strings.Split(d, "\n")

	var path string
	var maxX int
	maze := make([][]rune, 0)

	for i := range rows[:len(rows)-1] {
		if len(rows[i]) > maxX {
			maxX = len(rows[i])
		}
	}

	for i, r := range rows {
		if r == "" {
			continue
		}
		if i == len(rows)-1 {
			path = r
		}

		row := make([]rune, 0, maxX)
		for _, l := range r {
			row = append(row, l)
		}
		maze = append(maze, row)
	}

	pth := rawPathToSteps(path)

	pnt := findStart(maze)
	var dir point
	var move step

Outer:
	for len(pth) > 0 {
		move, pth = pth[0], pth[1:]
		dir = getDirection(dir, move.turn)
	Inner:
		for i := move.tiles; i > 0; i-- {
			x := pnt.x + dir.x
			y := pnt.y + dir.y

			if x > 0 && x < len(maze[y]) && y > 0 && y < len(maze) && maze[y][x] == '.' {
				pnt.x = x
				pnt.y = y
				continue Inner
			}
			if x > 0 && x < len(maze[y]) && y > 0 && y < len(maze) && maze[y][x] == '#' {
				continue Outer
			}

			if x != pnt.x {
				xx := findNotEmptyX(maze, dir, pnt)
				if xx == x {
					break Outer
				}
				pnt.x = xx
				continue Inner
			}
			if y != pnt.y {
				yy := findNotEmptyY(maze, dir, pnt)
				if yy == y {
					break Outer
				}
				pnt.y = yy
				continue Inner
			}
		}
	}

	var face int

	for i := range progress {
		if progress[i].x == dir.x && progress[i].y == dir.y {
			face = i
		}
	}

	return (pnt.y+1)*1000 + (pnt.x+1)*4 + face
}

func findNotEmptyY(maze [][]rune, dir, curr point) int {
	switch dir {
	case point{0, -1}:
		for y := curr.y; y < len(maze); y++ {
			if maze[y][curr.x] == 0 && maze[y-1][curr.x] != '#' {
				return y - 1
			}
			if maze[y][curr.x] == 0 && maze[y-1][curr.x] == '#' {
				return curr.y
			}
		}
		return len(maze) - 1
	case point{0, 1}:
		for y := curr.y; y >= 0; y-- {
			if maze[y][curr.x] == 0 && maze[y+1][curr.x] != '#' {
				return y + 1
			}
			if maze[y][curr.x] == 0 && maze[y+1][curr.x] == '#' {
				return curr.y
			}

		}
		return 0
	default:
		return curr.y
	}
}

func findNotEmptyX(maze [][]rune, dir, curr point) int {
	switch dir {
	case point{-1, 0}:
		for x := curr.x; x < len(maze[curr.y]); x++ {
			if maze[curr.y][x] == 0 && maze[curr.y][x-1] != '#' {
				return x - 1
			}
			if maze[curr.y][x] == 0 && maze[curr.y][x-1] == '#' {
				return curr.x
			}
		}
		return len(maze[curr.y]) - 1
	case point{1, 0}:
		for x := curr.x; x >= 0; x-- {
			if maze[curr.y][x] == 0 && maze[curr.y][x+1] != '#' {
				return x + 1
			}
			if maze[curr.y][x] == 0 && maze[curr.y][x+1] == '#' {
				return curr.x
			}

		}
		return 0
	default:
		return curr.x
	}
}

func findStart(maze [][]rune) point {
	for i := range maze[0] {
		if maze[0][i] == '.' {
			return point{x: i, y: 0}
		}
	}
	return point{}
}

func getDirection(curr point, turn rune) point {
	var currentIdx int
	for i, p := range progress {
		if p.x == curr.x && p.y == curr.y {
			currentIdx = i
		}
	}

	switch turn {
	case 'R':
		idx := currentIdx + 1
		if idx == len(progress) {
			idx = 0
		}
		return progress[idx]
	case 'L':
		idx := currentIdx - 1
		if idx == -1 {
			idx = len(progress) - 1
		}
		return progress[idx]
	default:
		return progress[0]
	}
}

func rawPathToSteps(raw string) []step {
	raw = "S" + raw
	i := 0
	j := i + 1
	steps := make([]step, 0)
	for {
		if raw[j] == 'L' || raw[j] == 'R' {

			tiles, err := strconv.Atoi(raw[i+1 : j])
			if err != nil {
				panic(err)
			}
			dir := rune(raw[i])
			stp := step{
				tiles: tiles,
				turn:  dir,
			}
			steps = append(steps, stp)
			i = j
			j = i + 1
		}
		j++

		if j == len(raw) {

			tiles, err := strconv.Atoi(raw[i+1 : j])
			if err != nil {
				panic(err)
			}
			dir := rune(raw[i])
			stp := step{
				tiles: tiles,
				turn:  dir,
			}
			steps = append(steps, stp)
			break
		}
	}

	return steps
}
