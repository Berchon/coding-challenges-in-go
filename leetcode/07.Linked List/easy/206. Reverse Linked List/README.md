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
