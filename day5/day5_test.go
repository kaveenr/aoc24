package day5_test

import (
	"os"
	"testing"

	"github.com/kaveenr/aoc24/day5"
)

func Test_Day5_Part1_Example(t *testing.T) {
	result, err := day5.Part1(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 143
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day5_Part2_Example(t *testing.T) {
	result, err := day5.Part2(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 123
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day5_Part1(t *testing.T) {
	result, err := day5.Part1(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 5747
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day5_Part2(t *testing.T) {
	result, err := day5.Part2(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 5502
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Benchmark_Day5_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day5.Part1(testInput)
	}
}

func Benchmark_Day5_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day5.Part2(testInput)
	}
}

var (
	myPuzzleInput, _ = os.ReadFile(`input.txt`)
	testInput        = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`
)
