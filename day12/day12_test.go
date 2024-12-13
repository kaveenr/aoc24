package day12_test

import (
	"os"
	"testing"

	"github.com/kaveenr/aoc24/day12"
)

func Test_Day12_Part1_Example(t *testing.T) {
	result, err := day12.Part1(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 1930
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

// func Test_Day12_Part2_Example(t *testing.T) {
// 	result, err := day12.Part2(testInput)
// 	if err != nil {
// 		t.Errorf("error: %s", err)
// 	}
// 	expected := 1206
// 	if result != expected {
// 		t.Errorf("expected %d, got %d", expected, result)
// 	}
// }

func Test_Day12_Part1(t *testing.T) {
	result, err := day12.Part1(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 1449902
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

// func Test_Day12_Part2(t *testing.T) {
// 	result, err := day12.Part2(string(myPuzzleInput))
// 	if err != nil {
// 		t.Errorf("error: %s", err)
// 	}
// 	expected := -1
// 	if result != expected {
// 		t.Errorf("expected %d, got %d", expected, result)
// 	}
// }

func Benchmark_Day12_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day12.Part1(testInput)
	}
}

// func Benchmark_Day12_Part2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		day12.Part2(testInput)
// 	}
// }

var (
	myPuzzleInput, _ = os.ReadFile(`input.txt`)
	testInput        = `
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
	`
)
