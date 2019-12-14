package main

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

type Direction int

const (
	DirectionUp Direction = iota + 1
	DirectionDown
	DirectionLeft
	DirectionRight
)

type Command struct {
	direction Direction
	amount    uint32
}

func NewCommand(dir Direction, amount uint32) Command {
	return Command{direction: dir, amount: amount}
}

type Commands []Command

func ParseCommands(commandsStr []string) Commands {
	logger := logrus.WithField("method", "parse")

	parsedCommands := make([]Command, 0)

	for _, commandStr := range commandsStr {
		var direction Direction

		switch commandStr[0] {
		case 'R':
			direction = DirectionRight
		case 'U':
			direction = DirectionUp
		case 'D':
			direction = DirectionDown
		case 'L':
			direction = DirectionLeft
		default:
			logger.Fatalf("Unknown direction: %v", commandStr[0])
		}

		amountStr := commandStr[1:]

		amountNum, err := strconv.ParseUint(amountStr, 10, 32)
		if err != nil {
			logger.Fatalf("Unable to parse amount as uint value: %v", err)
		}

		parsedCommands = append(parsedCommands, NewCommand(direction, uint32(amountNum)))
	}

	return parsedCommands
}
