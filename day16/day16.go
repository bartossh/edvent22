package day16

// This solution is using: https://en.wikipedia.org/wiki/Breadth-first_search

import (
	"fmt"
	"strings"
)

type valve struct {
	name        string
	rate        int
	connections []string
	paths       map[string]int
}

type valves map[string]*valve

type node struct {
	node string
	time int
	flow int
	open map[string]int
}

func CalcMaxPressureToRelease(d string, minutes int) int {
	rows := strings.Split(d, "\n")

	vls := make(valves, len(rows)-1)

	for _, row := range rows {
		if row == "" {
			continue
		}

		n := valve{}
		_, err := fmt.Sscanf(row, "Valve %s has flow rate=%d;", &n.name, &n.rate)
		if err != nil {
			panic(err)
		}

		split := "valve"
		if strings.Contains(row, "valves") {
			split = "valves"
		}
		arr := strings.Split(row, split)
		if len(arr) != 2 {
			panic("arr has wrong length")
		}

		conn := strings.Split(arr[1], ",")
		for _, c := range conn {
			cc := strings.ReplaceAll(c, " ", "")
			if cc == "" {
				continue
			}
			n.connections = append(n.connections, cc)
		}

		n.paths = make(map[string]int)

		vls[n.name] = &n
	}

	for _, vl := range vls {
		bfs(vls, vl)
		fmt.Printf("%v\n", vl)
	}

	return solver(vls, minutes)
}

func CalcMaxPressureToReleaseWithElephant(d string, minutes int) int {
	rows := strings.Split(d, "\n")

	vls := make(valves, len(rows)-1)

	for _, row := range rows {
		if row == "" {
			continue
		}

		n := valve{}
		_, err := fmt.Sscanf(row, "Valve %s has flow rate=%d;", &n.name, &n.rate)
		if err != nil {
			panic(err)
		}

		split := "valve"
		if strings.Contains(row, "valves") {
			split = "valves"
		}
		arr := strings.Split(row, split)
		if len(arr) != 2 {
			panic("arr has wrong length")
		}

		conn := strings.Split(arr[1], ",")
		for _, c := range conn {
			cc := strings.ReplaceAll(c, " ", "")
			if cc == "" {
				continue
			}
			n.connections = append(n.connections, cc)
		}

		n.paths = make(map[string]int)

		vls[n.name] = &n
	}

	for _, vl := range vls {
		bfs(vls, vl)
		fmt.Printf("%v\n", vl)
	}

	return solver(vls, minutes)
}

func bfs(vls valves, v *valve) {
	explored := make(map[string]struct{})
	q := make([]*valve, 0)
	explored[v.name] = struct{}{}
	q = append(q, v)

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, conn := range vls[cur.name].connections {
			if _, ok := explored[conn]; !ok {
				explored[conn] = struct{}{}
				connVL := v.paths[cur.name]
				v.paths[conn] = connVL + 1
				q = append(q, vls[conn])
			}
		}
	}
}

func solver(vls valves, minutes int) int {
	f := make([]node, 0)
	q := make([]node, 0)

	r := node{
		node: "AA",
		time: minutes,
		flow: 0,
		open: make(map[string]int),
	}
	q = append(q, r)

	for len(q) > 0 {
		var cur node
		cur, q = shift(q)
		if cur.time <= 0 {
			f = append(f, cur)
			continue
		}

		opts := make([]string, 0)

		for _, val := range vls {
			if _, ok := cur.open[val.name]; !ok && val.rate > 0 {
				opts = append(opts, val.name)
			}
		}

		if len(opts) == 0 {
			open := make(map[string]int, len(cur.open))
			copySet(open, cur.open)
			f = append(f, node{
				node: cur.node,
				time: 0,
				flow: cur.flow + cur.time*addFlow(vls, cur.open),
				open: cur.open,
			})
		}

		for _, val := range opts {
			stp := vls[cur.node].paths[val] + 1 // move and open
			if cur.time-stp < 0 {
				f = append(f, node{
					node: cur.node,
					time: 0,
					flow: cur.flow + cur.time*addFlow(vls, cur.open),
					open: cur.open,
				})
				continue
			}
			open := make(map[string]int, len(cur.open)+1)
			copySet(open, cur.open)
			open[val] = cur.time - stp
			q = append(q, node{
				node: val,
				time: cur.time - stp,
				flow: cur.flow + stp*addFlow(vls, cur.open),
				open: open,
			})
		}

	}

	var max int

	for _, pth := range f {
		if pth.flow > max {
			max = pth.flow
		}
	}

	return max
}

func copySet(dst, src map[string]int) {
	for k, v := range src {
		dst[k] = v
	}
}

func shift(arr []node) (node, []node) {
	return arr[0], arr[1:]
}

func addFlow(vls valves, open map[string]int) int {
	var sum int
	for k := range open {
		sum += vls[k].rate
	}
	return sum
}
