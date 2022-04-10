package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Query struct {
	writeNum int
	cnt      int
}

type queryQueue []Query

func (queue *queryQueue) empty() bool    { return len(*queue) == 0 }
func (queue *queryQueue) first() Query   { return (*queue)[0] }
func (queue *queryQueue) firstP() *Query { return &(*queue)[0] }
func (queue *queryQueue) last() Query    { return (*queue)[len(*queue)-1] }
func (queue *queryQueue) lastP() *Query  { return &(*queue)[len(*queue)-1] }
func (queue *queryQueue) push(i Query)   { *queue = append(*queue, i) }
func (queue *queryQueue) pop() Query {
	result := (*queue)[0]
	*queue = (*queue)[1:]
	return result
}

const QUEUE_PUSH = 1
const QUEUE_POP = 2

func main() {
	Q := ni()
	qQueue := queryQueue{}
	var ans []int

	// 計算パート
	for i := 0; i < Q; i++ {
		queryType := ni()
		switch queryType {
		case QUEUE_PUSH:
			q := Query{writeNum: ni(), cnt: ni()}
			qQueue.push(q)
		case QUEUE_POP:
			currentCnt := ni()
			sum := 0
			for currentCnt != 0 && !qQueue.empty() {
				if qQueue.first().cnt >= currentCnt { // 以下であればカウント分引いて終わり
					sum += qQueue.first().writeNum * currentCnt
					qQueue.firstP().cnt -= currentCnt // 使った分減らす

					// 使い切ったものは消す
					if qQueue.first().cnt == 0 {
						qQueue.pop()
					}

					ans = append(ans, sum)
					break
				} else { // より大きければ、ある分だけ使って次のqueueに
					sum += qQueue.first().writeNum * qQueue.first().cnt
					currentCnt -= qQueue.first().cnt // 現クエリの残回数を減らす
					qQueue.pop()
				}
			}
		}
	}

	// 表示パート
	for i := 0; i < len(ans); i++ {
		if ans[i] != 0 {
			fmt.Println(ans[i])
		}
	}
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
