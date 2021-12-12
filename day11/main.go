package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(content), "\n")

	size_x, size_y := len(rows[0]), len(rows)

	fmt.Println(size_x, size_y)

	var energy = make([][]int, size_y)
	var flashed = make([][]bool, size_y)

	for i, row := range rows {
		energy[i] = make([]int, size_x)
		flashed[i] = make([]bool, size_x)
		for j, c := range row {
			energy[i][j] = int(c - '0')
		}
	}
	fmt.Println(energy)

	tot_flashes := 0

	for step := 0; true; step++ {

		for i := range flashed {
			for j := range flashed[i] {
				flashed[i][j] = false
			}
		}

		for _, row := range energy {
			for i := range row {
				row[i] += 1
			}
		}

		flash := true

		for flash {

			flash = false

			for i := range energy {
				for j := range energy[i] {
					if (energy[i][j] > 9) && !flashed[i][j] {
						tot_flashes += 1
						flash = true
						flashed[i][j] = true
						for y := i - 1; y < i+2; y++ {
							for x := j - 1; x < j+2; x++ {
								if (x >= 0) && (y >= 0) && (x < size_x) && (y < size_y) && ((x != j) || (y != i)) {
									energy[y][x] += 1
								}
							}
						}

					}
				}
			}

		}

		all_flashed := true
		for i := range flashed {
			for j := range flashed[i] {

				if flashed[i][j] {
					energy[i][j] = 0
				} else {
					all_flashed = false
				}

			}
		}

		if step == 99 {
			fmt.Println("part 1 answer:", tot_flashes)
		}
		if all_flashed {
			fmt.Println("part 2 answer:", step+1)
			break
		}
	}

}
