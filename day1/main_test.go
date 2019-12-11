package main

import "testing"

func TestCalculateFuel(t *testing.T) {
	testCases := []struct {
		Mass               int
		ExpectedFuelAmount int
	}{
		{
			Mass:               12,
			ExpectedFuelAmount: 2,
		},
		{
			Mass:               14,
			ExpectedFuelAmount: 2,
		},
		{
			Mass:               1969,
			ExpectedFuelAmount: 654,
		},
		{
			Mass:               100756,
			ExpectedFuelAmount: 33583,
		},
	}

	for _, testCase := range testCases {
		result := CalculateFuel(testCase.Mass)

		if result != testCase.ExpectedFuelAmount {
			t.Fatalf("Extected %d, got %d", testCase.ExpectedFuelAmount, result)
		}
	}
}

func TestCalculateFuelTotal(t *testing.T) {
	testCases := []struct {
		Mass               int
		ExpectedFuelAmount int
	}{
		{
			Mass:               14,
			ExpectedFuelAmount: 2,
		},
		{
			Mass:               1969,
			ExpectedFuelAmount: 966,
		},
		{
			Mass:               100756,
			ExpectedFuelAmount: 50346,
		},
	}

	for _, testCase := range testCases {
		result := CalculateFuelTotal(testCase.Mass)

		if result != testCase.ExpectedFuelAmount {
			t.Fatalf("Extected %d, got %d", testCase.ExpectedFuelAmount, result)
		}
	}
}
