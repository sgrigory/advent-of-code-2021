package main

import (
	"testing"
)

func TestPart1(t *testing.T) {

	content := []byte("3,4,3,1,2")
	res := run_part1(content, 18)
	expected := 26
	if res != expected {
		t.Errorf("error in part 1: expected %d, got %d", expected, res)
	}

	res = run_part1(content, 80)
	expected = 5934
	if res != expected {
		t.Errorf("error in part 1: expected %d, got %d", expected, res)
	}

	res = run_part1(content, 256)
	expected = 26984457539
	if res != expected {
		t.Errorf("error in part 1: expected %d, got %d", expected, res)
	}

}
