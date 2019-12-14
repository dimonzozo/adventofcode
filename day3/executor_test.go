package main

import (
	"sort"
	"strings"
	"testing"
)

func TestExecuteCommand(t *testing.T) {
	testCases := []struct {
		Input1           string
		Input2           string
		ExpectedDistance int
		StepsCount       int
	}{
		{
			Input1:           "R8,U5,L5,D3",
			Input2:           "U7,R6,D4,L4",
			ExpectedDistance: 6,
			StepsCount:       30,
		},
		{
			Input1:           "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			Input2:           "U62,R66,U55,R34,D71,R55,D58,R83",
			ExpectedDistance: 159,
			StepsCount:       610,
		},
		{
			Input1:           "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			Input2:           "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			ExpectedDistance: 135,
			StepsCount:       410,
		},
	}

	for _, testCase := range testCases {
		commandsStr := strings.Split(testCase.Input1, ",")
		commands := ParseCommands(commandsStr)
		commandsStr2 := strings.Split(testCase.Input2, ",")
		commands2 := ParseCommands(commandsStr2)
		initPos := Coordinate{250, 250}

		wire1Grid := NewGrid(500, 500)
		executor1 := NewExecutor(initPos)
		executor1.ExecuteCommands(commands, SetGridAction(wire1Grid))

		wire2Grid := NewGrid(500, 500)
		executor2 := NewExecutor(initPos)
		executor2.ExecuteCommands(commands2, SetGridAction(wire2Grid))

		intersections := wire1Grid.FindIntersections(wire2Grid)

		result := FindClosestIntersectionDistance(intersections, initPos)

		if result != testCase.ExpectedDistance {
			t.Fatalf("Expected %d, got %d", testCase.ExpectedDistance, result)
		}

		intersectionDistances := make([]int, 0)

		for _, intersection := range intersections {
			wire1Distance := 0
			distanceExecutor1 := NewExecutor(initPos)
			distanceExecutor1.ExecuteCommands(commands, DistanceAction(intersection, &wire1Distance))

			wire2Distance := 0
			distanceExecutor2 := NewExecutor(initPos)
			distanceExecutor2.ExecuteCommands(commands2, DistanceAction(intersection, &wire2Distance))

			intersectionDistances = append(intersectionDistances, wire1Distance+wire2Distance)
		}

		sort.Ints(intersectionDistances)

		if intersectionDistances[0] != testCase.StepsCount {
			t.Fatalf("Expected %d, got %d", testCase.StepsCount, intersectionDistances[0])
		}
	}
}
