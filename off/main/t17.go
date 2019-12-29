package main

/*
输入两棵二叉树A，B，判断B是不是A的子结构。（ps：我们约定空树不是任意一个树的子结构）
*/

type TreeNode1 struct {
	val   int
	left  *TreeNode1
	right *TreeNode1
}

func subTree(t1, t2 * TreeNode1) bool {
	if t2 == nil {
		return true
	}

	if t1 == nil {
		return false
	}

	if t1.val != t2.val {
		return false
	}

	return subTree(t1.left, t2.left) && subTree(t1.right, t2.right)
}

func hasSubtree(t1, t2 *TreeNode1) bool {
	if t2 == nil || t1 == nil {
		return false
	}
	return subTree(t1, t2) || hasSubtree(t1.left, t2) || hasSubtree(t1.right, t2)
}

func main() {

}
