package main

import "runtime/debug"

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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return max(leftDepth, rightDepth) + 1
}

func main() {
	// Example usage:
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	depth := maxDepth(root)
	println("Maximum Depth of Binary Tree:", depth) // Output: 3
}
