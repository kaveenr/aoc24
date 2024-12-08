package day8_test

import (
	"os"
	"testing"

	"github.com/kaveenr/aoc24/day8"
)

func Test_Day8_Part1_Example(t *testing.T) {
	result, err := day8.Part1(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 14
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day8_Part2_Example(t *testing.T) {
	result, err := day8.Part2(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 34
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day8_Part1(t *testing.T) {
	result, err := day8.Part1(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 299
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day8_Part2(t *testing.T) {
	result, err := day8.Part2(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 1032
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Benchmark_Day8_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day8.Part1(testInput)
	}
}

func Benchmark_Day8_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day8.Part2(testInput)
	}
}

var (
	myPuzzleInput, _ = os.ReadFile(`input.txt`)
	testInput        = `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
	`
)
