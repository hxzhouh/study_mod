package test

func Quicksort(a []int) []int {
	if len(a) < 1 {
		return a
	}
	quickSortHelp(a, 0, len(a)-1)
	return a
}

func quickSortHelp(a []int, start, end int) {
	if start >= end {
		return
	}
	p := partition(a, start, end)
	quickSortHelp(a, start, p-1)
	quickSortHelp(a, p+1, end)

}

func partition(a []int, start, end int) int { //
	pivot := a[end] //暂时默认队尾为 pivot
	i := start - 1
	for j := start; j < end; j++ {
		if a[j] < pivot { // 将所有小于 pivot 的数据交换到最左边，原来的数据先丢到后边。
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[end] = a[end], a[i+1] //因为选取的最后的那个值为pivot，所以，将值放到正确的位置上。
	return i + 1                    // 第一个确定的pivot
}
