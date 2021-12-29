package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type Tree struct {
	left   *Tree
	right  *Tree
	parent *Tree
	val    int
}

func run_part1(content []byte) int {

	rows := strings.Split(string(content), "\n")
	inputs := parse_many_trees(rows)
	total := sum_many_trees(inputs)

	return total.get_magnitude()
}

func parse_many_trees(rows []string) []*Tree {

	inputs := make([]*Tree, len(rows))
	for i, row := range rows {
		inputs[i] = tree_from_string(row, nil)
	}

	return inputs

}

func sum_many_trees(inputs []*Tree) *Tree {
	total := inputs[0]
	for _, input := range inputs[1:] {
		total = add_trees(total, input)
	}
	return total
}

func tree_from_string(input string, parent *Tree) *Tree {
	var node Tree
	if input[0] == '[' {
		split_pos := find_split(input)
		lhs := tree_from_string(input[1:split_pos], &node)
		rhs := tree_from_string(input[split_pos+1:len(input)-1], &node)
		node = Tree{lhs, rhs, parent, 0}
	} else {
		val, _ := strconv.Atoi(input)
		node = Tree{nil, nil, parent, val}
	}
	return &node
}

func find_split(input string) int {
	stack := make([]byte, 0)
	for i, c := range input {
		switch c {
		case '[':
			stack = append(stack, '[')
		case ']':
			stack = stack[:len(stack)-1]
		case ',':
			if len(stack) == 1 {
				return i
			}
		}
	}
	log.Fatal("haven't found a split ", input)
	return -1
}

func (self Tree) get_magnitude() int {
	if (self.left == nil) && (self.right == nil) {
		return self.val
	}
	return 3*self.left.get_magnitude() + 2*self.right.get_magnitude()
}

func (self Tree) String() string {
	if self.left == nil {
		return strconv.Itoa(self.val)
	}
	return "[" + self.left.String() + "," + self.right.String() + "]"
}

func add_trees(lhs *Tree, rhs *Tree) *Tree {
	root := Tree{lhs, rhs, nil, 0}
	lhs.parent = &root
	rhs.parent = &root
	action := true
	for action {
		action = reduce_tree_both(&root, 0)
	}
	return &root
}

func (t *Tree) is_left_child() bool {
	return t.parent.left == t
}

func reduce_tree_both(t *Tree, depth int) bool {

	action := reduce_tree(t, 0, true)
	action = action || reduce_tree(t, 0, false)
	return action
}

func reduce_tree(t *Tree, depth int, is_explode bool) bool {
	if t == nil {
		return false
	}
	is_pair := (t.left != nil) && (t.left.left == nil) && (t.right.left == nil)
	if (depth == 4) && is_pair && is_explode {
		explode(t)
		return true
	}

	is_number := (t.left == nil)
	if is_number && (t.val >= 10) && !is_explode {
		split(t)
		return true
	}

	action_left := reduce_tree(t.left, depth+1, is_explode)

	if action_left {
		return true
	}

	action_right := reduce_tree(t.right, depth+1, is_explode)

	if action_right {
		return true
	}

	return false
}

func explode(t *Tree) {

	val_left := t.left.val
	val_right := t.right.val
	t.left = nil
	t.right = nil
	t.val = 0

	next_right := find_next_right(t)
	if next_right != nil {
		next_right.val += val_right
	}

	next_left := find_next_left(t)
	if next_left != nil {
		next_left.val += val_left
	}

}

func split(t *Tree) {
	val_left := int(math.Floor(float64(t.val) / 2))
	val_right := int(math.Ceil(float64(t.val) / 2))
	t.left = &Tree{nil, nil, t, val_left}
	t.right = &Tree{nil, nil, t, val_right}
}

func find_next_left(t *Tree) *Tree {
	t0 := t
	for (t != nil) && (t.parent != nil) && t.is_left_child() {
		t = t.parent
	}

	if t.parent == nil {
		return nil
	}
	t = t.parent.left

	if t == nil {
		return nil
	}
	for t.right != nil {
		t = t.right
	}
	if t == t0 {
		return nil
	}
	return t
}

func find_next_right(t *Tree) *Tree {
	t0 := t
	for (t != nil) && (t.parent != nil) && !t.is_left_child() {
		t = t.parent
	}

	if t.parent == nil {
		return nil
	}
	t = t.parent.right

	if t == nil {
		return nil
	}
	for t.left != nil {
		t = t.left
	}
	if t == t0 {
		return nil
	}
	return t
}

func (t *Tree) copy() *Tree {
	new_tree := Tree{nil, nil, t.parent, t.val}

	var right *Tree = nil
	if t.right != nil {
		right = t.right.copy()
		right.parent = &new_tree
	}

	var left *Tree = nil
	if t.left != nil {
		left = t.left.copy()
		left.parent = &new_tree
	}

	new_tree.left = left
	new_tree.right = right

	return &new_tree
}

func run_part2(content []byte) int {

	rows := strings.Split(string(content), "\n")
	inputs := parse_many_trees(rows)

	max_mag := 0
	for i := range inputs {
		for j := range inputs {
			if i != j {
				sum_pair := add_trees(inputs[i].copy(), inputs[j].copy())
				mag_sum_pair := sum_pair.get_magnitude()
				if mag_sum_pair > max_mag {
					max_mag = mag_sum_pair
				}
			}
		}
	}

	return max_mag
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
