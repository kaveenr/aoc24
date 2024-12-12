package day11_test

import (
	"os"
	"testing"

	"github.com/kaveenr/aoc24/day11"
)

func Test_Day11_Part1_Example(t *testing.T) {
	result, err := day11.Part1(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 55312
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day11_Part2_Example(t *testing.T) {
	result, err := day11.Part2(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 65601038650482
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day11_Part1(t *testing.T) {
	result, err := day11.Part1(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 217443
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day11_Part2(t *testing.T) {
	result, err := day11.Part2(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := -1
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Benchmark_Day11_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day11.Part1(testInput)
	}
}

func Benchmark_Day11_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day11.Part2(testInput)
	}
}

var (
	myPuzzleInput, _ = os.ReadFile(`input.txt`)
	testInput        = `125 17`
)
