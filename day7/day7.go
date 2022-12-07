package day7

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type node struct {
	parent   *node
	name     string
	size     int
	children []*node // if children is nil than it is a file,
}

func CalcFoldersSize(cap int, msgs string) int {
	sizes := make([]int, 0)
	n := splitTerminalMsgs(msgs)
	var rootSum int
	rootSum, sizes = n.calcTotalSizeSum(cap, sizes)
	if rootSum <= cap {
		sizes = append(sizes, rootSum)
	}
	sum := 0
	for _, size := range sizes {
		sum += size
	}
	return sum
}

func CalcSmallestSizeDirectoryOverThreshold(space, disk int, msgs string) int {
	results := make([]int, 0)
	var sum int
	n := splitTerminalMsgs(msgs)

	taken := n.totalSizeTaken()

	cap := space - (disk - taken)

	sum, results = n.calcTheSmallestToFree(cap, results)
	if sum >= cap {
		results = append(results, sum)
	}

	smallest := math.MaxInt

	if len(results) == 0 {
		return -1
	}

	for _, v := range results {
		if v < smallest {
			smallest = v
		}
	}

	return smallest
}

func splitTerminalMsgs(s string) *node {
	rows := strings.Split(s, "\n")

	var n *node
	var root *node
	var isLS bool

	for _, r := range rows {
		switch {
		case r == "":
			continue
		case strings.Contains(r, "$ cd"):
			isLS = false
			if len(r) < 6 {
				panic("unexpected length")
			}
			if strings.Contains(r, "..") {
				if n == nil || n.parent == nil {
					break
				}
				n = n.parent
				continue
			}
			d := r[5:]
			if d == "/" {
				n = &node{
					parent: nil,
					name:   d,
					size:   -1,
				}
				root = n
				continue
			}
			for _, child := range n.children {
				if child.name == d {
					n = child
					x := n
					_ = x
					continue
				}
			}
			child := &node{
				name:   d,
				size:   -1,
				parent: n,
			}
			n.children = append(n.children, child)
			n = child
			x := n
			_ = x
		case strings.Contains(r, "$ ls"):
			isLS = true
		case isLS && !strings.Contains(r, "$"):
			if strings.Contains(r, "dir") {
				if len(r) < 5 {
					panic("dir has no name")
				}
				continue
			}
			file := strings.Split(r, " ")
			if len(file) != 2 {
				panic("file has wrong prompt format")
			}
			size, err := strconv.Atoi(file[0])
			if err != nil {
				panic("unexpected variable")
			}
			n.children = append(n.children, &node{
				parent: n,
				name:   file[1],
				size:   size,
			})
		}
	}
	return root
}

func (n *node) printStructure(prefix string) {
	if n.size > 0 {
		fmt.Printf("%sFILE %s SIZE %v\n", prefix, n.name, n.size)
		return
	}
	fmt.Printf("%sFOLDER %v\n", prefix, n.name)

	for _, child := range n.children {
		child.printStructure("	" + prefix)
	}
}

func (n *node) calcTotalSizeSum(threshold int, sizes []int) (int, []int) {
	if n.size > -1 {
		return n.size, sizes
	}
	sum := 0
	for _, child := range n.children {
		var v int
		v, sizes = child.calcTotalSizeSum(threshold, sizes)
		sum += v
	}
	if sum <= threshold {
		sizes = append(sizes, sum)
	}
	return sum, sizes
}

func (n *node) calcTheSmallestToFree(threshold int, res []int) (int, []int) {
	if n.size > -1 {
		return n.size, res
	}
	sum := 0
	for _, child := range n.children {
		var v int
		v, res = child.calcTheSmallestToFree(threshold, res)
		sum += v
	}
	if sum >= threshold {
		res = append(res, sum)
	}
	return sum, res
}

func (n *node) totalSizeTaken() int {
	if n.size > -1 {
		return n.size
	}
	sum := 0

	for _, c := range n.children {
		sum += c.totalSizeTaken()
	}
	return sum
}
