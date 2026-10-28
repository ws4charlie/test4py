
def find2nd(nums):
    if len(nums) == 0:
        return
    # sort and return
    nums.sort();
    if len(nums) == 1:
        return nums[0]
    else:
        return nums[-2]

def find2nd_lop(nums):
    if len(nums) == 0:
        return
    if len(nums) == 1:
        return nums[0]

    i, j = (0, 1) if nums[0] > nums[1] else (1, 0)
    for k in range(2, len(nums)):
        if nums[k] > nums[i]:
            j = i
            i = k
        elif nums[k] > nums[j]:
            j = k
    return j

if __name__ == "__main__":
    nums = [4,7,3,6,8,0,2,6]
    res = find2nd_lop(nums)
    print(res)