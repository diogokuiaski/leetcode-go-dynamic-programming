package day03

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type RobIndex struct {
	index, cost int
}
