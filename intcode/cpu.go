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
	OpCodeAdd         OpCode = 1
	OpCodeMul         OpCode = 2
	OpCodeInput       OpCode = 3
	OpCodeOutput      OpCode = 4
	OpCodeJumpIfTrue  OpCode = 5
	OpCodeJumpIfFalse OpCode = 6
	OpCodeLessThan    OpCode = 7
	OpCodeEquals      OpCode = 8
	OpCodeHalt        OpCode = 99
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
			c.add(c.buildParameters(data, parameterModes))
		case OpCodeMul:
			data := c.memory.GetRange(c.ip+1, c.ip+4)
			c.mul(c.buildParameters(data, parameterModes))
		case OpCodeInput:
			data := c.memory.GetRange(c.ip+1, c.ip+2)
			c.input(c.buildParameters(data, parameterModes))
		case OpCodeOutput:
			data := c.memory.GetRange(c.ip+1, c.ip+2)
			c.output(c.buildParameters(data, parameterModes))
		case OpCodeJumpIfTrue:
			data := c.memory.GetRange(c.ip+1, c.ip+3)
			c.jumpIfTrue(c.buildParameters(data, parameterModes))
		case OpCodeJumpIfFalse:
			data := c.memory.GetRange(c.ip+1, c.ip+3)
			c.jumpIfFalse(c.buildParameters(data, parameterModes))
		case OpCodeLessThan:
			data := c.memory.GetRange(c.ip+1, c.ip+4)
			c.lessThan(c.buildParameters(data, parameterModes))
		case OpCodeEquals:
			data := c.memory.GetRange(c.ip+1, c.ip+4)
			c.equals(c.buildParameters(data, parameterModes))
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

func (c *Cpu) buildParameters(data []int, parameterModes []ParameterMode) []Parameter {
	parameters := make([]Parameter, len(data))
	for i := 0; i < len(data); i++ {
		parameters[i] = NewParamenter(data[i], parameterModes[i])
	}
	return parameters
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
	c.ip += 4
}

func (c *Cpu) mul(params []Parameter) {
	a := params[0].Get(c.memory)
	b := params[1].Get(c.memory)
	params[2].Set(c.memory, a*b)
	c.ip += 4
}

func (c *Cpu) input(params []Parameter) {
	params[0].Set(c.memory, int(c.inputs[c.currentInputIndex]))
	c.currentInputIndex++
	c.ip += 2
}

func (c *Cpu) output(params []Parameter) {
	value := params[0].Get(c.memory)
	c.outputs = append(c.outputs, Output(value))
	c.ip += 2
}

func (c *Cpu) jumpIfTrue(params []Parameter) {
	value := params[0].Get(c.memory)
	if value != 0 {
		c.ip = uint32(params[1].Get(c.memory))
	} else {
		c.ip += 3
	}
}

func (c *Cpu) jumpIfFalse(params []Parameter) {
	value := params[0].Get(c.memory)
	if value == 0 {
		c.ip = uint32(params[1].Get(c.memory))
	} else {
		c.ip += 3
	}
}

func (c *Cpu) lessThan(params []Parameter) {
	p1 := params[0].Get(c.memory)
	p2 := params[1].Get(c.memory)
	if p1 < p2 {
		params[2].Set(c.memory, 1)
	} else {
		params[2].Set(c.memory, 0)
	}
	c.ip += 4
}

func (c *Cpu) equals(params []Parameter) {
	p1 := params[0].Get(c.memory)
	p2 := params[1].Get(c.memory)
	if p1 == p2 {
		params[2].Set(c.memory, 1)
	} else {
		params[2].Set(c.memory, 0)
	}
	c.ip += 4
}
