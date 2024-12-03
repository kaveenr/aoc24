package day3_test

import (
	"os"
	"testing"

	"github.com/kaveenr/aoc24/day3"
)

func Test_Day3_Part1_Example(t *testing.T) {
	result, err := day3.Part1(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 161
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day3_Part2_Example(t *testing.T) {
	result, err := day3.Part2(testInput1)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 48
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day3_Part1(t *testing.T) {
	result, err := day3.Part1(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 173517243
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day3_Part2(t *testing.T) {
	result, err := day3.Part2(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 100450138
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Benchmark_Day3_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day3.Part1(testInput)
	}
}

func Benchmark_Day3_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day3.Part2(testInput)
	}
}

var (
	myPuzzleInput, _ = os.ReadFile(`input.txt`)
	testInput        = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	testInput1       = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
)
