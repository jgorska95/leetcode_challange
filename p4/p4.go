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
	copy(workSlice, nums1)

	//TODO: fix proper sorted array joining
	for _, val := range nums2 {
		workSlice = append(workSlice, val)
	}

	elem := len(workSlice) / 2

	if len(workSlice)%2 != 0 {
		return float64(workSlice[elem+1])
	}

	retVal = float64(workSlice[elem-1]) + float64(workSlice[elem])
	retVal = retVal / 2

	return retVal
}
