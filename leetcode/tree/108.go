/*
@Time : 2020/12/29 14:33
@Author : ZhouHui2
@File : 108
@Software: GoLand
*/
package tree

/**
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1。

示例:

给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return &TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
	}
	mid := len(nums) / 2
	result := &TreeNode{
		Val:   nums[mid],
		Left:  nil,
		Right: nil,
	}
	result.Left = sortedArrayToBST(nums[0:mid])
	result.Right = sortedArrayToBST(nums[mid+1:])
	return result
}
