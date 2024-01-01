package day12

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/utils"
	"slices"
	"strings"
)

func SolvePart1(s string) int {
	return solve(s, 0)
}

func SolvePart2(s string) int {
	return solve(s, 1)
}

func solve(s string, numRevisitsAllowed int) int {
	connections := input.Read(s, func(s string) utils.Pair[string] {
		return utils.Pair[string]{Left: strings.Split(s, "-")[0], Right: strings.Split(s, "-")[1]}
	})
	graph := utils.NewUndirectedGraph(connections)

	path := Path{
		[]*utils.GraphNode{
			utils.Single(graph.Nodes(), func(node *utils.GraphNode) bool {
				return node.Name == "start"
			}),
		},
		numRevisitsAllowed,
	}

	paths := findAllEndingPaths(path)

	return len(paths)
}

func findAllEndingPaths(path Path) []Path {
	if path.Last().Name == "end" {
		return []Path{path}
	}
	r := make([]Path, 0)
	for _, connection := range path.Last().Connections {
		if path.canVisit(connection) {
			r = append(r, findAllEndingPaths(path.plus(connection))...)
		}
	}
	return r
}

type Path struct {
	nodes        []*utils.GraphNode
	revisitsLeft int
}

func (p *Path) hasVisited(node *utils.GraphNode) bool {
	return slices.Contains(p.nodes, node)
}

func (p *Path) plus(node *utils.GraphNode) Path {
	var revisitsLeftAfter int
	if isBigCave(node) || !p.hasVisited(node) {
		revisitsLeftAfter = p.revisitsLeft
	} else {
		revisitsLeftAfter = p.revisitsLeft - 1
	}
	return Path{append(p.nodes, node), revisitsLeftAfter}
}

func (p *Path) Last() *utils.GraphNode {
	return p.nodes[len(p.nodes)-1]
}

func (p *Path) canVisit(connection *utils.GraphNode) bool {
	if isBigCave(connection) {
		return true
	}
	if connection.Name == "start" {
		return false
	}
	if !p.hasVisited(connection) {
		return true
	}
	return p.revisitsLeft > 0
}

func isBigCave(node *utils.GraphNode) bool {
	return utils.FirstChar(node.Name) >= 'A' && utils.FirstChar(node.Name) <= 'Z'
}
