package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	return help129(root, 0)
}

func help129(root *TreeNode, value int) int {
	if root == nil {
		return 0
	}
	temp := value*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return temp
	}
	return help129(root.Left, temp) + help129(root.Right, temp)
}

//func clacSliceToInt(values []int) int {
//	var result []byte
//	//length := len(values) - 1
//	for i := range values {
//		result = append(result, []byte(strconv.Itoa(i))...)
//	}
//	a, _ := strconv.Atoi(string(result))
//	//fmt.Println(a)
//	return a
//}
