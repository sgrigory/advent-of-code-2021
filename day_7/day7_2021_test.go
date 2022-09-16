package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	test_input := "16,1,2,0,4,2,7,1,2,14"

	input := strings.Split(test_input, ",")
	res := part1(input)
	corr_res := 37
	if res != corr_res {
		t.Fatalf("Part 1 test failed: %d vs %d", res, corr_res)
	}

}

func TestPart2(t *testing.T) {
	test_input := "16,1,2,0,4,2,7,1,2,14"

	input := strings.Split(strings.Trim(test_input, " \n"), ",")
	res := part2(input)
	corr_res := 168
	if res != corr_res {
		t.Fatalf("Part 2 test failed: %d vs %d", res, corr_res)
	}

}
