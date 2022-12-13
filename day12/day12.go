package day12

import (
	"math"
	"sort"
	"strings"
)

type point struct {
	x, y int
}

func CalcDistance(d string) int {
	rowsNotCleaned := strings.Split(d, "\n")
	rows := make([]string, 0)
	for _, r := range rowsNotCleaned {
		if r == "" {
			continue
		}
		rows = append(rows, r)
	}

	arr := make([][]rune, 0)
	var start, end point

	for _, l := range rows {
		var line []rune
		for i, elevation := range l {
			if elevation == 'S' {
				start = point{i, len(arr)}
				elevation = 'a'
			}
			if elevation == 'E' {
				end = point{i, len(arr)}
				elevation = 'z'
			}
			line = append(line, elevation)
		}
		arr = append(arr, line)
	}

	visited := make(map[point]bool)
	toGo := []point{start}
	traveled := map[point]int{start: 0}

	for {
		currentPoint := toGo[0]
		visited[currentPoint] = true
		toGo = toGo[1:]

		if currentPoint == end {
			return traveled[end]
		}

		for _, neighbor := range [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}} {
			j, i := neighbor[1], neighbor[0]
			nextPoint := point{currentPoint.x + j, currentPoint.y + i}
			if !visited[nextPoint] && nextPoint.x >= 0 && nextPoint.y >= 0 &&
				nextPoint.x < len(arr[0]) && nextPoint.y < len(arr) &&
				(arr[nextPoint.y][nextPoint.x]-arr[currentPoint.y][currentPoint.x] <= 1) {
				if traveled[nextPoint] == 0 {
					toGo = append(toGo, nextPoint)
					traveled[nextPoint] = traveled[currentPoint] + 1
				}
				if traveled[nextPoint] >= traveled[currentPoint]+1 {
					traveled[nextPoint] = traveled[currentPoint] + 1
				}
			}
		}
		sort.Slice(toGo, func(i, j int) bool {
			return traveled[toGo[i]] < traveled[toGo[j]]
		})
	}
}

func CalcDistanceAnyLowerElevationStart(d string) int {
	rowsNotCleaned := strings.Split(d, "\n")
	rows := make([]string, 0)
	for _, r := range rowsNotCleaned {
		if r == "" {
			continue
		}
		rows = append(rows, r)
	}

	arr := make([][]rune, 0)
	var end point
	starts := make([]point, 0)

	for _, l := range rows {
		var line []rune
		for i, elevation := range l {
			if elevation == 'S' {
				elevation = 'a'
			}
			if elevation == 'a' {
				starts = append(starts, point{i, len(arr)})
			}
			if elevation == 'E' {
				end = point{i, len(arr)}
				elevation = 'z'
			}
			line = append(line, elevation)
		}
		arr = append(arr, line)
	}

	shortestPath := math.MaxInt
	for _, start := range starts {
		visited := make(map[point]bool)
		toGo := []point{start}
		traveled := map[point]int{start: 0}

		for {
			if len(toGo) == 0 {
				break
			}
			currentPoint := toGo[0]
			visited[currentPoint] = true
			toGo = toGo[1:]

			if currentPoint == end {
				shortest := traveled[end]
				if shortest < shortestPath {
					shortestPath = shortest
				}
				break
			}

			for _, neighbor := range [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}} {
				j, i := neighbor[1], neighbor[0]
				nextPoint := point{currentPoint.x + j, currentPoint.y + i}
				if !visited[nextPoint] && nextPoint.x >= 0 && nextPoint.y >= 0 &&
					nextPoint.x < len(arr[0]) && nextPoint.y < len(arr) &&
					(arr[nextPoint.y][nextPoint.x]-arr[currentPoint.y][currentPoint.x] <= 1) {
					if traveled[nextPoint] == 0 {
						toGo = append(toGo, nextPoint)
						traveled[nextPoint] = traveled[currentPoint] + 1
					}
					if traveled[nextPoint] >= traveled[currentPoint]+1 {
						traveled[nextPoint] = traveled[currentPoint] + 1
					}
				}
			}
			sort.Slice(toGo, func(i, j int) bool {
				return traveled[toGo[i]] < traveled[toGo[j]]
			})
		}
	}
	return shortestPath
}
