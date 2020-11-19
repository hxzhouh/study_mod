/*
@Time : 2020/9/11 17:21
@Author : ZhouHui2
@File : InsertionSort
@Software: GoLand
*/
package test

func insertionSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	length := len(a)
	for i := 1; i < length; i++ { //外面的一层
		value := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			}
		}
		a[j+1] = value
	}
	return a
}
