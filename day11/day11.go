package day11

import (
	"maps"
	"math"
	"strconv"
	"strings"
)

func Part1(input string) (result int, err error) {

	stones := parseInput(input)
	result = blink(stones, 25)
	return
}

func Part2(input string) (result int, err error) {
	stones := parseInput(input)
	result = blink(stones, 75)
	return
}

func blink(stonesMap map[int]int, count int) (result int) {

	for blinkIdx := 0; blinkIdx < count; blinkIdx++ {

		snapshot := maps.Clone(stonesMap)

		for stone := range maps.Keys(snapshot) {
			// Rule 1
			if stone == 0 {
				stonesMap[1] += snapshot[0]
				stonesMap[0] -= snapshot[0]
				continue
			}
			// Rule 2
			digitCount := numDigits(stone)
			if digitCount%2 == 0 {
				valueStr := strconv.Itoa(stone)
				leftVal, _ := strconv.Atoi(valueStr[:digitCount/2])
				rightVal, _ := strconv.Atoi(valueStr[digitCount/2:])
				stonesMap[leftVal] += snapshot[stone]
				stonesMap[rightVal] += snapshot[stone]
				stonesMap[stone] -= snapshot[stone]
				continue
			}
			// Rule 3
			stonesMap[stone*2024] += snapshot[stone]
			stonesMap[stone] -= snapshot[stone]
		}
	}
	for _, v := range stonesMap {
		result += v
	}
	return
}

func parseInput(input string) (stones map[int]int) {
	stones = make(map[int]int)
	input = strings.TrimSpace(input)
	for _, stone := range strings.Split(input, " ") {
		val, _ := strconv.Atoi(stone)
		stones[val] += 1
	}
	return
}

func numDigits(n int) int {
	if n == 0 {
		return 1
	}
	return int(math.Floor(math.Log10(math.Abs(float64(n)))) + 1)
}
