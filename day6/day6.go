package day6

import (
	"fmt"
	"slices"
	"strings"
	"sync"
)

func Part1(input string) (result int, err error) {

	area, guard, obstacles := parseInput(input)

	visits := make(map[vec]bool)
	// Start position needs to be there
	// DUH: Thanks Reddit
	visits[guard] = true

	cursor := guard
	heading := vec{
		row: -1,
		col: 0,
	}
	for true {
		nextStep := vec{
			row: cursor.row + heading.row,
			col: cursor.col + heading.col,
		}
		if outOfBounds(nextStep, area) {
			break
		}
		if slices.Contains(obstacles, nextStep) {
			heading = vec{
				row: heading.col,
				col: -heading.row,
			}
			continue
		}
		cursor = nextStep
		visits[cursor] = true
	}
	result = len(visits)

	return
}

func Part2(input string) (result int, err error) {

	area, guard, obstacles := parseInput(input)

	// Step 1
	// Optimization 1: Walk it once, note down points of interest
	pointsOfInterest := make(map[vec]bool)

	cursor := guard
	heading := vec{
		row: -1,
		col: 0,
	}
	for true {
		nextStep := vec{
			row: cursor.row + heading.row,
			col: cursor.col + heading.col,
		}
		if outOfBounds(nextStep, area) {
			break
		}
		if slices.Contains(obstacles, nextStep) {
			heading = vec{
				row: heading.col,
				col: -heading.row,
			}
			continue
		}
		cursor = nextStep
		if cursor != guard {
			pointsOfInterest[cursor] = true
		}
	}

	// Step 2
	// Use visited places as basis for new obstacle placement

	// Goroutines go brrr
	var wg sync.WaitGroup
	var m sync.Mutex

	for newObstacle := range pointsOfInterest {

		wg.Add(1)
		go func() {
			defer wg.Done()
			// Collision Check Map string with(loc+heading)
			newObstacleVisits := make(map[string]bool)
			cursor := guard
			heading := vec{
				row: -1,
				col: 0,
			}
			for true {
				nextStep := vec{
					row: cursor.row + heading.row,
					col: cursor.col + heading.col,
				}
				if outOfBounds(nextStep, area) {
					break
				}
				if slices.Contains(obstacles, nextStep) || newObstacle == nextStep {
					// Check if already visited
					if _, ok := newObstacleVisits[nextStep.String()+heading.String()]; ok {
						m.Lock()
						result += 1
						m.Unlock()
						break
					}
					newObstacleVisits[nextStep.String()+heading.String()] = true

					heading = vec{
						row: heading.col,
						col: -heading.row,
					}
					continue
				}
				cursor = nextStep
			}
		}()
	}
	wg.Wait()
	return
}

type vec struct {
	row int
	col int
}

func (v vec) String() string {
	return fmt.Sprintf("vec{row: %d, col: %d}", v.row, v.col)
}

func outOfBounds(loc vec, area vec) bool {
	if loc.row < 0 || loc.row >= area.row {
		return true
	}
	if loc.col < 0 || loc.col >= area.col {
		return true
	}
	return false
}

func parseInput(input string) (area vec, guard vec, obstacles []vec) {
	input = strings.TrimSpace(input)
	rows := strings.Split(input, "\n")
	for rowIdx, row := range rows {
		for colIdx, char := range row {
			switch char {
			case '^':
				guard = vec{
					row: rowIdx,
					col: colIdx,
				}
				area = vec{
					row: len(rows),
					col: len(row),
				}
			case '#':
				obstacles = append(obstacles, vec{
					row: rowIdx,
					col: colIdx,
				})
			}
		}
	}
	return
}
