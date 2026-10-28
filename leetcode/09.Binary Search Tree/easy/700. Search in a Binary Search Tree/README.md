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
