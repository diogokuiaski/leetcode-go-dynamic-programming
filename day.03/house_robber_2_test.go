package day03

import "testing"

type houseRobber2Test struct {
	n        []int
	expected int
}

var houseRobber2Tests = []houseRobber2Test{
	{[]int{2, 3, 2}, 3},
	{[]int{1, 2, 3, 1}, 4},
	{[]int{1, 2, 3}, 3},
	{[]int{1}, 1},
	{[]int{1, 2}, 2},
	{[]int{1, 10, 1, 1, 15, 10}, 25},
	{[]int{10, 1, 1, 10, 11}, 20},
	{[]int{6, 6, 4, 8, 4, 3, 3, 10}, 27},
	{[]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, 16},
}

func TestHouseRobber2(t *testing.T) {
	t.Log("[DynamicPrograming Day 3 Testing HouseRobber2")
	for _, test := range houseRobber2Tests {
		if output := rob2(test.n); output != test.expected {
			t.Errorf("Input %v :: output %d not equal to expected %d", test.n, output, test.expected)
		}
	}
}
