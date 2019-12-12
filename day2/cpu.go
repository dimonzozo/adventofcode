package main

import "github.com/sirupsen/logrus"

type OpCode uint16

const (
	OpCodeAdd  OpCode = 1
	OpCodeMul  OpCode = 2
	OpCodeHalt OpCode = 99
)

type Cpu struct {
	ip     uint32
	memory *Memory
}

func NewCpu(memory *Memory) *Cpu {
	return &Cpu{memory: memory}
}

func (c *Cpu) Run() {
	logger := logrus.WithField("method", "Run")

LOOP:
	for {
		operation := c.memory.Get(c.ip)

		switch OpCode(operation) {
		case OpCodeAdd:
			data := c.memory.GetRange(c.ip+1, c.ip+4)
			c.add(data[0], data[1], data[2])
			c.ip += 4
		case OpCodeMul:
			data := c.memory.GetRange(c.ip+1, c.ip+4)
			c.mul(data[0], data[1], data[2])
			c.ip += 4
		case OpCodeHalt:
			break LOOP
		default:
			logger.Fatalf("Unknown OpCode: %d", operation)
		}
	}
}

func (c *Cpu) add(ap uint32, bp uint32, dp uint32) {
	a := c.memory.Get(ap)
	b := c.memory.Get(bp)
	c.memory.Set(dp, a+b)
}

func (c *Cpu) mul(ap uint32, bp uint32, dp uint32) {
	a := c.memory.Get(ap)
	b := c.memory.Get(bp)
	c.memory.Set(dp, a*b)
}
