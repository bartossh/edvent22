package day15

import (
	"sort"
	"strconv"
	"strings"
)

type point2d struct{ x, y int }
type pair struct{ sensor, beacon point2d }
type span struct{ start, end int }
type byStart []span

func (a byStart) Len() int           { return len(a) }
func (a byStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byStart) Less(i, j int) bool { return a[i].start < a[j].start }

func CountPositionWhereBeaconNotPresent(d string, y int) int {
	sensors := dataToPositions(d)

	spans, count := spanY(sensors, y), 0
	for _, s := range spans {
		count += s.end - s.start
	}

	return count
}

func CountPositionWhereBeaconTuningFrequency(d string, yd int) int {
	sensors := dataToPositions(d)

	var x, y int
	for y = 0; y <= yd; y++ {
		spans := spanY(sensors, y)
		if len(spans) > 1 {
			x = spans[0].end + 1
			break
		}
	}
	return x*yd + y
}

func spanY(sensors []pair, y int) (spans []span) {
	for _, s := range sensors {
		sbDist, srDist := abs(s.sensor.x-s.beacon.x)+abs(s.sensor.y-s.beacon.y), abs(s.sensor.y-y)
		dist := sbDist - srDist
		if dist > 0 {
			spans = append(spans, span{s.sensor.x - dist, s.sensor.x + dist})
		}
	}
	return merge(spans)
}

func merge(spans []span) (res []span) {
	sort.Sort(byStart(spans))
	res = append(res, spans[0])
	for _, next := range spans[1:] {
		prev := res[len(res)-1]
		if prev.end >= next.start-1 {
			res[len(res)-1] = span{prev.start, max(prev.end, next.end)}
		} else {
			res = append(res, next)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func dataToPositions(d string) (sensorsWithBeacons []pair) {
	rows := strings.Split(d, "\n")

	for _, row := range rows {
		if row == "" {
			continue
		}
		pack := strings.Split(row, ":")
		if len(pack) != 2 {
			panic("expecting sensor and beacon parts")
		}
		s := readPoint(pack[0])
		b := readPoint(pack[1])

		sensorsWithBeacons = append(sensorsWithBeacons, pair{s, b})
	}

	return
}

func readPoint(d string) point2d {
	pack := strings.Split(d, ",")
	if len(pack) != 2 {
		panic("expecting two coords in text")
	}

	numX := strings.Split(pack[0], "x=")[1]
	numY := strings.Split(pack[1], "y=")[1]

	x, err := strconv.Atoi(numX)
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(numY)
	if err != nil {
		panic(err)
	}

	return point2d{
		x: x,
		y: y,
	}
}
