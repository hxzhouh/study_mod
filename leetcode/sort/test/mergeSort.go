package test

var temp []int

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	m := len(a) / 2
	l := MergeSort(a[:m])
	r := MergeSort(a[m:])
	return merge(l, r)
}

func merge(a, b []int) (c []int) {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}
	c = append(c, a[i:]...) // 将 a[] 剩下的元素拷贝进最终结果
	c = append(c, b[j:]...) // 将 b[] 剩下的元素拷贝进最终结果
	return
}
