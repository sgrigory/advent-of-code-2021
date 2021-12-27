package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func get_rules(content []byte) (map[string]string, string) {

	rows := strings.Split(string(content), "\n")
	rules := map[string]string{}
	polymer := rows[0]
	for i, row := range rows[2:] {
		split_row := strings.Split(row, " -> ")
		fmt.Println(i, split_row)
		rules[split_row[0]] = split_row[1]
	}

	return rules, polymer

}

func run_part1(content []byte, steps int) int {

	rules, polymer := get_rules(content)

	max_num := make([]int, 0)
	min_num := make([]int, 0)

	for i := 0; i < steps; i++ {
		//fmt.Println(i, polymer)
		j := 0
		for j < len(polymer)-1 {
			if repl, ok := rules[polymer[j:j+2]]; ok {
				polymer = polymer[:j+1] + repl + polymer[j+1:]
				j++
			}
			j++
		}
		max_num_last, min_num_last := get_nums(polymer)
		max_num = append(max_num, max_num_last)
		min_num = append(min_num, min_num_last)
		if len(max_num) > 1 {
			fmt.Println(i, max_num[len(max_num)-1], min_num[len(max_num)-1],
				float64(max_num[len(max_num)-1])/float64(max_num[len(max_num)-2]),
				float64(min_num[len(min_num)-1])/float64(min_num[len(min_num)-2]),
				float64(max_num[len(min_num)-1]-min_num[len(min_num)-1])/float64(max_num[len(min_num)-2]-min_num[len(min_num)-2]))
		}
	}
	return max_num[len(max_num)-1] - min_num[len(min_num)-1]
}

func run_part2(content []byte, steps int) int {

	rules, polymer := get_rules(content)

	first_letter := polymer[0]
	last_letter := polymer[len(polymer)-1]

	pair_counts := map[string]int{}

	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i : i+2]
		pair_counts[pair]++
	}

	for i := 0; i < steps; i++ {
		new_pairs_count := map[string]int{}
		for lhs := range pair_counts {
			if rhs, ok := rules[lhs]; ok {
				//fmt.Printf("pair %s will be replaced with % and %")
				new_pairs_count[string(lhs[0])+rhs] += pair_counts[lhs]
				new_pairs_count[rhs+string(lhs[1])] += pair_counts[lhs]
			}
		}
		pair_counts = new_pairs_count
		//fmt.Println("step ", i, " counts ", pair_counts)
	}
	counts_f := map[rune]float64{}
	for key := range pair_counts {
		counts_f[rune(key[0])] += float64(pair_counts[key]) / 2
		counts_f[rune(key[1])] += float64(pair_counts[key]) / 2
	}

	counts_f[rune(first_letter)] -= 0.5
	counts_f[rune(last_letter)] -= 0.5

	counts_i := map[rune]int{}
	for key := range counts_f {
		counts_i[key] = int(counts_f[key])
	}

	max_num, min_num := get_min_max(counts_i)

	return max_num - min_num + 1
}

func get_nums(polymer string) (int, int) {
	var max_num int
	var min_num int
	counts := map[rune]int{}
	for _, c := range polymer {
		if _, ok := counts[c]; ok {
			counts[c]++
		} else {
			counts[c] = 1
		}

	}

	max_num, min_num = get_min_max(counts)

	return max_num, min_num
}

func get_min_max(counts map[rune]int) (int, int) {
	max_num := 0
	min_num := int(^uint(0) >> 1)
	for key := range counts {
		if counts[key] > max_num {
			max_num = counts[key]
		}
		if counts[key] < min_num {
			min_num = counts[key]
		}
	}
	return max_num, min_num
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	res_part1 := run_part2(content, 10)
	fmt.Println("part 1 answer:", res_part1)

	res_part2 := run_part2(content, 40)
	fmt.Println("part 2 answer:", res_part2)

}
