package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func fill_field(content []byte) [][]byte {

	rows := strings.Split(string(content), "\n")
	field := init_field(rows)
	for i, row := range rows {
		for j, c := range row {
			field[i][j] = byte(c)
		}
	}

	return field
}

func init_field(rows []string) [][]byte {

	field := make([][]byte, len(rows))
	for i := range field {
		field[i] = make([]byte, len(rows[0]))
		for j := range field[i] {
			field[i][j] = rows[i][j]
		}
	}
	return field
}

func copy_field(content []byte, field [][]byte) [][]byte {

	old_field := fill_field(content)
	for i := range old_field {
		for j := range old_field[i] {
			old_field[i][j] = field[i][j]
		}
	}
	return old_field
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

func run_part1(content []byte) int {

	field := fill_field(content)
	changes := true
	step := 0
	for changes {
		old_field := copy_field(content, field)

		changes = false
		for i := range old_field {
			for j := range old_field[i] {
				if old_field[i][j] == '>' {
					new_j := (j + 1) % len(field[i])
					if old_field[i][new_j] == '.' {
						field[i][new_j] = old_field[i][j]
						field[i][j] = '.'
						changes = true
					}
				}
			}
		}

		old_field = copy_field(content, field)

		for i := range old_field {
			for j := range old_field[i] {
				if old_field[i][j] == 'v' {
					new_i := (i + 1) % len(field)
					if old_field[new_i][j] == '.' {
						field[new_i][j] = old_field[i][j]
						field[i][j] = '.'
						changes = true
					}
				}
			}
		}
		step++
		// if step > 100 {
		// 	break
		// }
	}
	return step
}

func run_part2(content []byte) int {
	return 0
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
