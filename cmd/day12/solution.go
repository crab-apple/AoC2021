package day12

import (
	"fmt"
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/utils"
	"slices"
	"strings"
)

func SolvePart1(s string) int {
	connections := input.Read(s, func(s string) utils.Pair[string] {
		return utils.Pair[string]{Left: strings.Split(s, "-")[0], Right: strings.Split(s, "-")[1]}
	})
	graph := utils.NewUndirectedGraph(connections)

	fmt.Println(graph)

	path := Path{utils.Single(graph.Nodes(), func(node *utils.GraphNode) bool {
		return node.Name == "start"
	})}

	paths := findAllEndingPaths(path)

	return len(paths)
}

func findAllEndingPaths(path Path) []Path {
	if path.Last().Name == "end" {
		return []Path{path}
	}
	r := make([]Path, 0)
	for _, connection := range path.Last().Connections {
		if isBigCave(connection) || !slices.Contains(path, connection) {
			r = append(r, findAllEndingPaths(append(path, connection))...)
		}
	}
	return r
}

func SolvePart2(s string) int {
	return 0
}

type Path []*utils.GraphNode

func (p Path) Last() *utils.GraphNode {
	return p[len(p)-1]
}

func isBigCave(node *utils.GraphNode) bool {
	return utils.FirstChar(node.Name) >= 'A' && utils.FirstChar(node.Name) <= 'Z'
}
