package main

import (
	"testing"
)

const INPUT = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

func TestPart1(t *testing.T) {

	expected := 17
	content := []byte(INPUT)
	res := run_part1(content, 1)
	if res != expected {
		t.Errorf("error in part 1: expected %d, got %d", expected, res)
	}
}
