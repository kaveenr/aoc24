package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
)

func Part1(input string) (result int, err error) {

	size, robots := parseInput(input)

	for seconds := 0; seconds < 100; seconds++ {
		tick(robots, size)
	}

	quadScore := make([]int, 4)
	for _, robot := range robots {
		if robot.pos.col == (size.col / 2) {
			continue
		}
		if robot.pos.row == (size.row / 2) {
			continue
		}
		var quad int
		if robot.pos.col < (size.col / 2) {
			quad = 1
		} else {
			quad = 2
		}
		if robot.pos.row > (size.row / 2) {
			quad += 2
		}
		quadScore[quad-1] += 1
	}
	return quadScore[0] * quadScore[1] * quadScore[2] * quadScore[3], nil
}

var (
	moveVectors = []vec{
		{-1, 0}, // Up
		{1, 0},  // Down
		{0, -1}, // Left
		{0, 1},  // Right
	}
)

func Part2(input string) (result int, err error) {

	size, robots := parseInput(input)

	cellIndex := func(v vec) int {
		return size.col*v.row + v.col
	}

	for seconds := 0; true; seconds++ {
		tick(robots, size)

		// Check with union find for clusters > 200
		robotMap := make(map[vec]bool)
		for _, robot := range robots {
			robotMap[robot.pos] = true
		}
		uf := newUnionFind(size.row * size.col)
		for pos := range robotMap {
			for _, move := range moveVectors {
				next := pos.add(move)
				if _, ok := robotMap[next]; !ok {
					continue
				}
				uf.union(cellIndex(pos), cellIndex(next))
			}
		}
		for _, robot := range robots {
			if uf.getSize(cellIndex(robot.pos)) > 200 {
				result = seconds + 1
				viz(size, robots, result)
				return result, err
			}
		}
	}

	return
}

func tick(robots []*entry, size vec) {
	for _, robot := range robots {
		nextLoc := robot.pos.add(robot.vel)

		if nextLoc.row < 0 {
			nextLoc.row = size.row + nextLoc.row
		} else if nextLoc.row >= size.row {
			nextLoc.row = nextLoc.row - size.row
		}

		if nextLoc.col < 0 {
			nextLoc.col = size.col + nextLoc.col
		} else if nextLoc.col >= size.col {
			nextLoc.col = nextLoc.col - size.col
		}
		robot.pos = nextLoc
	}
}

func parseInput(input string) (size vec, robots []*entry) {
	input = strings.TrimSpace(input)
	for _, row := range strings.Split(input, "\n") {
		var cur entry
		parts := strings.Split(row, " ")
		posParts := strings.Split(parts[0][2:], ",")
		vecParts := strings.Split(parts[1][2:], ",")
		cur.pos.row, _ = strconv.Atoi(posParts[1])
		cur.pos.col, _ = strconv.Atoi(posParts[0])
		cur.vel.row, _ = strconv.Atoi(vecParts[1])
		cur.vel.col, _ = strconv.Atoi(vecParts[0])
		robots = append(robots, &cur)
		if cur.pos.row > size.row {
			size.row = cur.pos.row
		}
		if cur.pos.col > size.col {
			size.col = cur.pos.col
		}
	}
	size.col += 1
	size.row += 1
	return
}

func viz(size vec, robots []*entry, second int) {
	dc := gg.NewContext(size.col, size.row)
	for _, robot := range robots {
		dc.DrawPoint(float64(robot.pos.col), float64(robot.pos.row), 1)
		dc.SetRGB(0, 1, 0)
		dc.Fill()
	}
	dc.SavePNG(fmt.Sprintf("%d.png", second))
}

type entry struct {
	pos vec
	vel vec
}

type vec struct {
	row int
	col int
}

func (v vec) add(x vec) (n vec) {
	return vec{
		row: v.row + x.row,
		col: v.col + x.col,
	}
}

// Union find from day 12 with size now!

type unionFind struct {
	parent []int
	size   []int
}

func newUnionFind(size int) (uf *unionFind) {
	uf = &unionFind{
		parent: make([]int, size),
		size:   make([]int, size),
	}
	for i := 0; i < size; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}

	return
}

func (u *unionFind) find(x int) (parent int) {
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

func (u *unionFind) union(x, y int) {
	rootX, rootY := u.find(x), u.find(y)
	if rootX == rootY {
		return
	}

	u.parent[rootX] = rootY
	u.size[rootY] += u.size[rootX]
}

func (u *unionFind) getSize(x int) int {
	root := u.find(x)
	return u.size[root]
}
