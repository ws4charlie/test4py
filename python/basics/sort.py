import heapq
import copy
import random

#---------------------------------------------------------------
#洗牌 Shuffle poker
#---------------------------------------------------------------
def rand_ij(inds):
    range = len(inds)
    i = random.randrange(range)
    j = random.randrange(range)
    while i == j:
        j = random.randrange(range)
    i, j = inds.pop(i), inds.pop(j)
    return i, j

def swap_shuffle_poker(deck):
    random.seed()

    inds = list(range(54))
    while inds:
        i, j = rand_ij(inds)
        deck[i], deck[j] = deck[j], deck[i]

# select a random piece and put to end
def select_shuffle_poker(deck):
    random.seed()
    for i in range(53):
        j = random.randrange(54 - i)
        deck.append(deck.pop(j))

deck = list(range(1, 55))
select_shuffle_poker(deck)

#---------------------------------------------------------------
# 冒泡排序 Bubble Sort: stable O(n^2)
#---------------------------------------------------------------
def bubble_sort(nums):
    n = len(nums)
    for i in range(n - 1):
        for j in range(n - 1, i, -1):
            if nums[j] < nums[j-1]:
                nums[j], nums[j-1] = nums[j-1], nums[j]

# stable O(n^2)
def bubble_sort_optmz(nums):
    n = len(nums)
    for i in range(n - 1):
        swapped = False
        for j in range(n - 1, i, -1):
            if nums[j] < nums[j-1]:
                nums[j], nums[j-1] = nums[j-1], nums[j]
                swapped = True
        if not swapped:
            break

#---------------------------------------------------------------
# 插入排序 Insertion Sort: stable O(n^2)
#---------------------------------------------------------------
def insertion_sort(nums):
    n = len(nums)
    for i in range(1, n):
        for j in range(i, 0, -1):#go backwards and swap
            if nums[j] < nums[j-1]:
                nums[j-1], nums[j] = nums[j], nums[j-1]
            else:
                break

# stable O(n^2)
def insertion_sort_1(nums):
    for i in range(1, len(nums)):
        for j in range(i):
            if nums[i] < nums[j]:
                val = nums.pop(i)
                nums.insert(j, val)
                break

#---------------------------------------------------------------
# 选择排序 Selection Sort: unstable O(n^2)
#---------------------------------------------------------------
def selection_sort(nums):
    n = len(nums)
    for i in range(n-1):
        min = i
        for j in range(i+1, n):
            if nums[j] < nums[min]:
                min = j
        if min != i:#swap breaks stability
            nums[i], nums[min] = nums[min], nums[i]

#---------------------------------------------------------------
# 归并排序 Merge Sort: stable O(nlogn), extra space
#---------------------------------------------------------------
def merge(left, right):
    res = []
    while left and right:
        if left[0] < right[0]:
            res.append(left.pop(0))
        else:
            res.append(right.pop(0))

    res += left + right # append remained elems
    return res

def merge_sort(nums):
    n = len(nums)
    if n <= 1:
        return nums

    mid = n // 2
    left = merge_sort(nums[:mid])
    right = merge_sort(nums[mid:])
    res = merge(left, right)

    return res

#---------------------------------------------------------------
# 归并排序-原地 Merge Sort: stable O(nlogn)
#---------------------------------------------------------------
def merge_inplace(nums, beg, mid, end):
    for i in range(mid+1, end+1):
        swapped = False
        while i > beg and nums[i] < nums[i-1]:
            nums[i], nums[i-1] = nums[i-1], nums[i]
            swapped = True
            i -= 1
        if not swapped:
            break

def merge_sort_inplace(nums, beg, end):
    if beg >= end:
        return

    mid = (beg + end) // 2
    merge_sort_inplace(nums, beg, mid)
    merge_sort_inplace(nums, mid+1, end)
    merge_inplace(nums, beg, mid, end)

#---------------------------------------------------------------
# 归并排序-非递归-原地 Merge Sort: stable O(nlogn)
#---------------------------------------------------------------
def merge_sort_loop_inplace(nums):
    n = len(nums)
    if n <= 1:
        return

    step = 2
    while step // 2 < n:
        steps = (n-1) // step + 1 # number of steps we need to go
        for i in range(steps):
            beg = step * i
            mid = beg + step // 2 - 1
            end = step * (i+1) - 1
            end = min(end, n-1) # end must be <= n-1
            merge_inplace(nums, beg, mid, end)
        step *= 2

#---------------------------------------------------------------
# 堆排序 Heap Sort: stable O(nlogn)
#---------------------------------------------------------------
# sift node k down to end
def siftdown(heap, k, end):
    lson = k*2 + 1
    while lson < end: # until k has no child
        max = lson
        rson = k*2 + 2
        if rson < end and heap[rson] > heap[lson]: # compare right son as well if it exists
            max = rson
        if heap[max] < heap[k]: # break if k is already the biggest
            break

        heap[max], heap[k] = heap[k], heap[max]  # swap with max-son
        k = max # goes down to max-son's position and sift again
        lson = k*2 + 1

def heap_sort(nums):
    heapq._heapify_max(nums)

    n = len(nums)
    for k in range(n - 1, 0, -1):
        nums[0], nums[k] = nums[k], nums[0]# swap max and end
        siftdown(nums, 0, k)

#---------------------------------------------------------------
# 快速排序 Quick Sort: unstable O(nlogn)
#---------------------------------------------------------------
def partition(nums, beg, end):
    pivot = nums[end]

    i, j = beg, end
    while i < j:
        while nums[i] < pivot:
            i += 1
        while j > i and nums[j] >= pivot:
            j -= 1
        if i < j:
            nums[i], nums[j] = nums[j], nums[i]

    # swap leftmost number(nums[i] >= pivot) with the pivot
    if i != end:
        nums[i], nums[end] = nums[end], nums[i]
    return i

def qsort(nums, beg, end):
    if beg < end:
        pos = partition(nums, beg, end)
        qsort(nums, beg, pos-1)
        qsort(nums, pos+1, end)

def quick_sort(nums):
    n = len(nums)
    qsort(nums, 0, n - 1)

#---------------------------------------------------------------
# 快速排序-V2 Quick Sort: unstable O(nlogn)
#---------------------------------------------------------------
def partition_v2(nums, beg, end):
    pivot = nums[end]

    i = beg - 1
    for j in range(beg, end):
        if nums[j] < pivot:
            i += 1
            nums[i], nums[j] = nums[j], nums[i]# i, j can be equal
    i += 1
    nums[i], nums[end] = nums[end], nums[i]
    return i

def qsort_v2(nums, beg, end):
    if beg < end:
        pos = partition_v2(nums, beg, end)
        qsort_v2(nums, beg, pos-1)
        qsort_v2(nums, pos+1, end)

def quick_sort_v2(nums):
    n = len(nums)
    qsort_v2(nums, 0, n - 1)

#---------------------------------------------------------------
# 快速排序-非递归 Quick Sort: unstable O(nlogn)
#---------------------------------------------------------------
def qsort_loop(nums, beg, end):
    if beg >= end:
        return

    s = [(beg, end)]
    while s:
        beg, end = s.pop()
        pos = partition_v2(nums, beg, end)
        if beg < pos - 1:
            s.append((beg, pos-1))
        if pos + 1 < end:
            s.append((pos+1, end))

def quick_sort_loop(nums):
    n = len(nums)
    qsort_loop(nums, 0, n - 1)

#---------------------------------------------------------------
# 快速排序-简单版 Quick Sort: unstable O(nlogn)
#---------------------------------------------------------------
def quick_sort_out_of_space(nums):
    n = len(nums)
    if n <= 1:
        return nums

    les = []
    gtr = []
    pivot = nums.pop()
    while nums:
        v = nums.pop()
        if v < pivot:
            les.append(v)
        else:
            gtr.append(v)
    left = quick_sort(les)
    right = quick_sort(gtr)
    return left + [pivot] + right

#---------------------------------------------------------------
# Test
#---------------------------------------------------------------
SAMPLE = list(range(100, 0, -1))
SAMPLE1 = [2,7,9,4,7,5,1]

if __name__ == '__main__':
    #冒泡排序
    nums = copy.deepcopy(SAMPLE)
    bubble_sort(nums)
    print(nums)

    #冒泡排序，优化
    nums = copy.deepcopy(SAMPLE)
    bubble_sort_optmz(nums)
    print(nums)

    #插入排序
    nums = copy.deepcopy(SAMPLE)
    insertion_sort(nums)
    print(nums)

    #选择排序
    nums = copy.deepcopy(SAMPLE)
    selection_sort(nums)
    print(nums)

    #归并排序
    nums = copy.deepcopy(SAMPLE)
    res = merge_sort(nums)
    print(res)

    #归并排序-原地
    nums = copy.deepcopy(SAMPLE)
    merge_sort_inplace(nums, 0, len(nums)-1)
    print(nums)

    #归并排序-原地-非递归
    nums = copy.deepcopy(SAMPLE)
    merge_sort_loop_inplace(nums)
    print(nums)

    #堆排序
    nums = copy.deepcopy(SAMPLE)
    heap_sort(nums)
    print(nums)

    #快速排序
    nums = copy.deepcopy(SAMPLE)
    quick_sort(nums)
    print(nums)

    #快速排序-非递归
    nums = copy.deepcopy(SAMPLE)
    quick_sort_loop(nums)
    print(nums)