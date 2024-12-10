package day10_test

import (
	"os"
	"testing"

	"github.com/kaveenr/aoc24/day10"
)

func Test_Day10_Part1_Example(t *testing.T) {
	result, err := day10.Part1(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 36
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day10_Part2_Example(t *testing.T) {
	result, err := day10.Part2(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 81
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day10_Part1(t *testing.T) {
	result, err := day10.Part1(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 496
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day10_Part2(t *testing.T) {
	result, err := day10.Part2(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 1120
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Benchmark_Day10_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day10.Part1(testInput)
	}
}

func Benchmark_Day10_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day10.Part2(testInput)
	}
}

var (
	myPuzzleInput, _ = os.ReadFile(`input.txt`)
	testInput        = `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
	`
)
