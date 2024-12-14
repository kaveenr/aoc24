package day14_test

import (
	"os"
	"testing"

	"github.com/kaveenr/aoc24/day14"
)

func Test_Day14_Part1_Example(t *testing.T) {
	result, err := day14.Part1(testInput)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 12
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day14_Part1(t *testing.T) {
	result, err := day14.Part1(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 228690000
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Test_Day14_Part2(t *testing.T) {
	result, err := day14.Part2(string(myPuzzleInput))
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 7093
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func Benchmark_Day14_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day14.Part1(testInput)
	}
}

var (
	myPuzzleInput, _ = os.ReadFile(`input.txt`)
	testInput        = `
p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3
	`
)
