# 746. Min Cost Climbing Stairs

**Link:** https://leetcode.com/problems/min-cost-climbing-stairs/description/?envType=study-plan-v2&envId=leetcode-75

**Difficulty:** Easy

## Description

You are given an integer array `cost` where `cost[i]` is the cost of `i`th step on a staircase.

Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0, or the step with index 1.

Return the minimum cost to reach the top of the staircase, which is the position just past the last step (index `cost.length`).

## Examples

### Example 1

**Input:**

```text
cost = [10,15(*),20]
```

**Output:**

```text
15
```

**Explanation:**

You will start at index 1.

* Pay 15 and climb two steps to reach the top.

The total cost is 15.

### Example 2

**Input:**

```text
cost = [1(*), 100, 1(*), 1(*), 1(*), 100, 1(*), 1(*), 100, 1(*)]
```

**Output:**

```text
6
```

**Explanation:**

You will start at index 0.

* Pay 1 and climb two steps to reach index 2.
* Pay 1 and climb two steps to reach index 4.
* Pay 1 and climb two steps to reach index 6.
* Pay 1 and climb one step to reach index 7.
* Pay 1 and climb two steps to reach index 9.
* Pay 1 and climb one step to reach the top.

The total cost is 6.

## Constraints

* `2 <= cost.length <= 1000`
* `0 <= cost[i] <= 999`
