package intcode

import (
	"strconv"
	"strings"
	"testing"
)

func TestCpu_Run(t *testing.T) {
	testCases := []struct {
		Instructions string
		Result       int
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
		cells := make([]int, len(dataArr))
		for k, v := range dataArr {
			cell, _ := strconv.ParseInt(v, 10, 32)
			cells[k] = int(cell)
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

func TestJumpsAndCompares(t *testing.T) {
	prog1 := "3,9,8,9,10,9,4,9,99,-1,8"
	prog2 := "3,9,7,9,10,9,4,9,99,-1,8"
	prog3 := "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"

	testCases := []struct {
		Instructions    string
		Inputs          []Input
		ExpectedOutputs []Output
	}{
		{
			Instructions:    prog1,
			Inputs:          []Input{8},
			ExpectedOutputs: []Output{1},
		},
		{
			Instructions:    prog1,
			Inputs:          []Input{6},
			ExpectedOutputs: []Output{0},
		},
		{
			Instructions:    prog2,
			Inputs:          []Input{7},
			ExpectedOutputs: []Output{1},
		},
		{
			Instructions:    prog2,
			Inputs:          []Input{9},
			ExpectedOutputs: []Output{0},
		},
		{
			Instructions:    prog3,
			Inputs:          []Input{7},
			ExpectedOutputs: []Output{999},
		},
		{
			Instructions:    prog3,
			Inputs:          []Input{8},
			ExpectedOutputs: []Output{1000},
		},
		{
			Instructions:    prog3,
			Inputs:          []Input{55},
			ExpectedOutputs: []Output{1001},
		},
	}

	for _, testCase := range testCases {
		dataArr := strings.Split(testCase.Instructions, ",")
		cells := make([]int, len(dataArr))
		for k, v := range dataArr {
			cell, _ := strconv.ParseInt(v, 10, 32)
			cells[k] = int(cell)
		}

		memory := NewMemory()
		memory.Load(cells)

		cpu := NewCpu(memory)
		cpu.SetInputs(testCase.Inputs)
		cpu.Run()

		for i, o := range cpu.Outputs() {
			if o != testCase.ExpectedOutputs[i] {
				t.Fatalf("Expected output to be %v, but got %v", testCase.ExpectedOutputs, cpu.Outputs())
			}
		}
	}
}
