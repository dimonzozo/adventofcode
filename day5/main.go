package main

import (
	"github.com/dimonzozo/adventofcode/common"
	"github.com/dimonzozo/adventofcode/intcode"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func main() {
	logger := logrus.WithField("method", "main")

	dataArr := strings.Split(common.Input()[0], ",")
	cells := make([]int, len(dataArr))
	for k, v := range dataArr {
		cell, _ := strconv.ParseInt(v, 10, 32)
		cells[k] = int(cell)
	}

	memory := intcode.NewMemory()
	memory.Load(cells)
	cpu := intcode.NewCpu(memory)
	cpu.SetInputs([]intcode.Input{1})
	cpu.Run()

	logger.Infof("Result: %v", cpu.Outputs())
}
