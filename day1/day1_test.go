package day1_test

import (
	"os"
	"testing"

	"github.com/kaveenr/aoc24/day1"
)

func Test_Day1_Part1_Example(t *testing.T) {
	result, err := day1.Part1(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 11
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day1_Part2_Example(t *testing.T) {
	result, err := day1.Part2(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 31
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day1_Part1(t *testing.T) {
	result, err := day1.Part1(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 1722302
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day1_Part2(t *testing.T) {
	result, err := day1.Part2(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 20373490
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Benchmark_Day1_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day1.Part1(testInput)
	}
}

func Benchmark_Day1_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day1.Part2(testInput)
	}
}

var (
	myPuzzleInput, _ = os.ReadFile(`input.txt`)
	testInput        = `3   4
4   3
2   5
1   3
3   9
3   3`
)
