package leetCode

func intersection(nums1 []int, nums2 []int) []int {
	a := make(map[int]struct{})
	for _, v := range nums1 {
		a[v] = struct{}{}
	}
	result := make(map[int]struct{})
	for _, v := range nums2 {
		if _, ok := a[v]; ok {
			result[v] = struct{}{}
		}
	}
	aa := make([]int, 0)
	for k, _ := range result {
		aa = append(aa, k)
	}
	return aa
}
