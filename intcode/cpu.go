package intcode

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
)

type Input int
type Output int

type OpCode uint16

const (
	OpCodeAdd    OpCode = 1
	OpCodeMul    OpCode = 2
	OpCodeInput  OpCode = 3
	OpCodeOutput OpCode = 4
	OpCodeHalt   OpCode = 99
)

type Cpu struct {
	ip     uint32
	memory *Memory

	inputs            []Input
	currentInputIndex uint32

	outputs []Output
}

func NewCpu(memory *Memory) *Cpu {
	return &Cpu{
		memory:  memory,
		inputs:  make([]Input, 0),
		outputs: make([]Output, 0),
	}
}

func (c *Cpu) Run() {
	logger := logrus.WithField("method", "Run")

LOOP:
	for {
		operation := c.memory.Get(c.ip)

		opCode, parameterModes := c.parseOperation(operation)

		switch opCode {
		case OpCodeAdd:
			data := c.memory.GetRange(c.ip+1, c.ip+4)
			parameters := []Parameter{
				NewParamenter(data[0], parameterModes[0]),
				NewParamenter(data[1], parameterModes[1]),
				NewParamenter(data[2], parameterModes[2]),
			}
			c.add(parameters)
			c.ip += 4
		case OpCodeMul:
			data := c.memory.GetRange(c.ip+1, c.ip+4)
			parameters := []Parameter{
				NewParamenter(data[0], parameterModes[0]),
				NewParamenter(data[1], parameterModes[1]),
				NewParamenter(data[2], parameterModes[2]),
			}
			c.mul(parameters)
			c.ip += 4
		case OpCodeInput:
			data := c.memory.GetRange(c.ip+1, c.ip+2)
			parameters := []Parameter{
				NewParamenter(data[0], parameterModes[0]),
			}
			c.input(parameters)
			c.ip += 2
		case OpCodeOutput:
			data := c.memory.GetRange(c.ip+1, c.ip+2)
			parameters := []Parameter{
				NewParamenter(data[0], parameterModes[0]),
			}
			c.output(parameters)
			c.ip += 2
		case OpCodeHalt:
			break LOOP
		default:
			logger.Fatalf("Unknown OpCode: %d", operation)
		}
	}
}

func (c *Cpu) Outputs() []Output {
	return c.outputs
}

func (c *Cpu) SetInputs(inputs []Input) {
	c.inputs = inputs
}

func (c *Cpu) parseOperation(operation int) (OpCode, []ParameterMode) {
	operationStr := fmt.Sprintf("%05d", operation)
	opCodeNum, _ := strconv.ParseUint(string(operationStr[len(operationStr)-2:]), 10, 16)
	paramModes := make([]ParameterMode, 3)

	for i := 0; i <= 2; i++ {
		modeNum, _ := strconv.ParseUint(string(operationStr[i]), 10, 16)
		paramModes[2-i] = ParameterMode(modeNum)
	}

	return OpCode(opCodeNum), paramModes
}

func (c *Cpu) add(params []Parameter) {
	a := params[0].Get(c.memory)
	b := params[1].Get(c.memory)
	params[2].Set(c.memory, a+b)
}

func (c *Cpu) mul(params []Parameter) {
	a := params[0].Get(c.memory)
	b := params[1].Get(c.memory)
	params[2].Set(c.memory, a*b)
}

func (c *Cpu) input(params []Parameter) {
	params[0].Set(c.memory, int(c.inputs[c.currentInputIndex]))
	c.currentInputIndex++
}

func (c *Cpu) output(params []Parameter) {
	value := params[0].Get(c.memory)
	c.outputs = append(c.outputs, Output(value))
}
