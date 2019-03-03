package ds

import "fmt"

// TreeNode - Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildItem for create Tree
type BuildItem struct {
	Parent   *TreeNode
	ChildVal interface{}
	IsLeft   bool
}

// BuildTree with level order traversal nums
/*
 * 根据树的层序遍历数组，构建一棵二叉树。例如：
 * 输入：[3,2,3,null,3,null,1]
 *
 * 构建的二叉树是：
 *     3
 *    / \
 *   2   3
 *    \   \
 *     3   1
 *
 * 注意：如果某个位置上的节点为 null，那么它左右子树的 null 不需要写出来。
 * 比如对于下面这棵树：
 *     1
 *      \
 *       2
 *        \
 *         4
 * 应该表示为：[1,null,2,null,4]
 * 而不是：[1,null,2,null,null,null,4]
 *
 */
func BuildTree(nums []interface{}) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	dummy := &TreeNode{Val: 0}
	item := BuildItem{Parent: dummy, ChildVal: nums[0], IsLeft: true}
	q := []BuildItem{item}
	p := 1

	for len(q) != 0 {
		it := q[0]
		q = q[1:]

		var child *TreeNode
		if it.ChildVal != nil {
			child = &TreeNode{Val: it.ChildVal.(int)}
		}
		if it.IsLeft {
			it.Parent.Left = child
		} else {
			it.Parent.Right = child
		}

		if child != nil && p < len(nums) {
			q = append(q, BuildItem{Parent: child, ChildVal: nums[p], IsLeft: true})
			p++
			q = append(q, BuildItem{Parent: child, ChildVal: nums[p], IsLeft: false})
			p++
		}
	}

	return dummy.Left
}

// NewTree - new a tree point
func NewTree(v int) *TreeNode {
	return &TreeNode{v, nil, nil}
}

// Insert - insert node to tree
func (t *TreeNode) Insert(v int, left bool) *TreeNode {
	if t == nil {
		return &TreeNode{v, nil, nil}
	}
	if left {
		t.Left = t.Left.Insert(v, true)
		return t.Left
	}

	t.Right = t.Right.Insert(v, false)
	return t.Right
}

// Desc - description of tree
func (t *TreeNode) Desc() {
	if t == nil {
		return
	}

	fmt.Println(t.Val)
	t.Left.Desc()
	t.Right.Desc()
}
