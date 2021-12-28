package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

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

const BORDER = 400

func show_field(field [][]byte) {
	for _, row := range field {
		for _, c := range row {
			fmt.Print(string(c))
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func parse_input(content []byte) ([]byte, [][]byte) {

	rows := strings.Split(string(content), "\n")

	alg_aug := []byte(rows[0])
	rows = rows[2:]

	field := init_empty_field(len(rows)+BORDER, len(rows[0])+BORDER)
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			field[i+BORDER/2][j+BORDER/2] = rows[i][j]
		}
	}

	return alg_aug, field

}

func run_part1(content []byte, steps int) int {

	aug_alg, field := parse_input(content)

	for step := 0; step < steps; step++ {
		new_field := init_empty_field(len(field[0]), len(field))
		pad := 1
		for i := pad; i < len(field)-pad; i++ {
			for j := pad; j < len(field[i])-pad; j++ {

				bits := make([]byte, 3)
				copy(bits, field[i-1][j-1:j+2])
				bits = append(bits, field[i][j-1:j+2]...)
				bits = append(bits, field[i+1][j-1:j+2]...)
				idx := get_idx(bits)
				new_field[i][j] = aug_alg[idx]

			}
		}

		field = new_field
		field = field[1 : len(field)-1]
		for i := range field {
			field[i] = field[i][1 : len(field[i])-1]
		}
	}

	return count_hashes(field)
}

func get_idx(bits []byte) int {
	res := 0
	for _, c := range bits {
		dig := 0
		if c == '#' {
			dig = 1
		}
		res = 2*res + dig
	}
	return res
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

	res_part1 := run_part1(content, 2)
	fmt.Println("part 1 answer:", res_part1)

	res_part2 := run_part1(content, 50)
	fmt.Println("part 2 answer:", res_part2)

}
