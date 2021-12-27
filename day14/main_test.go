package main

import (
	"testing"
)

const INPUT = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

func TestPart1(t *testing.T) {

	expected := 1588
	content := []byte(INPUT)
	res := run_part2(content, 10)
	if res != expected {
		t.Errorf("error in part 1: expected %d, got %d", expected, res)
	}
}

func TestPart2(t *testing.T) {

	content := []byte(INPUT)
	res := run_part2(content, 40)
	expected := 2188189693529
	if res != expected {
		t.Errorf("error in part 2: expected %d, got %d", expected, res)
	}
}
