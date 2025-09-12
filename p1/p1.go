package p1

/*
Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]
*/

func Run(nums []int, target int) []int {
	return twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for key, val := range nums {
		if hm, ok := hashMap[val]; ok {
			return []int{key, hm}
		}

		hashMap[target-val] = key
	}

	return []int{}
}
