package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	test_input := `199
	200
	208
	210
	200
	207
	240
	269
	260
	263
	`

	input := strings.Split(test_input, "\n")
	res := part1(input)
	if res != 7 {
		t.Fatalf("Part 1 test failed: %d vs %d", res, 7)
	}

}

func TestPart2(t *testing.T) {
	test_input := `199
	200
	208
	210
	200
	207
	240
	269
	260
	263
	`

	input := strings.Split(test_input, "\n")
	res := part2(input)
	if res != 5 {
		t.Fatalf("Part 2 test failed: %d vs %d", res, 7)
	}

}
