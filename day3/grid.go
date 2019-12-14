package main

type Grid struct {
	grid [][]byte
}

func NewGrid(width uint32, height uint32) Grid {
	grid := make([][]byte, height)
	for i := range grid {
		grid[i] = make([]byte, width)
	}

	return Grid{grid: grid}
}

func (g *Grid) Get(x uint32, y uint32) byte {
	return g.grid[y][x]
}

func (g *Grid) Set(x uint32, y uint32, value byte) {
	g.grid[y][x] = value
}

func (g *Grid) FindIntersections(other Grid) []Coordinate {
	coordinates := make([]Coordinate, 0)

	for y, xes := range g.grid {
		for x, val := range xes {
			if val == 1 && other.Get(uint32(x), uint32(y)) == 1 {
				coordinates = append(coordinates, Coordinate{X: x, Y: y})
			}
		}
	}

	return coordinates
}
