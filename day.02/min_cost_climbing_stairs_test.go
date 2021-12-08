package day02

import "testing"

type minCostStairsTest struct {
	n        []int
	expected int
}

var minCostStairsTests = []minCostStairsTest{
	{[]int{10, 15, 20}, 15},
	{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	{[]int{0, 0, 1, 1}, 1},
}

func TestMinCostClimbingStairs(t *testing.T) {
	t.Log("[DynamicPrograming Day 2 Testing MinCostClimbingStairs")
	for _, test := range minCostStairsTests {
		if output := minCostClimbingStairs(test.n); output != test.expected {
			t.Errorf("Input %v :: output %d not equal to expected %d", test.n, output, test.expected)
		}
	}
}
