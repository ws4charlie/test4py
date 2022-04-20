
class Solution_backtrace:
    def Fibonacci(self, n):# n > 0
        if n == 1:
            return 1
        if n == 2:
            return 1

        return self.Fibonacci(n-1) + self.Fibonacci(n-2)

class Solution_Topdown:
    def fib(self, n, mem):
        if n in mem:
            print("in mem, mem[{}] = {}".format(n, mem[n]))
            return mem[n]

        res_n_2 = self.fib(n-2, mem)
        res_n_1 = self.fib(n-1, mem)
        mem[n] = res_n_2 + res_n_1
        return mem[n]

    def Fibonacci(self, n):# n > 0
        mem = {}
        mem[1] = 1
        mem[2] = 1

        return self.fib(n, mem)

class Solution_Buttomup:
    def Fibonacci(self, n):# n > 0
        if n == 1 or n == 2:
            return 1

        mem_i_2 = 1
        mem_i_1 = 1
        for i in range(3, n+1):
            mem_i = mem_i_2 + mem_i_1
            mem_i_2 = mem_i_1
            mem_i_1 = mem_i

        return mem_i

if __name__ == "__main__":
    sln1 = Solution_backtrace()
    sln2 = Solution_Topdown()
    sln3 = Solution_Buttomup()
    n = 9
    print(sln1.Fibonacci(n))
    print(sln2.Fibonacci(n))
    print(sln3.Fibonacci(n))
