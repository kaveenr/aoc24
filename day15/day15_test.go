package day15_test

import (
	"os"
	"testing"

	"github.com/kaveenr/aoc24/day15"
)

func Test_Day15_Part1_Example(t *testing.T) {
	result, err := day15.Part1(testInput, viz)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 10092
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

// func Test_Day15_Part2_Example(t *testing.T) {
// 	result, err := day15.Part2(testInput, false)
// 	if err != nil {
// 		t.Errorf("error: %s", err)
// 	}
// 	expected := 9021
// 	if result != expected {
// 		t.Errorf("expected %d, got %d", expected, result)
// 	}
// }

func Test_Day15_Part1(t *testing.T) {
	result, err := day15.Part1(string(myPuzzleInput), viz)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expected := 1415498
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

// func Test_Day15_Part2(t *testing.T) {
// 	result, err := day15.Part2(string(myPuzzleInput))
// 	if err != nil {
// 		t.Errorf("error: %s", err)
// 	}
// 	expected := -1
// 	if result != expected {
// 		t.Errorf("expected %d, got %d", expected, result)
// 	}
// }

func Benchmark_Day15_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day15.Part1(testInput, false)
	}
}

// func Benchmark_Day15_Part2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		day15.Part2(testInput)
// 	}
// }

var (
	viz              = false
	myPuzzleInput, _ = os.ReadFile(`input.txt`)
	testInput        = `
##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^
	`
)
