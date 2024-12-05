package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) (result int, err error) {

	rules, updates, err := parseInput(input)
	if err != nil {
		return result, fmt.Errorf("Input parsing error: %w", err)
	}
	for _, update := range updates {
		if isValid(update, rules) {
			midNum := update[len(update)/2]
			result += midNum
		}
	}
	return
}

func Part2(input string) (result int, err error) {
	rules, updates, err := parseInput(input)
	if err != nil {
		return result, fmt.Errorf("Input parsing error: %w", err)
	}
	for _, update := range updates {
		if !isValid(update, rules) {
			corrected := correctUpdate(update, rules)
			midNum := corrected[len(corrected)/2]
			result += midNum
		}
	}
	return
}

func isValid(update []int, rules map[int][]int) (valid bool) {
	return slices.IsSortedFunc(update, func(i, j int) int {
		if slices.Contains(rules[i], j) {
			return 1
		} else {
			return -1
		}
	})
}

func correctUpdate(update []int, rules map[int][]int) (corrected []int) {

	corrected = slices.Clone(update)
	slices.SortFunc(corrected, func(i, j int) int {
		if slices.Contains(rules[i], j) {
			return 1
		} else {
			return -1
		}
	})
	return
}

func parseInput(input string) (rules map[int][]int, updates [][]int, err error) {

	rules = make(map[int][]int)
	input = strings.TrimSpace(input)
	sections := strings.Split(input, "\n\n")
	if len(sections) != 2 {
		return rules, updates, fmt.Errorf("Found %d sections, only 2 allowed", len(sections))
	}
	for _, section := range strings.Split(sections[0], "\n") {
		var cur []int
		for _, sectionVal := range strings.Split(section, "|") {
			val, err := strconv.Atoi(sectionVal)
			if err != nil {
				return rules, updates, fmt.Errorf("Unable to parse number %s", sectionVal)
			}
			cur = append(cur, val)
		}
		if len(cur) != 2 {
			return rules, updates, fmt.Errorf("Found %d rule sections, only 2 allowed in '%s'", len(cur), section)
		}
		if _, ok := rules[cur[1]]; !ok {
			rules[cur[1]] = make([]int, 0)
		}
		rules[cur[1]] = append(rules[cur[1]], cur[0])
	}
	for _, section := range strings.Split(sections[1], "\n") {
		var cur []int
		for _, sectionVal := range strings.Split(section, ",") {
			val, err := strconv.Atoi(sectionVal)
			if err != nil {
				return rules, updates, fmt.Errorf("Unable to parse number %s", sectionVal)
			}
			cur = append(cur, val)
		}
		updates = append(updates, cur)
	}
	return
}
