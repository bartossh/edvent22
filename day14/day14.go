package day14

import (
	"fmt"
	"strconv"
	"strings"
)

type point2D struct {
	x, y int
}

func CalcSandDropComingToRest(d string, start point2D) int {
	cave := createCaveSystem(d, point2D{})
	pos := start
	counter := 0

	for {
		if pos.y == len(cave)-1 { // last line is a void space that a sand will start falling to infinity
			for i := range cave {
				fmt.Printf("%v\n", string(cave[i][490:len(cave[i])-499]))
			}
			break
		}
		var move bool
		pos, move = drop(pos, cave)
		if !move {
			cave[pos.y][pos.x] = 'o'
			pos = start
			counter++
			// for i := range cave {
			// 	fmt.Printf("%v\n", string(cave[i][490:]))
			// }
			// time.Sleep(time.Millisecond * 300)
			continue
		}
	}
	return counter
}

func CalcSandDropComingToReachDropPoint(d string, start point2D) int {
	cave := createCaveSystem(d, point2D{})
	pos := start
	counter := 0
	seg := make([]rune, len(cave[0]))
	cave = append(cave, seg)
	for i := range cave[len(cave)-1] {
		cave[len(cave)-1][i] = '#'
	}

	for {
		var move bool
		pos, move = drop(pos, cave)
		if pos.y == start.y && pos.x == start.x { // last line is a void space that a sand will start falling to infinity
			for i := range cave {
				fmt.Printf("%v\n", string(cave[i][490:len(cave[i])-499]))
			}
			counter++
			break
		}
		if !move {
			cave[pos.y][pos.x] = 'o'
			pos = start
			counter++
			// for i := range cave {
			// 	fmt.Printf("%v\n", string(cave[i][490:]))
			// }
			// time.Sleep(time.Millisecond * 300)
			continue
		}
	}
	return counter
}

func createCaveSystem(d string, start point2D) (cave [][]rune) {
	rows := strings.Split(d, "\n")

	var max point2D

	segments := make([][]point2D, 0, len(rows))

	for _, r := range rows {
		if r == "" {
			continue
		}
		points := strings.Split(r, "->")
		segment := make([]point2D, 0, len(points))
		for _, point := range points {
			cleaned := strings.ReplaceAll(point, " ", "")
			coords := strings.Split(cleaned, ",")

			if len(coords) != 2 {
				panic("point has wrong set of coordinates")
			}

			x, err := strconv.Atoi(coords[0])
			if err != nil {
				panic("cannot init variable")
			}
			y, err := strconv.Atoi(coords[1])
			if err != nil {
				panic("cannot init variable")
			}

			if x > max.x {
				max.x = x
			}

			if y > max.y {
				max.y = y
			}

			segment = append(segment, point2D{x: x, y: y})
		}
		segments = append(segments, segment)
	}

	if start.x > max.x {
		max.x = start.x
	}
	cave = make([][]rune, max.y+2) // make space for drop outside of a cave grid
	for i := range cave {
		cave[i] = make([]rune, max.x+500) // make space for move sideways
		for j := range cave[i] {
			cave[i][j] = '.'
		}
	}

	for _, segment := range segments {
		for i := 0; i < len(segment)-1; i++ {
			draw(segment[i], segment[i+1], cave)
		}
	}

	for i := range cave {
		fmt.Printf("%v\n", string(cave[i][490:len(cave[i])-499])) // just show the part with obstacles
	}

	return
}

func draw(from, to point2D, on [][]rune) {
	if from.x > to.x {
		from.x, to.x = to.x, from.x
	}

	if from.y > to.y {
		from.y, to.y = to.y, from.y
	}

	switch {
	case from.x == to.x:
		for y := range on {
			if y >= from.y && y <= to.y {
				on[y][from.x] = '#'
			}
		}
	case from.y == to.y:
		for x := range on[from.y] {
			if x >= from.x && x <= to.x {
				on[from.y][x] = '#'
			}
		}
	}
}

func drop(pos point2D, cave [][]rune) (point2D, bool) {
	next := pos
	next.y += 1
	switch {
	case cave[next.y][next.x] == '.':
		return next, true
	case cave[next.y][next.x-1] == '.':
		next.x--
		return next, true
	case cave[next.y][next.x+1] == '.':
		next.x++
		return next, true
	default:
		return pos, false
	}
}
