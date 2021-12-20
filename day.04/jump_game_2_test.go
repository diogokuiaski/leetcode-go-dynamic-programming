package day04

import "testing"

type jumpGame2Test struct {
	n        []int
	expected int
}

var jumpGame2Tests = []jumpGame2Test{
	{[]int{2, 3, 1, 1, 4}, 2},
	{[]int{2, 3, 1, 0, 4}, 2},
}

func TestJumpGame2(t *testing.T) {
	t.Log("[DynamicPrograming Day 4 Testing JumpGame2")
	for _, test := range jumpGame2Tests {
		if output := jump(test.n); output != test.expected {
			t.Errorf("Input %v :: output %d not equal to expected %d", test.n, output, test.expected)
		}
		if output := jump2Best(test.n); output != test.expected {
			t.Errorf("Input %v :: output %d not equal to expected %d", test.n, output, test.expected)
		}
	}
}
