package main

import "fmt"

// Build a tree that has structure as following
//                A
//           /    |   \
//         /      |     \
//        B       C       D
//     /  | \    / \     / | \
//    E   F  G  H   I   J  K  L
func buildTree() *TreeNode {

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

	return &root
}

func main() {

	root := buildTree()
	fmt.Println("Depth First Search by recursive call (System Stack):")
	fmt.Print("  PreOrder:  ")
	root.PreOrderTraval(func(root *TreeNode) { fmt.Printf("%s ", (*root).Name) })

	fmt.Println()
	fmt.Print("  PostOrder: ")
	root.PostOrderTraval(showNode)

	fmt.Println()

	fmt.Print("Broad First Seach: ")
	root.BoradFirstSearch(showNode)
	fmt.Println()

	fmt.Print("Depth First Seach (By simulated stack): ")
	root.DepthFirstSearch(showNode)
}

func showNode(root *TreeNode) {
	fmt.Printf("%s ", (*root).Name)
}
