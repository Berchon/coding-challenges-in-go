package main

import (
	"fmt"
	"runtime/debug"
)

/*
# 700. Search in a Binary Search Tree

**Link:** https://leetcode.com/problems/search-in-a-binary-search-tree/description/?envType=study-plan-v2&envId=leetcode-75

**Difficulty:** Easy

## Problem Description

You are given the `root` of a binary search tree (BST) and an integer `val`.

Find the node in the BST that the node's value equals `val` and return the subtree rooted with that node. If such a node does not exist, return `null`.

## Example 1

**Input:**

```text
root = [4,2,7,1,3], val = 2
```

**Tree representation:**

The binary search tree represented by the input is:

```text
        4
       / \
      2   7
     / \
    1   3
```

The root node has value `4`. Its left child has value `2`, and its right child has value `7`. The node with value `2` has two children: `1` on the left and `3` on the right.

**Output:**

```text
[2,1,3]
```

**Explanation:**

The target value is `2`. The node containing `2` exists in the tree, so we return the subtree rooted at that node, including its left child (`1`) and right child (`3`).

The returned subtree is:

```text
    2
   / \
  1   3
```

## Example 2

**Input:**

```text
root = [4,2,7,1,3], val = 5
```

**Tree representation:**

The input tree is the same as in Example 1:

```text
        4
       / \
      2   7
     / \
    1   3
```

**Output:**

```text
[]
```

**Explanation:**

The target value is `5`. No node in the tree contains this value. Therefore, the function returns `null`.

In the output representation used by LeetCode, a null tree is serialized as an empty array (`[]`).

## Constraints

* The number of nodes in the tree is in the range `[1, 5000]`.
* `1 <= Node.val <= 10^7`
* `root` is a binary search tree.
* `1 <= val <= 10^7`

*/

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
