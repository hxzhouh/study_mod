package test

// 查找第一个value等于给定值的下标

func lastValueBisection(a []int, value int) int {

	return lastValueBisection_Helper(a, 0, len(a)-1, value)
}
func lastValueBisection_Helper(a []int, start, end, value int) int {
	if end-start <= 1 {
		if a[start] == value {
			return start
		} else if a[end] == value {
			return end
		} else {
			return -1
		}
	}
	mid := (start + end) / 2
	if a[mid] == value {
		if mid == end || a[mid+1] != value {
			return mid
		} else {
			return lastValueBisection_Helper(a, start+1, end, value)
		}
		//return mid
	} else if a[mid] < value {
		return lastValueBisection_Helper(a, mid+1, end, value)
	} else {
		return lastValueBisection_Helper(a, start, mid-1, value)
	}
}
