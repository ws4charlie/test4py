
class Solution:
    def isPalindrome(self, x: 'int') -> 'bool':
        if x < 0:
            return False
        
        x_str = str(abs(x))
        l, r = 0, len(x_str) - 1
        while l < r:
            if x_str[l] != x_str[r]:
                return False
            l += 1
            r -= 1
            
        return True

if __name__ == "__main__":
    sln = Solution()
    res = sln.isPalindrome(121)
