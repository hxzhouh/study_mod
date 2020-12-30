/*
@Time : 2020/12/30 16:38
@Author : ZhouHui2
@File : 589
@Software: GoLand
*/
package tree

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var result []int

	var f func(root *Node)
	f = func(root *Node) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		for _, v := range root.Children {
			f(v)
		}
	}
	f(root)
	return result
}
