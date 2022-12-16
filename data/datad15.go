package data

import (
	"fmt"
	"strings"
)

const TestBeaconsAndSensorsMap = `
Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
`

const PuzzleBeaconsAndSensorsMap = `
Sensor at x=2765643, y=3042538: closest beacon is at x=2474133, y=3521072
Sensor at x=2745662, y=2324735: closest beacon is at x=2491341, y=1883354
Sensor at x=2015742, y=2904055: closest beacon is at x=2474133, y=3521072
Sensor at x=3375262, y=3203288: closest beacon is at x=3321219, y=3415236
Sensor at x=3276468, y=3892409: closest beacon is at x=3321219, y=3415236
Sensor at x=952573, y=3147055: closest beacon is at x=-41010, y=2905006
Sensor at x=1823659, y=1779343: closest beacon is at x=1592718, y=2000000
Sensor at x=1156328, y=865741: closest beacon is at x=1592718, y=2000000
Sensor at x=3938443, y=271482: closest beacon is at x=4081274, y=1177185
Sensor at x=2815232, y=1641178: closest beacon is at x=2491341, y=1883354
Sensor at x=3984799, y=3424711: closest beacon is at x=3321219, y=3415236
Sensor at x=1658825, y=3999931: closest beacon is at x=2474133, y=3521072
Sensor at x=3199859, y=1285962: closest beacon is at x=4081274, y=1177185
Sensor at x=3538649, y=2788193: closest beacon is at x=3725736, y=2414539
Sensor at x=3522208, y=3336284: closest beacon is at x=3321219, y=3415236
Sensor at x=3093758, y=3492396: closest beacon is at x=3321219, y=3415236
Sensor at x=2464979, y=562119: closest beacon is at x=2491341, y=1883354
Sensor at x=3665010, y=1556840: closest beacon is at x=3735739, y=2128164
Sensor at x=207525, y=3893957: closest beacon is at x=-41010, y=2905006
Sensor at x=3894678, y=1974599: closest beacon is at x=3735739, y=2128164
Sensor at x=2185146, y=3822275: closest beacon is at x=2474133, y=3521072
Sensor at x=31166, y=1467978: closest beacon is at x=-41010, y=2905006
Sensor at x=3242364, y=3335961: closest beacon is at x=3321219, y=3415236
Sensor at x=3773718, y=3999789: closest beacon is at x=3321219, y=3415236
Sensor at x=423046, y=2227938: closest beacon is at x=-41010, y=2905006
Sensor at x=1600225, y=2529059: closest beacon is at x=1592718, y=2000000
Sensor at x=3291752, y=2241389: closest beacon is at x=3735739, y=2128164
Sensor at x=2741333, y=3984346: closest beacon is at x=2474133, y=3521072
Sensor at x=3935288, y=2292902: closest beacon is at x=3725736, y=2414539
Sensor at x=291635, y=140996: closest beacon is at x=212146, y=-1154950
Sensor at x=3966296, y=2600346: closest beacon is at x=3725736, y=2414539
Sensor at x=2228916, y=1461096: closest beacon is at x=2491341, y=1883354
`

const example = `
-2 ..........#.................
-1 .........###................
 0 ....S...#####...............
 1 .......#######........S.....
 2 ......#########S............
 3 .....###########SB..........
 4 ....#############...........
 5 ...###############..........
 6 ..#################.........
 7 .#########S#######S#........
 8 ..#################.........
 9 ...###############..........
10 ....B############...........
11 ..S..###########............
12 ......#########.............
13 .......#######..............
14 ........#####.S.......S.....
15 B........###................
16 ..........#SB...............
17 ................S..........B
18 ....S.......................
19 ............................
20 ............S......S........
21 ............................
22 .......................B....
`

type xy struct{ x, y int }

type sensor struct{ loc, beacon xy }

func readInput(input string) (res []sensor) {
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		var s sensor
		fmt.Sscanf(l, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&s.loc.x, &s.loc.y, &s.beacon.x, &s.beacon.y)
		res = append(res, s)
	}
	return res
}
