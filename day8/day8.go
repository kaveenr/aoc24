package day8

import (
	"strings"
)

func Part1(input string) (result int, err error) {
	area, antennas := parseInput(input)
	eligiblePairs := makeAntennaPairs(antennas)

	uniqueAntinode := make(map[vec]bool)
	for _, resonantPair := range eligiblePairs {
		nodeA, nodeB := resonantPair[0], resonantPair[1]
		antinode := vec{
			row: nodeA.row + (nodeA.row - nodeB.row),
			col: nodeA.col + (nodeA.col - nodeB.col),
		}
		if antinode.outOfBounds(area) {
			continue
		}
		uniqueAntinode[antinode] = true
	}

	return len(uniqueAntinode), nil
}

func Part2(input string) (result int, err error) {
	area, antennas := parseInput(input)
	eligiblePairs := makeAntennaPairs(antennas)

	uniqueAntinode := make(map[vec]bool)
	for _, resonantPair := range eligiblePairs {
		nodeA, nodeB := resonantPair[0], resonantPair[1]
		// Keep making antinodes until out of bounds
		antinode := vec{
			row: nodeA.row,
			col: nodeA.col,
		}
		// Origin is an antinode as well
		uniqueAntinode[antinode] = true
		for true {
			antinode = vec{
				row: antinode.row + (nodeA.row - nodeB.row),
				col: antinode.col + (nodeA.col - nodeB.col),
			}
			if antinode.outOfBounds(area) {
				break
			}
			uniqueAntinode[antinode] = true
		}
	}

	return len(uniqueAntinode), nil
}

type vec struct {
	row int
	col int
}

func (loc *vec) outOfBounds(area vec) bool {
	if loc.row < 0 || loc.row >= area.row {
		return true
	}
	if loc.col < 0 || loc.col >= area.col {
		return true
	}
	return false
}

func parseInput(input string) (area vec, antennas map[string][]vec) {
	antennas = make(map[string][]vec)
	input = strings.TrimSpace(input)
	rows := strings.Split(input, "\n")
	for rowIdx, row := range rows {
		for colIdx, char := range row {
			switch char {
			case '.':
				continue
			default:
				area = vec{
					row: len(rows),
					col: len(row),
				}
				frq := string(char)
				if _, ok := antennas[frq]; !ok {
					antennas[frq] = make([]vec, 0)
				}
				antennas[frq] = append(antennas[frq], vec{
					row: rowIdx,
					col: colIdx,
				})
			}
		}
	}
	return
}

func makeAntennaPairs(antennas map[string][]vec) (eligiblePairs [][]vec) {
	// Iterate and make eligible channel pairs
	for _, antennas := range antennas {
		for curAntennaIdx, curAntenna := range antennas {
			for cmpAntennaIdx, cmpAntenna := range antennas {
				// TODO: check if in line of sight (didn't need it)
				// Don't add self
				if curAntennaIdx == cmpAntennaIdx {
					continue
				}
				eligiblePairs = append(eligiblePairs, []vec{curAntenna, cmpAntenna})
			}
		}
	}
	return
}
