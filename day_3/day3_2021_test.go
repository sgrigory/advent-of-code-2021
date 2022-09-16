package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	test_input := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

	input := strings.Split(test_input, "\n")
	res := part1(input)
	corr_res := 198
	if res != corr_res {
		t.Fatalf("Part 1 test failed: %d vs %d", res, corr_res)
	}

}

func TestPart2(t *testing.T) {
	test_input := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`

	input := strings.Split(strings.Trim(test_input, " \n"), "\n")
	res := part2(input)
	corr_res := 230
	if res != corr_res {
		t.Fatalf("Part 2 test failed: %d vs %d", res, corr_res)
	}

}
