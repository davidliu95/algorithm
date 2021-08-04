package main

//冗余连接
func findRedundantConnection(edges [][]int) []int {
	var res = make([]int, 2)
	var p = make([]int, len(edges)+1)
	for i := 1; i <= len(edges); i++ {
		p[i] = i
	}

	for i := 0; i < len(edges); i++ {
		a := find(edges[i][0], p)
		b := find(edges[i][1], p)
		if a == b {
			res = []int{edges[i][0], edges[i][1]}
		} else {
			// 合并
			//p[a] = b
			union(a, b, p)
		}
	}
	return res
}

// 包含了树压缩动作，将每个子集合树变为二级
func find(x int, p []int) int {
	if p[x] != x {
		p[x] = find(p[x], p)
	}

	return p[x]
}

// 对两个集合按树根节点（子集的领导节点）合并， x, y为两个集合的任意两个节点
func union(x, y int, p []int) {
	x_root := find(x, p)
	y_root := find(y, p)
	if x_root != y_root {
		p[x_root] = y_root
	}
}

//岛屿数量
type UnionFindSet struct {
	Parents  []int // 每个结点的顶级节点
	SetCount int   // 连通分量的个数
}

func (u *UnionFindSet) Init(grid [][]byte) {
	row := len(grid)
	col := len(grid[0])
	count := row * col
	u.Parents = make([]int, count)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			u.Parents[i*col+j] = i*col + j
			if grid[i][j] == '1' {
				u.SetCount++
			}
		}
	}
}

func (u *UnionFindSet) Find(node int) int {
	if u.Parents[node] == node {
		return node
	}
	root := u.Find(u.Parents[node])
	u.Parents[node] = root
	return root
}

func (u *UnionFindSet) Union(node1 int, node2 int) {
	root1 := u.Find(node1)
	root2 := u.Find(node2)
	if root1 == root2 {
		return
	}
	if root1 < root2 {
		u.Parents[root1] = root2
	} else {
		u.Parents[root2] = root1
	}
	u.SetCount--
}

// 心得：并查集是一种搜索算法（针对聚合的）
func numIslands(grid [][]byte) int {
	// 创建并初始化并查集
	u := &UnionFindSet{}
	row := len(grid)
	col := len(grid[0])
	u.Init(grid)
	// 根据grid建立相应的并查集，并统计连通分量个数【每连接一次进行减一】
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				// 如果周边四个方向也是1就进行union
				if i-1 >= 0 && grid[i-1][j] == '1' {
					u.Union(i*col+j, (i-1)*col+j)
				}
				if i+1 < row && grid[i+1][j] == '1' {
					u.Union(i*col+j, (i+1)*col+j)
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					u.Union(i*col+j, i*col+(j-1))
				}
				if j+1 < col && grid[i][j+1] == '1' {
					u.Union(i*col+j, i*col+(j+1))
				}
				grid[i][j] = '0'
			}
		}
	}
	// 返回结果
	return u.SetCount
}
