package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func run_part1(content []byte) int {

	rows := strings.Split(string(content), "\n")

	field := init_field(rows)
	total_risk := 0

	for i, row := range field {
		for j := range row {

			lower_then_above := (i == 0) || (field[i-1][j] > field[i][j])
			lower_then_below := (i == len(field)-1) || (field[i+1][j] > field[i][j])
			lower_then_left := (j == 0) || (field[i][j-1] > field[i][j])
			lower_then_right := (j == len(row)-1) || (field[i][j+1] > field[i][j])

			if lower_then_above && lower_then_below && lower_then_left && lower_then_right {
				total_risk += field[i][j] + 1
			}

		}
	}

	return total_risk
}

func run_part2(content []byte) int {

	rows := strings.Split(string(content), "\n")

	field := init_field(rows)
	var basin_sizes []int

	num_not_9 := 0

	for i, row := range field {
		for j := range row {
			if row[j] != 9 {
				num_not_9++
			}

			lower_then_above := (i == 0) || (field[i-1][j] > field[i][j])
			lower_then_below := (i == len(field)-1) || (field[i+1][j] > field[i][j])
			lower_then_left := (j == 0) || (field[i][j-1] > field[i][j])
			lower_then_right := (j == len(row)-1) || (field[i][j+1] > field[i][j])

			if lower_then_above && lower_then_below && lower_then_left && lower_then_right {
				basin_size := find_basin_size(i, j, field)

				basin_sizes = append(basin_sizes, -basin_size)

			}

		}
	}

	s := 0
	for _, sz := range basin_sizes {
		s += sz
	}

	fmt.Printf("sum of basin sizes = %d, number of non-9 = %d\n", -s, num_not_9)

	sort.Ints(basin_sizes)

	return -basin_sizes[0] * basin_sizes[1] * basin_sizes[2]
}

func init_field(rows []string) [][]int {

	field := make([][]int, len(rows))
	for i := range field {
		field[i] = make([]int, len(rows[0]))
		for j := range field[i] {
			field[i][j] = int(rows[i][j]) - int('0')
		}
	}
	return field
}

type Coords struct {
	i int
	j int
}

type CoordsSet map[Coords]struct{}

func find_basin_size(i int, j int, field [][]int) int {

	basin := CoordsSet{}
	new_cells := CoordsSet{Coords{i, j}: struct{}{}}

	for {

		for cell := range new_cells {
			if _, ok := basin[cell]; ok {
				delete(new_cells, cell)
			}
		}

		for cell := range new_cells {
			one_cell_set := CoordsSet{cell: struct{}{}}
			cell_neighbours := find_neighbours(one_cell_set, len(field), len(field[0]))
			drop := field[cell.i][cell.j] == 9
			for neighbour := range cell_neighbours {
				_, in_basin := basin[neighbour]
				drop = drop || ((!in_basin) && (field[neighbour.i][neighbour.j] < field[cell.i][cell.j]))
			}
			if drop {
				delete(new_cells, cell)
			}

		}

		if len(new_cells) == 0 {
			break
		}

		for cell := range new_cells {
			basin[cell] = struct{}{}
		}

		new_cells = find_neighbours(new_cells, len(field), len(field[0]))

	}
	return len(basin)
}

func find_neighbours(cells CoordsSet, i_size int, j_size int) CoordsSet {
	neighbours := make(CoordsSet)
	for cell := range cells {
		if cell.i > 0 {
			neighbours[Coords{cell.i - 1, cell.j}] = struct{}{}
		}
		if cell.i < i_size-1 {
			neighbours[Coords{cell.i + 1, cell.j}] = struct{}{}
		}
		if cell.j > 0 {
			neighbours[Coords{cell.i, cell.j - 1}] = struct{}{}
		}
		if cell.j < j_size-1 {
			neighbours[Coords{cell.i, cell.j + 1}] = struct{}{}
		}
	}
	return neighbours
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
