package main

import (
	"fmt"
	"runtime/debug"
)

// The setMemoryLimit function sets a limit on the amount of memory that the Go program can use.
// In this case, it is set to 6 MiB (megabytes). This means that if the program tries to allocate
// more than 6 MiB of memory, it will be terminated by the runtime. This is useful for controlling
// resource usage and preventing excessive memory consumption in applications.
// Isso melhora as métricas de desempenho do LeetCode, pois limita o uso de memória do programa Go a 6 MiB,
// evitando que ele consuma recursos excessivos e garantindo que ele funcione dentro de limites aceitáveis.
// Isso é especialmente importante em ambientes de execução restritos, como plataformas de avaliação de
// código online, onde o uso eficiente de recursos é crucial para manter a estabilidade e o desempenho do sistema.
func init() {
	debug.SetMemoryLimit(3*1024*1024 + 9*1024) // 3.9 MiB
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
        if val < root.Val {
            root = root.Left
        } else {
            root = root.Right
        }
    }

    return root
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("nil")
		return
	}

	fmt.Printf("Ponteiro root: %p\n", root)                      // Endereço do nó
	fmt.Printf("Endereço de Val: %p\n", &root.Val)               // Endereço do campo Val
	fmt.Printf("Valor de Val: %d\n", root.Val)                   // Valor armazenado
	fmt.Printf("Valor em Left: %d\n", root.Left.Val)						 // Valor do filho esquerdo
	fmt.Printf("Valor em Right: %d\n", root.Right.Val)					 // Valor do filho direito
	fmt.Printf("Imprimindo o conteúdo do ponteiro: %+v\n", root) // Imprime o conteúdo do ponteiro root, mostrando os campos Val, Left e Right
}

func main() {
	// Example usage:
	// Construct a binary search tree and search for a value.
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}

	val := 2
	result := searchBST(root, val)
	if result != nil {
		printTree(result)
		println("Found node with value:", result.Val)
	} else {
		println("Node with value", val, "not found.")
	}
}
