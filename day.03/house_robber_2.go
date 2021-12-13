package day03

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return Max(nums[0], nums[1])
	} else if len(nums) == 3 {
		return Max(Max(nums[0], nums[1]), nums[2])
	}
	memoRob2 := make(map[RobIndex]int)
	a := maxRob2(nums[:len(nums)-1], 0, memoRob2)

	memoRob2 = make(map[RobIndex]int)
	b := maxRob2(nums, 1, memoRob2)

	memoRob2 = make(map[RobIndex]int)
	c := maxRob2(nums, 2, memoRob2)

	return Max(Max(a, b), c)
}

func maxRob2(cost []int, index int, memoRob2 map[RobIndex]int) int {
	if index >= len(cost) {
		return 0
	}
	i := RobIndex{index, cost[index]}
	val, ok := memoRob2[i]
	if ok {
		return val
	} else {
		b := Max(
			maxRob2(cost, index+2, memoRob2),
			maxRob2(cost, index+3, memoRob2),
		)
		memoRob2[i] = cost[index] + b
		return memoRob2[i]
	}
}
