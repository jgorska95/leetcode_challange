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
	var retVal float64
	workSlice := []int{}
	ptr1, ptr2 := 0, 0

	for ptr1 < len(nums1) && ptr2 < len(nums2) {
		if nums1[ptr1] < nums2[ptr2] {
			workSlice = append(workSlice, nums1[ptr1])
			ptr1++
		} else {
			workSlice = append(workSlice, nums2[ptr2])
			ptr2++
		}
	}

	if ptr1 < len(nums1) {
		workSlice = append(workSlice, nums1[ptr1:]...)
	}

	if ptr2 < len(nums2) {
		workSlice = append(workSlice, nums2[ptr2:]...)
	}

	elem := len(workSlice) / 2

	if len(workSlice)%2 != 0 {
		return float64(workSlice[elem])
	}

	retVal = float64(workSlice[elem-1]) + float64(workSlice[elem])
	retVal = retVal / 2

	return retVal
}
