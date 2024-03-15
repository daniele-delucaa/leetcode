# easy solution (doesn't work on very large input)
# brute force method (but when we find a number that is 
# > than the target, we skip to the next index)
class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        res = []
        for i, n in enumerate(numbers):
            for j, n2 in enumerate(numbers):
                if i == j:
                    continue 
                if n+n2 > target:
                    break
                if n+n2 == target:
                    res.append(i+1)
                    res.append(j+1)
                    return res
        return res

# best solution: two pointers approach, one at the beginning,
# one at the end, if the sum is == to target return, 
# if sum is > than target we decrement the right pointer, 
# if sum is < than target we increment left pointer
class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        res = []
        l, r = 0, len(numbers)-1
        while len(res) == 0:
            if numbers[l] + numbers[r] == target:
                res.append(l+1)
                res.append(r+1)
            elif numbers[l] + numbers[r] > target:
                r -= 1
            else:
                l += 1
        return res