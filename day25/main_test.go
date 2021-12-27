package main

import (
	"testing"
)

const INPUT = `v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`

func TestPart1(t *testing.T) {

	expected := 58
	content := []byte(INPUT)
	res := run_part1(content)
	if res != expected {
		t.Errorf("error in part 1: expected %d, got %d", expected, res)
	}
}

// func TestPart2(t *testing.T) {

// 	content := []byte(INPUT)
// 	res := run_part2(content, 40)
// 	expected := 2188189693529
// 	if res != expected {
// 		t.Errorf("error in part 2: expected %d, got %d", expected, res)
// 	}
// }
