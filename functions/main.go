package main

import "fmt"

// TreeNode struct represent a general tree structure
type TreeNode struct {
	Name     string
	Children []TreeNode
}

func main() {
	root := TreeNode{
		Name: "A",
		Children: []TreeNode{
			TreeNode{
				Name: "B",
				Children: []TreeNode{
					TreeNode{
						Name: "E",
					},
					TreeNode{
						Name: "F",
					},
					TreeNode{
						Name: "G",
					},
				},
			},
			TreeNode{
				Name: "C",
				Children: []TreeNode{
					TreeNode{
						Name: "H",
					},
					TreeNode{
						Name: "I",
					},
				}},
			TreeNode{
				Name: "D",
				Children: []TreeNode{
					TreeNode{
						Name: "J",
					},
					TreeNode{
						Name: "K",
					},
					TreeNode{
						Name: "L",
					},
				}},
		},
	}

	fmt.Println("Depth First Search by recursive call (System Stack):")
	fmt.Print("  PreOrder:  ")
	preOrderTraval(&root, func(root *TreeNode) { fmt.Printf("%s ", (*root).Name) })
	fmt.Println()
	fmt.Print("  PostOrder: ")
	postOrderTraval(&root, func(root *TreeNode) { fmt.Printf("%s ", (*root).Name) })

	fmt.Println()

	fmt.Print("Broad Fist Seach: ")
	boradFirstSearch(&root, func(root *TreeNode) { fmt.Printf("%s ", (*root).Name) })
	fmt.Println()
}

func preOrderTraval(root *TreeNode, action func(*TreeNode)) {
	if root == nil || action == nil {
		return
	}

	action(root)

	if (*root).Children == nil || len((*root).Children) == 0 {
		return
	}

	for _, v := range (*root).Children {
		preOrderTraval(&v, action)
	}
}

func postOrderTraval(root *TreeNode, action func(*TreeNode)) {
	if root == nil || action == nil {
		return
	}

	if (*root).Children != nil {
		for _, v := range (*root).Children {
			postOrderTraval(&v, action)
		}
	}

	action(root)
}

func boradFirstSearch(root *TreeNode, action func(*TreeNode)) {
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

func (node *TreeNode) hasChildren() bool {
	if node == nil {
		return true
	}
	return len(node.Children) > 0
}
