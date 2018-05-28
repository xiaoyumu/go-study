package main

import "fmt"

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

	fmt.Println("DepthFirstTraval by recursive call (System Stack):")
	depthFirstTraval(&root, func(root *TreeNode) { fmt.Printf("%s ", (*root).Name) })
	fmt.Println()
}

// TreeNode struct represent a general tree structure
type TreeNode struct {
	Name     string
	Children []TreeNode
}

func depthFirstTraval(root *TreeNode, action func(*TreeNode)) {
	if root == nil {
		return
	}

	if action == nil {
		return
	}

	action(root)

	if (*root).Children == nil || len((*root).Children) == 0 {
		return
	}

	for _, v := range (*root).Children {
		depthFirstTraval(&v, action)
	}
}

func broadFistTraval(root *TreeNode, action func(*TreeNode)) {
	// to be implemented
}
