package day1

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) (result int, err error) {

	list1, list2, err := parseInput(input)
	if err != nil {
		return result, fmt.Errorf("input parsing error: %w", err)
	}
	slices.Sort(list1)
	slices.Sort(list2)

	for i := 0; i < len(list1); i++ {
		diff := math.Abs(float64(list1[i] - list2[i]))
		result += int(diff)
	}
	return
}

func Part2(input string) (result int, err error) {
	list1, list2, err := parseInput(input)
	if err != nil {
		return result, fmt.Errorf("input parsing error: %w", err)
	}
	for i := 0; i < len(list1); i++ {
		cur := list1[i]
		multiplier := 0
		for j := 0; j < len(list2); j++ {
			if cur == list2[j] {
				multiplier += 1
			}
		}
		result += cur * multiplier
	}
	return
}

func parseInput(input string) (list1, list2 []int, err error) {
	input = strings.TrimSpace(input)
	for _, line := range strings.Split(input, "\n") {
		numbers := strings.Split(line, "   ")
		num1, err := strconv.Atoi(strings.TrimSpace(numbers[0]))
		if err != nil {
			return list1, list2, fmt.Errorf("err parsing n1: %w", err)
		}
		list1 = append(list1, num1)
		num2, err := strconv.Atoi(strings.TrimSpace(numbers[1]))
		if err != nil {
			return list1, list2, fmt.Errorf("err parsing n2: %w", err)
		}
		list2 = append(list2, num2)
	}
	return
}
