# https://leetcode.com/problems/group-anagrams/

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
