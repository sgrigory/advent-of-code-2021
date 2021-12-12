package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Node struct {
	id          string
	connections map[string]bool
	visited     int
}

type NodeMap map[string]*Node

func add_connection(nodes NodeMap, v1 string, v2 string) NodeMap {
	p, ok := nodes[v1]
	if ok {
		p.connections[v2] = true
	} else {
		nodes[v1] = &Node{id: v1, connections: map[string]bool{v2: true}}
	}
	return nodes
}

func search_paths_1(nodes NodeMap, start_node *Node) int {

	if start_node.id == "end" {
		return 1
	}

	if (start_node.visited > 0) && (start_node.id[0] >= 'a') && (start_node.id[0] <= 'z') {
		return 0
	}

	start_node.visited = 1
	s := 0
	for key, _ := range start_node.connections {
		s += search_paths_1(nodes, nodes[key])
	}
	start_node.visited = 0

	return s

}

func search_paths_2(nodes NodeMap, start_node *Node, used_twice bool) int {

	if start_node.id == "end" {
		return 1
	}

	is_lowercase := (start_node.id[0] >= 'a') && (start_node.id[0] <= 'z')

	if (start_node.visited > 0) && ((start_node.id == "start") || (is_lowercase && used_twice)) {
		return 0
	}
	if (start_node.visited > 1) && is_lowercase {
		return 0
	}

	used_twice_new := used_twice || ((start_node.visited == 1) && is_lowercase)

	start_node.visited += 1
	s := 0
	for key, _ := range start_node.connections {
		s += search_paths_2(nodes, nodes[key], used_twice_new)
	}
	start_node.visited -= 1
	return s

}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(content), "\n")

	nodes := NodeMap{}

	for _, row := range rows {
		rows_parts := strings.Split(row, "-")
		v1, v2 := rows_parts[0], rows_parts[1]

		// Add nodes v1 and v2 if necessary
		nodes = add_connection(nodes, v1, v2)
		nodes = add_connection(nodes, v2, v1)

	}

	fmt.Println(nodes)

	res_part1 := search_paths_1(nodes, nodes["start"])
	fmt.Println("part 1 answer:", res_part1)

	res_part2 := search_paths_2(nodes, nodes["start"], false)
	fmt.Println("part 2 answer:", res_part2)

}
