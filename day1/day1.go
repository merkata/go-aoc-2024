package day1

import (
	"strings"
	"strconv"
	"sort"
)

func Solve1(input []string) int {
	var left []int
	var right []int
	var dif int


	for _, line := range input {
		parts := strings.Fields(line)
		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[1])
		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	sort.Ints(left)
	sort.Ints(right)

	for i := range left {
		dif += abs(left[i] - right[i])
	}

	return dif
}

func Solve2(input []string) int {
	var left []int
	right := make(map[int]int)
	var dif int

	for _, line := range input {
		parts := strings.Fields(line)
		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[1])
		left = append(left, leftNum)
		right[rightNum]++
	}

	for i := range left {
		dif += left[i] * right[left[i]]
	}

	return dif
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
