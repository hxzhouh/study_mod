/*
@Time : 2020/10/30 14:38
@Author : ZhouHui2
@File : 905
@Software: GoLand
*/
package leetCode

func sortArrayByParity(A []int) []int {
	begin := 0
	end := len(A) - 1
	for begin < end {
		if A[begin]&1 == 0 { // 偶数
			begin++
		} else if A[end]&1 == 1 { //基数
			end--
		} else {
			A[begin] += A[end]
			A[end] = A[begin] - A[end]
			A[begin] = A[begin] - A[end]
		}
	}
	return A
}
