package main

import (
	"testing"
)

const INPUT1 = `Player 1 starting position: 4
Player 2 starting position: 8`

func TestPart1(t *testing.T) {

	content := []byte(INPUT1)
	res := run_part1(content)
	var expected int
	expected = 739785
	if res != expected {
		t.Errorf("error in part 1: expected %d, got %d", expected, res)
	}
}

func TestPart2(t *testing.T) {

	content := []byte(INPUT1)
	res := run_part2(content)
	expected := 444356092776315
	if res != expected {
		t.Errorf("error in part 2: expected %d, got %d", expected, res)
	}
}
