package day01

import (
	"testing"
)

type tribTest struct {
	n        int
	expected int
}

var tribTests = []tribTest{
	{4, 4},
	{25, 1389537},
}

func TestTribonacci(t *testing.T) {
	t.Log("[DynamicPrograming Day 1] Testing tribonacci")
	for _, test := range tribTests {
		if output := tribonacci(test.n); output != test.expected {
			t.Errorf("Input %d :: output %d not equal to expected %d", test.n, output, test.expected)
		}
	}
}
