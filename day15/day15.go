package day15

import (
	"fmt"
	"slices"
	"strings"

	"github.com/fogleman/gg"
)

func Part1(input string, vizEnabled bool) (result int, err error) {
	size, area, robot, moves := parseInput(input, false)
	if vizEnabled {
		im := viz(size, area, robot)
		im.SavePNG("0.png")
	}
	for moveIdx, move := range moves {
		switch move {
		case '<':
			robot, area = simulate(robot, area, vec{0, -1})
		case '>':
			robot, area = simulate(robot, area, vec{0, 1})
		case '^':
			robot, area = simulate(robot, area, vec{-1, 0})
		case 'v':
			robot, area = simulate(robot, area, vec{1, 0})
		}
		if vizEnabled {
			im := viz(size, area, robot)
			im.SavePNG(fmt.Sprintf("%d.png", moveIdx))
		}
	}

	for loc, obsType := range area {
		if obsType == BOX {
			result += (100 * loc.row) + loc.col
		}
	}

	return
}

func Part2(input string, vizEnabled bool) (result int, err error) {

	return
}

func simulate(robot vec, area map[vec]int, move vec) (nextRobot vec, nextArea map[vec]int) {
	next := robot.add(move)
	if v, ok := area[next]; ok && v == WALL {
		return robot, area
	}
	if v, ok := area[next]; ok && v == BOX {
		moveStack := []vec{next}
		for true {
			nextMove := moveStack[len(moveStack)-1].add(move)
			// Hit a wall? No luck.
			if v, ok := area[nextMove]; ok && v == WALL {
				moveStack = make([]vec, 0)
				break
			}
			// Hit thin air? great
			if _, ok := area[nextMove]; !ok {
				break
			}
			// Hit a box? Move on
			if v, ok := area[nextMove]; ok && (v == BOX || v == WBOX) {
				moveStack = append(moveStack, nextMove)
				continue
			}
		}
		// Nothing to move? stay as is
		if len(moveStack) == 0 {
			return robot, area
		}
		// Make the moves
		slices.Reverse(moveStack)
		for _, moveBox := range moveStack {
			delete(area, moveBox)
			area[moveBox.add(move)] = BOX
		}
	}
	return next, area
}

const (
	WALL = iota
	BOX
	WBOX
)

func parseInput(input string, wide bool) (size vec, area map[vec]int, start vec, moves []rune) {
	area = make(map[vec]int)
	input = strings.TrimSpace(input)
	parts := strings.Split(input, "\n\n")

	if wide {
		parts[0] = strings.ReplaceAll(parts[0], "#", "##")
		parts[0] = strings.ReplaceAll(parts[0], "O", "[]")
		parts[0] = strings.ReplaceAll(parts[0], ".", "..")
		parts[0] = strings.ReplaceAll(parts[0], "@", "@.")
	}

	rows := strings.Split(parts[0], "\n")
	for rowIdx, row := range rows {
		for colIdx, char := range row {
			cur := vec{rowIdx, colIdx}
			switch char {
			case '#':
				area[cur] = WALL
			case 'O':
				area[cur] = BOX
			case '[':
				area[cur] = BOX
			case ']':
				area[cur] = WBOX
			case '@':
				start = cur
				size = vec{len(rows), len(row)}
			}
		}
	}
	moves = []rune(parts[1])
	return
}

func viz(size vec, area map[vec]int, robot vec) *gg.Context {
	bSize := 100
	dc := gg.NewContext(size.col*bSize, size.row*bSize)
	dc.SetRGB255(174, 214, 241)
	dc.DrawRectangle(0, 0, float64(size.col*bSize), float64(size.row*bSize))
	dc.Fill()
	for loc, obsType := range area {
		x, y := float64(loc.col*bSize), float64(loc.row*bSize)
		switch obsType {
		case WALL:
			dc.SetRGB255(217, 136, 128)
			dc.DrawRectangle(x, y, float64(bSize), float64(bSize))
		case BOX:
			dc.SetRGB255(247, 220, 111)
			dc.DrawRectangle(x, y, float64(bSize), float64(bSize))
		case WBOX:
			dc.SetRGB255(241, 196, 15)
			dc.DrawRectangle(x, y, float64(bSize), float64(bSize))
		}
		dc.Fill()
	}
	dc.SetRGB255(52, 73, 94)
	radius := (float64(bSize) / 2)
	dc.DrawCircle(float64(robot.col*bSize)+radius, float64(robot.row*bSize)+radius, radius-5)
	dc.Fill()

	return dc
}

type vec struct {
	row int
	col int
}

func (v vec) String() string {
	return fmt.Sprintf("(%d,%d)", v.row, v.col)
}

func (v vec) add(x vec) (n vec) {
	return vec{
		row: v.row + x.row,
		col: v.col + x.col,
	}
}

// func (area vec) outOfBounds(loc vec) bool {
// 	if loc.row < 0 || loc.row >= area.row {
// 		return true
// 	}
// 	if loc.col < 0 || loc.col >= area.col {
// 		return true
// 	}
// 	return false
// }
