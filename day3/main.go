package main

import (
	"github.com/dimonzozo/adventofcode/common"
	"github.com/sirupsen/logrus"
	"math"
	"sort"
	"strings"
)

var initPos = Point{0, 0}

func main() {
	logger := logrus.WithField("method", "main")

	commands1 := ParseCommands(strings.Split(common.Input()[0], ","))
	commands2 := ParseCommands(strings.Split(common.Input()[1], ","))

	wire1Grid := NewGrid()
	executor1 := NewExecutor(initPos)
	executor1.ExecuteCommands(commands1, SetGridAction(wire1Grid))

	wire2Grid := NewGrid()
	executor2 := NewExecutor(initPos)
	executor2.ExecuteCommands(commands2, SetGridAction(wire2Grid))

	intersections := wire1Grid.FindIntersections(wire2Grid)

	logger.Infof("Result part 1: %d", FindClosestIntersectionDistance(intersections, initPos))

	wireDistances := make([]int, 0)

	for _, intersection := range intersections {
		wire1Distance := 0
		distanceExecutor1 := NewExecutor(initPos)
		distanceExecutor1.ExecuteCommands(commands1, DistanceAction(intersection, &wire1Distance))

		wire2Distance := 0
		distanceExecutor2 := NewExecutor(initPos)
		distanceExecutor2.ExecuteCommands(commands2, DistanceAction(intersection, &wire2Distance))

		wireDistances = append(wireDistances, wire1Distance+wire2Distance)
	}

	sort.Ints(wireDistances)

	logger.Infof("Result part 2: %d", wireDistances[0])
}

func FindClosestIntersectionDistance(intersections []Point, pos Point) int {
	distances := make([]int, 0)

	for _, intersection := range intersections {
		distance := math.Abs(float64(intersection.X-pos.X)) + math.Abs(float64(intersection.Y-pos.Y))
		distances = append(distances, int(distance))
	}

	sort.Ints(distances)

	if len(distances) > 0 {
		return distances[0]
	} else {
		return 0
	}
}
