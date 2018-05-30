package main

// TreeNode struct represent a general tree structure
type TreeNode struct {
	Name     string
	Children []TreeNode
}

func (root *TreeNode) PreOrderTraval(action func(*TreeNode)) {
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

func (root *TreeNode) PostOrderTraval(action func(*TreeNode)) {
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

func (root *TreeNode) BoradFirstSearch(action func(*TreeNode)) {
	if root == nil || action == nil {
		return
	}
	// Create a queue
	queue := CreateQueue()

	// Enqueue the root node
	queue.enqueue(root)

	for !queue.isEmpty() {
		var node *TreeNode
		// Dequeue an element from the queue
		element, err := queue.dequeue()

		// err will contains value if queue is empty
		if err != nil {
			panic(err)
		}

		// Assert and Convert type of the element into *TreeNode
		node = element.(*TreeNode)

		// Perform the action on node
		action(node)

		// If the node is leaf node
		if !node.hasChildren() {
			continue
		}

		// Enqueue all children of the current node.
		for i := 0; i < len(node.Children); i++ {
			queue.enqueue(&node.Children[i])
		}
	}
}

func (root *TreeNode) hasChildren() bool {
	if root == nil {
		return true
	}
	return len(root.Children) > 0
}
