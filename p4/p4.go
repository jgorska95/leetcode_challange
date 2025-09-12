package p4

/*
Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
*/

func Run(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays(nums1, nums2)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1sum := 0
	nums2sum := 0

	for _, val := range nums1 {
		nums1sum += val
	}

	for _, val := range nums2 {
		nums2sum += val
	}

	sum := nums1sum + nums2sum
	numsItems := len(nums1) + len(nums2)

	retVal := float64(sum % numsItems)

	return retVal
}
