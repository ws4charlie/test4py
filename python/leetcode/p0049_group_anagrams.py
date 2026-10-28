"""
49. Group Anagrams  (Medium)
https://leetcode.com/problems/group-anagrams/

Group the strings that are anagrams of each other (same letters, any
order). Return the groups in any order.

Example 1:
    Input:  strs = ["eat","tea","tan","ate","nat","bat"]
    Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:
    Input:  strs = [""]
    Output: [[""]]

Example 3:
    Input:  strs = ["a"]
    Output: [["a"]]

Constraints:
    1 <= strs.length <= 10^4
    0 <= strs[i].length <= 100
    strs[i] consists of lowercase English letters.
"""

from typing import List

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        n = len(strs)
        if n == 0:
            return []
        
        res = {}
        for s in strs:
            key = ''.join(sorted(s))
            if not key in res:
                res[key] = []
            res[key].append(s)
        
        return list(res.values())
   
if __name__ == "__main__":
    strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
    sln = Solution()
    res = sln.groupAnagrams(strs)
