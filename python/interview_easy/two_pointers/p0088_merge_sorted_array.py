"""
88. Merge Sorted Array  (Easy)
https://leetcode.com/problems/merge-sorted-array/

nums1 and nums2 are sorted in non-decreasing order and hold m and n
elements. nums1 has length m + n: its last n slots are 0 placeholders.
Merge nums2 into nums1 in place so nums1 ends up sorted. Return nothing.

Example 1:
    Input:  nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
    Output: [1,2,2,3,5,6]

        nums1: 1 2 3 _ _ _
        nums2: 2 5 6
        ->     1 2 2 3 5 6

Example 2:
    Input:  nums1 = [1], m = 1, nums2 = [], n = 0
    Output: [1]

Example 3:
    Input:  nums1 = [0], m = 0, nums2 = [1], n = 1
    Output: [1]
    Explanation: nums1 has no real elements; the 0 is only a placeholder.

Constraints:
    nums1.length == m + n, nums2.length == n
    0 <= m, n <= 200, 1 <= m + n <= 200
    -10^9 <= nums1[i], nums2[j] <= 10^9

Follow-up: can you do it in O(m + n) time?
"""

class Solution:
    def merge(self, nums1, m, nums2, n):
        """
        Do not return anything, modify nums1 in-place instead.
        """
        # fill from the back so nothing in nums1 is overwritten before use
        i, j, k = m - 1, n - 1, m + n - 1
        while j >= 0:
            if i >= 0 and nums1[i] > nums2[j]:
                nums1[k] = nums1[i]
                i -= 1
            else:
                nums1[k] = nums2[j]
                j -= 1
            k -= 1

if __name__ == "__main__":
    sln = Solution()
    nums1 = [1, 2, 3, 0, 0, 0]
    sln.merge(nums1, 3, [2, 5, 6], 3)
    assert nums1 == [1, 2, 2, 3, 5, 6]
    nums1 = [0]
    sln.merge(nums1, 0, [1], 1)
    assert nums1 == [1]
    nums1 = [4, 5, 6, 0, 0, 0]
    sln.merge(nums1, 3, [1, 2, 3], 3)
    assert nums1 == [1, 2, 3, 4, 5, 6]
    print("ok")
