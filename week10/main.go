package main

// 定义方块的结构体
type Node struct {
	l, r, h, maxR int
	left, right   *Node // 指针类型，难难难（大学没学好C语言的后果，一不小心bu会用）
}

func fallingSquares(positions [][]int) []int {
	// 创建返回值 使用切片 (动态数组)
	var res = make([]int, 0)
	// 根节点
	var root *Node = new(Node) // 初始化，对应类型的零值
	// 目前最高的高度
	maxH := 0
	for _, position := range positions {
		l := position[0]               // 左横坐标
		r := position[0] + position[1] // 右横坐标
		e := position[1]               // 边长
		curH := query(root, l, r)      // 目前区间的最高的高度
		root = insert(root, l, r, curH+e)
		maxH = max(maxH, curH+e)
		res = append(res, maxH)
	}
	return res
}

func insert(root *Node, l int, r int, h int) *Node {
	if root == nil {
		return &Node{
			l:    l,
			r:    r,
			h:    h,
			maxR: r,
		}
	}
	if l <= root.l {
		root.left = insert(root.left, l, r, h)
	} else {
		root.right = insert(root.right, l, r, h)
	}
	root.maxR = max(r, root.maxR)
	return root
}

func query(root *Node, l int, r int) int {
	// reflect.ValueOf(root).IsValid() 表示判断root是否为空
	// 新节点的左边界大于等于目前的maxR的话，直接得到0，不需要遍历了
	if root == nil || l >= root.maxR {
		return 0
	}
	// 高度
	curH := 0
	if !(r <= root.l || root.r <= l) { // 是否跟这个节点相交
		curH = root.h
	}
	// 剪枝
	curH = max(curH, query(root.left, l, r))
	if r >= root.l {
		curH = max(curH, query(root.right, l, r))
	}
	return curH
}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}
