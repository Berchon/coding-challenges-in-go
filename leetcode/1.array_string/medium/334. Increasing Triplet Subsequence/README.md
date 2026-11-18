# 334. Increasing Triplet Subsequence

**Link:** https://leetcode.com/problems/increasing-triplet-subsequence/description/?envType=study-plan-v2&envId=leetcode-75

**Difficulty:** Medium

## Problem Description

You are given an integer array `nums`.

Return `true` if there exists a triple of indices `(i, j, k)` such that:

- `i < j < k`
- `nums[i] < nums[j] < nums[k]`

If no such indices exist, return `false`.

## Examples

**Example 1:**

```text
Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.
```

**Example 2:**

```text
Input: nums = [5,4,3,2,1]
Output: false
Explanation: No triplet exists.
```

**Example 3:**

```text
Input: nums = [2,1,5,0,4,6]
Output: true
Explanation: One of the valid triplets is (1, 4, 5),
because nums[1] == 1 < nums[4] == 4 < nums[5] == 6.
```

## Constraints

- `1 <= nums.length <= 5 * 10^5`
- `-2^31 <= nums[i] <= 2^31 - 1`

## Follow-up

Could you implement a solution that runs in:

- **Time Complexity:** `O(n)`
- **Space Complexity:** `O(1)`