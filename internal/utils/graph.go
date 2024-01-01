package utils

type Graph struct {
	nodes []*GraphNode
}

func (g *Graph) Nodes() []*GraphNode {
	return g.nodes
}

func NewUndirectedGraph(connections []Pair[string]) Graph {
	nodes := make(map[string]*GraphNode)

	getNode := func(name string) *GraphNode {
		node, ok := nodes[name]
		if !ok {
			node = &GraphNode{Name: name, Connections: make([]*GraphNode, 0)}
			nodes[name] = node
		}
		return node
	}

	addConnection := func(from, to string) {
		nodeFrom := getNode(from)
		nodeTo := getNode(to)
		nodeFrom.Connections = append(nodeFrom.Connections, nodeTo)
	}

	for _, conn := range connections {
		addConnection(conn.Left, conn.Right)
		addConnection(conn.Right, conn.Left)
	}

	return Graph{Values(nodes)}
}

type GraphNode struct {
	Name        string
	Connections []*GraphNode
}
