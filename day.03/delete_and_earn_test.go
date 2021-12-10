package day03

import "testing"

type deleteAndEarnTest struct {
	n        []int
	expected int
}

var deleteAndEarnTests = []deleteAndEarnTest{
	{[]int{1, 2, 3, 1}, 4},
	{[]int{2, 7, 9, 3, 1}, 12},
	{[]int{1, 10, 1, 1, 10, 1}, 20},
}

func TestDeleteAndEarn(t *testing.T) {
	t.Log("[DynamicPrograming Day 3 Testing deleteAndEarn")
	for _, test := range deleteAndEarnTests {
		if output := deleteAndEarn(test.n); output != test.expected {
			t.Errorf("Input %v :: output %d not equal to expected %d", test.n, output, test.expected)
		}
	}
}
