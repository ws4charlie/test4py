
# https://leetcode.com/problems/jump-game-iii/

class Solution_BFS:
    # i is the root
    def reach(self, arr, i, seen):
        if arr[i] == 0:
            return True
        else:
            seen.add(i)
        
        left = i - arr[i]
        if left >= 0 and left not in seen:
            if self.reach(arr, left, seen):
                return True
            
        right = i + arr[i]
        if right < len(arr) and right not in seen:
            if self.reach(arr, right, seen):
                return True
    
    # Treats start as the root a tree flatted in arr
    def canReach(self, arr: List[int], start: int) -> bool:
        seen = set()
        return self.reach(arr, start, seen)

class Solution_Topdown:
    # i is the root
    def reach(self, arr, i, mem, seen):
        if i in mem:
            return mem[i]
        
        seen.add(i)
        
        left = i - arr[i]
        if left >= 0 and left not in seen:
            if self.reach(arr, left, mem, seen):
                mem[i] = True
                return True
        
        right = i + arr[i]
        if right < len(arr) and right not in seen:
            if self.reach(arr, right, mem, seen):
                mem[i] = True
                return True
        
        mem[i] = False
        return False

    def canReach(self, arr: List[int], start: int) -> bool:
        if not arr:
            return False
        
        # init mem
        mem = {}
        for i, v in enumerate(arr):
            if v == 0:
                mem[i] = True
                
        # to avoid cycle search, keep posotions we have seen already
        seen = set()
        
        # start
        return self.reach(arr, start, mem, seen)
   
if __name__ == "__main__":
    sln = Solution()
    arr = [4,2,3,0,3,1,2]
    res = sln.canReach(arr)
