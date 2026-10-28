import sys
import os
from heapq import heappop, heappush

hq=[]
heappush(hq, 2)
heappush(hq, 1)
heappush(hq, 3)

a = [[2, None, 2, None], 
      [2, None, 2, None], 
      [None, None, None, None], 
      [None, None, None, None]]

def foo():
    for y in range(4):
        for x in range(3):
            if a[x+1][y] != None:
                if a[x+1][y] == a[x][y]:
                    a[x][y] = a[x][y] * 2
                    a[x+1][y] = None
                if a[x][y] == None:
                    a[x][y] = a[x+1][y]
                    a[x+1][y] = None

ar1d = [None]*4
ar2d = [[None]*4]*4

class A:
    def foo(self):
        self.bar()

    def bar(self):
        print("from A")

class B(A):
    def foo(self):
        self.bar()

    def bar(self):
        print("from B")

if __name__ == "__main__":
    a = B()
    a.bar()

    nums = [-1,-2,0]
    nums.sort()
    for i in range(len(nums)):
            if nums[i] <= 0:
                continue
            else:
                break
    print(i)
