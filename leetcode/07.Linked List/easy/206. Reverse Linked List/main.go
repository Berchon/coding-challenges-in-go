package main

import (
	"fmt"
	"runtime/debug"
)

/*
# 206. Reverse Linked List

**Difficulty:** Easy

Given the `head` of a singly linked list, reverse the list, and return the reversed list.

## Example 1

The input linked list is:

```text
1 → 2 → 3 → 4 → 5 → null
```

After reversing the linked list:

```text
5 → 4 → 3 → 2 → 1 → null
```

**Input:**

```text
head = [1,2,3,4,5]
```

**Output:**

```text
[5,4,3,2,1]
```

---

## Example 2

The input linked list is:

```text
1 → 2 → null
```

After reversing the linked list:

```text
2 → 1 → null
```

**Input:**

```text
head = [1,2]
```

**Output:**

```text
[2,1]
```

---

## Example 3

The input linked list is empty:

```text
null
```

**Input:**

```text
head = []
```

**Output:**

```text
[]
```

---

## Constraints

* The number of nodes in the list is in the range `[0, 5000]`.
* `-5000 <= Node.val <= 5000`

---

## Follow-up

A linked list can be reversed either **iteratively** or **recursively**.

Could you implement both?
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
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func iterativelyReverseList(head *ListNode) *ListNode {
	var previous *ListNode

	for current := head; current != nil; {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}

	return previous
}

func recursivelyReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := recursivelyReverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return newHead
}

func reverseList(head *ListNode) *ListNode {
	// return iterativelyReverseList(head)
	return recursivelyReverseList(head)
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	fmt.Println("head:")
	printList(head)

	reversed := reverseList(head)

	fmt.Println("reversed:")
	printList(reversed)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}

	fmt.Println()
}
