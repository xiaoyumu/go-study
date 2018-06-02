package tree

import (
	"github.com/xiaoyumu/go-study/datastructure"
)

// Node struct represent a Node element in general tree structure
type Node struct {
	Name     string
	Children []Node
}

func (root *Node) hasChildren() bool {
	if root == nil {
		return true
	}
	return len(root.Children) > 0
}

// PreOrderTraval is a Recursive implementation of depth first search
// Which visit the root node first
func (root *Node) PreOrderTraval(action func(*Node)) {
	if root == nil || action == nil {
		return
	}

	action(root)

	if (*root).Children == nil || len((*root).Children) == 0 {
		return
	}

	for _, v := range (*root).Children {
		v.PreOrderTraval(action)
	}
}

// PostOrderTraval is a Recursive implementation of depth first search
// Which visit the root node after all the child nodes
func (root *Node) PostOrderTraval(action func(*Node)) {
	if root == nil || action == nil {
		return
	}

	if (*root).Children != nil {
		for _, v := range (*root).Children {
			v.PostOrderTraval(action)
		}
	}

	action(root)
}

// DepthFirstSearch is a general implementation based on stack
func (root *Node) DepthFirstSearch(action func(*Node)) {
	if root == nil || action == nil {
		return
	}

	// Create a stack
	stack := datastructure.CreateStack()

	stack.Push(root)

	for !stack.IsEmpty() {
		element, err := stack.Pop()
		if err != nil {
			panic(err)
		}
		node := element.(*Node)

		// Call the action method with current node
		action(node)

		// If the node is a leaf node, move on to the next node in the stack
		if !node.hasChildren() {
			continue
		}

		// Backward pushing children nodes of current node into stack
		// Make the node order the same as preorder traversal
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack.Push(&node.Children[i])
		}
	}
}

// BoradFirstSearch of the given tree
func (root *Node) BoradFirstSearch(action func(*Node)) {
	if root == nil || action == nil {
		return
	}
	// Create a queue
	queue := datastructure.CreateQueue()

	// Enqueue the root node
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		var node *Node
		// Dequeue an element from the queue
		element, err := queue.Dequeue()

		// err will contains value if queue is empty
		if err != nil {
			panic(err)
		}

		// Assert and Convert type of the element into *Node
		node = element.(*Node)

		// Perform the action on node
		action(node)

		// If the node is leaf node
		if !node.hasChildren() {
			continue
		}

		// Enqueue all children of the current node.
		for i := range node.Children {
			queue.Enqueue(&node.Children[i])
		}
	}
}
