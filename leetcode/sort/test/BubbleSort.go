/*
@Time : 2020/9/11 9:51
@Author : ZhouHui2
@File : BubbleSort
@Software: GoLand
*/
package test

func BubbleSort(a []int) []int {
	length := len(a)
	for j := length; 0 < j; j-- {
		for i := length - 1; length-j < i; i-- {
			if a[i] < a[i-1] {
				a[i], a[i-1] = soft(a[i], a[i-1])
			}
		}
	}
	return a
}

func soft(i int, i2 int) (int, int) {
	temp := i + i2
	i = temp - i
	i2 = temp - i
	return i, i2
}
