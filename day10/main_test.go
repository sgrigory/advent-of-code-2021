package main

import (
	"testing"
)

const INPUT = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

func TestPart1(t *testing.T) {

	content := []byte(INPUT)
	res := run_part1(content)
	expected := 26397
	if res != expected {
		t.Errorf("error in part 1: expected %d, got %d", expected, res)
	}
}

func TestPart2(t *testing.T) {

	content := []byte(INPUT)
	res := run_part2(content)
	expected := 288957
	if res != expected {
		t.Errorf("error in part 2: expected %d, got %d", expected, res)
	}
}
