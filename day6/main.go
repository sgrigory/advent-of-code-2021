package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const SIZE = 1000
const N_STATES = 9

func run_part1(content []byte, n_iter int) int {

	counts := make([]int, N_STATES)
	for i := range counts {
		counts[i] = 0
	}
	cycles := strings.Split(string(content), ",")
	for _, c := range cycles {
		cycle, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal(err)
		}
		counts[cycle] += 1
	}

	for i := 0; i < n_iter; i++ {

		old_zero := counts[0]
		for j := 1; j < N_STATES; j++ {
			counts[j-1] = counts[j]
		}

		counts[N_STATES-1] = old_zero
		counts[N_STATES-1-2] += old_zero

	}

	return count_total(counts)
}

func count_total(counts []int) int {

	s := 0
	for _, count := range counts {
		s += count
	}
	return s
}

func run_part2(content []byte) int {

	return 0
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	res_part1 := run_part1(content, 80)
	fmt.Println("part 1 answer:", res_part1)

	res_part2 := run_part1(content, 256)
	fmt.Println("part 2 answer:", res_part2)

}
