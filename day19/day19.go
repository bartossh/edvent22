package day19

import (
	"fmt"
	"runtime"
	"strings"

	"golang.org/x/exp/constraints"
)

type node struct {
	t, ores, clays, obsidians, oreRobots, clayRobots, obsidianRobots, geodeRobots int
}

type obsidianCost struct {
	oreCost  int
	clayCost int
}

type geodeCost struct {
	oreCost      int
	obsidianCost int
}

type blueprint struct {
	id           int
	oreCost      int
	clayCost     int
	obsidianCost obsidianCost
	geodeCost    geodeCost
}

func CalcBestStrategy(d string, m int) int {
	rows := strings.Split(d, "\n")

	blueprints := make([]blueprint, 0)

	for _, row := range rows {
		if row == "" {
			continue
		}
		var id, oreRobotCostOre, clayRobotCostOre, obsidianRobotCostOre, obsidianRobotCostClay, geodeRobotCostsOre, geodeRobotCostsObsidian int
		fmt.Sscanf(row, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &id, &oreRobotCostOre, &clayRobotCostOre, &obsidianRobotCostOre, &obsidianRobotCostClay, &geodeRobotCostsOre, &geodeRobotCostsObsidian)

		blueprints = append(blueprints, blueprint{
			id:       id,
			oreCost:  oreRobotCostOre,
			clayCost: clayRobotCostOre,
			obsidianCost: obsidianCost{
				oreCost:  obsidianRobotCostOre,
				clayCost: obsidianRobotCostClay,
			},
			geodeCost: geodeCost{
				oreCost:      geodeRobotCostsOre,
				obsidianCost: geodeRobotCostsObsidian,
			},
		})

	}

	return execute(blueprints, m)
}

func CalcMaxGeodesOpened(d string, m int) int {
	rows := strings.Split(d, "\n")

	blueprints := make([]blueprint, 0)

	for _, row := range rows {
		if row == "" {
			continue
		}
		var id, oreRobotCostOre, clayRobotCostOre, obsidianRobotCostOre, obsidianRobotCostClay, geodeRobotCostsOre, geodeRobotCostsObsidian int
		fmt.Sscanf(row, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &id, &oreRobotCostOre, &clayRobotCostOre, &obsidianRobotCostOre, &obsidianRobotCostClay, &geodeRobotCostsOre, &geodeRobotCostsObsidian)

		blueprints = append(blueprints, blueprint{
			id:       id,
			oreCost:  oreRobotCostOre,
			clayCost: clayRobotCostOre,
			obsidianCost: obsidianCost{
				oreCost:  obsidianRobotCostOre,
				clayCost: obsidianRobotCostClay,
			},
			geodeCost: geodeCost{
				oreCost:      geodeRobotCostsOre,
				obsidianCost: geodeRobotCostsObsidian,
			},
		})

	}

	return execute2(blueprints, m)
}

func execute2(blps []blueprint, m int) int {

	multi := 1

	for i, blp := range blps {
		if i == 3 {
			break
		}
		nodes = make(map[node]int)
		runtime.GC()
		multi *= update(blp, node{m, 0, 0, 0, 1, 0, 0, 0}) * blp.id
	}

	return multi
}

func execute(blps []blueprint, m int) int {

	sum := 0

	for _, blp := range blps {
		nodes = make(map[node]int)
		runtime.GC()
		sum += update(blp, node{m, 0, 0, 0, 1, 0, 0, 0}) * blp.id
	}

	return sum
}

var nodes map[node]int
var hits = 0

func update(b blueprint, n node) int {
	if val, ok := nodes[n]; ok {
		hits++
		return val
	}
	if n.t == 0 {
		return 0
	}

	// Option don't make any robots
	best := n.geodeRobots

	switch {
	case n.ores >= b.geodeCost.oreCost && n.obsidians >= b.geodeCost.obsidianCost:
		// Option: make geode robot
		best = maxF(best, n.geodeRobots+update(b, node{
			n.t - 1,
			n.ores + n.oreRobots - b.geodeCost.oreCost, n.clays + n.clayRobots, n.obsidians + n.obsidianRobots - b.geodeCost.obsidianCost,
			n.oreRobots, n.clayRobots, n.obsidianRobots, n.geodeRobots + 1,
		}))
	default:
		if n.ores >= b.obsidianCost.oreCost && n.clays >= b.obsidianCost.clayCost {
			// Option: make obsidian robot
			best = maxF(best, n.geodeRobots+update(b, node{
				n.t - 1,
				n.ores + n.oreRobots - b.obsidianCost.oreCost, n.clays + n.clayRobots - b.obsidianCost.clayCost, n.obsidians + n.obsidianRobots,
				n.oreRobots, n.clayRobots, n.obsidianRobots + 1, n.geodeRobots,
			}))
		}

		if n.ores >= b.clayCost {
			// Option: make clay robot
			best = maxF(best, n.geodeRobots+update(b, node{
				n.t - 1,
				n.ores + n.oreRobots - b.clayCost, n.clays + n.clayRobots, n.obsidians + n.obsidianRobots,
				n.oreRobots, n.clayRobots + 1, n.obsidianRobots, n.geodeRobots,
			}))
		}

		// Option: make ore robot
		if n.ores >= b.oreCost {
			best = maxF(best, n.geodeRobots+update(b, node{
				n.t - 1,
				n.ores + n.oreRobots - b.oreCost, n.clays + n.clayRobots, n.obsidians + n.obsidianRobots,
				n.oreRobots + 1, n.clayRobots, n.obsidianRobots, n.geodeRobots,
			}))
		}

		maxOres := maxFList([]int{b.oreCost, b.clayCost, b.obsidianCost.oreCost, b.geodeCost.oreCost})
		if n.ores < maxOres || n.clays < b.clayCost || n.obsidians < b.geodeCost.obsidianCost {
			best = maxF(best, n.geodeRobots+update(b, node{
				n.t - 1,
				n.ores + n.oreRobots, n.clays + n.clayRobots, n.obsidians + n.obsidianRobots,
				n.oreRobots, n.clayRobots, n.obsidianRobots, n.geodeRobots,
			}))
		}
	}

	nodes[n] = best
	return best

}

func maxF[T constraints.Ordered](a, b T) T {
	if a >= b {
		return a
	} else {
		return b
	}
}

func minF[T constraints.Ordered](a, b T) T {
	if a <= b {
		return a
	} else {
		return b
	}
}

func maxFList[T constraints.Ordered](l []T) (max T) {
	_, max = minMaxListFn(l, func(x T) T { return x })
	return max
}

func minMaxListFn[T any, U constraints.Ordered](l []T, f func(T) U) (min, max U) {
	if len(l) == 0 {
		panic("cannot find min/max of empty list")
	}
	min = f(l[0])
	max = min
	for _, val := range l {
		v := f(val)
		min = minF(min, v)
		max = maxF(max, v)
	}
	return
}
