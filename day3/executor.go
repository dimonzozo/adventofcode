package main

import (
	"github.com/sirupsen/logrus"
)

type Action func(point Point) bool

func SetGridAction(grid Grid) Action {
	return func(point Point) bool {
		grid.Set(point)
		return true
	}
}

func DistanceAction(targetPoint Point, distance *int) Action {
	return func(point Point) bool {
		*distance++
		if point == targetPoint {
			return false
		} else {
			return true
		}
	}
}

type Executor struct {
	initPos    Point
	currentPos Point
}

func NewExecutor(initPos Point) *Executor {
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

			if !action(e.currentPos) {
				return
			}
		}
	}
}
