# https://leetcode.com/problems/letter-combinations-of-a-phone-number/

class Solution_recursive:
    hm = {
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz"}
    
    def back_trace(self, res, comb, digits):
        if not digits:
            res.append(comb)
            return
        
        for c in self.hm[digits[0]]:
            next_comb = comb + c
            self.back_trace(res, next_comb, digits[1:])
    
    def letterCombinations(self, digits: str) -> list:
        res = []
        if not digits:
            return res
        
        self.back_trace(res, '', digits)
        return res

class Solution_loop:
    hm = {
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz"}
    
    def letterCombinations(self, digits: str) -> List[str]:
        res = []
        if not digits:
            return res
        
        res.append("")
        while digits:
            next_res = []
            for s in res:
                for c in self.hm[digits[0]]:
                    next_res.append(s+c)
            digits = digits[1:]
            res = next_res
            
        return res

if __name__ == "__main__":
    sln1 = Solution17Solution_recursive()
    sln2 = Solution17Solution_recursive()
    res1 = sln1.letterCombinations('23')
    res2 = sln2.letterCombinations('23')
    print(res1)
    print(res2)