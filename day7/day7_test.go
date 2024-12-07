package day7_test

import (
	"os"
	"testing"

	"github.com/kaveenr/aoc24/day7"
)

func Test_Day7_Part1_Example(t *testing.T) {
	result, err := day7.Part1(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 3749
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day7_Part2_Example(t *testing.T) {
	result, err := day7.Part2(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 11387
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day7_Part1(t *testing.T) {
	result, err := day7.Part1(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 5540634308362
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day7_Part2(t *testing.T) {
	result, err := day7.Part2(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 472290821152397
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Benchmark_Day7_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day7.Part1(testInput)
	}
}

func Benchmark_Day7_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day7.Part2(testInput)
	}
}

var (
	myPuzzleInput, _ = os.ReadFile(`input.txt`)
	testInput        = `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`
)
