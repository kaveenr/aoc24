package day2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part1(input string) (result int, err error) {

	reports, err := parseInput(input)
	if err != nil {
		return result, fmt.Errorf("input parsing error: %w", err)
	}
	for _, report := range reports {
		if isSafe(report) {
			result += 1
		}
	}
	return
}

func Part2(input string) (result int, err error) {
	reports, err := parseInput(input)
	if err != nil {
		return result, fmt.Errorf("input parsing error: %w", err)
	}
	for _, report := range reports {
		if isSafe(report) {
			result += 1
		} else {
			for i := 0; i < len(report); i++ {
				modifiedReport := make([]int, 0, len(report)-1)
				modifiedReport = append(modifiedReport, report[:i]...)
				modifiedReport = append(modifiedReport, report[i+1:]...)
				if isSafe(modifiedReport) {
					result += 1
					break
				}
			}
		}

	}
	return
}

func isSafe(report []int) (result bool) {
	var safe bool
	var sign bool
	for i := 0; i < len(report); i++ {
		if i == 0 {
			continue
		}
		delta := float64(report[i] - report[i-1])
		deltaAbs := math.Abs(delta)
		if deltaAbs < 1 || deltaAbs > 3 {
			safe = false
			break
		}
		if i != 1 && !(sign == (delta > 0)) {
			safe = false
			break
		}
		sign = delta > 0
		safe = true
	}
	return safe
}

func parseInput(input string) (parsed [][]int, err error) {
	input = strings.TrimSpace(input)

	for x, row := range strings.Split(input, "\n") {
		cur := make([]int, 0)
		for y, sector := range strings.Split(row, " ") {
			val, err := strconv.Atoi(strings.TrimSpace(sector))
			if err != nil {
				return parsed, fmt.Errorf("Unable to parse '%s' on line %d sector %d: %w", sector, x, y, err)
			}
			cur = append(cur, val)
		}
		parsed = append(parsed, cur)
	}
	return
}
