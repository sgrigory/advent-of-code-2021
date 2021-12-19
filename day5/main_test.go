package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	content := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
	rows := strings.Split(string(content), "\n")

	res := run_part1(rows)
	expected := 5
	if res != expected {
		t.Errorf("error in part 1: expected %d, got %d", expected, res)
	}

}

func TestPart2(t *testing.T) {
	content := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
	rows := strings.Split(string(content), "\n")

	res := run_part2(rows)
	expected := 12
	if res != expected {
		t.Errorf("error in part 2: expected %d, got %d", expected, res)
	}

}
