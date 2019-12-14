package main

import (
	"github.com/sirupsen/logrus"
)

type Action func(x int, y int) bool

func SetGridAction(grid Grid) Action {
	return func(x int, y int) bool {
		grid.Set(uint32(x), uint32(y), 1)
		return true
	}
}

func DistanceAction(point Coordinate, distance *int) Action {
	return func(x int, y int) bool {
		*distance++
		if x == point.X && y == point.Y {
			return false
		} else {
			return true
		}
	}
}

type Executor struct {
	initPos    Coordinate
	currentPos Coordinate
}

func NewExecutor(initPos Coordinate) *Executor {
	return &Executor{
		initPos:    initPos,
		currentPos: initPos,
	}
}

func (e *Executor) ExecuteCommands(commands []Command, action Action) {
	logger := logrus.WithField("method", "ExecuteCommands")
	for _, command := range commands {
		for delta := uint32(0); delta < command.amount; delta++ {
			switch command.direction {
			case DirectionUp:
				e.currentPos.X += 1
			case DirectionDown:
				e.currentPos.X -= 1
			case DirectionRight:
				e.currentPos.Y += 1
			case DirectionLeft:
				e.currentPos.Y -= 1
			default:
				logger.Fatalf("Unknown direction: %d", command.direction)
			}

			if !action(e.currentPos.X, e.currentPos.Y) {
				return
			}
		}
	}
}
