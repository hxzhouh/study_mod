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
//https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
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

// 迭代 法
func preorderTraversal1(root *TreeNode) []int {
	for root == nil {
		return nil
	}
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		head := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		for head != nil {
			res = append(res, head.Val)
			if head.Right != nil {
				queue = append(queue, head.Right) // 右节点入栈，下一步处理左节点
			}
			head = head.Left
		}
	}
	return res
}
