package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const SIZE = 10
const TAREGT_SCORE = 1000
const DICE_SIZE = 100

const TAREGT_SCORE2 = 21

type Field [][][][]int

func get_pos(content []byte) (int, int) {
	rows := strings.Split(string(content), "\n")
	row1_split := strings.Split(rows[0], ":")
	row2_split := strings.Split(rows[1], ":")
	pos1, _ := strconv.Atoi(strings.TrimSpace(row1_split[1]))
	pos2, _ := strconv.Atoi(strings.TrimSpace(row2_split[1]))
	pos1--
	pos2--
	return pos1, pos2
}

func run_part1(content []byte) int {

	pos1, pos2 := get_pos(content)

	fmt.Println("initial positions ", pos1, pos2)
	score1 := 0
	score2 := 0

	dice := 1
	step := 0

	for (score1 < TAREGT_SCORE) && (score2 < TAREGT_SCORE) {

		a := dice
		dice = (dice + 1) % DICE_SIZE
		b := dice
		dice = (dice + 1) % DICE_SIZE
		c := dice
		dice = (dice + 1) % DICE_SIZE
		if step%2 == 0 {
			pos1 = (pos1 + a + b + c) % SIZE
			score1 += pos1 + 1
		} else {
			pos2 = (pos2 + a + b + c) % SIZE
			score2 += pos2 + 1
		}
		step += 1
		//fmt.Printf("step %d: %d+%d+%d. positions: %d, %d.scores: %d, %d\n", step, a, b, c, pos1, pos2, score1, score2)
	}

	losing_score := score1
	if score2 < score1 {
		losing_score = score2
	}

	return step * 3 * losing_score
}

func run_part2(content []byte) int {

	pos1, pos2 := get_pos(content)

	field := init_field()
	step := 0
	var new_won1 int
	var new_won2 int
	not_finished := true
	won1, won2 := 0, 0
	field[0][0][pos1][pos2] = 1
	for not_finished {
		field, new_won1, new_won2, not_finished = make_step(field, step)
		won1 += new_won1
		won2 += new_won2
		step += 1
	}
	if won1 > won2 {
		return won1
	}
	return won2
}

func init_field() Field {
	//  score1 x score2 x pos1 x pos2, 21x21x10x10
	field := make(Field, TAREGT_SCORE2)
	for i := range field {
		field[i] = make([][][]int, TAREGT_SCORE2)
		for j := range field[i] {
			field[i][j] = make([][]int, SIZE)
			for k := range field[i][j] {
				field[i][j][k] = make([]int, SIZE)
				for l := range field[i][j][k] {
					field[i][j][k][l] = 0
				}
			}
		}
	}
	return field
}

func make_step(field Field, step int) (Field, int, int, bool) {
	//  score1 x score2 x pos1 x pos2, 21x21x10x10
	new_field := init_field()

	won1, won2 := 0, 0
	not_finished := false

	for score1 := range field {
		for score2 := range field[score1] {
			for pos1 := range field[score1][score2] {
				for pos2 := range field[score1][score2][pos1] {
					val := field[score1][score2][pos1][pos2]
					if val == 0 {
						continue
					}
					not_finished = true
					for dice1 := 1; dice1 <= 3; dice1++ {
						for dice2 := 1; dice2 <= 3; dice2++ {
							for dice3 := 1; dice3 <= 3; dice3++ {
								total_dice := dice1 + dice2 + dice3
								if step%2 == 0 {
									pos1_new := (pos1 + total_dice) % SIZE
									score1_new := score1 + pos1_new + 1
									if score1_new < TAREGT_SCORE2 {
										new_field[score1_new][score2][pos1_new][pos2] += val
									} else {
										won1 += val
									}
								} else {

									pos2_new := (pos2 + total_dice) % SIZE
									score2_new := score2 + pos2_new + 1
									if score2_new < TAREGT_SCORE2 {
										new_field[score1][score2_new][pos1][pos2_new] += val
									} else {
										won2 += val
									}

								}
							}
						}
					}
				}
			}
		}
	}

	return new_field, won1, won2, not_finished
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
