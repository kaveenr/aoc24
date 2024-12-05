package day4

import (
	"fmt"
	"slices"
	"strings"
)

var (
	matchTerm  = []rune("XMAS")
	matchTerm2 = [][][]rune{
		{
			[]rune("M.S"),
			[]rune(".A."),
			[]rune("M.S"),
		},
		{
			[]rune("S.M"),
			[]rune(".A."),
			[]rune("S.M"),
		},
		{
			[]rune("M.M"),
			[]rune(".A."),
			[]rune("S.S"),
		},
		{
			[]rune("S.S"),
			[]rune(".A."),
			[]rune("M.M"),
		},
	}
)

func Part1(input string) (result int, err error) {

	puzzle, err := parse(input, matchTerm)
	if err != nil {
		return result, fmt.Errorf("Input parsing error: %w", err)
	}

	for rowIdx := 0; rowIdx < len(puzzle); rowIdx++ {
		for colIdx := 0; colIdx < len(puzzle[rowIdx]); colIdx++ {
			if puzzle[rowIdx][colIdx] == matchTerm[0] {
				result += matches(puzzle, rowIdx, colIdx)
			}
		}
	}
	return
}

func Part2(input string) (result int, err error) {

	puzzle, err := parse(input, []rune("MAS"))
	if err != nil {
		return result, fmt.Errorf("Input parsing error: %w", err)
	}

	for rowIdx := 0; rowIdx < len(puzzle); rowIdx++ {
		for colIdx := 0; colIdx < len(puzzle[rowIdx]); colIdx++ {
			if puzzle[rowIdx][colIdx] == []rune("A")[0] {
				for _, term := range matchTerm2 {
					if match2(puzzle, rowIdx, colIdx, term) {
						result += 1
					}
				}
			}
		}
	}
	return
}

func parse(input string, mask []rune) (result [][]rune, err error) {

	input = strings.TrimSpace(input)
	for _, row := range strings.Split(input, "\n") {
		cur := []rune(row)
		for i := 0; i < len(cur); i++ {
			if !slices.Contains(mask, cur[i]) {
				cur[i] = []rune(".")[0]
			}
		}
		result = append(result, cur)
	}
	return
}

func matches(p [][]rune, row, col int) (count int) {

	var matchSingle = func(p [][]rune, row, col int, rowMod, colMod int) (match bool) {

		endRow, endCol := row+(rowMod*(len(matchTerm)-1)), col+(colMod*(len(matchTerm)-1))
		if endRow < 0 || endRow >= len(p) {
			return
		}
		if endCol < 0 || endCol >= len(p[row]) {
			return
		}
		rowIdx, colIdx, matchIdx := row, col, 0
		for rowIdx != endRow || colIdx != endCol {
			rowIdx += rowMod
			colIdx += colMod
			matchIdx += 1
			if p[rowIdx][colIdx] != matchTerm[matchIdx] {
				return false
			}
		}
		return true
	}

	matchMods := [][]int{
		{0, -1},  // Left
		{0, 1},   // Right
		{-1, 0},  // Up
		{1, 0},   // Down
		{-1, -1}, // Left Up
		{-1, 1},  // Right Up
		{1, -1},  // Left Down
		{1, 1},   // Right Down
	}
	for _, mod := range matchMods {
		if matchSingle(p, row, col, mod[0], mod[1]) {
			count += 1
		}
	}
	return
}

func match2(p [][]rune, row, col int, term [][]rune) (match bool) {

	for rowIdx := 0; rowIdx < len(term); rowIdx++ {
		for colIdx := 0; colIdx < len(term[rowIdx]); colIdx++ {
			checkRowIdx, checkColIdx := (row-1)+rowIdx, (col-1)+colIdx
			if checkRowIdx < 0 || checkRowIdx >= len(p) {
				return
			}
			if checkColIdx < 0 || checkColIdx >= len(p[row]) {
				return
			}
			if term[rowIdx][colIdx] == []rune(".")[0] {
				continue
			}
			if term[rowIdx][colIdx] != p[checkRowIdx][checkColIdx] {
				return
			}
		}
	}
	return true
}
