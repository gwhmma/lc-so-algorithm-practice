package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

示例: 

你可以将以下二叉树：

    1
   / \
  2   3
     / \
    4   5

序列化为 "[1,2,3,null,null,4,5]"
提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的

*/

type Tree4 struct {
	val   int
	left  *Tree4
	right *Tree4
}

type queue2 struct {
	val []*Tree4
}

func (q *queue2) push(t *Tree4) {
	q.val = append(q.val, t)
}

func (q *queue2) pop() *Tree4 {
	t := q.val[0]
	q.val = q.val[1:]
	return t
}

// bfs   !表示一层结束  #表示为空
// 1,!,2,3,!,#,#,4,5,!
// 辅助使用队列来存储每一层的节点
func serialize(t *Tree4) string {
	var rs string
	var bfs func(t *Tree4)
	q := &queue2{}
	mark := &Tree4{}

	q.push(t)
	bfs = func(t *Tree4) {
		for len(q.val) > 0 {
			size := len(q.val)
			for i := 0; i < size; i++ {
				node := q.pop()
				if node == mark {
					rs += "#"
					continue
				} else {
					rs += strconv.Itoa(node.val)
				}

				if node.left != nil {
					q.push(node.left)
				} else {
					q.push(mark)
				}
				if node.right != nil {
					q.push(node.right)
				} else {
					q.push(mark)
				}
			}
			rs += "!"
		}
	}
	bfs(t)

	return rs
}

// 通过split将二叉树每一层的数据得到
// 第一次先将二叉树根节点放入队列
// 之后对于每一层的节点  从队列取出前一层的节点 按照顺序将当前的数据拼接到上一层的节点
func deserialize(s string) *Tree4 {
	if len(s) == 0 {
		return nil
	}

	nodes := strings.Split(s, "!")
	nodes = nodes[:len(nodes)-1]
	q := &queue2{}
	t := &Tree4{}

	for _, node := range nodes {
		if len(q.val) > 0 {
			times := len(q.val)
			bt := []byte(node) // 当前层的节点数据
			j := 0

			// 遍历的次数为队列的大小(上一层节点的数量)
			for i := 0; i < times; i++ {
				t := q.pop()

				// 只有不是#(不是空节点)才进行拼接操作
				// 如果不为#(不是空节点)那么在把当前节点放入队列  在下一次迭代时拼接上它的子节点
				if string(bt[j]) != "#" {
					left, _ := strconv.Atoi(string(bt[j]))
					t.left = &Tree4{val: left}
					q.push(t.left)
				}
				if string(bt[j+1]) != "#" {
					right, _ := strconv.Atoi(string(bt[j+1]))
					t.right = &Tree4{val: right}
					q.push(t.right)
				}

				j += 2
			}
		} else {
			// 第一次将根节点放入队列
			val, _ := strconv.Atoi(node)
			t.val = val
			q.push(t)
		}
	}

	return t
}

func main() {
	n1 := &Tree4{val: 1}
	n2 := &Tree4{val: 2}
	n3 := &Tree4{val: 3}
	n4 := &Tree4{val: 4}
	n5 := &Tree4{val: 5}
	n6 := &Tree4{val: 6}
	n7 := &Tree4{val: 7}

	n1.left = n2
	n1.right = n3
	n2.left = n4
	n2.right = n5
	n3.right = n6
	n4.left = n7

	fmt.Println(serialize(n1))
	// 1!23!45#6!7#####!##!
	tree := deserialize("1!23!45#6!7#####!##!")
	fmt.Println(serialize(tree))

}
