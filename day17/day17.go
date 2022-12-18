package day17

import (
	"fmt"
	"strings"
)

const (
	A = `
####
`

	B = `
.#.
###
.#.
`
	C = `
..#
..#
###
`

	D = `
#
#
#
#
`

	E = `
##
##
`
)
const (
	startX           = 2
	startY           = 4
	chamberRightWall = 7
)

type point2D struct {
	x, y int
}

func CalcHighOfTetrisTowel(d string, rocksToFall int) int {
	jetFlowCounter := 0
	rested := createRested()
	objects := createObjects()
	var maxY int
	var firstRecorded bool
	var bottomTopDist int
	var bottom int
	total := rocksToFall
	curFigure := 1

	for rocksToFall > 0 {

		var object []point2D

		object, objects = getNextObject(objects)
		startingPosition(maxY, object)

	Fall:
		for {
			var m rune
			m, jetFlowCounter = nextFlow(d, jetFlowCounter)

			switch {
			case m == '<':
				moveLeft(rested, object)
			case m == '>':
				moveRight(rested, object)
			}

			objectCopy := make([]point2D, len(object))
			copy(objectCopy, object)

			moveDown(object)
			if collision(rested, object) {
				rest(rested, objectCopy)
				maxY = hight(rested, maxY)
				if curFigure%len(objects) == 0 {
					switch {
					case !firstRecorded:
						firstRecorded = true
						bottomTopDist = maxY
					default:
						if maxY-bottom == bottomTopDist {
							// a :=(total - rocksToFall) * bottomTopDist
							if total%curFigure == 0 {
								fmt.Printf("AAAAAA %v: maxY: %v \n", curFigure, maxY)

							}
							// fmt.Printf("repetition: %v\n", curFigure)
							// return 0
						}
						bottom = maxY
					}

				}
				break Fall
			}

		}
		curFigure++

		rocksToFall--
	}
	// drawPattern(rested)
	return hight(rested, maxY)
}

func drawPattern(rest map[string]point2D) {
	y := hight(rest, 0)

	cave := make([][]string, y+1)

	for iy := range cave {
		xs := make([]string, chamberRightWall)
		cave[iy] = xs
	}

	for iy := range cave {
		for ix := range cave[iy] {
			h := hash(ix, iy)
			if _, ok := rest[h]; ok {
				cave[iy][ix] = "#"
				continue
			}
			cave[iy][ix] = "."
		}
	}

	fmt.Println()
	for iy := len(cave) - 1; iy >= 0; iy-- {
		fmt.Printf("%v\n", cave[iy])
	}
	fmt.Println()
}

func nextFlow(d string, idx int) (rune, int) {
	if idx == len(d) {
		idx = 0
	}
	r := rune(d[idx])
	idx++
	return r, idx
}

func createObjects() [][]point2D {

	objects := make([][]point2D, 0, 5)
	for _, elem := range []string{A, B, C, D, E} {
		obj := shapeToCoords(elem)
		objects = append(objects, obj)
	}

	return objects
}

func getNextObject(objects [][]point2D) ([]point2D, [][]point2D) {
	p := objects[0]
	objects = append(objects[1:], p)
	cp := make([]point2D, len(p))
	copy(cp, p)
	return cp, objects
}

func shapeToCoords(d string) []point2D {
	rows := strings.Split(d, "\n")

	for i, j := 0, len(rows)-1; i < j; i, j = i+1, j-1 {
		rows[i], rows[j] = rows[j], rows[i]
	}

	cleaned := make([]string, 0, len(rows)-1)

	for _, row := range rows {
		if row == "" {
			continue
		}
		cleaned = append(cleaned, row)
	}

	coords := make([]point2D, 0)

	for y := range cleaned {
		for x := range cleaned[y] {
			if cleaned[y][x] == '#' {
				coords = append(coords, point2D{x, y})
			}
		}
	}

	return coords
}

func moveLeft(rested map[string]point2D, coords []point2D) {
	for _, coord := range coords {
		if coord.x == 0 {
			return
		}
		h := hash(coord.x-1, coord.y)
		if _, ok := rested[h]; ok {
			return
		}
	}
	for i := range coords {
		coords[i].x--
	}
}

func moveRight(rested map[string]point2D, coords []point2D) {
	for _, coord := range coords {
		if coord.x == chamberRightWall-1 {
			return
		}
		h := hash(coord.x+1, coord.y)
		if _, ok := rested[h]; ok {
			return
		}
	}
	for i := range coords {
		coords[i].x++
	}
}

func moveDown(coords []point2D) {
	for i := range coords {
		coords[i].y--
	}
}

func rest(rested map[string]point2D, object []point2D) {
	for _, coord := range object {
		rested[hash(coord.x, coord.y)] = coord
	}
}

func hight(rested map[string]point2D, last int) int {

	max := 0

	if last != 0 {
		max = last - 1
	}

Last:
	for {

		for x := 0; x < chamberRightWall; x++ {
			h := hash(x, max)
			if _, ok := rested[h]; ok {
				max++
				continue Last
			}
		}
		break Last
	}

	return max - 1
}

func hash(x, y int) string {
	return fmt.Sprintf("%v:%v", x, y)
}

func collision(rested map[string]point2D, object []point2D) bool {
	for _, coord := range object {
		h := hash(coord.x, coord.y)
		if _, ok := rested[h]; ok {
			return true
		}
	}
	return false
}

func objectMaxY(object []point2D) int {
	var max int
	for _, coord := range object {
		if coord.y > max {
			max = coord.y
		}
	}
	return max
}

func startingPosition(y int, object []point2D) {
	for i := range object {
		object[i].x += startX
		object[i].y += y + startY
	}
}

func createRested() map[string]point2D {
	rested := make(map[string]point2D, chamberRightWall)
	for i := 0; i < chamberRightWall+1; i++ {
		rested[hash(i, 0)] = point2D{i, 0}
	}
	return rested
}
