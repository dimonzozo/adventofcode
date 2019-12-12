package main

import (
	"strconv"
	"strings"
)

type Memory struct {
	cells []uint32
}

func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) Load(cells []uint32) {
	m.cells = cells
}

func (m *Memory) Get(index uint32) uint32 {
	return m.cells[index]
}

func (m *Memory) GetRange(start uint32, end uint32) []uint32 {
	return m.cells[start:end]
}

func (m *Memory) Set(index uint32, value uint32) {
	m.cells[index] = value
}

func (m *Memory) Dump() string {
	cellsString := make([]string, len(m.cells))

	for k, v := range m.cells {
		cellsString[k] = strconv.FormatUint(uint64(v), 10)
	}

	return strings.Join(cellsString, ",")
}
