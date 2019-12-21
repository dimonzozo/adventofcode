package intcode

import "github.com/sirupsen/logrus"

type ParameterMode uint16

const (
	ParameterModePosition  ParameterMode = 0
	ParameterModeImmediate ParameterMode = 1
)

type Parameter struct {
	Value int
	Mode  ParameterMode
}

func NewParamenter(value int, mode ParameterMode) Parameter {
	return Parameter{Value: value, Mode: mode}
}

func (p *Parameter) Get(memory *Memory) int {
	logger := logrus.WithField("method", "Get")

	switch p.Mode {
	case ParameterModeImmediate:
		return p.Value
	case ParameterModePosition:
		return memory.Get(uint32(p.Value))
	default:
		logger.Fatalf("Unknown ParameterMode: %d", p.Mode)
	}

	return 0
}

func (p *Parameter) Set(memory *Memory, value int) {
	logger := logrus.WithField("method", "Set")

	switch p.Mode {
	case ParameterModeImmediate:
		logger.Fatal("Immediate parameter mode is not supported as set parameter")
	case ParameterModePosition:
		memory.Set(uint32(p.Value), value)
	default:
		logger.Fatalf("Unknown ParameterMode: %d", p.Mode)
	}
}
