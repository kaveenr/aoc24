package day7

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) (result int, err error) {
	equations, err := parseInput(input)
	if err != nil {
		return result, fmt.Errorf("Input parsing error: %w", err)
	}

	return solve(equations, []string{"+", "*"})
}

func Part2(input string) (result int, err error) {
	equations, err := parseInput(input)
	if err != nil {
		return result, fmt.Errorf("Input parsing error: %w", err)
	}

	return solve(equations, []string{"+", "*", "||"})
}

type equation struct {
	answer int
	parts  []int
}

func solve(equations []equation, ops []string) (result int, err error) {

	for _, curEquation := range equations {
		solveStack := []equation{
			{
				answer: curEquation.parts[0],
				parts:  curEquation.parts[1:],
			},
		}
		var cur equation
		for true {
			if len(solveStack) == 0 {
				break
			}
			cur, solveStack = solveStack[0], solveStack[1:]
			if len(cur.parts) != 0 {
				for _, op := range ops {
					var next equation
					switch op {
					case "+":
						next.answer = cur.answer + cur.parts[0]
					case "*":
						next.answer = cur.answer * cur.parts[0]
					case "||":
						concatNum, err := strconv.Atoi(fmt.Sprintf("%d%d", cur.answer, cur.parts[0]))
						if err != nil {
							return result, fmt.Errorf("Unable to concat %d, %d: %w", cur.answer, cur.parts[0], err)
						}
						next.answer = concatNum
					}
					next.parts = cur.parts[1:]
					solveStack = append(solveStack, next)
				}
				continue
			} else if cur.answer == curEquation.answer {
				result += curEquation.answer
				break
			}
		}

	}
	return
}

func parseInput(input string) (result []equation, err error) {
	input = strings.TrimSpace(input)
	for rowIdx, row := range strings.Split(input, "\n") {
		inputParts := strings.Split(row, ":")
		if len(inputParts) != 2 {
			return result, fmt.Errorf("found %d parts on line %d, only 2 allowed", len(inputParts), rowIdx)
		}
		answer, err := strconv.Atoi(strings.TrimSpace(inputParts[0]))
		if err != nil {
			return result, fmt.Errorf("error parsing answer on line %d: %w", rowIdx, err)
		}
		parts := make([]int, 0)
		for partIdx, val := range strings.Split(strings.TrimSpace(inputParts[1]), " ") {
			answer, err := strconv.Atoi(strings.TrimSpace(val))
			if err != nil {
				return result, fmt.Errorf("error parsing part on line %d col %d: %w", rowIdx, partIdx, err)
			}
			parts = append(parts, answer)
		}

		result = append(result, equation{
			answer: answer,
			parts:  parts,
		})
	}
	return
}
