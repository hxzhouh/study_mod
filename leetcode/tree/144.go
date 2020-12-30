/*
@Time : 2020/12/30 11:18
@Author : ZhouHui2
@File : 144
@Software: GoLand
*/
package tree

/**
144. 二叉树的前序遍历
*/
//var result []int

func preorderTraversal(root *TreeNode) []int {
	var vals []int
	var helper114 func(node *TreeNode)
	helper114 = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		helper114(node.Left)
		helper114(node.Right)
	}
	helper114(root)
	return vals
}
