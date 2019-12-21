package intcode

import (
	"strconv"
	"strings"
)

type Memory struct {
	cells []int
}

func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) Load(cells []int) {
	m.cells = make([]int, len(cells))
	copy(m.cells, cells)
}

func (m *Memory) Get(index uint32) int {
	return m.cells[index]
}

func (m *Memory) GetRange(start uint32, end uint32) []int {
	return m.cells[start:end]
}

func (m *Memory) Set(index uint32, value int) {
	m.cells[index] = value
}

func (m *Memory) Dump() string {
	cellsString := make([]string, len(m.cells))

	for k, v := range m.cells {
		cellsString[k] = strconv.FormatInt(int64(v), 10)
	}

	return strings.Join(cellsString, ",")
}
