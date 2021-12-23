package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const SIZE = 102
const SIZE2 = 1700

func run_part1(content []byte) int {

	pattern := regexp.MustCompile(`(?P<switch>on|off) x=(?P<x0>-?\d+)\.\.(?P<x1>-?\d+),y=(?P<y0>-?\d+)\.\.(?P<y1>-?\d+),z=(?P<z0>-?\d+)\.\.(?P<z1>-?\d+)`)

	field := init_field(SIZE)
	rows := strings.Split(string(content), "\n")
	for _, row := range rows {
		match := pattern.FindStringSubmatch(row)
		//fmt.Println(row, match)
		flag := match[1]
		x0, _ := strconv.Atoi(match[2])
		x1, _ := strconv.Atoi(match[3])
		y0, _ := strconv.Atoi(match[4])
		y1, _ := strconv.Atoi(match[5])
		z0, _ := strconv.Atoi(match[6])
		z1, _ := strconv.Atoi(match[7])
		if (x0 <= 50) && (x0 >= -50) {
			for x := x0; x <= x1; x++ {
				for y := y0; y <= y1; y++ {
					for z := z0; z <= z1; z++ {
						field[x+SIZE/2][y+SIZE/2][z+SIZE/2] = flag == "on"
					}
				}
			}
		}
	}

	s := 0
	for x := 0; x < SIZE; x++ {
		for y := 0; y < SIZE; y++ {
			for z := 0; z < SIZE; z++ {
				if field[x][y][z] {
					s++
				}
			}
		}
	}

	return s
}

func run_part2(content []byte) uint64 {

	pattern := regexp.MustCompile(`(?P<switch>on|off) x=(?P<x0>-?\d+)\.\.(?P<x1>-?\d+),y=(?P<y0>-?\d+)\.\.(?P<y1>-?\d+),z=(?P<z0>-?\d+)\.\.(?P<z1>-?\d+)`)

	field := init_field(SIZE2)
	rows := strings.Split(string(content), "\n")
	flag := make([]bool, len(rows))
	x0 := make([]int, len(rows))
	x1 := make([]int, len(rows))
	y0 := make([]int, len(rows))
	y1 := make([]int, len(rows))
	z0 := make([]int, len(rows))
	z1 := make([]int, len(rows))
	fmt.Println("parsing...")
	for i, row := range rows {
		match := pattern.FindStringSubmatch(row)
		//fmt.Println(row, match)
		flag[i] = match[1] == "on"
		x0[i], _ = strconv.Atoi(match[2])
		x1[i], _ = strconv.Atoi(match[3])
		y0[i], _ = strconv.Atoi(match[4])
		y1[i], _ = strconv.Atoi(match[5])
		z0[i], _ = strconv.Atoi(match[6])
		z1[i], _ = strconv.Atoi(match[7])
	}

	all_x := append(x0, x1...)
	all_y := append(y0, y1...)
	all_z := append(z0, z1...)
	for _, val := range all_x {
		all_x = append(all_x, val+1)
	}
	for _, val := range all_y {
		all_y = append(all_y, val+1)
	}
	for _, val := range all_z {
		all_z = append(all_z, val+1)
	}

	// -10 10 -> 0 1
	// 9 10 -> 0 1
	sort.Ints(all_x)
	sort.Ints(all_y)
	sort.Ints(all_z)

	coord_map_x := map[int]int{}
	coord_map_y := map[int]int{}
	coord_map_z := map[int]int{}

	for i := range all_x {
		//fmt.Println(i, all_x[i])
		coord_map_x[all_x[i]] = i
		coord_map_y[all_y[i]] = i
		coord_map_z[all_z[i]] = i
	}

	fmt.Println("lengths: ", len(coord_map_x), len(coord_map_y), len(coord_map_z))

	for i := range rows {
		for x := coord_map_x[x0[i]]; x <= coord_map_x[x1[i]]; x++ {
			for y := coord_map_y[y0[i]]; y <= coord_map_y[y1[i]]; y++ {
				for z := coord_map_z[z0[i]]; z <= coord_map_z[z1[i]]; z++ {
					field[x][y][z] = flag[i]
					if flag[i] {
						//fmt.Println(x, y, z, ": field", field[x][y][z])
					}
				}
			}
		}
	}

	fmt.Println("counting the volume...")

	var s uint64
	s = 0
	for x := 0; x < len(all_x); x++ {
		for y := 0; y < len(all_y); y++ {
			for z := 0; z < len(all_z); z++ {
				//fmt.Println("cell ", x, y, z, field[x][y][z])
				if field[x][y][z] {
					var x_size uint64
					var y_size uint64
					var z_size uint64
					if x < len(all_x)-1 {
						x_size = uint64(all_x[x+1] - all_x[x])
					} else {
						x_size = 1
					}
					if y < len(all_y)-1 {
						y_size = uint64(all_y[y+1] - all_y[y])
					} else {
						y_size = 1
					}
					if z < len(all_z)-1 {
						z_size = uint64(all_z[z+1] - all_z[z])
					} else {
						z_size = 1
					}
					//fmt.Println(x, y, z, "sizes: ", x_size, y_size, z_size)
					s += x_size * y_size * z_size
				}
			}
		}
	}

	return s
}

// type Rect struct {
// 	x0 int
// 	x1 int
// 	y0 int
// 	y1 int
// 	z0 int
// 	z1 int
// }

// func (r Rect) volume() int {
// 	return (r.x1 - r.x0 + 1) * (r.y1 - r.y0 + 1) * (r.z1 - r.z0 + 1)
// }

// func run_part2(content []byte) int {

// 	var rects []Rect

// 	rows := strings.Split(string(content), "\n")

// 	pattern := regexp.MustCompile(`(?P<switch>on|off) x=(?P<x0>-?\d+)\.\.(?P<x1>-?\d+),y=(?P<y0>-?\d+)\.\.(?P<y1>-?\d+),z=(?P<z0>-?\d+)\.\.(?P<z1>-?\d+)`)

// 	for _, row := range rows {
// 		match := pattern.FindStringSubmatch(row)
// 		//fmt.Println(row, match)
// 		flag := match[1] == "on"
// 		x0, _ := strconv.Atoi(match[2])
// 		x1, _ := strconv.Atoi(match[3])
// 		y0, _ := strconv.Atoi(match[4])
// 		y1, _ := strconv.Atoi(match[5])
// 		z0, _ := strconv.Atoi(match[6])
// 		z1, _ := strconv.Atoi(match[7])

// 		fmt.Println(x0, x1, y0, y1, z0, z1)

// 		new_rect := Rect{x0, x1, y0, y1, z0, z1}

// 		new_rects := []Rect{new_rect}
// 		for _, rect := range rects {
// 			generated_rects := get_intersection(rect, new_rect, flag)
// 			fmt.Println("got generated rects ", generated_rects)
// 			new_rects = append(new_rects, generated_rects...)
// 		}
// 		rects = new_rects
// 	}

// 	s := 0
// 	for _, r := range rects {
// 		fmt.Println("final rect ", r, " volume ", r.volume())
// 		s += r.volume()
// 	}

// 	return s
// }

// func get_intersection(rect1 Rect, rect2 Rect, flag bool) []Rect {
// 	// |   | |   |
// 	return []Rect{rect1, rect2}
// }

func init_field(size int) [][][]bool {

	field := make([][][]bool, size)
	for i := 0; i < size; i++ {
		field[i] = make([][]bool, size)
		for j := 0; j < size; j++ {
			field[i][j] = make([]bool, size)
			for k := 0; k < size; k++ {
				field[i][j][k] = false
			}
		}
	}
	return field
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
