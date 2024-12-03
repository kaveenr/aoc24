package day3

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	rePattern = `(?<mul>mul\((?<x>[0-9]{1,3}),(?<y>[0-9]{1,3})\))|(?<do>do\(\))|(?<don>don't\(\))`
	re        = regexp.MustCompile(rePattern)
)

func Part1(input string) (result int, err error) {

	return solve(input, false)
}

func Part2(input string) (result int, err error) {

	return solve(input, true)
}

func solve(input string, conditionals bool) (result int, err error) {
	groupNames := re.SubexpNames()
	e, x, y := 1, 0, 0
	for _, match := range re.FindAllStringSubmatch(input, -1) {
		for groupIdx, group := range match {
			name := groupNames[groupIdx]
			if group == "" {
				continue
			}
			switch name {
			case "do":
				if !conditionals {
					continue
				}
				e = 1
			case "don":
				if !conditionals {
					continue
				}
				e = 0
			case "x":
				x, err = strconv.Atoi(group)
				if err != nil {
					return result, fmt.Errorf("Unable to parse x '%s': %w", group, err)
				}
			case "y":
				y, err = strconv.Atoi(group)
				if err != nil {
					return result, fmt.Errorf("Unable to parse y '%s': %w", group, err)
				}
				result += (x * y) * e
			}
		}
	}
	return
}
