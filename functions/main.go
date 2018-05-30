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

	fmt.Println("Depth First Search by recursive call (System Stack):")
	fmt.Print("  PreOrder:  ")
	root.PreOrderTraval(func(root *TreeNode) { fmt.Printf("%s ", (*root).Name) })

	fmt.Println()
	fmt.Print("  PostOrder: ")
	root.PostOrderTraval(func(root *TreeNode) { fmt.Printf("%s ", (*root).Name) })

	fmt.Println()

	fmt.Print("Broad First Seach: ")
	root.BoradFirstSearch(func(root *TreeNode) { fmt.Printf("%s ", (*root).Name) })
	fmt.Println()

	stack := CreateStack()

	listValue := []int{1, 2, 3, 4, 5, 6}

	for i := 0; i < len(listValue); i++ {
		stack.Push(&listValue[i])
	}

	for !stack.isEmpty() {
		v, err := stack.Pop()
		if err != nil {
			panic(err)
		}

		value := v.(*int)
		fmt.Print(*value)
	}
	fmt.Println()
}
