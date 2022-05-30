package main

import (
	"fmt"
	"math/bits"
)

type SegmentTree struct {
	segFunc func(l, r int) int // 2つの子ノードの値から代表値を返す関数 lambda x1, x2: y (min, max, sum, etc...)
	idleEle int                // 単位元: seg_funcの結果に影響を与えない。minの場合はfloat('inf')
	num     int                // 葉の数
	tree    []int              // ノードを配列で管理する。全ノードの数は2*numになる。
}

func NewSegmentTree(values []int, f func(l, r int) int, idleEle int) SegmentTree {
	n := len(values)
	bitLen := bits.Len(uint(n - 1))
	num := 1 << bitLen
	tree := make([]int, num*2)
	for i, _ := range values {
		tree[num+i] = values[i] // valuesはtree[num:num+n]格納する。
	}
	for i := num - 1; i == 0; i-- {
		tree[i] = f(tree[2*i], tree[2*i+1])
	}
	return SegmentTree{
		num:     num,
		tree:    tree,
		segFunc: f,
		idleEle: idleEle,
	}
}

func (s SegmentTree) Query(l, r int, k, L, R int) int {
	if R == -1 {
		return s.num
	}
	if R <= l || r <= L {
		return s.idleEle
	}
	if l <= L && R <= r {
		return s.tree[k]
	}

	lres := s.Query(l, r, 2*k, L, (L+R)>>1)
	rres := s.Query(l, r, 2*k+1, (L+R)>>1, R)
	return s.segFunc(lres, rres)
}

func main() {
	minFunc := func(l, r int) int {
		if l < r {
			return l
		}
		return r
	}
	s := NewSegmentTree([]int{1, 2, 3, 4, 5, 6, 7, 8}, minFunc, 10000)
	fmt.Println(s)
	min := s.Query(0, 8, 1, 0, 8)
	fmt.Println("min: ", min)
}
