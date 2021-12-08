package day02

import "testing"

type stairsTest struct {
	n        int
	expected int
}

var stairsTests = []stairsTest{
	{2, 2},
	{3, 3},
}

func TestClimbingStairs(t *testing.T) {
	t.Log("[DynamicPrograming Day 2 Testing ClimbingStairs")
	for _, test := range stairsTests {
		if output := climbStairs(test.n); output != test.expected {
			t.Errorf("Input %d :: output %d not equal to expected %d", test.n, output, test.expected)
		}
	}
}
