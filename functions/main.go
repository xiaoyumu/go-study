package main

import (
	"fmt"

	"github.com/xiaoyumu/go-study/datastructure/tree"
)

// Build a tree that has structure as following
//                A
//           /    |   \
//         /      |     \
//        B       C       D
//     /  | \    / \     / | \
//    E   F  G  H   I   J  K  L
func buildTree() *tree.Node {

	root := tree.Node{
		Name: "A",
		Children: []tree.Node{
			tree.Node{
				Name: "B",
				Children: []tree.Node{
					tree.Node{
						Name: "E",
					},
					tree.Node{
						Name: "F",
					},
					tree.Node{
						Name: "G",
					},
				},
			},
			tree.Node{
				Name: "C",
				Children: []tree.Node{
					tree.Node{
						Name: "H",
					},
					tree.Node{
						Name: "I",
					},
				}},
			tree.Node{
				Name: "D",
				Children: []tree.Node{
					tree.Node{
						Name: "J",
					},
					tree.Node{
						Name: "K",
					},
					tree.Node{
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
	root.PreOrderTraval(func(root *tree.Node) { fmt.Printf("%s ", (*root).Name) })

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

func showNode(root *tree.Node) {
	fmt.Printf("%s ", (*root).Name)
}
