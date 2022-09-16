package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	test_input := `forward 5
down 5
forward 8
up 3
down 8
forward 2
`

	input := strings.Split(test_input, "\n")
	res := part1(input)
	if res != 150 {
		t.Fatalf("Part 1 test failed: %d vs %d", res, 7)
	}

}

func TestPart2(t *testing.T) {
	test_input := `forward 5
down 5
forward 8
up 3
down 8
forward 2
`

	input := strings.Split(test_input, "\n")
	res := part2(input)
	if res != 900 {
		t.Fatalf("Part 2 test failed: %d vs %d", res, 900)
	}

}
