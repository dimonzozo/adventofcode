package main

import (
	"fmt"
	"strconv"
)

const RangeStart = 231832
const RangeEnd = 767346
const Length = 6

func part1() {
	matchNums := make([]string, 0)

	for num := uint32(RangeStart); num < uint32(RangeEnd); num++ {
		numString := strconv.FormatUint(uint64(num), 10)

		digits := make([]int, 0)

		for _, digitRune := range numString {
			digit, _ := strconv.Atoi(string(digitRune))
			digits = append(digits, digit)
		}

		sameNums := make(map[int]int)
		twoSame := false
		for _, d := range digits {
			count, exists := sameNums[d]
			if exists {
				sameNums[d] = count + 1
			} else {
				sameNums[d] = 1
			}
		}

		for _, count := range sameNums {
			if count >= 2 {
				twoSame = true
			}
		}

		neverDecreasing := true
		for i := 1; i < Length; i++ {
			if digits[i] < digits[i-1] {
				neverDecreasing = false
			}
		}

		if twoSame && neverDecreasing {
			matchNums = append(matchNums, numString)
		}
	}

	fmt.Printf("Result (part 1): %d\n", len(matchNums))
}

func part2() {
	matchNums := make([]string, 0)

	for num := uint32(RangeStart); num < uint32(RangeEnd); num++ {
		numString := strconv.FormatUint(uint64(num), 10)

		digits := make([]int, 0)

		for _, digitRune := range numString {
			digit, _ := strconv.Atoi(string(digitRune))
			digits = append(digits, digit)
		}

		sameNums := make(map[int]int)
		twoSame := false
		for _, d := range digits {
			count, exists := sameNums[d]
			if exists {
				sameNums[d] = count + 1
			} else {
				sameNums[d] = 1
			}
		}

		for _, count := range sameNums {
			if count == 2 {
				twoSame = true
			}
		}

		neverDecreasing := true
		for i := 1; i < Length; i++ {
			if digits[i] < digits[i-1] {
				neverDecreasing = false
			}
		}

		if twoSame && neverDecreasing {
			matchNums = append(matchNums, numString)
		}
	}

	fmt.Printf("Result (part 2): %d\n", len(matchNums))
}

func main() {
	part1()
	part2()
}
