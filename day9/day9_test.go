package day9_test

import (
	"os"
	"testing"

	"github.com/kaveenr/aoc24/day9"
)

func Test_Day9_Part1_Example(t *testing.T) {
	result, err := day9.Part1(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 1928
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day9_Part2_Example(t *testing.T) {
	result, err := day9.Part2(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 2858
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

// https://old.reddit.com/r/adventofcode/comments/1hamyyn/2024_day_9_part_2_python/m19rlcx/
func Test_Day9_Part2_Example_Reddit(t *testing.T) {
	result, err := day9.Part2("2333133121414131499")
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 6204
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day9_Part1(t *testing.T) {
	result, err := day9.Part1(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 6259790630969
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day9_Part2(t *testing.T) {
	result, err := day9.Part2(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := -1
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Benchmark_Day9_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day9.Part1(testInput)
	}
}

func Benchmark_Day9_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day9.Part2(testInput)
	}
}

var (
	myPuzzleInput, _ = os.ReadFile(`input.txt`)
	testInput        = `2333133121414131402`
)
