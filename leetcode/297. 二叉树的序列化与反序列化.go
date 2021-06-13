package main

import (
	"fmt"
	"strconv"
	"strings"
)

var NULL = "#"

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {

	queue := []*TreeNode{root}
	var str []string

	for len(queue) != 0 {
		node := queue[0]
		if node == nil {
			str = append(str, NULL)
			continue
		}
		str = append(str, fmt.Sprintf("%d", node.Val))
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
		queue = queue[1:]
	}

	return strings.Join(str, ",")
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" || data == "#" {
		return nil
	}

	queue := make([]*TreeNode, 0)
	queI := 0
	strS := strings.Split(data, ",")
	queue = append(queue, str2Node(strS[0]))

	for strI := 0; strI < len(strS); queI++ {
		node := queue[queI]

		strI++
		if strI < len(strS) {
			node.Left = str2Node(strS[strI])
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		strI++
		if strI < len(strS) {
			node.Right = str2Node(strS[strI])
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

	}

	return queue[0]
}

func str2Node(str string) *TreeNode {
	if str == NULL {
		return nil
	}
	val, _ := strconv.ParseInt(str, 10, 64)
	return &TreeNode{Val: int(val)}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
