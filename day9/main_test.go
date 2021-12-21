package main

import (
	"testing"
)

func TestPart1(t *testing.T) {

	content := []byte(`2199943210
3987894921
9856789892
8767896789
9899965678`)
	res := run_part1(content)
	expected := 15
	if res != expected {
		t.Errorf("error in part 1: expected %d, got %d", expected, res)
	}
}

func TestPart2(t *testing.T) {

	content := []byte(`2199943210
3987894921
9856789892
8767896789
9899965678`)
	res := run_part2(content)
	expected := 1134
	if res != expected {
		t.Errorf("error in part 2: expected %d, got %d", expected, res)
	}
}
