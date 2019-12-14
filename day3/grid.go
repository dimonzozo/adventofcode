package main

type Grid struct {
	grid map[Point]struct{}
}

func NewGrid() Grid {
	return Grid{grid: make(map[Point]struct{})}
}

func (g *Grid) Set(point Point) {
	g.grid[point] = struct{}{}
}

func (g *Grid) FindIntersections(other Grid) []Point {
	intersections := make([]Point, 0)

	for point := range g.grid {
		if other.Exists(point) {
			intersections = append(intersections, point)
		}
	}

	return intersections
}

func (g *Grid) Exists(point Point) bool {
	_, ok := g.grid[point]
	return ok
}
