package main

import (
	"fmt"
	"math"
)

// Node struct
type node struct {
	value int
	left  *node
	right *node
}

// Tree definition
var Tree = &node{
	value: 1,
	left: &node{
		value: 2,
		left:  &node{value: 3},
		right: &node{value: 4},
	},
	right: &node{
		value: 5,
		left:  &node{value: 6},
		right: &node{value: 7},
	},
}

// sumIterative
func sumIterative(root *node) int {
	queue := make(chan *node, 10)
	queue <- root
	var sum int
	for {
		select {
		case node := <-queue:
			sum += node.value
			if node.left != nil {
				queue <- node.right
			}
			if node.right != nil {
				queue <- node.right
			}
		default:
			return sum
		}
	}
}

func sumRecursive(node *node) int {
	if node == nil {
		return 0
	}
	return node.value + sumRecursive(node.left) + sumRecursive(node.right)
}

var maximum = 0

func MaxGlobalVariable(node *node) {
	if node == nil {
		return
	}
	if node.value > maximum {
		maximum = node.value
	}
	MaxGlobalVariable(node.left)
	MaxGlobalVariable(node.right)
}

func maxInline(node *node, maxValue int) int {
	if node == nil {
		return maxValue
	}
	if node.value > maxValue {
		maxValue = node.value
	}
	maxLeft := maxInline(node.left, maxValue)
	maxRight := maxInline(node.left, maxValue)
	if maxLeft > maxRight {
		return maxLeft
	}
	return maxRight
}

func main() {
	maximum = int(math.MinInt)
	MaxGlobalVariable(Tree)
	fmt.Println(maximum)
	fmt.Println(maxInline(Tree, 0))
}
