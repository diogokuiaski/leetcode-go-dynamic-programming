package day03

import "testing"

type deleteAndEarnTest struct {
	n        []int
	expected int
}

var deleteAndEarnTests = []deleteAndEarnTest{
	{[]int{3, 4, 2}, 6},
	{[]int{2, 2, 3, 3, 3, 4}, 9},
	{[]int{1, 10, 1, 1, 10, 1}, 24},
	{[]int{1, 1, 1, 2, 4, 5, 5, 5, 6}, 18},
	{[]int{8, 7, 3, 8, 1, 4, 10, 10, 10, 2}, 52},
}

func TestDeleteAndEarn(t *testing.T) {
	t.Log("[DynamicPrograming Day 3 Testing deleteAndEarn")
	for _, test := range deleteAndEarnTests {
		if output := deleteAndEarn(test.n); output != test.expected {
			t.Errorf("Input %v :: output %d not equal to expected %d", test.n, output, test.expected)
		}
	}
}
