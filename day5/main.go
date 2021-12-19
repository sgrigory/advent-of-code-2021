package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const SIZE = 1000

func init_field() [][]int {

	field := make([][]int, SIZE)
	for i := 0; i < SIZE; i++ {
		field[i] = make([]int, SIZE)
		for j := range field[i] {
			field[i][j] = 0
		}
	}
	return field
}

func count_overlaps(field [][]int) int {
	s := 0
	for _, row := range field {
		for _, c := range row {
			if c > 1 {
				s += 1
			}
		}
	}
	return s
}

func get_expr() *regexp.Regexp {
	return regexp.MustCompile(`(?P<x0>\d+),(?P<y0>\d+) -> (?P<x1>\d+),(?P<y1>\d+)`)
}

func run_part1(rows []string) int {

	field := init_field()
	expr := get_expr()

	for _, row := range rows {
		match := expr.FindStringSubmatch(row)
		x0, _ := strconv.Atoi(match[1])
		y0, _ := strconv.Atoi(match[2])
		x1, _ := strconv.Atoi(match[3])
		y1, _ := strconv.Atoi(match[4])
		if y0 > y1 {
			y0, y1 = y1, y0
		}
		if x0 > x1 {
			x0, x1 = x1, x0
		}

		if x0 == x1 {
			for y := y0; y <= y1; y++ {
				field[x0][y] += 1
			}
		} else if y0 == y1 {
			for x := x0; x <= x1; x++ {
				field[x][y0] += 1
			}
		}

	}
	return count_overlaps(field)
}

func run_part2(rows []string) int {

	field := init_field()
	expr := get_expr()

	for _, row := range rows {
		match := expr.FindStringSubmatch(row)
		x0, _ := strconv.Atoi(match[1])
		y0, _ := strconv.Atoi(match[2])
		x1, _ := strconv.Atoi(match[3])
		y1, _ := strconv.Atoi(match[4])

		step_y := 1
		if y0 > y1 {
			step_y = -1
		}
		if y0 == y1 {
			step_y = 0
		}

		step_x := 1
		if x0 > x1 {
			step_x = -1
		}
		if x0 == x1 {
			step_x = 0
		}

		range_x := (x1 - x0) * step_x
		range_y := (y1 - y0) * step_y

		max_range := range_x
		if range_y > range_x {
			max_range = range_y
		}

		for i := 0; i <= max_range; i++ {
			field[x0+i*step_x][y0+i*step_y] += 1
		}

	}
	return count_overlaps(field)
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(content), "\n")

	res_part1 := run_part1(rows)
	fmt.Println("part 1 answer:", res_part1)

	res_part2 := run_part2(rows)
	fmt.Println("part 2 answer:", res_part2)

}
