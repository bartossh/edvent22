package day9

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	L = iota
	R
	U
	D
)

type move struct {
	dir   int
	steps int
}

func dataToMoves(d string) []move {
	rows := strings.Split(d, "\n")

	moves := make([]move, 0)
	for _, r := range rows {
		if r == "" || len(r) < 3 {
			continue
		}
		var dir int
		switch string(r[0]) {
		case "L":
			dir = L
		case "R":
			dir = R
		case "U":
			dir = U
		case "D":
			dir = D
		default:
			panic("direction unknown")
		}

		steps, err := strconv.Atoi(string(r[2:]))
		if err != nil {
			panic("expecting integer")
		}

		moves = append(moves, move{
			dir:   dir,
			steps: steps,
		})
	}
	return moves
}

func hash(x, y int) string {
	return fmt.Sprintf("%v-%v", x, y)
}

func moveToCoords(dir int) (x, y int) {
	switch dir {
	case L:
		x, y = -1, 0
	case R:
		x, y = 1, 0
	case U:
		x, y = 0, 1
	case D:
		x, y = 0, -1
	}
	return
}

func tailFollowMove(xH, yH, xT, yT, xM, yM int) (xh, yh, xt, yt int) {
	xh = xH + xM
	yh = yH + yM
	xDif := xh - xT
	yDif := yh - yT
	switch {
	case xDif == -1 && yDif == 2 || xDif == 0 && yDif == 2 || xDif == 1 && yDif == 2:
		xt, yt = xh, yh-1
	case xDif == -2 && yDif == 1 || xDif == -2 && yDif == 0 || xDif == -2 && yDif == -1:
		xt, yt = xh+1, yh
	case xDif == -1 && yDif == -2 || xDif == 0 && yDif == -2 || xDif == 1 && yDif == -2:
		xt, yt = xh, yh+1
	case xDif == 2 && yDif == 1 || xDif == 2 && yDif == 0 || xDif == 2 && yDif == -1:
		xt, yt = xh-1, yh
	case xDif == -2 && yDif == 2:
		xt, yt = xh+1, yh-1
	case xDif == -2 && yDif == -2:
		xt, yt = xh+1, yh+1
	case xDif == 2 && yDif == -2:
		xt, yt = xh-1, yh+1
	case xDif == 2 && yDif == 2:
		xt, yt = xh-1, yh-1
	default:
		xt, yt = xT, yT
	}
	return
}

func CountVisitedByTheTailAtLeastOnce(d string) int {
	moves := dataToMoves(d)
	visited := make(map[string]struct{})
	xh, yh, xt, yt := 0, 0, 0, 0
	visited[hash(xt, yt)] = struct{}{}
	for _, m := range moves {
		for i := 0; i < m.steps; i++ {
			xM, yM := moveToCoords(m.dir)
			xh, yh, xt, yt = tailFollowMove(xh, yh, xt, yt, xM, yM)

			h := hash(xt, yt)
			if _, ok := visited[h]; !ok {
				visited[h] = struct{}{}
			}

		}
	}
	return len(visited)
}

type coords struct {
	x, y int
}

func CountVisitedByTheTailAtLeastOnceMultiKnots(d string, knotsNum int) int {
	moves := dataToMoves(d)
	visited := make(map[string]struct{})
	knots := make([]coords, knotsNum)

	for _, m := range moves {
		for i := 0; i < m.steps; i++ {
			var xM, yM int
			for k := 0; k < len(knots)-1; k++ {
				if k == 0 {
					xM, yM = moveToCoords(m.dir)
				}

				knots[k].x, knots[k].y, knots[k+1].x, knots[k+1].y = tailFollowMove(knots[k].x, knots[k].y, knots[k+1].x, knots[k+1].y, xM, yM)

				xM, yM = 0, 0
				if k == len(knots)-2 {
					h := hash(knots[k+1].x, knots[k+1].y)
					if _, ok := visited[h]; !ok {
						visited[h] = struct{}{}
					}
				}

			}

		}
	}
	return len(visited)
}
