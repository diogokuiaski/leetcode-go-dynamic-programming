package day04

import "testing"

type jumpGameTest struct {
	n        []int
	expected bool
}

var jumpGameTests = []jumpGameTest{
	{[]int{2, 3, 1, 1, 4}, true},
	{[]int{3, 2, 1, 0, 4}, false},
}

func TestJumpGame(t *testing.T) {
	t.Log("[DynamicPrograming Day 3 Testing JumpGame")
	for _, test := range jumpGameTests {
		if output := canJump(test.n); output != test.expected {
			t.Errorf("Input %v :: output %t not equal to expected %t", test.n, output, test.expected)
		}
	}
}
