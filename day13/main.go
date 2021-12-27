package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	vert bool
	pos  int
}

func fill_field(content []byte) ([][]byte, []Rule) {

	input := strings.Split(string(content), "\n\n")

	rows := strings.Split(input[0], "\n")
	field := init_empty_field(2000, 2000)
	for _, row := range rows {
		row_split := strings.Split(row, ",")
		x, _ := strconv.Atoi(row_split[0])
		y, _ := strconv.Atoi(row_split[1])
		field[y][x] = '#'
	}

	fold_rules := strings.Split(input[1], "\n")

	pattern := regexp.MustCompile(`fold along (?P<dir>x|y)\=(?P<val>\d+)`)

	rules := make([]Rule, len(fold_rules))

	for i, rule := range fold_rules {

		match := pattern.FindStringSubmatch(rule)

		rules[i].vert = match[1] == "y"
		rules[i].pos, _ = strconv.Atoi(match[2])

	}

	return field, rules
}

func init_empty_field(size_x int, size_y int) [][]byte {

	field := make([][]byte, size_y)
	for i := range field {
		field[i] = make([]byte, size_x)
		for j := range field[i] {
			field[i][j] = '.'
		}
	}
	return field
}

type coords struct {
	i int
	j int
}

func show_field(field [][]byte) {
	for _, row := range field {
		for _, c := range row {
			fmt.Print(string(c))
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func fold_vert(field [][]byte, pos int) [][]byte {

	size_x := len(field[0])
	size_y := pos

	new_field := init_empty_field(size_x, size_y)
	for x := 0; x < size_x; x++ {
		for y := 0; y < size_y; y++ {
			if (field[y][x] == '#') || (field[2*pos-y][x] == '#') {
				new_field[y][x] = '#'
			} else {
				new_field[y][x] = '.'
			}
		}
	}
	return new_field
}

func fold_horiz(field [][]byte, pos int) [][]byte {

	size_x := pos
	size_y := len(field)
	new_field := init_empty_field(size_x, size_y)
	for x := 0; x < size_x; x++ {
		for y := 0; y < size_y; y++ {
			if (field[y][x] == '#') || (field[y][2*pos-x] == '#') {
				new_field[y][x] = '#'
			} else {
				new_field[y][x] = '.'
			}
		}
	}
	return new_field
}

func run_part1(content []byte, steps int) int {

	field, rules := fill_field(content)

	for i, rule := range rules {

		if rule.vert {
			field = fold_vert(field, rule.pos)
		} else {
			field = fold_horiz(field, rule.pos)
		}

		if i >= steps-1 {
			break
		}

	}
	if (len(field) < 100) && (len(field[0]) < 100) {
		show_field(field)
	}
	return count_hashes(field)
}

func count_hashes(field [][]byte) int {
	s := 0
	for _, row := range field {
		for _, c := range row {
			if c == '#' {
				s++
			}
		}
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

	res_part1 := run_part1(content, 1)
	fmt.Println("part 1 answer:", res_part1)

	res_part2 := run_part1(content, 100)
	fmt.Println("part 2 answer:", res_part2)

}
