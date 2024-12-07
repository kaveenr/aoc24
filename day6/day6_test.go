package day6_test

import (
	"os"
	"testing"

	"github.com/kaveenr/aoc24/day6"
)

func Test_Day6_Part1_Example(t *testing.T) {
	result, err := day6.Part1(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 41
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day6_Part2_Example(t *testing.T) {
	result, err := day6.Part2(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 6
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day6_Part1(t *testing.T) {
	result, err := day6.Part1(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 5531
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day6_Part2(t *testing.T) {
	result, err := day6.Part2(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 2165
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Benchmark_Day6_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day6.Part1(testInput)
	}
}

func Benchmark_Day6_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day6.Part2(testInput)
	}
}

var (
	myPuzzleInput, _ = os.ReadFile(`input.txt`)
	testInput        = `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
	`
)
