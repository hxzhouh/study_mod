package test

func selectionSort(a []int) []int {
	length := len(a) // 长度
	if length <= 1 {
		return a
	}
	head := 0
	for head < length {
		tail := length - 1
		for tail > head {
			if a[head] > a[tail] {
				a[head], a[tail] = soft(a[head], a[tail])
			}
			tail--
		}
		head++
	}
	return a
}
