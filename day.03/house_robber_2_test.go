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
}

func TestHouseRobber2(t *testing.T) {
	t.Log("[DynamicPrograming Day 3 Testing HouseRobber2")
	for _, test := range houseRobber2Tests {
		if output := rob2(test.n); output != test.expected {
			t.Errorf("Input %v :: output %d not equal to expected %d", test.n, output, test.expected)
		}
	}
}
