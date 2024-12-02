package day2_test

import (
	"os"
	"testing"

	"github.com/kaveenr/aoc24/day2"
)

func Test_Day2_Part1_Example(t *testing.T) {
	result, err := day2.Part1(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 2
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day2_Part2_Example(t *testing.T) {
	result, err := day2.Part2(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 4
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day2_Part1(t *testing.T) {
	result, err := day2.Part1(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 356
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day2_Part2(t *testing.T) {
	result, err := day2.Part2(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 20373490
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Benchmark_Day2_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day2.Part1(testInput)
	}
}

func Benchmark_Day2_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day2.Part2(testInput)
	}
}

var (
	myPuzzleInput, _ = os.ReadFile(`input.txt`)
	testInput        = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`
)
