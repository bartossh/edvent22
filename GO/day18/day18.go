package day18

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type point struct{ x, y, z int }

func CalcSurfaceArea(d string) int {
	rows := strings.Split(d, "\n")
	pts := make([]point, 0)

	for _, r := range rows {
		if r == "" {
			continue
		}

		var x, y, z int

		fmt.Sscanf(r, "%d,%d,%d", &x, &y, &z)

		pts = append(pts, point{x, y, z})
	}

	return len(pts)*6 - calcCoveredFaces(pts)

}

func calcCoveredFaces(pts []point) int {
	ptsC := make([]point, len(pts))
	copy(ptsC, pts)
	var covered int

	for len(ptsC) > 1 {
		cur := ptsC[0]
		ptsC = ptsC[1:]

		for _, next := range ptsC {
			one := cur.x == next.x && cur.y == next.y && (cur.z == next.z+1 || cur.z == next.z-1)
			two := cur.z == next.z && cur.y == next.y && (cur.x == next.x+1 || cur.x == next.x-1)
			three := cur.x == next.x && cur.z == next.z && (cur.y == next.y+1 || cur.y == next.y-1)
			if one || two || three {
				covered += 2
			}
		}

	}
	return covered
}

func CalcTouching(d string) int {
	rows := strings.Split(d, "\n")

	lines := make([]string, 0)
	for _, row := range rows {
		if row == "" {
			continue
		}
		lines = append(lines, row)
	}
	return bytesApproach(lines)
}

func calcWhenTrapped(pts []point) int {
	fmt.Printf("all points: %v\n", len(pts))

	ext := make(map[string]point)
	all := make(map[string]point)

	polygonXY := make([][]float64, 0, len(pts))
	polygonXZ := make([][]float64, 0, len(pts))
	polygonYZ := make([][]float64, 0, len(pts))

	for _, p := range pts {
		polygonXY = append(polygonXY, []float64{float64(p.x), float64(p.y)})
		polygonXZ = append(polygonXY, []float64{float64(p.x), float64(p.z)})
		polygonYZ = append(polygonXY, []float64{float64(p.y), float64(p.z)})
	}

	holderXY := Polygon{polygonXY}
	holderXZ := Polygon{polygonXZ}
	holderYZ := Polygon{polygonYZ}

	var countFaces int

	for _, p := range pts {
		pxy := []float64{float64(p.x), float64(p.y)}
		pxz := []float64{float64(p.x), float64(p.z)}
		pyz := []float64{float64(p.y), float64(p.z)}
		one, _ := holderXY.Within(Polygon{[][]float64{pxy}})
		two, _ := holderXZ.Within(Polygon{[][]float64{pxz}})
		three, _ := holderYZ.Within(Polygon{[][]float64{pyz}})

		if !one || !two || !three {
			ext[hash(p.x, p.y, p.z)] = point{p.x, p.y, p.z}
		}
		if !one && !two && !three {
			countFaces += 6
		}
		all[hash(p.x, p.y, p.z)] = point{p.x, p.y, p.z}

	}

	fmt.Printf("number of ext: %v\n", len(ext))

	for _, v := range ext {
		yp := hash(v.x, v.y-1, v.z)
		ym := hash(v.x, v.y+1, v.z)
		xp := hash(v.x-1, v.y, v.z)
		xm := hash(v.x+1, v.y, v.z)
		zp := hash(v.x, v.y, v.z-1)
		zm := hash(v.x, v.y, v.z+1)

		counter := 6
		for _, h := range []string{yp, ym, xp, xm, zp, zm} {

			if _, ok := all[h]; ok {
				counter--
			}
		}
		countFaces += counter

	}

	return countFaces
}

func hash(x, y, z int) string {
	return fmt.Sprintf("%v:%v:%v", x, y, z)
}

type orientation int8

const (
	collinear        orientation = iota
	clockwise        orientation = iota
	counterclockwise orientation = iota
)

// Point is 2d point representing edge of polygon or end of a line
// order is {lon, lat}
type Point []float64

// Line is 2d line representing side of polygon or end of a line
type Line [][]float64

// Polygon is a multi vertices object
type Polygon [][][]float64

func (p Polygon) ToPrimitive() [][][]float64 {
	return p
}

// Within verifies if given q Polygon is within the Polygon.
// Will not work for Arctic and Antarctic as it makes some simplifications.
// Intersection is calculated in latitude direction,
// assuming that extreme value of latitude equals always 90 deg.
func (p Polygon) Within(q Polygon) (bool, error) {
	if len(p) == 0 || len(q) == 0 {
		return false, errors.New("no polygons present")
	}
	if len(p[0]) < 3 {
		return false, errors.New("receiver struct is not a polygon")
	}
	if len(q[0]) == 0 {
		return false, errors.New("given polygon is empty")
	}

	intersections := 0
	for _, qPoint := range q[0] {
		if len(qPoint) != 2 {
			return false, errors.New("point should have to coordinates")
		}
		polygon := p[0]
		for i := 0; i < len(polygon)-1; i++ {
			j := i + 1
			l1 := Line{polygon[i], polygon[j]}
			l2 := Line{qPoint, {qPoint[0], 90}}
			ok, err := doIntersect(l1, l2)
			if err != nil {
				return false, err
			}
			if ok {
				intersections++
			}
		}
	}
	if intersections%2 != 0 {
		return true, nil
	}

	return false, nil
}

func onLine(p, q, r Point) bool {
	if len(p) != 2 || len(q) != 2 || len(r) != 2 {
		return false
	}

	distLon := r[0] - p[0]
	distLat := r[1] - p[1]

	dLon := distLon / (q[0] - p[0])
	dLat := distLat / (q[1] - p[1])

	absDistLon := distLon
	if distLon < 0 {
		absDistLon = -distLon
	}
	absDistLat := distLat
	if distLat < 0 {
		absDistLat = -distLat
	}

	isOn := dLon == dLat

	betweenLon := 0.0 <= dLon && dLon <= absDistLon
	betweenLat := 0.0 <= dLat && dLat <= absDistLat

	return isOn && betweenLon && betweenLat
}

func calculateOrientation(p, q, r Point) (orientation, error) {
	if len(p) != 2 || len(q) != 2 || len(r) != 2 {
		return collinear, errors.New("line should have two points, longitude and latitude")
	}
	v := (q[1]-p[1])*(r[0]-q[0]) - (q[0]-p[0])*(r[1]-q[1])
	if v == 0 {
		return collinear, nil
	}
	if v > 0 {
		return clockwise, nil
	}
	return counterclockwise, nil
}

func doIntersect(l1, l2 [][]float64) (bool, error) {
	if len(l1) != 2 || len(l2) != 2 {
		return false, errors.New("line should have two points, longitude and latitude")
	}

	o1, err := calculateOrientation(l1[0], l1[1], l2[0])
	if err != nil {
		return false, err
	}
	o2, err := calculateOrientation(l1[0], l1[1], l2[1])
	if err != nil {
		return false, err
	}
	o3, err := calculateOrientation(l2[0], l2[1], l1[0])
	if err != nil {
		return false, err
	}
	o4, err := calculateOrientation(l2[0], l2[1], l1[1])
	if err != nil {
		return false, err
	}

	if o1 != o2 && o3 != o4 {
		return true, nil
	}

	if o1 == 0 && onLine(l1[0], l2[0], l1[1]) {
		return true, nil
	}
	if o2 == 0 && onLine(l1[0], l2[1], l1[1]) {
		return true, nil
	}
	if o3 == 0 && onLine(l2[0], l1[0], l2[1]) {
		return true, nil
	}
	if o4 == 0 && onLine(l2[0], l1[1], l2[1]) {
		return true, nil
	}

	return false, nil
}

// we modelize the 3D grid space, as a linear array of positions,
// a position p is: x + y*size + z*area
var size, area, volume int // dims of world: cube enclosing the lava droplet
var cube []bool            // is there a rock bit at position?
var bits []int             // the positions of rock bits in the droplet
var adjacent [6]int        // the relative position offsets of the 6 adjacent bits
var vapor []bool           // part2: is there vapor at position?

var verbose bool

func bytesApproach(lines []string) (outside int) {
	parse(lines)
	vapor = make([]bool, volume, volume)  // is filled with vapor?
	for z := 0; z < size; z += size - 1 { // top & bottom sides
		for x := 0; x < size; x++ {
			for y := 0; y < size; y++ {
				injectVapor(x + y*size + z*area)
			}
		}
	}
	for z := 1; z < size-1; z++ { // middle slices: perimeter x,y
		for x := 0; x < size; x += size - 1 { // x edges
			for y := 0; y < size; y++ {
				injectVapor(x + y*size + z*area)
			}
		}
		for x := 1; x < size-1; x++ { // y edges
			for y := 0; y < size; y += size {
				injectVapor(x + y*size + z*area)
			}
		}
	}

	for _, bit := range bits {
		for _, adj := range adjacent {
			p := bit + adj
			if p < 0 || p >= volume || vapor[p] {
				outside++
			}
		}
	}
	return
}

func parse(lines []string) {
	var x, y, z, max int
	bits = make([]int, 0)
	for lineno, line := range lines {
		_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if err != nil {
			log.Fatalf("Syntax error line %d: %s\n", lineno, line)
		}
		if x > max {
			max = x
		}
		if y > max {
			max = y
		}
		if z > max {
			max = z
		}
	}
	size = max + 1
	area = size * size
	volume = area * size
	cube = make([]bool, volume)
	adjacent[0] = 1
	adjacent[1] = -1
	adjacent[2] = size
	adjacent[3] = -size
	adjacent[4] = area
	adjacent[5] = -area

	for _, line := range lines {
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		cube[x+y*size+z*area] = true
		bits = append(bits, x+y*size+z*area)
	}
}

func injectVapor(bit int) {
	if !cube[bit] && !vapor[bit] {
		vapor[bit] = true              // fill it
		for _, adj := range adjacent { // expand to neighbors
			p := bit + adj
			if p >= 0 && p < volume && !cube[p] && !vapor[p] {
				injectVapor(p)
			}
		}
	}
}
