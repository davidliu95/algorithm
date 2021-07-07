package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//从中序与后序遍历序列构造二叉树
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	//根节点为后序的最后一个索引
	l := len(postorder) - 1
	for k := range inorder {
		if inorder[k] == postorder[l] {
			return &TreeNode{
				Val:   inorder[k],
				Left:  buildTree(inorder[:k], postorder[:k]),
				Right: buildTree(inorder[k+1:], postorder[k:l]),
			}
		}
	}
	return nil
}

//课程表 II
func findOrder(numCourses int, prerequisites [][]int) []int {
	//入度表，记录本课程的先修课程有几门
	inDegree := make([]int, numCourses)
	//出边表，记录本课程是哪些课程的先修
	outEdge := make([][]int, numCourses)
	// 结果集
	var res []int

	for i := 0; i < len(prerequisites); i++ {
		// 更新入度
		inDegree[prerequisites[i][0]]++
		// 更新出边
		outEdge[prerequisites[i][1]] = append(outEdge[prerequisites[i][1]], prerequisites[i][0])
	}

	// 队列用于记录入度为0的课程
	var queue []int

	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 没有不依赖其它课程的课程，返回空
	if len(queue) == 0 {
		return res
	}

	for len(queue) > 0 {
		// 取出队头，加入结果集
		node := queue[0]
		queue = queue[1:]
		res = append(res, node)

		// 处理队头的出边, 把入度为0的入队
		for i := 0; i < len(outEdge[node]); i++ {
			out := outEdge[node][i]
			inDegree[out]--
			if inDegree[out] == 0 {
				queue = append(queue, out)
			}
		}
	}

	// 检查是否所有的课程都被安排，如果没有，说明有环，无法安排，返回空
	if len(res) == numCourses {
		return res
	}

	return []int{}
}
