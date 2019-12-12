package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestCpu_Run(t *testing.T) {
	testCases := []struct {
		Instructions string
		Result       uint32
		ResultIndex  uint32
	}{
		{
			Instructions: "1,9,10,3,2,3,11,0,99,30,40,50",
			Result:       3500,
			ResultIndex:  0,
		},
		{
			Instructions: "1,0,0,0,99",
			Result:       2,
			ResultIndex:  0,
		},
		{
			Instructions: "2,3,0,3,99",
			Result:       6,
			ResultIndex:  3,
		},
		{
			Instructions: "2,4,4,5,99,0",
			Result:       9801,
			ResultIndex:  5,
		},
		{
			Instructions: "1,1,1,4,99,5,6,0,99",
			Result:       30,
			ResultIndex:  0,
		},
	}

	for _, testCase := range testCases {
		dataArr := strings.Split(testCase.Instructions, ",")
		cells := make([]uint32, len(dataArr))
		for k, v := range dataArr {
			cell, _ := strconv.ParseUint(v, 10, 32)
			cells[k] = uint32(cell)
		}

		memory := NewMemory()
		memory.Load(cells)

		cpu := NewCpu(memory)
		cpu.Run()

		result := memory.Get(testCase.ResultIndex)
		expectedResult := testCase.Result

		if result != expectedResult {
			t.Fatalf("Expected result to be %d, got %d", expectedResult, result)
		}
	}
}
