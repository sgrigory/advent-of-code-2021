package main

import (
	"fmt"
	"os"
	"strings"
)

func get_freqs(input_split []string) []int {

	n_cols := len(input_split[0])
	freqs := make([]int, n_cols)
	for i := range freqs {
		freqs[i] = 0
	}
	for _, row := range input_split {
		if len(row) == 0 {
			continue
		}
		for j := range row {
			if row[j] == '1' {
				freqs[j] += 1
			}
		}

	}

	return freqs

}

func part1(input_split []string) int {
	n_rows := len(input_split)

	freqs := get_freqs(input_split)

	most_freq := 0
	least_freq := 0
	for i := range freqs {
		curr := 0
		if freqs[i] > n_rows/2 {
			curr = 1
		}

		most_freq = 2*most_freq + curr
		least_freq = 2*least_freq + 1 - curr
	}

	return most_freq * least_freq
}

func part2(input_split []string) int {

	most_freq := input_split
	least_freq := input_split
	pos := 0
	for len(most_freq) > 1 {
		most_freq = filter_most_freq(most_freq, pos, true)
		pos += 1
	}
	pos = 0
	for len(least_freq) > 1 {
		least_freq = filter_most_freq(least_freq, pos, false)
		pos += 1
	}
	return bin_2_int(most_freq[0]) * bin_2_int(least_freq[0])
}

func bin_2_int(s string) int {
	ret := 0
	for _, el := range s {
		val := 0
		if el == '1' {
			val = 1
		}
		ret = 2*ret + val
	}
	return ret
}

func filter_most_freq(rows []string, pos int, sign bool) []string {
	zeros := make([]string, 0)
	ones := make([]string, 0)
	for _, row := range rows {
		if row[pos] == '0' {
			zeros = append(zeros, row)
		} else {
			ones = append(ones, row)
		}

	}
	if (len(ones) >= len(zeros)) == sign {
		return ones
	} else {
		return zeros
	}
}

func main() {
	file, _ := os.ReadFile("day3_2021_input.txt")
	input := strings.Trim(string(file), " \n")

	input_split := strings.Split(input, "\n")

	res_part1 := part1(input_split)
	fmt.Println("result of part 1: ", res_part1)

	res_part2 := part2(input_split)
	fmt.Println("result of part 2: ", res_part2)

}
