package main

import (
	"math/rand"
)

type list struct {
	r, c int // 行、列
	dist int // 路径距离
}

//二进制矩阵中的最短路径
func shortestPathBinaryMatrix(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 || grid[0][0] == 1 {
		return -1
	}

	direction := [][]int{{1, -1}, {1, 0}, {1, 1}, {0, -1}, {0, 1}, {-1, -1}, {-1, 0}, {-1, 1}}
	rows, cols := len(grid), len(grid[0])
	var queue []list
	queue = append(queue, list{0, 0, 1})
	for len(queue) > 0 {
		curList := queue[0]
		queue = queue[1:]
		// 走到右下角，返回路径距离
		if curList.r == rows-1 && curList.c == cols-1 {
			return curList.dist
		}
		for _, d := range direction {
			nr := curList.r + d[0]
			nc := curList.c + d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 0 {
				queue = append(queue, list{nr, nc, curList.dist + 1})
				// 标记已走过
				grid[nr][nc] = 1
			}
		}
	}
	return -1
}

//设计跳表
type Node struct {
	val         int
	Right, Down *Node
}

func NewNode(v int, r, d *Node) *Node {
	return &Node{val: v, Right: r, Down: d}
}

type Skiplist struct {
	level int
	Head  *Node
}

func Constructor() Skiplist {
	return Skiplist{level: 1, Head: NewNode(0, nil, nil)}
}

func (this *Skiplist) Search(target int) bool {
	// 申请临时变量指向头部
	cur := this.Head
	// 直接while到死
	for cur != nil {
		// 如果cur的右边不为空，并且右边的值小于target...
		for cur.Right != nil && cur.Right.val < target {
			// ... 向右边移动
			cur = cur.Right
		} // 此时已经到了最右边（边界，或者小于右边的值了，准备向下走）

		// 如果右边的值还等于目标值..
		if cur.Right != nil && cur.Right.val == target {
			// ...直接返回
			return true
		}
		// 否则向下走
		cur = cur.Down
	}
	return false
}

func (this *Skiplist) Add(num int) {
	rLevel := 1
	// 如果小于等于总的层数，就掷***，rand.Int31()%2会得到一个1或者0的数，如果是0就给这个值的所在层总数加1
	for rLevel <= this.level && ((rand.Int31() & 1) == 0) {
		rLevel++ // 本值的越大概率越小
	}

	// 如果总层数已经超过了最大层数...
	if rLevel > this.level {
		// ...反向赋值，增长最大层数
		this.level = rLevel
		// 头结点换成新的，值就是插入值，右边是nil，下面是原来的头
		this.Head = NewNode(num, nil, this.Head)
	}
	cur := this.Head
	var last *Node = nil
	// 从最大层开始递减到第一层
	for l := this.level; l >= 1; l-- {
		// cur的右不为nil，并且右方的值小于要插入的值...(找到一个合适的区间)
		for cur.Right != nil && cur.Right.val < num {
			// ... 向右走
			cur = cur.Right
		}

		// 如果l小于rLevel，也就是本层依然需要这个节点..
		if l <= rLevel { // 在这个语句前已经找到了最大的节点，要么右边为空，要么右边比本值大。
			// ... 创建新的节点，值是要插入的值，右边就是原来的右边，下面为空
			cur.Right = NewNode(num, cur.Right, nil)
			// last是之前添加的点，具体在下方 last = cur.Right
			// 如果last存在，就应该把last（上层的cur.right）的下，指向当前新插入的节点
			if last != nil {
				// last的下就是cur的右
				last.Down = cur.Right
			}
			// last是要等于cur的右边的
			last = cur.Right
		}
		cur = cur.Down
	}
}

func (this *Skiplist) Erase(num int) bool {
	cur := this.Head
	var seen = false
	for l := this.level; l >= 1; l-- {
		for cur.Right != nil && cur.Right.val < num {
			cur = cur.Right
		}
		if cur.Right != nil && cur.Right.val == num {
			seen = true
			cur.Right = cur.Right.Right
		}
		cur = cur.Down
	}
	return seen
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
