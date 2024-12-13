package day12

import (
	"fmt"
	"strings"
)

var (
	moveVectors = []vec{
		{-1, 0}, // Up
		{1, 0},  // Down
		{0, -1}, // Left
		{0, 1},  // Right
	}
)

func Part1(input string) (result int, err error) {
	area, garden := parseInput(input)

	cellIndex := func(v vec) int {
		return area.col*v.row + v.col
	}
	uf := newUnionFind(area.row * area.col)

	perimeter := make(map[vec][]vec)
	locByIndex := make(map[int]vec)
	for rowIdx := 0; rowIdx < area.row; rowIdx++ {
		for colIdx := 0; colIdx < area.col; colIdx++ {
			curVec := vec{rowIdx, colIdx}
			for _, moveVector := range moveVectors {
				movedVec := curVec.add(moveVector)
				if area.outOfBounds(movedVec) {
					perimeter[curVec] = append(perimeter[curVec], movedVec)
					continue
				}
				if garden[curVec] != garden[movedVec] {
					perimeter[curVec] = append(perimeter[curVec], movedVec)
					continue
				}
				uf.union(cellIndex(curVec), cellIndex(movedVec))
			}
			locByIndex[cellIndex(curVec)] = curVec
		}
	}

	totalAreaMap := make(map[vec]int)
	totalPerimeterMap := make(map[vec]int)
	for rowIdx := 0; rowIdx < area.row; rowIdx++ {
		for colIdx := 0; colIdx < area.col; colIdx++ {
			curVec := vec{rowIdx, colIdx}
			parentIndex := uf.find(cellIndex(curVec))
			totalAreaMap[locByIndex[parentIndex]] += 1
			totalPerimeterMap[locByIndex[parentIndex]] += len(perimeter[curVec])
		}
	}

	for parent, totalArea := range totalAreaMap {
		result += totalArea * totalPerimeterMap[parent]
	}

	return
}

func Part2(input string) (result int, err error) {

	return
}

func parseInput(input string) (area vec, data map[vec]string) {
	data = make(map[vec]string)
	input = strings.TrimSpace(input)
	rows := strings.Split(input, "\n")
	for rowIdx, row := range rows {
		for colIdx, char := range row {
			data[vec{
				row: rowIdx,
				col: colIdx,
			}] = string(char)
		}
		if rowIdx == 0 {
			area = vec{
				row: len(rows),
				col: len(row),
			}
		}
	}
	return
}

type vec struct {
	row int
	col int
}

func (v vec) String() string {
	return fmt.Sprintf("%d,%d", v.row, v.col)
}

func (v vec) add(x vec) (n vec) {
	return vec{
		row: v.row + x.row,
		col: v.col + x.col,
	}
}

func (area vec) outOfBounds(loc vec) bool {
	if loc.row < 0 || loc.row >= area.row {
		return true
	}
	if loc.col < 0 || loc.col >= area.col {
		return true
	}
	return false
}

type unionFind struct {
	parent []int
}

func newUnionFind(size int) (uf *unionFind) {
	uf = &unionFind{
		parent: make([]int, size),
	}
	for parentIdx := 0; parentIdx < size; parentIdx++ {
		uf.parent[parentIdx] = parentIdx
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
	u.parent[rootY] = rootX
}
