// atc001 b
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MERGE = 0
const SAME = 1

func solve(N, Q int, queries []Query) {
	uf := NewUnionFind(N)

	// クエリ回実施
	for i := 0; i < Q; i++ {
		q := &queries[i]

		if q.queryType == MERGE {
			uf.Merge(q.a, q.b)
		} else { // JUDGE
			if uf.IsSame(q.a, q.b) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}

type Query struct {
	queryType int // 0:連結クエリ, 1:判定クエリ
	a         int // 頂点A
	b         int // 頂点B
}

func main() {
	N, Q := ni(), ni()
	queries := make([]Query, Q)
	for i := 0; i < Q; i++ {
		q := &queries[i]
		q.queryType, q.a, q.b = ni(), ni(), ni()
	}
	solve(N, Q, queries)
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

// 変数の中身を入れ替える
func swap(a, b *int) {
	*b, *a = *a, *b
}

type UnionFind struct {
	parent []int // n個目の要素がどの親に属するか
	height []int // 木の高さ-1
	rSize  []int // 親配下の要素数(親自身を除く)
	uSize  int   // UnionFind全体の要素数
}

func NewUnionFind(size int) UnionFind {
	uf := UnionFind{}
	uf.uSize = size
	uf.parent = make([]int, size)
	uf.rSize = make([]int, size)
	uf.height = make([]int, size)
	for i := 0; i < uf.uSize; i++ {
		uf.parent[i] = i
	}
	return uf
}

// 木の根を求める
func (u *UnionFind) Root(x int) int {
	// speed 優先
	if u.parent[x] == x {
		return x
	}
	u.parent[x] = u.Root(u.parent[x]) // 経路圧縮
	return u.parent[x]
}

// xとyの属する集合をマージ
func (u *UnionFind) Merge(x, y int) {
	rootX, rootY := u.Root(x), u.Root(y)

	// 同じ親であれば既にマージ済み
	if rootX == rootY {
		return
	}

	heightX, heightY := u.height[rootX], u.height[rootY]
	// 後続処理でXにYを付けるため、Yを小さくしておく
	if heightX < heightY {
		swap(&rootX, &rootY)
	}
	u.parent[rootY] = rootX
	u.rSize[rootX] = 1 + u.rSize[rootX] + u.rSize[rootY]
	if heightX == heightY {
		u.height[rootX]++
	}
}

// 自身が属する集合に何要素あるか(自身を含む)
func (u *UnionFind) Size(x int) int {
	return u.rSize[u.Root(x)] + 1
}

// aとbが同じ集合に属するか
func (u *UnionFind) IsSame(a, b int) bool {
	return u.Root(a) == u.Root(b)
}

func (u *UnionFind) Groups() map[int][]int {
	hash := make(map[int][]int)
	for i := 0; i < u.uSize; i++ {
		r := u.Root(i)
		hash[r] = append(hash[r], i)
	}
	return hash
}

func init() {
	sc.Split(bufio.ScanWords)
}
