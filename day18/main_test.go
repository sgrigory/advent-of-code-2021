package main

import (
	"strings"
	"testing"
)

const INPUT1 = `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`

func TestPart1(t *testing.T) {

	expected := 4140
	content := []byte(INPUT1)
	res := run_part1(content)
	if res != expected {
		t.Errorf("error in part 1: expected %d, got %d", expected, res)
	}
}

func TestAddManyTrees(t *testing.T) {

	expected := []string{"[[[[1,1],[2,2]],[3,3]],[4,4]]",
		"[[[[3,0],[5,3]],[4,4]],[5,5]]",
		"[[[[5,0],[7,4]],[5,5]],[6,6]]",
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"}
	input := []string{`[1,1]
[2,2]
[3,3]
[4,4]`,
		`[1,1]
[2,2]
[3,3]
[4,4]
[5,5]`,
		`[1,1]
[2,2]
[3,3]
[4,4]
[5,5]
[6,6]`,
		`[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`}

	for i := range input {

		rows := strings.Split(input[i], "\n")
		input_trees := parse_many_trees(rows)
		total := sum_many_trees(input_trees)

		if total.String() != expected[i] {
			t.Errorf("error in part 1: expected %s, got %s", expected[i], total)
		}

	}

}

func TestAddTrees(t *testing.T) {

	expected := []string{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"}
	input := [][]string{[]string{"[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]"},
		[]string{"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]", "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]"}}
	for i := range expected {
		tree1 := tree_from_string(input[i][0], nil)
		tree2 := tree_from_string(input[i][1], nil)
		res := add_trees(tree1, tree2)
		if res.String() != expected[i] {
			t.Errorf("error in part 1: expected %s, got %s", expected[i], res)
		}
	}
}

func TestReduceTree(t *testing.T) {

	inputs := []string{"[[[[[9,8],1],2],3],4]", "[7,[6,[5,[4,[3,2]]]]]", "[[6,[5,[4,[3,2]]]],1]",
		"[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]", "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"}
	expected := []string{"[[[[0,9],2],3],4]", "[7,[6,[5,[7,0]]]]", "[[6,[5,[7,0]]],3]",
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"}
	for i := range inputs {
		tree := tree_from_string(inputs[i], nil)
		reduce_tree_both(tree, 0)
		res := tree.String()
		if res != expected[i] {
			t.Errorf("error in part 1: started from %s expected %s, got %s", inputs[i], expected[i], res)
		}
	}

}

func TestGetMagnitude(t *testing.T) {
	inputs := []string{"[[1,2],[[3,4],5]]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", "[[[[1,1],[2,2]],[3,3]],[4,4]]"}
	expected := []int{143, 1384, 445}
	for i := range inputs {
		//content := []byte(inputs[i])
		tree := tree_from_string(inputs[i], nil)
		res := tree.get_magnitude()
		if res != expected[i] {
			t.Errorf("error in part 1: expected %d, got %d", expected, res)
		}
	}
}

func TestFindSplit(t *testing.T) {
	inputs := []string{"[[1,2],[[3,4],5]]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", "[[[[1,1],[2,2]],[3,3]],[4,4]]"}
	expected := []int{6, 26, 22}
	for i := range inputs {
		//content := []byte(inputs[i])
		res := find_split(inputs[i])
		if res != expected[i] {
			t.Errorf("error in part 1: expected %d, got %d", expected, res)
		}
	}
}

func TestPart2(t *testing.T) {

	expected := 3993
	content := []byte(INPUT1)
	res := run_part2(content)
	if res != expected {
		t.Errorf("error in part 2: expected %d, got %d", expected, res)
	}
}
