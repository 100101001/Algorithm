package main

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

type UF struct {
	parent []int
	size   []int //
	count  int   // 连通分量
}

func newUF(n int) *UF {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < len(size); i++ {
		size[i] = 1
		parent[i] = i
	}

	return &UF{
		count:  n,
		parent: parent,
		size:   size,
	}
}

func (u *UF) connected(p, q int) bool {
	rootP := u.find(p)
	rootQ := u.find(q)
	return rootP == rootQ
}

//
func (u *UF) find(p int) int {
	for u.parent[p] != p {
		// 进一步压缩到 O(1)
		u.parent[p] = u.parent[u.parent[p]]
		p = u.parent[p]
	}
	return p
}

func (u *UF) union(p, q int) {
	rootP := u.find(p)
	rootQ := u.find(q)
	if rootQ == rootP {
		return
	}

	// 小的挂到大的下面，降低到 O(logN)
	if u.size[rootQ] > u.size[rootP] {
		u.parent[rootP] = rootQ
		u.size[rootQ] = u.size[rootP] + u.size[rootQ]
	} else {
		u.parent[rootQ] = rootP
		u.size[rootP] = u.size[rootP] + u.size[rootQ]
	}

	u.count--
}
