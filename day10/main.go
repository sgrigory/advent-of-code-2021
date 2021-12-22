package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

var SCORES = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
var CLOSE = map[rune]rune{'}': '{', '>': '<', ']': '[', ')': '('}
var SCORES_PART2 = map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}

func run_part1(content []byte) int {

	rows := strings.Split(string(content), "\n")
	score := 0
	for _, row := range rows {
		score += parse_row(row)
	}

	return score
}

func parse_row(row string) int {

	stack := make([]rune, 0)
	for _, c := range row {

		if (c == '{') || (c == '(') || (c == '[') || (c == '<') {
			stack = append(stack, c)
		} else {

			if (len(stack) > 0) && (stack[len(stack)-1] == CLOSE[c]) {
				stack = stack[:len(stack)-1]
			} else {
				return SCORES[c]
			}
		}

	}
	return 0
}

func run_part2(content []byte) int {

	rows := strings.Split(string(content), "\n")
	scores := make([]int, 0)
	for _, row := range rows {
		row_score := parse_row_part2(row)
		if row_score > 0 {
			scores = append(scores, row_score)
		}
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

func parse_row_part2(row string) int {

	stack := make([]rune, 0)
	for _, c := range row {

		if (c == '{') || (c == '(') || (c == '[') || (c == '<') {
			stack = append(stack, c)
		} else {

			if (len(stack) > 0) && (stack[len(stack)-1] == CLOSE[c]) {
				stack = stack[:len(stack)-1]
			} else {
				return 0
			}
		}

	}
	score := 0
	for len(stack) > 0 {
		last_idx := len(stack) - 1
		score = 5*score + SCORES_PART2[stack[last_idx]]
		stack = stack[:last_idx]
	}
	return score
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	res_part1 := run_part1(content)
	fmt.Println("part 1 answer:", res_part1)

	res_part2 := run_part2(content)
	fmt.Println("part 2 answer:", res_part2)

}
