package graph

type Node struct {
	Value string
}

type ItemGraph struct {
	nodes []*Node
	edges map[Node][]*Node
}

func (g *ItemGraph) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (g *ItemGraph) AddEdge(n1, n2 *Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}

func (g *ItemGraph) Traverse(f func(*Node)) {
	q := NewQueue()
	n := g.nodes[0]
	q.Enqueue(*n)
	visited := make(map[*Node]bool)
	for {
		if q.IsEmpty() {
			break
		}
		node := q.Dequeue()
		visited[node] = true
		near := g.edges[*node]

		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[j] {
				q.Enqueue(*j)
				visited[j] = true
			}
		}
		if f != nil {
			f(node)
		}
	}
}
