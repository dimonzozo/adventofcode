package main

import (
	"github.com/dimonzozo/adventofcode/common"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func main() {
	logger := logrus.WithField("method", "main")

	instructions := strings.Split(common.Input(), "\n")[0]
	dataArr := strings.Split(instructions, ",")
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
}
