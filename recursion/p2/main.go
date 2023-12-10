// Tree
package main

type node struct {
	value int
	left  *node
	right *node
}

var (
	ExampleTree = &node{
		value: 1,
		left: &node{
			value: 2,
			left: &node{
				value: 3,
			},
			right: &node{
				value: 4,
			},
		},
		right: &node{
			value: 5,
		},
	}
)

func sumIterative(root *node) int {
	queue := make(chan *node, 10)
	queue <- root
	var sum int
	for {
		select {
		case node := <-queue:
			sum += node.value
			if node.left != nil {
				queue <- node.left
			}
			if node.right != nil {
				queue <- node.right
			}
		default:
			return sum
		}
	}
}
