package graph

type NodeQueue struct {
	items []Node
}

func NewQueue() *NodeQueue {
	return &NodeQueue{items: make([]Node, 0)}
}

func (s *NodeQueue) Enqueue(t Node) {
	s.items = append(s.items, t)
}

func (s *NodeQueue) Dequeue() *Node {
	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	return &item
}

func (s *NodeQueue) Front() *Node {
	item := s.items[0]
	return &item
}

func (s *NodeQueue) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *NodeQueue) Size() int {
	return len(s.items)
}
