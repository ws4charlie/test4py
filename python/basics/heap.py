
# sift each node if its value > parent value
def siftup(heap, k):
    while k > 0:
        p = (k - 1) // 2      # get parent node
        if heap[k] < heap[p]: #swap k and parent if k < parent
            heap[k], heap[p] = heap[p], heap[k]
            k = p
        else:
            break

# sift each node down (from second to last level) if its value > min(lson, rson)
def siftdown(heap, k):
    n = len(heap)
    lson = k*2 + 1
    while lson < n: # until k has no child
        min = lson
        rson = k*2 + 2
        if rson < n and heap[rson] < heap[lson]: # compare as well if right son exists
            min = rson
        if heap[min] > heap[k]: # break if k is already the smallest
            break

        heap[min], heap[k] = heap[k], heap[min]  # swap value
        k = min # goes down to min-son's position and sift again
        lson = k*2 + 1

# heapify: sift up version
def heapify_up(heap):
    n = len(heap)
    for k in range(1, n):
        siftup(heap, k)

# heapify: sift down version
def heapify_down(heap):
    n = len(heap)
    sec2last = (n-1) // 2
    for k in range(sec2last, -1, -1):
        siftdown(heap, k)

def heappush(heap, v):
    heap.append(v)
    siftup(heap, len(heap) - 1) # sift it up

def heappop(heap):
    last = heap.pop() # throw exception if heap is empty
    if heap:
        ret = heap[0]
        heap[0] = last
        siftdown(heap, 0) # sift heap[0] down
        return ret
    return last # return last when size of heap is 1