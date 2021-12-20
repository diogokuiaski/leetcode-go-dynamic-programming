package day04

func canJump(nums []int) bool {
	visited := make([]bool, len(nums))
	return letsJump(nums, visited, 0)
}

func letsJump(nums []int, visited []bool, index int) bool {
	if index >= len(nums)-1 {
		return true
	}
	if visited[index] {
		return false
	}
	visited[index] = true
	for i := 1; i <= nums[index]; i++ {
		if letsJump(nums, visited, index+i) {
			return true
		}
	}
	return false
}

func bestCanJump(nums []int) bool {
	lastReachable := 0
	sz := len(nums)
	for i := 0; i < sz; i++ {
		if i > lastReachable {
			break
		}
		if i+nums[i] > lastReachable {
			lastReachable = i + nums[i]
		}
		if lastReachable >= sz-1 {
			return true
		}
	}
	return false
}
