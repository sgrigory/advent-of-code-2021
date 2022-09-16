package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(input_split []string) int {
	var parsed_input = parse_input(input_split)
	sort.Ints(parsed_input)
	var s = 0.0
	med := float64(parsed_input[len(input_split)/2])
	for v := range parsed_input {
		s += math.Abs(med - float64(parsed_input[v]))
	}
	return int(s)
}

func part2(input_split []string) int {
	var parsed_input = parse_input(input_split)
	sort.Ints(parsed_input)
	min_vals := 1e10
	for i := parsed_input[0]; i <= parsed_input[len(parsed_input)-1]; i++ {
		vals := 0.0
		for _, s := range parsed_input {
			abs_dist := math.Abs(float64(s - i))
			addition := abs_dist * (abs_dist + 1) / 2
			vals += addition
		}
		if vals < min_vals {
			min_vals = vals
		}

	}

	return int(min_vals)
}

func parse_input(input_split []string) []int {

	vals := make([]int, len(input_split))

	for i := 0; i < len(input_split); i++ {
		res, _ := strconv.Atoi(strings.TrimSpace(input_split[i]))
		vals[i] = res
	}
	return vals
}

func main() {
	file, _ := os.ReadFile("input.txt")
	input := strings.Trim(string(file), " \n")

	input_split := strings.Split(input, ",")

	res_part1 := part1(input_split)
	fmt.Println("result of part 1: ", res_part1)

	res_part2 := part2(input_split)
	fmt.Println("result of part 2: ", res_part2)

}
