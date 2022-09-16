package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(input_split []string) int {
	x, y := 0, 0
	for _, command := range input_split {
		fmt.Println(command)
		if len(command) == 0 {
			continue
		}
		split_command := strings.Split(command, " ")
		steps, _ := strconv.Atoi(split_command[1])
		switch split_command[0] {
		case "forward":
			x += steps
		case "up":
			y -= steps
		case "down":
			y += steps
		}

	}
	return x * y
}

func part2(input_split []string) int {
	x, y, a := 0, 0, 0
	for _, command := range input_split {
		fmt.Println(command)
		if len(command) == 0 {
			continue
		}
		split_command := strings.Split(command, " ")
		steps, _ := strconv.Atoi(split_command[1])
		switch split_command[0] {
		case "forward":
			{
				x += steps
				y += a * steps
			}
		case "up":
			a -= steps
		case "down":
			a += steps
		}

	}
	return x * y
}

func main() {
	file, _ := os.ReadFile("day2_2021_input.txt")
	input := string(file)

	input_split := strings.Split(input, "\n")

	res_part1 := part1(input_split)
	fmt.Println("result of part 1: ", res_part1)

	res_part2 := part2(input_split)
	fmt.Println("result of part 1: ", res_part2)

}
