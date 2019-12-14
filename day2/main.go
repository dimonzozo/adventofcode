package main

import (
	"github.com/dimonzozo/adventofcode/common"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func main() {
	logger := logrus.WithField("method", "main")

	dataArr := strings.Split(common.Input()[0], ",")
	cells := make([]uint32, len(dataArr))
	for k, v := range dataArr {
		cell, _ := strconv.ParseUint(v, 10, 32)
		cells[k] = uint32(cell)
	}

	memory := NewMemory()
	memory.Load(cells)

	memory.Set(1, 12)
	memory.Set(2, 2)

	cpu := NewCpu(memory)
	cpu.Run()

	logger.Infof("Result: %d", memory.Get(0))

	targetResult := 19690720

LOOP:
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			logger.Debugf("Checking noun %d, and verb %d", noun, verb)

			memory := NewMemory()
			memory.Load(cells)

			memory.Set(1, uint32(noun))
			memory.Set(2, uint32(verb))

			cpu := NewCpu(memory)
			cpu.Run()

			if memory.Get(0) == uint32(targetResult) {
				logger.Infof("Result is %d with noun %d, and verb %d", targetResult, noun, verb)
				logger.Infof("Part 2 result: %d", 100 * noun + verb)
				break LOOP
			}
		}
	}
}
