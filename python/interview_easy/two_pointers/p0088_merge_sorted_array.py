# https://leetcode.com/problems/merge-sorted-array/

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
