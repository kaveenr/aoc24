package day4_test

import (
	"os"
	"testing"

	"github.com/kaveenr/aoc24/day4"
)

func Test_Day4_Part1_Example(t *testing.T) {
	result, err := day4.Part1(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 18
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day4_Part2_Example(t *testing.T) {
	result, err := day4.Part2(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 9
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day4_Part1(t *testing.T) {
	result, err := day4.Part1(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 2599
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day4_Part2(t *testing.T) {
	result, err := day4.Part2(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 1948
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Benchmark_Day4_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day4.Part1(testInput)
	}
}

func Benchmark_Day4_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day4.Part2(testInput)
	}
}

var (
	myPuzzleInput, _ = os.ReadFile(`input.txt`)
	testInput        = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
	`
)
