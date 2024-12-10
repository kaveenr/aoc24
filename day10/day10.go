package day10

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var (
	moveVectors = [][]int{
		{-1, 0}, // Up
		{1, 0},  // Down
		{0, -1}, // Left
		{0, 1},  // Right
	}
)

func Part1(input string) (result int, err error) {

	scores := calculateTrailScores(input)
	for _, v := range scores {
		result += len(v)
	}

	return
}

func Part2(input string) (result int, err error) {

	scores := calculateTrailScores(input)
	for _, v := range scores {
		for _, vi := range v {
			result += vi
		}
	}

	return
}

func calculateTrailScores(input string) (trailheadScores map[string]map[string]int) {

	topo := parseInput(input)
	size := []int{
		len(topo), len(topo[0]),
	}
	solve := make([][][]int, 0)

	// Populate solve stack with trailheads
	for rowIdx := 0; rowIdx < len(topo); rowIdx++ {
		for colIdx := 0; colIdx < len(topo[rowIdx]); colIdx++ {
			if topo[rowIdx][colIdx] == 0 {
				solve = append(solve, [][]int{
					{rowIdx, colIdx},
				})
			}
		}
	}

	trailheadScores = make(map[string]map[string]int)
	for true {
		// Stop when nothing else to solve
		if len(solve) == 0 {
			break
		}
		// Pop from solve stack
		var cur [][]int
		cur, solve = solve[0], solve[1:]

		// Check if trail lead to a peak
		last := cur[len(cur)-1]
		lastElevation := topo[last[0]][last[1]]
		if lastElevation == 9 {

			trailKey := fmt.Sprintf("%d,%d", cur[0][0], cur[0][1])
			destinationKey := fmt.Sprintf("(%d,%d)", last[0], last[1])
			if _, ok := trailheadScores[trailKey]; !ok {
				trailheadScores[trailKey] = make(map[string]int)
			}
			trailheadScores[trailKey][destinationKey] += 1
			continue
		}

		// Try traverse ahead
		for _, vector := range moveVectors {
			next := []int{
				last[0] + vector[0],
				last[1] + vector[1],
			}
			if outOfBounds(next, size) {
				continue
			}
			// Elation check
			if (lastElevation + 1) != topo[next[0]][next[1]] {
				continue
			}
			nextPath := slices.Clone(cur)
			nextPath = append(nextPath, next)
			solve = append(solve, nextPath)
		}
	}
	return
}

func parseInput(input string) (topo [][]int) {
	input = strings.TrimSpace(input)
	for _, row := range strings.Split(input, "\n") {
		var curRow []int
		for _, col := range row {
			val, _ := strconv.Atoi(string(col))
			curRow = append(curRow, val)
		}
		topo = append(topo, curRow)
	}
	return
}

func outOfBounds(loc []int, area []int) bool {
	if loc[0] < 0 || loc[0] >= area[0] {
		return true
	}
	if loc[1] < 0 || loc[1] >= area[1] {
		return true
	}
	return false
}
