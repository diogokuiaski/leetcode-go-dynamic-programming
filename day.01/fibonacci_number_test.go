package day01

import (
	"testing"
)

type fibTest struct {
	n        int
	expected int
}

var fibTests = []fibTest{
	{2, 1},
	{3, 2},
	{4, 3},
	{0, 0},
	{30, 832040},
}

func TestFibonacci(t *testing.T) {
	t.Log("[DynamicPrograming Day 1] Testing fibonacci")
	for _, test := range fibTests {
		if output := fib(test.n); output != test.expected {
			t.Errorf("Input %d :: output %d not equal to expected %d", test.n, output, test.expected)
		}
	}
}
