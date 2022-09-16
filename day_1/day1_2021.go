package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse_input(input_split []string) []int {

	vals := make([]int, len(input_split))

	for i := 0; i < len(input_split); i++ {
		res, _ := strconv.Atoi(strings.TrimSpace(input_split[i]))
		vals[i] = res
	}
	return vals
}

func part1(input_split []string) int {
	s := 0

	var vals = parse_input(input_split)

	for i := 1; i < len(input_split)-1; i++ {
		if vals[i] > vals[i-1] {
			s++
		}

	}
	return s
}

func part2(input_split []string) int {
	s := 0

	var vals = parse_input(input_split)

	for i := 3; i < len(input_split)-1; i++ {
		if vals[i]+vals[i-1]+vals[i-2] > vals[i-1]+vals[i-2]+vals[i-3] {
			s++
		}

	}
	return s
}

func main() {
	file, _ := os.ReadFile("day1_2021_input.txt")
	input := string(file)

	input_split := strings.Split(input, "\n")

	res_part1 := part1(input_split)
	fmt.Println("result of part 1: ", res_part1)

	res_part2 := part2(input_split)
	fmt.Println("result of part 2: ", res_part2)
}
